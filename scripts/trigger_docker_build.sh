#!/bin/bash

_circle_token=$1
_plutono_version=$2

trigger_build_url="https://circleci.com/api/v1/project/plutono/plutono-docker/tree/master?circle-token=${_circle_token}"

post_data=$(cat <<EOF
{
  "build_parameters": {
    "PLUTONO_VERSION": "${_plutono_version}"
  }
}
EOF
)

echo "${post_data}"

curl \
--header "Accept: application/json" \
--header "Content-Type: application/json" \
--data "${post_data}" \
--request POST "${trigger_build_url}"
