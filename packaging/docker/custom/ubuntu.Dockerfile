ARG PLUTONO_VERSION="latest"

FROM plutono/plutono:${PLUTONO_VERSION}-ubuntu

USER root

# Set DEBIAN_FRONTEND=noninteractive in environment at build-time
ARG DEBIAN_FRONTEND=noninteractive

ARG PL_INSTALL_IMAGE_RENDERER_PLUGIN="false"

ARG PL_GID="0"
ENV PL_PATHS_PLUGINS="/var/lib/plutono-plugins"

RUN mkdir -p "$PL_PATHS_PLUGINS" && \
    chown -R plutono:${PL_GID} "$PL_PATHS_PLUGINS"

RUN if [ $PL_INSTALL_IMAGE_RENDERER_PLUGIN = "true" ]; then \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y gdebi-core && \
    cd /tmp && \
    curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    gdebi --n google-chrome-stable_current_amd64.deb && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/*; \
fi

USER plutono

ENV PL_RENDERER_PLUGIN_CHROME_BIN="/usr/bin/google-chrome"

RUN if [ $PL_INSTALL_IMAGE_RENDERER_PLUGIN = "true" ]; then \
    plutono-cli \
        --pluginsDir "$PL_PATHS_PLUGINS" \
        --pluginUrl https://github.com/grafana/grafana-image-renderer/releases/latest/download/plugin-linux-x64-glibc-no-chromium.zip \
        plugins install plutono-image-renderer; \
fi

ARG PL_INSTALL_PLUGINS=""

RUN if [ ! -z "${PL_INSTALL_PLUGINS}" ]; then \
    OLDIFS=$IFS; \
        IFS=','; \
    for plugin in ${PL_INSTALL_PLUGINS}; do \
        IFS=$OLDIFS; \
        plutono-cli --pluginsDir "$PL_PATHS_PLUGINS" plugins install ${plugin}; \
    done; \
fi
