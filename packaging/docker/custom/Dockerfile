ARG PLUTONO_VERSION="latest"

FROM plutono/plutono:${PLUTONO_VERSION}

USER root

ARG PL_INSTALL_IMAGE_RENDERER_PLUGIN="false"

ARG PL_GID="0"
ENV PL_PATHS_PLUGINS="/var/lib/plutono-plugins"

RUN mkdir -p "$PL_PATHS_PLUGINS" && \
    chown -R plutono:${PL_GID} "$PL_PATHS_PLUGINS"

RUN if [ $PL_INSTALL_IMAGE_RENDERER_PLUGIN = "true" ]; then \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk --no-cache  upgrade && \
    apk add --no-cache udev ttf-opensans chromium && \
    rm -rf /tmp/* && \
    rm -rf /usr/share/plutono/tools/phantomjs; \
fi

USER plutono

ENV PL_RENDERER_PLUGIN_CHROME_BIN="/usr/bin/chromium-browser"

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
