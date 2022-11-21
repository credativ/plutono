ARG BASE_IMAGE=alpine:3.13
FROM ${BASE_IMAGE}

ARG PLUTONO_TGZ="plutono-latest.linux-x64-musl.tar.gz"

# Make sure we have Gnu tar
RUN apk add --no-cache tar

COPY ${PLUTONO_TGZ} /tmp/plutono.tar.gz

# Change to tar xfzv to make tar print every file it extracts
RUN mkdir /tmp/plutono && tar xzf /tmp/plutono.tar.gz --strip-components=1 -C /tmp/plutono

FROM ${BASE_IMAGE}

ARG PL_UID="472"
ARG PL_GID="0"

ENV PATH=/usr/share/plutono/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin \
    PL_PATHS_CONFIG="/etc/plutono/plutono.ini" \
    PL_PATHS_DATA="/var/lib/plutono" \
    PL_PATHS_HOME="/usr/share/plutono" \
    PL_PATHS_LOGS="/var/log/plutono" \
    PL_PATHS_PLUGINS="/var/lib/plutono/plugins" \
    PL_PATHS_PROVISIONING="/etc/plutono/provisioning"

WORKDIR $PL_PATHS_HOME

RUN apk add --no-cache ca-certificates bash tzdata && \
    apk add --no-cache openssl musl-utils

# Oracle Support for x86_64 only
RUN if [ `arch` = "x86_64" ]; then \
      apk add --no-cache libaio libnsl && \
      ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
      wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.30-r0/glibc-2.30-r0.apk \
        -O /tmp/glibc-2.30-r0.apk && \
      wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.30-r0/glibc-bin-2.30-r0.apk \
        -O /tmp/glibc-bin-2.30-r0.apk && \
      apk add --no-cache --allow-untrusted /tmp/glibc-2.30-r0.apk /tmp/glibc-bin-2.30-r0.apk && \
      rm -f /tmp/glibc-2.30-r0.apk && \
      rm -f /tmp/glibc-bin-2.30-r0.apk && \
      rm -f /lib/ld-linux-x86-64.so.2 && \
      rm -f /etc/ld.so.cache; \
    fi

COPY --from=0 /tmp/plutono "$PL_PATHS_HOME"

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

EXPOSE 3000

COPY ./run.sh /run.sh

USER "$PL_UID"
ENTRYPOINT [ "/run.sh" ]
