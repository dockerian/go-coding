#!/usr/bin/env bash
############################################################
# Run "make ${1:-test}" inside a Docker container
#
# Injectable environment variables:
#   GITHUB_REPO - the github repository name (default: go-coding)
#   GITHUB_USER - the github user (default: dockerian)
#   DOCKER_USER - the docker user (default: ${GITHUB_USER})
#   DOCKER_NAME - default: go-coding, as in ${DOCKER_USER}/${DOCKER_NAME}
#   DOCKER_PORT - default: 8080, http port mapping to API service
#   PROJECT_DIR - the base path to the project, default: ..
#   DEBUG       - debug/verbose flag, default: 0
#
############################################################
set -e
script_file="${BASH_SOURCE[0]##*/}"
script_base="$( cd "$( echo "${BASH_SOURCE[0]%/*}/.." )" && pwd )"
script_path="$( cd "$( echo "${BASH_SOURCE[0]%/*}" )" && pwd )"

GITHUB_REPO="${GITHUB_REPO:-go-coding}"
GITHUB_USER="${GITHUB_USER:-dockerian}"
DOCKER_USER="${DOCKER_USER:-${GITHUB_USER}}"
DOCKER_NAME="${DOCKER_NAME:-go-coding}"
DOCKER_IMAG="${DOCKER_USER}/${DOCKER_NAME}"
DOCKER_HOST="${DOCKER_HOST:-${DOCKER_NAME}}"
DOCKER_HOST_NAME="${DOCKER_HOST_NAME:-${DOCKER_NAME}}"
DOCKER_PORT="${DOCKER_PORT:-8080}"
DOCKER_PORT_TEST="$(((8080 + 10000) % 65535))"
DOCKER_EXEC="run"

DOCKER_TAGS=$(docker images 2>&1|grep ${DOCKER_IMAG}|awk '{print $1;}')
# detect if the process running inside the container
DOCKER_PROC="$(cat /proc/1/cgroup 2>&1|grep -e "/docker/[0-9a-z]\{64\}"|head -1)"
SOURCE_PATH="/go/src/github.com/${GITHUB_USER}/${GITHUB_REPO}"

GO_PATH_SRC="${GOPATH}/src"
PROJECT_DIR="${PROJECT_DIR:-${script_base}}"
DEBUG="${DEBUG:-${VERBOSE:-0}}"


# main function
function main() {
  ARGS="$@"

  shopt -s nocasematch
  if [[ "${DEBUG}" =~ (1|enable|on|true|yes) ]]; then DEBUG="1"; fi

  cd -P "${PROJECT_DIR}" && pwd
  check_args || return
  check_docker_image || return
  docker_run
}

function check_args() {
  MAKE_ARGS="${ARGS:-test}"
  # get a list of make targets
  MAKE_LIST="$(make -qp|awk -F':' '/^[a-zA-Z0-9][^$#\/\t=]*:([^=]|$)/ {split($1,A,/ /);for(i in A)print A[i]}'|sort)"
  MAKE_BASH=""
  # checking target(s) from command line
  for target in ${MAKE_ARGS}; do
    echo -e "\nChecking target: ${target} ..."
    t="`echo ${MAKE_LIST}|xargs -n1 echo|grep -e \"^${target}$\"`"
    if [[ ! -n "$t" ]]; then
      echo "Makefile does not have target: ${target}"
      return 1
    elif [[ -e "/.dockerenv" ]] || [[ "${DOCKER_PROC}" != "" ]]; then
      if [[ "${target}" =~ (show) ]] || [[ "${target}" =~ (cover) ]]; then
        echo "Cannot open test coverage in the container."
        echo "See: cover/index.html"
        return 2
      elif [[ "${target}" == "sql" ]]; then
        echo "Cannot start MySQLWorkbench in the container."
        return 3
      fi
    elif [[ "${target}" =~ (check|only|test) ]]; then
      DOCKER_HOST_NAME="${DOCKER_HOST_NAME}-test"
      DOCKER_PORT="${DOCKER_PORT_TEST}"
      echo -e "\nUsing docker container: ${DOCKER_HOST_NAME}:${DOCKER_PORT}"
    elif [[ "${target}" == "cmd" ]]; then
      if [[ "$(docker ps -q -f name=${DOCKER_NAME})" != "" ]]; then
        DOCKER_EXEC=exec
      fi
      MAKE_BASH="; /bin/bash"
    fi
  done

  # run make directly if already inside the container
  if [[ -e "/.dockerenv" ]] || [[ "${DOCKER_PROC}" != "" ]]; then
    make ${MAKE_ARGS}
    return 9
  fi

  return 0
}

