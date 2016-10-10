#!/usr/bin/env bash
# This script will run "make ${1:-test}" inside a Docker container
set -e
script_file="${BASH_SOURCE[0]##*/}"
script_path="$( cd "$( echo "${BASH_SOURCE[0]%/*}" )" && pwd )"

# main function
function main() {
  getOS

  ARGS="$@"
  GITHUB_REPO="go-coding"
  GITHUB_USER="dockerian"
  DOCKER_IMAG="${GITHUB_USER}/${GITHUB_REPO}"
  DOCKER_TAGS=$(docker images 2>&1|grep ${DOCKER_IMAG}|awk '{print $1;}')
  # detect if running inside the container
  DOCKER_PROC="$(cat /proc/1/cgroup 2>&1|grep -e "/docker/[0-9a-z]\{64\}"|head -1)"
  SOURCE_PATH="src/github.com/${GITHUB_USER}/${GITHUB_REPO}"
  BUILD_OS=${BUILD_OS:-${OS}}

  if [[ "${BUILD_OS}" == "" ]]; then
    echo -e "\nFailed: Unset/unknown BUILD_OS in environment variables"
    exit 1
  fi

  cd -P "${script_path}"

  RUN_TARGET="${ARGS:-test}"
  if [[ -e "/.dockerenv" ]] || [[ "${DOCKER_PROC}" != "" ]]; then
    if [[ "${RUN_TARGET}" != "cmd" ]]; then
      make ${RUN_TARGET}
    else
      echo "Env in the container:"
      make show-env
    fi
    return
  fi

  echo -e "\nChecking docker image [${DOCKER_IMAG}] for '${GITHUB_REPO}'"
  if [[ "${DOCKER_TAGS}" != "${DOCKER_IMAG}" ]]; then
    echo -e "\nBuilding docker image [${DOCKER_IMAG}] for '${GITHUB_REPO}'"
    echo "------------------------------------------------------------"
    docker build -t ${DOCKER_IMAG} .
    echo "------------------------------------------------------------"
  fi

  CMD="docker run -it --rm
  --hostname ${GITHUB_REPO}
  --name ${GITHUB_REPO}
  -e DEBUG=${DEBUG}
  -e BUILD_OS=${BUILD_OS}
  -e BUILD_MASTER_VERSION
  -e BUILD_VERSION
  -e BINARY
  -e TEST_VERBOSE=${TEST_VERBOSE}
  -e TEST_DIR
  -e TEST_MATCH
  -e TEST_TAGS
  -e GITHUB_USERNAME
  -e GITHUB_PASSWORD
  -e VERBOSE=${VERBOSE}
  -v "${PWD}":/go/${SOURCE_PATH}
  ${DOCKER_IMAG} "

  echo -e "\nRunning 'make ${RUN_TARGET}' in docker container"
  echo "${CMD} bash -c \"make ${RUN_TARGET}\""
  echo "............................................................"
  if [[ "${RUN_TARGET}" != "cmd" ]]; then
    ${CMD} bash -c "make ${RUN_TARGET}"
  else
    ${CMD}
  fi
  echo "............................................................"
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
