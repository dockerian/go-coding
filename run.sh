#!/usr/bin/env bash
# This script will run "make ${1:-test}" inside a Docker container
set -e
script_file="${BASH_SOURCE[0]##*/}"
script_path="$( cd "$( echo "${BASH_SOURCE[0]%/*}" )" && pwd )"

# main function
function main() {
  getOS

  ARGS="$@"
  PROJECT="go-coding"
  PROJECT_IMAGE=$(docker images | grep ${PROJECT} | awk '{ print $1; }')
  GITHUB_USER="dockerian"
  SOURCE_PATH="src/github.com/${GITHUB_USER}/${PROJECT}"
  BUILD_OS=${BUILD_OS:-${OS}}

  if [[ "${BUILD_OS}" == "" ]]; then
    echo -e "\nFailed: Unset/unknown BUILD_OS in environment variables"
    exit 1
  fi

  cd -P "${script_path}"

  echo -e "\nLooking docker image [${PROJECT_IMAGE}] for ${PROJECT} ..."
  if [[ "${PROJECT_IMAGE}" != "${PROJECT}" ]]; then
    echo -e "\nBuilding docker container for ${PROJECT} ..."
    docker build -t ${PROJECT} .
  fi

  RUN_TARGET="${ARGS:-test}"

  echo -e "\nRunning 'make ${RUN_TARGET}' in docker container ..."
  docker run --rm \
    --hostname ${PROJECT} \
    --name ${PROJECT} \
    -e DEBUG=${DEBUG} \
    -e BUILD_OS=${BUILD_OS} \
    -e BUILD_MASTER_VERSION \
    -e BUILD_VERSION \
    -e BINARY \
    -e TEST_VERBOSE \
    -e TEST_DIR \
    -e TEST_MATCH \
    -e TEST_TAGS \
    -e GITHUB_USERNAME \
    -e GITHUB_PASSWORD \
    -e VERBOSE \
    -v "${PWD}":/go/${SOURCE_PATH} \
    ${PROJECT} bash -c "make ${RUN_TARGET}"
}

# getOS function sets OS environment variable in runtime
function getOS() {
  # Detect the platform (similar to $OSTYPE)
  UNAME="$(uname)"
  case ${UNAME} in
    'Darwin')
      OS="darwin"
      ;;
    'FreeBSD')
      OS="bsd"
      alias ls='ls -G'
      ;;
    'Linux')
      OS="linux"
      alias ls='ls --color=auto'
      ;;
    'SunOS')
      OS="solaris"
      ;;
    'WindowsNT')
      OS="windows"
      ;;
    'AIX') ;;
    *) ;;
  esac

  if [[ "${OS}" != "" ]]; then return; fi

  case "${OSTYPE}" in
    bsd*)     OS="bsd" ;;
    darwin*)  OS="darwin" ;;
    linux*)   OS="linux" ;;
    solaris*) OS="soloris" ;;
    *)        OS="" ;;
  esac
}

[[ $0 != "${BASH_SOURCE}" ]] || main "$@"