function check_docker_image() {
  # check existing docker image
  DOCKER_IMAG_STATUS="not available"
  echo -e "\nChecking docker image [${DOCKER_IMAG}] for '${GITHUB_REPO}'"
  if [[ "${DOCKER_TAGS}" != "${DOCKER_IMAG}" ]]; then
    GITHUB_SSH="ssh -v git@github.com"
    GITHUB_SSH_KEYF="$(${GITHUB_SSH} 2>&1|grep 'Server accepts key:'|awk '{print $5}')"
    LOG_PAT="GitHub Personal Access Token in $$GITHUB_ACCESS_TOKEN"
    LOG_SSH="SSH key by ${GITHUB_SSH}"

    # Checking embeded SSH key or GitHub password/token
    # Caution: use docker staging for production. see https://bit.ly/2oY3pCn.
    if [[ -e "${GITHUB_SSH_KEYF}" ]]; then
      GITHUB_PRIVATE_KEY="$(cat ${GITHUB_SSH_KEYF} 2>/dev/null)"
    else
      # SSH key may be required for GitHub access to private repository
      log_trace "No ${LOG_SSH}"
    fi
    if [[ "${GITHUB_ACCESS_TOKEN}" == "" ]]; then
      # GITHUB_ACCESS_TOKEN is not set for GitHub personcal access token
      log_trace "No ${LOG_PAT}"
    fi
    if [[ "${GITHUB_ACCESS_TOKEN}${GITHUB_PRIVATE_KEY}" == "" ]]; then
      log_trace "Cannot find either ${LOG_PAT}, or ${LOG_SSH}" "WARNING"
    fi

    echo -e "\nBuilding docker image [${DOCKER_IMAG}] for '${GITHUB_REPO}'"
    if [[ "${GITHUB_ACCESS_TOKEN}" == "" ]]; then
      echo -e "\n- Embeded GitHub Personal Access Token (for private repository)"
    fi
    if [[ "${GITHUB_PRIVATE_KEY}" == "" ]]; then
      echo -e "\n- Embeded SSH key: ${GITHUB_SSH_KEYF}"
    fi
    echo "-----------------------------------------------------------------------"
    GITHUB_PRIVATE_KEY="${GITHUB_PRIVATE_KEY}" \
    GITHUB_ACCESS_TOKEN="${GITHUB_ACCESS_TOKEN}" \
    docker build -f "${DOCKER_FILE}" -t ${DOCKER_IMAG} \
      --build-arg GITHUB_ACCESS_TOKEN \
      --build-arg GITHUB_PRIVATE_KEY \
      .
    echo "-----------------------------------------------------------------------"
    DOCKER_IMAG_STATUS="ready"
  else
    DOCKER_IMAG_STATUS="up"
  fi
  if [[ "${MAKE_ARGS}" == "docker" ]]; then
    echo -e "\nDocker image [${DOCKER_IMAG}] is ${DOCKER_IMAG_STATUS}.\n"
    return 1
  fi

  return 0
}

# check_tools() func verifies prerequsite tools
function check_tools()
{
  local tool_set="${@:-aws}"
  for tool in ${tool_set}; do
    if ! [[ -x "$(which ${tool})" ]]; then
      log_error "Cannot find command '${tool}'"
    fi
  done
}

