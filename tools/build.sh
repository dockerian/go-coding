#!/usr/bin/env bash
# This script will run "make build" inside a Docker container
# Note: setting 'BUILD_OS' to get build for specific platform
set -e
SOURCE_FILE="${BASH_SOURCE[0]##*/}"
SOURCE_PATH="$( cd "$( echo "${BASH_SOURCE[0]%/*}" )" && pwd )"

if [[ ! -x "${SOURCE_PATH}/run.sh" ]]; then
  echo "Cannot find '${SOURCE_PATH}/run.sh' to run build."
else
  "${SOURCE_PATH}/run.sh" build
fi
