FROM node:14.15.5-alpine3.13 as js-builder

RUN apk add --no-cache git

WORKDIR /usr/src/app/

COPY package.json yarn.lock ./
COPY packages packages

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

ENV NODE_ENV production
RUN yarn build

FROM golang:1.16.1-alpine3.13 as go-builder

RUN apk add --no-cache gcc g++

WORKDIR $GOPATH/src/github.com/credativ/plutono

COPY go.mod go.sum build.go package.json ./
COPY pkg pkg

RUN go mod verify
RUN go run build.go build

# Final stage
FROM alpine:3.13

LABEL maintainer="Plutono team <hello@grafana.com>"

ARG GF_UID="472"
ARG GF_GID="0"

ENV PATH="/usr/share/plutono/bin:$PATH" \
    GF_PATHS_CONFIG="/etc/plutono/plutono.ini" \
    GF_PATHS_DATA="/var/lib/plutono" \
    GF_PATHS_HOME="/usr/share/plutono" \
    GF_PATHS_LOGS="/var/log/plutono" \
    GF_PATHS_PLUGINS="/var/lib/plutono/plugins" \
    GF_PATHS_PROVISIONING="/etc/plutono/provisioning"

WORKDIR $GF_PATHS_HOME

RUN apk add --no-cache ca-certificates bash tzdata && \
    apk add --no-cache openssl musl-utils

COPY conf ./conf

RUN if [ ! $(getent group "$GF_GID") ]; then \
      addgroup -S -g $GF_GID plutono; \
    fi

RUN export GF_GID_NAME=$(getent group $GF_GID | cut -d':' -f1) && \
    mkdir -p "$GF_PATHS_HOME/.aws" && \
    adduser -S -u $GF_UID -G "$GF_GID_NAME" plutono && \
    mkdir -p "$GF_PATHS_PROVISIONING/datasources" \
             "$GF_PATHS_PROVISIONING/dashboards" \
             "$GF_PATHS_PROVISIONING/notifiers" \
             "$GF_PATHS_PROVISIONING/plugins" \
             "$GF_PATHS_LOGS" \
             "$GF_PATHS_PLUGINS" \
             "$GF_PATHS_DATA" && \
    cp "$GF_PATHS_HOME/conf/sample.ini" "$GF_PATHS_CONFIG" && \
    cp "$GF_PATHS_HOME/conf/ldap.toml" /etc/plutono/ldap.toml && \
    chown -R "plutono:$GF_GID_NAME" "$GF_PATHS_DATA" "$GF_PATHS_HOME/.aws" "$GF_PATHS_LOGS" "$GF_PATHS_PLUGINS" "$GF_PATHS_PROVISIONING" && \
    chmod -R 777 "$GF_PATHS_DATA" "$GF_PATHS_HOME/.aws" "$GF_PATHS_LOGS" "$GF_PATHS_PLUGINS" "$GF_PATHS_PROVISIONING"

COPY --from=go-builder /go/src/github.com/credativ/plutono/bin/linux-amd64/plutono-server /go/src/github.com/credativ/plutono/bin/linux-amd64/plutono-cli ./bin/
COPY --from=js-builder /usr/src/app/public ./public
COPY --from=js-builder /usr/src/app/tools ./tools

EXPOSE 3000

COPY ./packaging/docker/run.sh /run.sh

USER plutono
ENTRYPOINT [ "/run.sh" ]