function docker_run() {
  # configure and start the container
  CMD_OPT="-it
    -e DEBUG=${DEBUG}
    -e PROJECT="${DOCKER_NAME}"
    -e PROJECT_DIR="${SOURCE_PATH}"
    -e ALL_PACKAGES
    -e AWS_ACCESS_KEY_ID
    -e AWS_DEFAULT_REGION
    -e AWS_SECRET_ACCESS_KEY
    -e BUILD_ENV
    -e BUILD_ARTIFACT
    -e BUILD_NUMBER
    -e BUILD_OS="${BUILD_OS:-linux}"
    -e BUILD_VERSION="${BUILD_VERSION:-1.0}"
    -e BUILDS_DIR
    -e LOCAL_USER_ID=${LOCAL_USER_ID:-$(id -u)}
    -e LOCAL_GROUP_ID=${LOCAL_GROUP_ID:-$(id -g)}
    -e USER
    -e MYSQL_HOST
    -e MYSQL_PORT
    -e MYSQL_DATABASE
    -e MYSQL_USERNAME
    -e S3_BUCKET
    -e S3_PREFIX
    -e SOCKEYE_API_KEY
    -e TEST_DIR
    -e TEST_BENCH
    -e TEST_COVER_MODE
    -e TEST_COVERAGES="${TEST_COVERAGES}"
    -e TEST_MATCH="${TEST_MATCH:-.}"
    -e TEST_TAGS="${TEST_TAGS:-all}"
    -e TEST_VERBOSE=${TEST_VERBOSE}
    -e VERBOSE=${VERBOSE}
    "
  CMD="docker run --rm ${CMD_OPT}
    --hostname ${DOCKER_HOST_NAME}
    --name ${DOCKER_HOST_NAME} --net="bridge"
    --expose ${DOCKER_PORT} -p 0.0.0.0:${DOCKER_PORT}:${DOCKER_PORT}
    -v "${PWD}":${SOURCE_PATH}
    ${DOCKER_IMAG}"

  if [[ "${DOCKER_EXEC}" == "exec" ]]; then
    CMD="docker exec ${CMD_OPT} ${DOCKER_HOST_NAME} /bin/bash"
  fi

  echo -e "\nRunning 'make ${MAKE_ARGS}' in docker container ${DOCKER_HOST_NAME}:${DOCKER_PORT}"
  if [[ "${DEBUG}" == "1" ]]; then
    echo "${CMD} bash -c \"make ${MAKE_ARGS} ${MAKE_BASH}\""
  fi
  echo "......................................................................."
  if [[ "${MAKE_ARGS}" != "cmd" ]]; then
    ${CMD} bash -c "make ${MAKE_ARGS} ${MAKE_BASH}"
  else
    ${CMD}
  fi
  echo "......................................................................."
}

# log_error() func: exits with non-zero code on error unless $2 specified
function log_error() {
  log_trace "$1" "ERROR" $2
}

# log_trace() func: print message at level of INFO, DEBUG, WARNING, or ERROR
function log_trace() {
  local err_text="${1:-Here}"
  local err_name="${2:-INFO}"
  local err_code="${3:-1}"

  if [[ "${err_name}" == "CAUTION" ]] || [[ "${err_name}" == "WARNING" ]]; then
    echo ""
    echo "**********************************************************************"
    echo -e "\n${err_name}: ${err_text}"
    echo "**********************************************************************"
    echo ""
  elif [[ "${err_name}" == "ERROR" ]] || [[ "${err_name}" == "FATAL" ]]; then
    HAS_ERROR="true"
    echo ''
    echo '                                                      \\\^|^///  '
    echo '                                                     \\  - -  // '
    echo '                                                      (  @ @  )  '
    echo '----------------------------------------------------oOOo-(_)-oOOo-----'
    echo -e "\n${err_name}: ${err_text}" >&2
    echo '                                                            Oooo '
    echo '-----------------------------------------------------oooO---(   )-----'
    echo '                                                     (   )   ) / '
    echo '                                                      \ (   (_/  '
    echo '                                                       \_)       '
    echo ''
    exit ${err_code}
  else
    echo -e "\n${err_name}: ${err_text}"
  fi
}


[[ $0 != "${BASH_SOURCE}" ]] || main "$@"
