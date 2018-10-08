#!/usr/bin/env bash
##############################################################################
# git clone from a repository by username/password, or ssh
# Author: jzhu@infoblox.com
#
# Command line arguments:
#   $1 : git repository url
#        default = https://github.com/Infoblox-CTO/cyberint-sng-api.git
#   $2 : check out location
#
# Expecting key in ~/.ssh or the following variables defined:
#   GITHUB_USERNAME
#   GITHUB_PASSWORD
#
##############################################################################
script_file="${BASH_SOURCE[0]##*/}"
# assuming this script is under ${script_base}/tools
script_base="$( cd "$( echo "${BASH_SOURCE[0]%/*}/.." )" && pwd )"
script_path="${script_base}/tools/${script_file}"

# git@github.com:Infoblox-CTO/cyberint-sng-api.git
GITHUB_REPO="${1:-https://github.com/Infoblox-CTO/cyberint-sng-api.git}"
GIT_CLONE_TARGET="$2"
GITHUB_REPO_OKAY="false"
GITHUB_REPO_TYPE=""
GITHUB_REPO_ORGZ=""
GITHUB_REPO_NAME=""

# main entry
function main() {
  REVISION="${1:-None}"

  shopt -s nocasematch
  for arg in $@ ; do
    if [[ "${arg}" =~ (help|/h|-\?|\/\?) ]] || [[ "${arg}" == "-h" ]]; then
      usage; return
    fi
  done
  if [[ "$@" =~ (--help|/help|-\?|/\?) ]]; then
    usage; return
  fi

  echo ""
  cd -P "${script_base}" && echo "PWD: ${PWD}"
  do_check_url "${GITHUB_REPO}" "${GIT_CLONE_TARGET}"
  do_clone

  echo ""
  echo "DONE."
}

# check validation on github repo url
function do_check_url() {
  local repo="$1"
  local dest="$2"

  log_trace "Checking repo url: $1"

  if [[ "$1" =~ ^(https://|git@)github.com[:/](.*)/(.*)\.git$ ]]; then
    GITHUB_REPO_TYPE="${BASH_REMATCH[1]}"
    GITHUB_REPO_ORGZ="${BASH_REMATCH[2]}"
    GITHUB_REPO_NAME="${BASH_REMATCH[3]}"
    GITHUB_REPO_PART="${GITHUB_REPO_ORGZ}/${GITHUB_REPO_NAME}.git"

    if [[ "${GITHUB_REPO_TYPE}" == "" ]]; then return; fi
    if [[ "${GITHUB_REPO_ORGZ}" == "" ]]; then return; fi
    if [[ "${GITHUB_REPO_NAME}" == "" ]]; then return; fi

    GITHUB_REPO_HTTP="https://github.com/${GITHUB_REPO_PART}"
    GITHUB_REPO_GSSH="git@github.com:${GITHUB_REPO_PART}"
    GIT_CLONE_TARGET="${2:-${GITHUB_REPO_NAME}}"
    GITHUB_REPO_OKAY="true"
  fi

  if [[ "${GITHUB_REPO_OKAY}" != "true" ]]; then
    log_error "Invalid repository url: ${GITHUB_REPO}"
  fi
}

# check out git commit sha or revision tag
function do_clone() {
  echo ""
  echo "......................................................................."
  echo "git clone: ${GITHUB_REPO_HTTP} => ${GIT_CLONE_TARGET}"
  echo ""
  git clone "${GITHUB_REPO_HTTP}" "${GIT_CLONE_TARGET}"

  if [[ "$?" == "0" ]]; then return; fi

  if [[ "${GITHUB_USERNAME}" != "" ]] && [[ "${GITHUB_PASSWORD}" != "" ]]; then
    GITHUB_REPO_HTTP="https://${GITHUB_USERNAME}:${GITHUB_PASSWORD}@github.com/${GITHUB_REPO_PART}"
  fi
  echo ""
  echo "......................................................................."
  echo "git clone [with user/pass]: ${GITHUB_REPO_HTTP} => ${GIT_CLONE_TARGET}"
  echo ""
  git clone "${GITHUB_REPO_HTTP}" "${GIT_CLONE_TARGET}"

  if [[ "$?" == "0" ]]; then return; fi

  echo ""
  echo "......................................................................."
  echo "git clone [ssh]: ${GITHUB_REPO_GSSH} => ${GIT_CLONE_TARGET}"
  echo ""
  git clone "${GITHUB_REPO_GSSH}" "${GIT_CLONE_TARGET}"

  if [[ "$?" == "0" ]]; then return; fi

  echo ""
  echo "======================================================================="
  log_error "Cannot clone: ${GITHUB_REPO}"
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
