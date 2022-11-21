ARG BASE_IMAGE=ubuntu:20.04
FROM ${BASE_IMAGE} AS plutono-builder

ARG PLUTONO_TGZ="plutono-latest.linux-x64.tar.gz"

COPY ${PLUTONO_TGZ} /tmp/plutono.tar.gz

RUN mkdir /tmp/plutono && tar xzf /tmp/plutono.tar.gz --strip-components=1 -C /tmp/plutono

FROM ${BASE_IMAGE}

EXPOSE 3000

# Set DEBIAN_FRONTEND=noninteractive in environment at build-time
ARG DEBIAN_FRONTEND=noninteractive
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

# Install dependencies
# We need curl in the image
RUN apt-get update && apt-get install -y ca-certificates curl tzdata && \
    apt-get autoremove -y && rm -rf /var/lib/apt/lists/*;

COPY --from=plutono-builder /tmp/plutono "$PL_PATHS_HOME"

RUN if [ ! $(getent group "$PL_GID") ]; then \
      addgroup --system --gid $PL_GID plutono; \
    fi

RUN export PL_GID_NAME=$(getent group $PL_GID | cut -d':' -f1) && \
    mkdir -p "$PL_PATHS_HOME/.aws" && \
    adduser --system --uid $PL_UID --ingroup "$PL_GID_NAME" plutono && \
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

COPY ./run.sh /run.sh

USER "$PL_UID"
ENTRYPOINT [ "/run.sh" ]
