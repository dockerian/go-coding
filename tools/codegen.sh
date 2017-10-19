#!/usr/bin/env bash
########################################################################
# Swagger codegen script
#
# Injectable environment variables:
#   CODEGEN_VER   - codegen-cli version for download
#   CODEGEN_IMAG  - codegen-cli docker image (:latest)
#
#   CODEGEN_LANG  - specify language for codegen cli
#   CODEGEN_TYPE  - specify either 'client' or 'server'
#   CODEGEN_CONF  - path to config file for codegen cli
#
#   SWAGGER_SPEC  - e.g. v1, or use CODEGEN_PATH, e.g. app/v1
#   SWAGGER_YAML  - path to swagger.yaml
#
#   DELETE_JAR    - boolean flag to delete codegen jar (default: yes)
#   USE_DOCKER    - boolean flag to enable/disable using docker
#
# Note: codegen-cli docker image has no version tags, e.g. 2.2.3,
#       so that it is generating different code than java jar cli
#
########################################################################
script_file="${BASH_SOURCE[0]##*/}"
script_base="$( cd "$( echo "${BASH_SOURCE[0]%/*}/.." )" && pwd )"
script_path="${script_base}/tools/${script_file}"

CODEGEN_BIN=swagger-codegen
CODEGEN_CLI=swagger-codegen-cli
CODEGEN_MVN=http://central.maven.org/maven2/io/swagger/swagger-codegen-cli
CODEGEN_VER="${CODEGEN_VER:-2.2.3}"
CODEGEN_URL="${CODEGEN_MVN}/${CODEGEN_VER}/${CODEGEN_CLI}-${CODEGEN_VER}.jar"
CODEGEN_JAR="${CODEGEN_CLI}.jar"

CODEGEN_IMAG="swaggerapi/swagger-codegen-cli:latest"
CODEGEN_LANG="${CODEGEN_LANG:-go}"
CODEGEN_TYPE="${CODEGEN_TYPE:-client}"
SWAGGER_SPEC="${SWAGGER_SPEC:-v1}"
CODEGEN_PATH="${CODEGEN_PATH:-app/${SWAGGER_SPEC}}"
CODEGEN_CONF="${CODEGEN_CONF:-${CODEGEN_PATH}/swagger.config}"
SWAGGER_YAML="${SWAGGER_YAML:-${CODEGEN_PATH}/swagger.yaml}"
PACKAGE_NAME="${PACKAGE_NAME:-${CODEGEN_TYPE}}"

# set flag to delete downloaded codegen jar file
DELETE_JAR="${DELETE_JAR:-yes}"
# set flag to enable/disable using docker container
USE_DOCKER="${USE_DOCKER:-true}"


# main entrance
function main() {
  shopt -s nocasematch
  if [[ "$@" =~ (help|-h|/h|-\?|/\?) ]]; then
    usage; return
  fi
  if [[ "$@" =~ (keep-jar) ]]; then
    DELETE_JAR="no"
  fi
  if [[ "${DELETE_JAR}" =~ (1|enabled|on|true|yes) ]]; then
    DELETE_JAR="yes"
  fi
  if [[ "${USE_DOCKER}" =~ (0|disable|off|false|no) ]]; then
    echo "Using docker container is disabled."
    USE_DOCKER=0
  fi
  cd -P "${script_base}"

  if [[ -x "$(which docker)" ]] && [[ "${USE_DOCKER}" != "0" ]]; then
    codegen
  elif [[ -x "$(which ${CODEGEN_BIN})" ]]; then
    codegen_cli
  else
    check_depends
    codegen_cli
  fi
  echo "-----------------------------------------------------------------------"
  echo "- DONE: ${script_path}"
}

# check_depends(): verifies preset environment variables exist
function check_depends() {
  local tool_set="java wget rm"
  set +u
  echo "......................................................................."
  echo "Checking dependencies: ${tool_set}"
  for tool in ${tool_set}; do
    if ! [[ -x "$(which ${tool})" ]]; then
      log_error "Cannot find command '${tool}'"
    fi
  done
  set -u
}

# check_return_code(): checks exit code from last command
function check_return_code() {
  local return_code="${1:-0}"
  local action_name="${2:-AWS CLI}"

  if [[ "${return_code}" != "0" ]]; then
    log_fatal "${action_name} [code: ${return_code}]" ${return_code}
  else
    echo ""
    echo "SUCCESS: ${action_name}"
  fi
}

