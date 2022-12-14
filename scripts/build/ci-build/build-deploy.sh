#!/bin/bash
set -eo pipefail

_version="1.4.1"
_tag="plutono/build-container:${_version}"

_dpath=$(dirname "${BASH_SOURCE[0]}")
cd "$_dpath"

docker build -t $_tag .
docker push $_tag
