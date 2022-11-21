ARG BASE_IMAGE=ubuntu:20.04
FROM ${BASE_IMAGE} AS plutono-builder

ARG PLUTONO_TGZ="plutono-latest.linux-x64.tar.gz"

COPY ${PLUTONO_TGZ} /tmp/plutono.tar.gz

RUN mkdir /tmp/plutono && tar xzf /tmp/plutono.tar.gz --strip-components=1 -C /tmp/plutono

FROM ${BASE_IMAGE}

EXPOSE 3000

# Set DEBIAN_FRONTEND=noninteractive in environment at build-time
ARG DEBIAN_FRONTEND=noninteractive
ARG GF_UID="472"
ARG GF_GID="0"

ENV PATH=/usr/share/plutono/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin \
    GF_PATHS_CONFIG="/etc/plutono/plutono.ini" \
    GF_PATHS_DATA="/var/lib/plutono" \
    GF_PATHS_HOME="/usr/share/plutono" \
    GF_PATHS_LOGS="/var/log/plutono" \
    GF_PATHS_PLUGINS="/var/lib/plutono/plugins" \
    GF_PATHS_PROVISIONING="/etc/plutono/provisioning"

WORKDIR $GF_PATHS_HOME

# Install dependencies
# We need curl in the image
RUN apt-get update && apt-get install -y ca-certificates curl tzdata && \
    apt-get autoremove -y && rm -rf /var/lib/apt/lists/*;

COPY --from=plutono-builder /tmp/plutono "$GF_PATHS_HOME"

RUN if [ ! $(getent group "$GF_GID") ]; then \
      addgroup --system --gid $GF_GID plutono; \
    fi

RUN export GF_GID_NAME=$(getent group $GF_GID | cut -d':' -f1) && \
    mkdir -p "$GF_PATHS_HOME/.aws" && \
    adduser --system --uid $GF_UID --ingroup "$GF_GID_NAME" plutono && \
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

COPY ./run.sh /run.sh

USER "$GF_UID"
ENTRYPOINT [ "/run.sh" ]
