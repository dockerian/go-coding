#!/usr/bin/env bash
########################################################################
# Check go test log
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


# main entrance
function main() {
  TEST_LOGS="${TEST_LOGS:-tests.log}"
  DEBUG="${DEBUG:-${VERBOSE:-${TEST_VERBOSE}}}"
  COUNT_FAIL=0

  check_depends

  cd -P "${PROJECT_DIR}" && echo "PWD: ${PWD:-$(pwd)}"

  if [[ ! -f "${TEST_LOGS}" ]]; then
    echo "Cannot find ${TEST_LOGS}"
    return
  fi

  if [[ "${DEBUG}" =~ (0|disable|off|false|no) ]]; then
    DEBUG="false"
  fi
  # NOTE: use 'grep --text' (or 'grep -a', processing as text file)
  #       to avoid from 'grep' error: Binary file (standard input) matches
  #  see: man page at http://ss64.com/bash/grep.html
  #
  COUNT_COMP="$(cat "${TEST_LOGS}"|grep -a "\: cannot find package"|wc -l|xargs echo)"
  COUNT_FAIL="$(cat "${TEST_LOGS}"|grep -a "\--- FAIL:\|^FAIL"|wc -l|xargs echo)"
  COUNT_PASS="$(cat "${TEST_LOGS}"|grep -a "\--- PASS:"|wc -l|xargs echo)"
  COUNT_SKIP="$(cat "${TEST_LOGS}"|grep -a "\--- SKIP:"|wc -l|xargs echo)"

  if [[ "${COUNT_COMP}" != "0" ]]; then
    printf "\n*** Build errors  : %2d ***\n" ${COUNT_COMP}
    echo "cannot find:"
    (cat "${TEST_LOGS}"|grep -a ": cannot find package"|awk '{print $5}'|sort|uniq)
    exit ${COUNT_COMP}
  fi
  if [[ "${COUNT_FAIL}${COUNT_PASS}${COUNT_SKIP}" != "000" ]]; then
    echo ""
    echo "============================= TEST SUMMARY ============================"

    if [[ "${COUNT_PASS}" != "" ]] || [[ "${DEBUG}" != "false" ]]; then
      printf "\n*** Passed tests  : %2d ***\n" ${COUNT_PASS}
      (cat "${TEST_LOGS}" | grep -e "\--- PASS:" | cut -d':' -f2 | sort)
    fi
    if [[ "${COUNT_SKIP}" != "" ]] || [[ "${DEBUG}" != "false" ]]; then
      printf "\n*** Skipped tests : %2d ***\n" ${COUNT_SKIP}
      (cat "${TEST_LOGS}" | grep -e "\--- SKIP:" | cut -d':' -f2 | sort)
    fi
    if [[ "${COUNT_FAIL}" != "" ]] || [[ "${DEBUG}" != "false" ]]; then
      printf "\n*** Failed tests  : %2d ***\n" ${COUNT_FAIL}
      (cat "${TEST_LOGS}" | grep -e "\--- FAIL:\|^FAIL" | cut -d':' -f2 | sort)
    fi
    echo ""
    echo "======================================================================="
  elif [[ "${DEBUG}" == "" ]]; then
    echo "No failed test (TEST_VERBOSE is unset)"
  fi

  echo ""
  # The exit code is 0 if there are no test failures.
  echo "exit code: ${COUNT_FAIL}  (see ${TEST_LOGS})"
  exit ${COUNT_FAIL}
}

# check_depends(): verifies preset environment variables exist
function check_depends() {
  local tool_set="cat cut grep sort uniq wc xargs"
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

# log_error() func: exits with non-zero code on error unless $2 specified
function log_error() {
  log_trace "$1" "ERROR" $2
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


# main entrance, preventing from source
[[ $0 != "${BASH_SOURCE}" ]] || main ${ARGS}