# codegen(): run swagger-codegen-cli in docker container
function codegen() {
  local opt="-c /local/${CODEGEN_CONF}"
  local out="${CODEGEN_PATH}/${CODEGEN_TYPE}"

  local tmp="${script_base}/${CODEGEN_CONF}"
  echo "Configuring package name to '${CODEGEN_TYPE}'"
  echo '{ "packageName": "'${CODEGEN_TYPE}'" }' > "${script_base}/${CODEGEN_CONF}"

  local cmd="docker run --rm
    -v ${script_base}:/local
    ${CODEGEN_IMAG} generate
    -i /local/${SWAGGER_YAML} -l ${CODEGEN_LANG}
    -o /local/${out} --api-package ${CODEGEN_TYPE}
    ${opt}"
  echo "......................................................................."
  echo "Generating '${CODEGEN_LANG}' from '${SWAGGER_YAML}' by ${CODEGEN_IMAG} "
  echo "cli: ${cmd}"
  $cmd

  check_return_code $? "docker run ${CODEGEN_IMAG}"
  rm -rf "${tmp}"
  echo ""
}

# codegen(): call swagger-codegen-cli
function codegen_cli() {
  local opt="-c ${CODEGEN_CONF}"
  local out="${CODEGEN_PATH}/${CODEGEN_TYPE}"
  local cli="$(which ${CODEGEN_BIN})"

  echo ""
  cd -P "${script_base}"
  if [[ ! -x "${cli}" ]]; then
    if [[ ! -e "${CODEGEN_JAR}" ]]; then
      echo "Downloading ${CODEGEN_JAR}"
      echo "       from ${CODEGEN_URL}"
      echo ""
      wget ${CODEGEN_URL} -O "${CODEGEN_JAR}"

      if [[ ! -e "${CODEGEN_JAR}" ]]; then
        log_error "Unable to download '${CODEGEN_JAR}': ${CODEGEN_URL}"
      fi
    fi
    cli="java -jar ${CODEGEN_JAR}"
  fi
  echo "Using ${cli}"

  local tmp="${CODEGEN_CONF}"
  echo "Configuring package name to '${CODEGEN_TYPE}'"
  echo '{ "packageName": "'${CODEGEN_TYPE}'" }' > "${CODEGEN_CONF}"

  local cmd="${cli} generate
    -i ${SWAGGER_YAML} -l ${CODEGEN_LANG}
    -o ${out} --api-package ${CODEGEN_TYPE}
    ${opt}"

  echo "......................................................................."
  echo "Generating '${CODEGEN_LANG}' from '${SWAGGER_YAML}' by ${cli}"
  echo "cli: ${cmd}"
  $cmd

  check_return_code $? "${CODEGEN_CLI}"
  if [[ "${DELETE_JAR}" == "yes" ]]; then
    rm -rf "${CODEGEN_JAR}"
  fi
  rm -rf "${tmp}"
  echo ""
}

# log_error() func: exits with non-zero code on error unless $2 specified
function log_error() {
  log_trace "$1" "ERROR" $2
}

# log_fatal() func: exits with non-zero code on fatal failure unless $2 specified
function log_fatal() {
  log_trace "$1" "FATAL" $2
}

# log_trace() func: print message at level of INFO, DEBUG, WARNING, or ERROR
function log_trace() {
  local err_text="${1:-Here}"
  local err_name="${2:-INFO}"
  local err_code="${3:-1}"

  if [[ "${err_name}" == "ERROR" ]] || [[ "${err_name}" == "FATAL" ]]; then
    HAS_ERROR="true"
    echo -e "\n${err_name}: ${err_text}" >&2
    exit ${err_code}
  else
    echo -e "\n${err_name}: ${err_text}"
  fi
}

# usage() func: show help
function usage() {
  local headers="0"
  echo ""
  echo "USAGE: ${script_file} --help"
  echo ""
  # echo "$(cat ${script_path} | grep -e '^#   \$[1-9] - ')"
  while IFS='' read -r line || [[ -n "${line}" ]]; do
    if [[ "${headers}" == "0" ]] && [[ "${line}" =~ (^#[#=-\\*]{59}) ]]; then
      headers="1"
      echo "${line}"
    elif [[ "${headers}" == "1" ]] && [[ "${line}" =~ (^#[#=-\\*]{59}) ]]; then
      headers="0"
      echo "${line}"
    elif [[ "${headers}" == "1" ]]; then
      echo "${line}"
    fi
  done < "${script_path}"
}



[[ $0 != "${BASH_SOURCE}" ]] || main "$@"
