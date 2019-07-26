#!/usr/bin/env bash
########################################################################
# Check go packages in all package management lists
#   - fixing "HEAD detached" issue in each package path
#   - exectuing 'go get' package if not yet installed
########################################################################
set -e +x
script_file="${BASH_SOURCE[0]##*/}"
script_base="$( cd "$( echo "${BASH_SOURCE[0]%/*}/.." )" && pwd )"
script_path="${script_base}/tools/${script_file}"

PROJECT_DIR="${PROJECT_DIR:-${script_base}}"
SYNC_VENDOR="${SYNC_VENDOR:-no}"

GLIDE="./glide.yaml"
GODEP="./Godeps/Godeps.json"
GOVEN="./vendor/vendor.json"
GOMOD="./go.mod"
GLOCK="./Gopkg.lock"
GOPKG="./Gopkg.toml"


# main entrance
function main() {

  check_depends

  cd -P "${PROJECT_DIR}" && echo "PWD: ${PWD:-$(pwd)}"

  PROJECT_PKG="${PROJECT_DIR//$GOPATH\/src\//}"
  log_trace "Current project package: ${PROJECT_PKG}"

  if [[ "$@" =~ (vendor) ]] || [[ "${SYNC_VENDOR}" =~ (1|enable|on|true|yes) ]]; then
    SYNC_VENDOR="yes"
  fi

# if [[ -f "${GODEP}" ]]; then check_godep; fi
  if [[ -f "${GOMOD}" ]]; then check_gomod; fi
# if [[ -f "${GLIDE}" ]]; then check_glide; fi
# if [[ -f "${GOVEN}" ]]; then check_govendor; fi
  if [[ -f "${GLOCK}" ]]; then check_gopkg; fi

}

# check_depends(): verifies preset environment variables exist
function check_depends() {
  local tool_set="awk git grep jq sort uniq"
  set +u
  echo "......................................................................."
  echo "Checking dependencies: ${tool_set}"
  for tool in ${tool_set}; do
    if ! [[ -x "$(which ${tool})" ]]; then
      log_error "Cannot find command '${tool}'"
    fi
  done

  if [[ ! -d "${GOPATH}" ]] || [[ ! -d "${GOPATH}/src" ]]; then
    local ERROR_TAG="$([[ -f "${GOMOD}" ]] && echo "WARNING" || echo "ERROR")"
    log_trace "Cannot find \$GOPATH or \$GOPATH/src" ${ERROR_TAG}
  fi

  local govendor_info="$(which govendor 1>/dev/null && govendor -version 2>&1|grep v)"
  local govendor_path="$(which govendor)"
  echo ""
  echo "Golang package managers"
  echo "-----------------------------------------------------------------------"
  echo -e "dep      : $(which dep)    \t- $(dep version|grep -e '^ version'|awk '{print $3}')"
  echo -e "glide    : $(which glide)  \t- $(glide -version 2>/dev/null)"
  echo -e "godep    : $(which godep)  \t- $(godep version 2>/dev/null)"
  echo -e "govendor : ${govendor_path}\t- ${govendor_info}"
  echo ""

  set -u
}

# check_glide
function check_glide() {
  log_trace "Checking go packages in ${GLIDE} ..."
  local packages="$(cat ${GLIDE}|grep -e '- package:'|awk '{print $3}')"
  for pkg in $packages; do check_package $pkg; done
}

# check_godep
function check_godep() {
  log_trace "Checking go packages in ${GODEP} ..."
  local packages="$(jq -r '.Deps[]|.ImportPath' ${GODEP})"
  for pkg in $packages; do check_package $pkg; done
}

# check_gomod
function check_gomod() {
  log_trace "Checking go packages in ${GOMOD} ..."
  local packages="$(cat ${GOMOD}|grep -e '^\t'|awk '{print $1}')"
  for pkg in $packages; do check_package $pkg; done
  # GO111MODULE=on go mod tidy
}

# check_gopkg (using packages list from go tool dep)
function check_gopkg() {
  log_trace "Checking go packages in ${GOPKG} and ${GLOCK} ..."
  local packages="$(cat ${GOPKG} ${GLOCK}|grep -e '^  name = '|awk '{print $3}'|sort|uniq)"
  for pkg in $packages; do check_package $pkg; done

  local DEP="$(which dep)"
  if [[ "${DEP}" == "${GOPATH}/bin/dep" ]] && [[ "${SYNC_VENDOR}" == "yes" ]]; then
    log_trace "Populating ./vendor from ${GLOCK} ..."
    dep ensure
    git checkout -- ${GOVEN} 2>&1 >/dev/null
  fi
}

# check_govendor
function check_govendor() {
  log_trace "Checking go packages in ${GOVEN} ..."
  local packages="$(jq -r '.package[]|.path' ${GOVEN})"
  for pkg in $packages; do check_package $pkg; done
}

# check go package in $GOPATH
function check_package() {
  # can also strip quotes by "$(eval echo $1)"
  local package="${1//\"/}"
  local pkg_dir="$GOPATH/src/$package"

  # skipping vendor and the project's own package
  if [[ "${package}" =~ (/vendor/) ]] || [[ "${package}" =~ (${PROJECT_PKG}) ]]; then
    log_trace "Skipping ${package}"
    return
  fi

  if [[ -d "${pkg_dir}" ]]; then
    log_trace "Checking go package in ${pkg_dir} ..."
    pushd "${pkg_dir}" > /dev/null
    # fixing "HEAD detached" issue if any
    local detached="$(git branch | grep 'HEAD detached')"
    if [[ "${detached}" != "" ]]; then
      git checkout master || true
    fi
    popd > /dev/null
  elif [[ "${package}" != "" ]]; then
    log_trace "Downloading and installing ${package} ..."
    GO111MODULE=on \
    go get -u ${package} || true
  fi
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


# main entrance, preventing from source
[[ $0 != "${BASH_SOURCE}" ]] || main $@
