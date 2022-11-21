#!/bin/bash -e

PERMISSIONS_OK=0

if [ ! -r "$PL_PATHS_CONFIG" ]; then
    echo "PL_PATHS_CONFIG='$PL_PATHS_CONFIG' is not readable."
    PERMISSIONS_OK=1
fi

if [ ! -w "$PL_PATHS_DATA" ]; then
    echo "PL_PATHS_DATA='$PL_PATHS_DATA' is not writable."
    PERMISSIONS_OK=1
fi

if [ ! -r "$PL_PATHS_HOME" ]; then
    echo "PL_PATHS_HOME='$PL_PATHS_HOME' is not readable."
    PERMISSIONS_OK=1
fi

if [ $PERMISSIONS_OK -eq 1 ]; then
    echo "You may have issues with file permissions, more information here: http://docs.plutono.org/installation/docker/#migrate-to-v51-or-later"
fi

if [ ! -d "$PL_PATHS_PLUGINS" ]; then
    mkdir "$PL_PATHS_PLUGINS"
fi

if [ ! -z ${PL_AWS_PROFILES+x} ]; then
    > "$PL_PATHS_HOME/.aws/credentials"

    for profile in ${PL_AWS_PROFILES}; do
        access_key_varname="PL_AWS_${profile}_ACCESS_KEY_ID"
        secret_key_varname="PL_AWS_${profile}_SECRET_ACCESS_KEY"
        region_varname="PL_AWS_${profile}_REGION"

        if [ ! -z "${!access_key_varname}" -a ! -z "${!secret_key_varname}" ]; then
            echo "[${profile}]" >> "$PL_PATHS_HOME/.aws/credentials"
            echo "aws_access_key_id = ${!access_key_varname}" >> "$PL_PATHS_HOME/.aws/credentials"
            echo "aws_secret_access_key = ${!secret_key_varname}" >> "$PL_PATHS_HOME/.aws/credentials"
            if [ ! -z "${!region_varname}" ]; then
                echo "region = ${!region_varname}" >> "$PL_PATHS_HOME/.aws/credentials"
            fi
        fi
    done

    chmod 600 "$PL_PATHS_HOME/.aws/credentials"
fi

# Convert all environment variables with names ending in __FILE into the content of
# the file that they point at and use the name without the trailing __FILE.
# This can be used to carry in Docker secrets.
for VAR_NAME in $(env | grep '^PL_[^=]\+__FILE=.\+' | sed -r "s/([^=]*)__FILE=.*/\1/g"); do
    VAR_NAME_FILE="$VAR_NAME"__FILE
    if [ "${!VAR_NAME}" ]; then
        echo >&2 "ERROR: Both $VAR_NAME and $VAR_NAME_FILE are set (but are exclusive)"
        exit 1
    fi
    echo "Getting secret $VAR_NAME from ${!VAR_NAME_FILE}"
    export "$VAR_NAME"="$(< "${!VAR_NAME_FILE}")"
    unset "$VAR_NAME_FILE"
done

export HOME="$PL_PATHS_HOME"

if [ ! -z "${PL_INSTALL_PLUGINS}" ]; then
  OLDIFS=$IFS
  IFS=','
  for plugin in ${PL_INSTALL_PLUGINS}; do
    IFS=$OLDIFS
    if [[ $plugin =~ .*\;.* ]]; then
        pluginUrl=$(echo "$plugin" | cut -d';' -f 1)
        pluginWithoutUrl=$(echo "$plugin" | cut -d';' -f 2)
        plutono-cli --pluginUrl "${pluginUrl}" --pluginsDir "${PL_PATHS_PLUGINS}" plugins install ${pluginWithoutUrl}
    else
        plutono-cli --pluginsDir "${PL_PATHS_PLUGINS}" plugins install ${plugin}
    fi
  done
fi

exec plutono-server                                         \
  --homepath="$PL_PATHS_HOME"                               \
  --config="$PL_PATHS_CONFIG"                               \
  --packaging=docker                                        \
  "$@"                                                      \
  cfg:default.log.mode="console"                            \
  cfg:default.paths.data="$PL_PATHS_DATA"                   \
  cfg:default.paths.logs="$PL_PATHS_LOGS"                   \
  cfg:default.paths.plugins="$PL_PATHS_PLUGINS"             \
  cfg:default.paths.provisioning="$PL_PATHS_PROVISIONING"
