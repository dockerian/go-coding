#!/usr/bin/env bash
########################################################################
# Check test coverage
#
# Arguments:
#       $1 - test coverage thresholds
#       $2 - test coverage percentage (without %), or a log, or "test"
#
# Options (after arguments)
#       --bypass  : always bypass the check/thresholds, return zero
#       --test    : only to run all tests, same as "$2" == "test"
#
# Injectable environment variables
#     TEST_ARGS            go test args
#     TEST_LOGS            go test output
#     COVER_MODE           go test cover mode
#     NO_THRESHOLDS        bypass thresholds check, default: false
#     RUN_ALL_TESTS        run all tests before check, default: true
#     ALL_PACKAGES         specified packages for test coverage
#     COVERAGE_PERCENTAGE  reported cover percent if missing $2
#     COVERAGE_THRESHOLDS  code coverage thresholds, default 90
#     COVER_ALL_OUT        coverage profile output, default: coverage.txt
#     PROJECT_DIR          project root path
#
# Exit code:
#       1  - fail
#       0  - pass
#
########################################################################
set -e +x
script_file="${BASH_SOURCE[0]##*/}"
script_base="$( cd "$( echo "${BASH_SOURCE[0]%/*}/.." )" && pwd )"
script_path="${script_base}/tools/${script_file}"

PROJECT_DIR="${PROJECT_DIR:-${script_base}}"

NO_THRESHOLDS="${NO_THRESHOLDS:-0}"
RUN_ALL_TESTS="${RUN_ALL_TESTS:-0}"

COVERAGE_PERCENTAGE="${2:-${COVERAGE_PERCENTAGE:-0}}"
COVERAGE_THRESHOLDS="${1:-${COVERAGE_THRESHOLDS:-90}}"
COVER_ALL_OUT="${COVER_ALL_OUT:-${PROJECT_DIR}/cover.out}"
COVER_FUNC="${COVER_FUNC:-${PROJECT_DIR}/cover-func.out}"
COVER_MODE="${COVER_MODE:-atomic}"

TEST_ARGS="${TEST_ARGS}"
TEST_LOGS="${TEST_LOGS:-${PROJECT_DIR}/tests.log}"


# main entrance
function main() {
  shopt -s nocasematch
  if [[ "$@" =~ (help|-h|/h|-\?|/\?) ]]; then
    usage; return
  fi

  cd -P "${PROJECT_DIR}" && echo "PWD: ${PWD:-$(pwd)}"

  PACKAGES="$(go list ./... 2>/dev/null|grep -v -E '/v[0-9]+/client|/v[0-9]+/server|/vendor/'||true)"
  ALL_PACKAGES="${ALL_PACKAGES:-${PACKAGES}}"

  if [[ "${RUN_ALL_TESTS}" =~ (1|true|yes) ]]; then RUN_ALL_TESTS="true"; fi
  if [[ "${NO_THRESHOLDS}" =~ (1|true|yes) ]]; then NO_THRESHOLDS="true"; fi

  if [[ "${RUN_ALL_TESTS}" == "true" ]]; then
    if [[ "${TEST_ARGS}" == "" ]]; then
      log_error "Missing go test arguments"
    fi
    run_all_tests
  else
    check "${COVERAGE_THRESHOLDS}" "${COVERAGE_PERCENTAGE}"
  fi
}

# run tests for all go packages
function run_all_tests() {
  local pkg_out="cover-pkg.out"

  echo "mode: ${COVER_MODE}"
  echo "mode: ${COVER_MODE}" > "${COVER_ALL_OUT}"

  log_trace "Generating coverage profile, mode = ${COVER_MODE}"
  if [[ -s "${COVER_ALL_OUT}" ]]; then
    touch "${pkg_out}"
    for pkg in ${ALL_PACKAGES}; do
      echo "Run: go test ${TEST_ARGS} --coverprofile=${pkg_out} ${pkg}"
      go test ${TEST_ARGS} --coverprofile="${pkg_out}" ${pkg} 2>&1|tee -a "${TEST_LOGS}"
      check_return_code $?
      tail -n +2 "${pkg_out}" >> "${COVER_ALL_OUT}"
      echo ""
    done
    go tool cover -func="${COVER_ALL_OUT}" | tee "${COVER_FUNC}"
  fi
}

# check code coverage thresholds
function check() {
  # detecting if $1 is a cov file (http://tldp.org/LDP/abs/html/fto.html)
  if [[ -s "$2" ]] && [[ "$2" == *.out ]]; then
    local COV="${COVERAGE_PERCENTAGE}"
    local LOF=$(grep -e '(statements)' "$2"|awk -F '[%]' '{print $1}')
    # log_trace "Parsing testing coverage: ${COV}"
    COVERAGE_PERCENTAGE=$(echo $LOF|awk '{print $3}')
  elif [[ -e "$2" ]]; then
    log_error "Invalid cover \*.out: $2"
  fi

  # checking coverage number validation
  if [[ "${COVERAGE_PERCENTAGE}" == "" ]]; then
    # resetting to zero in case there is no tests in log
    COVERAGE_PERCENTAGE="0"
  elif [[ ! "${COVERAGE_PERCENTAGE}" =~ ([0-9]+.?[0-9]*) ]]; then
    log_error "Invalid float number: '${COVERAGE_PERCENTAGE}' for coverage percentage"
  fi
  if [[ ! "${COVERAGE_THRESHOLDS}" =~ ([0-9]+.?[0-9]*) ]]; then
    log_error "Invalid float number: '${COVERAGE_THRESHOLDS}' for coverage thresholds"
  fi

  # setting 1 if percentage >= thresholds; otherwise 0
  PASS=$(awk "BEGIN {print (${COVERAGE_PERCENTAGE} >= ${COVERAGE_THRESHOLDS})}")

  if [[ "${PASS}" != "1" ]]; then
    log_error "Coverage: ${COVERAGE_PERCENTAGE} % is below ${COVERAGE_THRESHOLDS} %"
  fi
  echo ""
  echo "- PASS: ${COVERAGE_PERCENTAGE} % >= ${COVERAGE_THRESHOLDS} % (thresholds)"
}

# check_return_code(): checks exit code from last command
function check_return_code() {
  local return_code="${1:-0}"
  local action_name="${2:-go test}"

  if [[ "${return_code}" != "0" ]]; then
    log_fatal "${action_name} [code: ${return_code}]" ${return_code}
  else
    echo "Success: ${action_name}"
    echo ""
  fi
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
    echo -e "\n${err_name}: ${err_text}\n" >&2
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


ARGS=""
# step 1: pre-processing optional arguments
if [[ "$2" == "test" ]]; then RUN_ALL_TESTS="true"; fi
for arg in $@; do
  if [[ "${arg}" == "--pass" ]]; then
    NO_THRESHOLDS="true"
  elif [[ "${arg}" == "--test" ]]; then
    RUN_ALL_TESTS="true"
  else
    ARGS="${ARGS} "${arg}""
  fi
done

# main entrance, preventing from source
[[ $0 != "${BASH_SOURCE}" ]] || main ${ARGS}
