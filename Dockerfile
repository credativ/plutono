# Frontend builder stage
FROM --platform=$BUILDPLATFORM docker.io/library/node:14.15.5-alpine3.13 as js-builder

RUN apk add --no-cache git

WORKDIR /usr/src/app/

COPY package.json yarn.lock ./
COPY packages packages
COPY patches patches

RUN --mount=type=secret,id=NODE_AUTH_TOKEN \
    echo "//npm.pkg.github.com/:_authToken=$(cat /run/secrets/NODE_AUTH_TOKEN)" > .npmrc && \
    echo "@credativ:registry=https://npm.pkg.github.com" >> .npmrc && \
    yarn install --pure-lockfile --no-progress --network-concurrency 1 && \
    rm .npmrc

COPY tsconfig.json .eslintrc .editorconfig .browserslistrc .prettierrc.js ./
COPY public public
COPY tools tools
COPY scripts scripts
COPY emails emails

ENV NODE_ENV=production \
    NODE_OPTIONS="--max_old_space_size=4096"
RUN yarn build

# Backend builder stage
FROM docker.io/library/golang:1.23.4-alpine3.20 as go-builder

RUN apk add --no-cache gcc g++

WORKDIR $GOPATH/src/github.com/credativ/plutono

COPY go.mod go.sum build.go package.json ./
COPY pkg pkg

ARG SOURCE_GIT_BRANCH=""
ENV SOURCE_GIT_BRANCH=$SOURCE_GIT_BRANCH
ARG SOURCE_GIT_REV_SHORT=""
ENV SOURCE_GIT_REV_SHORT=$SOURCE_GIT_REV_SHORT

RUN go mod verify
RUN go run build.go build && \
    mv /go/src/github.com/credativ/plutono/bin/linux-$(go env GOARCH)/plutono-server /go/src/github.com/credativ/plutono/bin/ && \
    mv /go/src/github.com/credativ/plutono/bin/linux-$(go env GOARCH)/plutono-cli /go/src/github.com/credativ/plutono/bin/

# Final stage
FROM docker.io/library/alpine:3.21.0 as final

ARG PL_UID="472"
ARG PL_GID="0"

ENV PATH="/usr/share/plutono/bin:$PATH" \
    PL_PATHS_CONFIG="/etc/plutono/plutono.ini" \
    PL_PATHS_DATA="/var/lib/plutono" \
    PL_PATHS_HOME="/usr/share/plutono" \
    PL_PATHS_LOGS="/var/log/plutono" \
    PL_PATHS_PLUGINS="/var/lib/plutono/plugins" \
    PL_PATHS_PROVISIONING="/etc/plutono/provisioning"

WORKDIR $PL_PATHS_HOME

RUN apk add --no-cache ca-certificates bash tzdata && \
    apk add --no-cache openssl musl-utils

COPY conf ./conf

RUN if [ ! $(getent group "$PL_GID") ]; then \
      addgroup -S -g $PL_GID plutono; \
    fi

RUN export PL_GID_NAME=$(getent group $PL_GID | cut -d':' -f1) && \
    mkdir -p "$PL_PATHS_HOME/.aws" && \
    adduser -S -u $PL_UID -G "$PL_GID_NAME" plutono && \
    mkdir -p "$PL_PATHS_PROVISIONING/datasources" \
             "$PL_PATHS_PROVISIONING/dashboards" \
             "$PL_PATHS_PROVISIONING/notifiers" \
             "$PL_PATHS_PROVISIONING/plugins" \
             "$PL_PATHS_LOGS" \
             "$PL_PATHS_PLUGINS" \
             "$PL_PATHS_DATA" && \
    cp "$PL_PATHS_HOME/conf/sample.ini" "$PL_PATHS_CONFIG" && \
    cp "$PL_PATHS_HOME/conf/ldap.toml" /etc/plutono/ldap.toml && \
    chown -R "plutono:$PL_GID_NAME" "$PL_PATHS_DATA" "$PL_PATHS_HOME/.aws" "$PL_PATHS_LOGS" "$PL_PATHS_PLUGINS" "$PL_PATHS_PROVISIONING" && \
    chmod -R 777 "$PL_PATHS_DATA" "$PL_PATHS_HOME/.aws" "$PL_PATHS_LOGS" "$PL_PATHS_PLUGINS" "$PL_PATHS_PROVISIONING"

COPY --from=go-builder /go/src/github.com/credativ/plutono/bin/plutono-server /go/src/github.com/credativ/plutono/bin/plutono-cli ./bin/
COPY --from=js-builder /usr/src/app/public ./public
COPY --from=js-builder /usr/src/app/tools ./tools

EXPOSE 3000

COPY ./packaging/docker/run.sh /run.sh

USER plutono
ENTRYPOINT [ "/run.sh" ]
