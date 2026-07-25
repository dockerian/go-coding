#!/usr/bin/env bash
####################
# termux : ~/.bashrc
####################
shopt -s promptvars

PS1='\n\[\033[0;36m\]\h\[\033[0m\]:\[\033[0;35m\]\u\[\033[0m\] \W [\#]:\n\$ '
PS1='$(printf "%$((COLUMNS-1))s\r")'$PS1

export BASE="$HOME/storage/shared"
export HISTCONTROL=ignoredups

echo ""
echo "Loading bash aliases ..."
alias a="alias|cut -d' ' -f 2- "
alias bashrc='source ~/.bashrc; title ${PWD##*/};'
alias cdbase="cd $HOME/storage/shared/"
alias cddcim="cd $HOME/storage/shared/DCIM"
alias cdpics="cd $HOME/storage/shared/DCIM/Camera"
alias dcim="cd $HOME/storage/shared/DCIM/Camera; mvp"
alias csd="cd $HOME/storage/shared/_SD_"
alias cdp='cd -P .'
alias cls='clear && printf "\e[3J"'
alias dater='date +"%Y-%m-%d %H:%M:%S" -r'
alias dated='date +"%Y-%m-%d %H:%M:%S" -d'
alias dh='du -hs'
alias dir='ls -al '
alias envi='env | grep -i '
alias envs='env | sort'
alias hs='history | grep'
alias ip='echo $(ipconfig getifaddr en0) $(dig +short myip.opendns.com @resolver1.opendns.com)'
alias ll='ls -al'
alias lll='ls -al -T | sort -f -k9,9'  # --time-style=full-iso
alias lln='ls -al | sort -f -k9,9'
alias llo='ls -l --time-style=long-iso'
alias llt='ls -al -rt'
alias path='echo $PATH|tr ":" "\n"'
alias setp='(set -o posix; set|grep -v _xspec)'
alias tss='pkg install root-repo && termux-setup-storage'
alias ver='echo -e "$(uname -a)"; echo ""; echo -e "$(bash --version)"'



############################################################
# function: Main
############################################################
function main() {
  # find . -name '.DS_Store' -type f -delete 2>/dev/null &
  # ipy_autoreload
  echo fixpath
}

############################################################
# function: Output a relative path to absolute path
############################################################
function abspath() {
  set +u
  local thePath
  if [[ ! "$1" =~ ^/ ]]; then thePath="$PWD/$1"; else thePath="$1"; fi
  echo "$thePath"|(
  IFS=/
  read -a parr
  declare -a outp
  for i in "${parr[@]}";do
    case "$i" in
    ''|.) continue ;;
    ..)
      len=${#outp[@]}
      if ((len!=0));then unset outp[$((len-1))]; else continue; fi
      ;;
    *)
      len=${#outp[@]}
      outp[$len]="$i"
      ;;
    esac
  done
  echo /"${outp[*]}"
  )
  set -u
}

############################################################
# function: Print file info
############################################################
function fileinfo() {
  for file in "$@"; do
    path=$(abspath "$file")

    # strip longest match of */ from start
    name="${file##*/}"

    # substring from 0 thru pos of filename
    dir_="${file:0:${#file} - ${#name}}"

    # strip shortest match of . plus at least one non-dot char from end
    base="${name%.[^.]*}"

    # substring from len of base thru end
    ext_="${name:${#base} + 1}"

    size=$((
      du --apparent-size --block-size=1 "$file" 2>/dev/null ||
      gdu --apparent-size --block-size=1 "$file" 2>/dev/null ||
      find "$file" -printf "%s" 2>/dev/null ||
      gfind "$file" -printf "%s" 2>/dev/null ||
      stat --printf="%s" "$file" 2>/dev/null ||
      stat -f%z "$file" 2>/dev/null ||
      wc -c <"$file" 2>/dev/null
    ) | awk '{print $1}')

    # in case of an extension without base, it's really the base
    if [[ -z "$base" && -n "$ext_" ]]; then
      base=".$ext_"
      ext=""
    fi
    if [[ "${dir_}" == "" ]]; then
      dir_=${path%/*}
    fi
    echo -e "------------------------------------------------------------"
    echo -e "\t file : $file"
    echo -e "\t path : $path"
    echo -e "\t  dir : $dir_"
    echo -e "\t base : $base"
    echo -e "\t  ext : $ext_"
    echo -e "\t size : $size"
  done
}

############################################################
# function: Remove duplicates in $PATH variable
############################################################
function fixpath() {
  # remove duplicates in $PATH
  export PATH=$(perl -e 'print join ":", grep {!$h{$_}++} split ":", $ENV{PATH}')
  export PATH=$(printf %s "$PATH" | awk -v RS=: -v ORS=: '!arr[$0]++')
  export PATH=$(printf %s "$PATH" | awk -v RS=: -v ORS=: '{ if (!arr[$0]++) { print $0 } }')
  export PATH=$(printf %s "$PATH" | awk -v RS=: '{ if (!arr[$0]++) {printf("%s%s",!ln++?"":":",$0)}}')
  echo "........................................................................"
  echo $PATH|tr ":" "\n"
  echo "........................................................................"
}

############################################################
# function: Rename camera photo IMG* files
############################################################
function mvp () {
  mvp_oppo
  mvp_vivo
  echo ""
  echo "DONE. `date +"%Y-%m-%d %H:%M:%S"`"
  echo ""
}
#----------------------------------------
# function internal to fix duplicate name
function _mvps_ () {
  if [[ $1 == "" || $2 == "" || $3 == "" ]]; then
    return;
  fi
  local _file_="$1"
  local _part_="$2"
  local _secs_="$3"
  local _date_="$4"
  local _orgf_="${_file_%.*}"
  local _type_="${_file_##*.}"
  local _name_="${_part_}_${_secs_}0"
  local _newf_="${_name_}"
  local _raws_="dng nef raw CR2 CR3 DNG NEF NRW ARW RAF ORF RAW PEF PTX"
  local _secf_=""
  local _cmdl_=""

  if [[ "${_type_,,}" == "mp4" ]]; then
    _name_="${_part_}_0${_secs_}"
    _newf_="${_name_}"
  fi
  # adding to seconds for duplicate name
  for n in {1..9}; do
    if [[ ! -e "${_newf_}.${_type_}" ]]; then break; fi
    if [[ "${_type_}" == "jpg" ]]; then
      _secf_=$((10 * ${_secs_##+(0)} + ${n##+(0)}))
    else
      _secf_=$((${_secs_##+(0)} + ${n##+(0)}))
    fi
    _newf_=`printf "%s_%03d" ${_part_} ${_secf_}`
  done
  # appending dash for duplicate name
  for n in {1..5}; do
    if [[ ! -e "${_newf_}.${_type_}" ]]; then break; fi
    _newf_="${_name_}$(for i in $(seq 1 $n); do printf "~"; done)"
  done

  _cmdl_="mv $1 ${_newf_}.${_type_}"
  echo "- [${_date_}] ${_cmdl_}"
  ${_cmdl_}

  for x in ${_raws_}; do
    if [[ -e "${_orgf_}.$x" ]]; then
      _cmdl_="mv ${_orgf_}.$x ${_newf_}.$x"
      echo "- [${_date_}] ${_cmdl_}"
      ${_cmdl_}
    fi
  done
}
####################################
# function: Rename OPPO camera files
####################################
function mvp_oppo () {
  local _pattern_='IMG??????????????'
  local _pat_mp4_='VID??????????????.mp4'
  local _date_=''
  local _part_=''
  local _secs_=''
  local _orgf_=''
  local _done_='NONE'

  for x in jpg JPG dng DNG RAW; do
    if ls ${_pattern_}.$x 1>/dev/null 2>&1; then
      echo ""
      echo "Processing OPPO ${_pattern_} ..."
      for f in ${_pattern_}.$x; do
        _date_="${f:3:4}-${f:7:2}-${f:9:2} ${f:11:2}:${f:13:2}"
        _part_="${f:3:8}_${f:11:4}" # in format yyyymmdd_HHMM
        _secs_="${f:15:2}" # 2-digit seconds at offset 15
        # do renaming IMG file
        _mvps_ "${f}" "${_part_}" "${_secs_}" "${_date_}"
        _done_="DONE"

        # checking sequenced files
        _orgf_="${f%.*}"
        for n in $(printf "%02d\n" {1..20}); do
          if [[ -e ${_orgf_}_$n.jpg ]]; then
            _mvps_ "${_orgf_}_$n.jpg" "${_part_}" "${_secs_}" "${_date_}"
          fi
        done
      done
    fi
  done

  if ls ${_pat_mp4_} 1>/dev/null 2>&1; then
    echo ""
    echo "Processing OPPO ${_pattern_} ..."
    for f in ${_pat_mp4_}; do
      _date_="${f:3:4}-${f:7:2}-${f:9:2} ${f:11:2}:${f:13:2}"
      _part_="${f:3:8}_${f:11:4}" # in format yyyymmdd_HHMM
      _secs_="${f:15:2}" # 2-digit seconds at offset 15
      _mvps_ "${f}" "${_part_}" "${_secs_}" "${_date_}"
      _done_="DONE"
    done
  fi

  echo ""
  echo "${_done_} (OPPO) `date +"%Y-%m-%d %H:%M:%S"`"
}
####################################
# function: Rename vivo camera files
####################################
function mvp_vivo () {
  local _raws_="dng DNG RAW"
  local _pattern_='IMG_????????_??????'
  local _pat_vid_='VID_?????????????????.mp4' # VID_yyyymmddHHMMsssnnnnnn.mp4
  local _pat_mp4_='video_????????_??????.mp4'
  local _date_=''
  local _part_=''
  local _secs_=''
  local _orgf_=''
  local _done_='NONE'

  for x in jpg JPG dng DNG RAW; do
    if ls ${_pattern_}.$x 1>/dev/null 2>&1; then
      echo ""
      echo "Processing vivo ${_pattern_} ..."
      for f in ${_pattern_}.$x; do
        _date_="${f:4:4}-${f:8:2}-${f:10:2} ${f:13:2}:${f:15:2}"
        _part_="${f:4:8}_${f:13:4}" # in format yyyymmdd_HHMM
        _secs_="${f:17:2}" # 2-digit seconds at offset 17
        # do renaming IMG file
        _mvps_ "${f}" "${_part_}" "${_secs_}" "${_date_}"
        _done_="DONE"

        # checking sequenced files
        _orgf_="${f%.*}"
        for n in $(printf "%02d\n" {1..20}); do
          if [[ -e ${_orgf_}~$n.jpg ]]; then
            _mvps_ "${_orgf_}~$n.jpg" "${_part_}" "${_secs_}" "${_date_}"
          fi
        done
        for n in {1..20}; do
          if [[ -e ${_orgf_}_$n.jpg ]]; then
            _mvps_ "${_orgf_}_$n.jpg" "${_part_}" "${_secs_}" "${_date_}"
          fi
        done
      done
    fi
  done

  if ls ${_pat_vid_} 1>/dev/null 2>&1; then
    echo ""
    echo "Processing vivo ${_pat_vid_} ..."
    for f in ${_pat_vid_}; do
      _date_="${f:4:4}-${f:8:2}-${f:10:2} ${f:12:2}:${f:14:2}"
      _part_="${f:4:8}_${f:12:4}" # in format yyyymmdd_HHMM
      _secs_="${f:16:2}" # 2-digit seconds at offset 19
      _mvps_ "${f}" "${_part_}" "${_secs_}" "${_date_}"
      _done_="DONE"
    done
  fi

  if ls ${_pat_mp4_} 1>/dev/null 2>&1; then
    echo ""
    echo "Processing vivo ${_pat_mp4_} ..."
    for f in ${_pat_mp4_}; do
      _date_="${f:6:4}-${f:10:2}-${f:12:2} ${f:15:2}:${f:17:2}"
      _part_="${f:6:8}_${f:15:4}" # in format yyyymmdd_HHMM
      _secs_="${f:19:2}" # 2-digit seconds at offset 19
      _mvps_ "${f}" "${_part_}" "${_secs_}" "${_date_}"
      _done_="DONE"
    done
  fi

  echo ""
  echo "${_done_} (vivo) `date +"%Y-%m-%d %H:%M:%S"`"
}

############################################################
# function: Add title to current terminal
############################################################
function title() {
  if [ $# -eq 0 ]; then
    # export PROMPT_COMMAND='echo -ne "\033]0;${PWD/#$HOME/~}\007"'
    export PROMPT_COMMAND='echo -ne "\033]0;${PWD##*/~}\007"'
  else
    TITLE=$*;
    export PROMPT_COMMAND='echo -ne "\033]0;$TITLE\007"'
    # echo -ne "\033]0;"$*"\007"
  fi
}

############################################################
# function: Executes command with a timeout
# Params:
#   $1 timeout in seconds
#   $2 command
# Returns 1 if timed out 0 otherwise
############################################################
function timeout() {
    time=$1
    # start the command in a subshell to avoid problem with pipes
    # (spawn accepts one command)
    command="/bin/sh -c \"$2\""
    expect -c "set echo \"-noecho\"; set timeout $time; spawn -noecho $command; expect timeout { exit 1 } eof { exit 0 }"
    if [ $? = 1 ] ; then
        echo "Timeout after ${time} seconds"
    fi
}

############################################################
# function: Use `touch -d` to apply all sub-dirs recursively
# Params: $1 a source dir path
#         $2 optional depth
#         $3 options
############################################################
function touchdbyfile() {
  if [[ ! -d "$1" ]]; then return 1; fi
  # echo "touchdbyfile $@"
  local _dir_=${1%/}
  local _lvl_=$((${2:-0} - 1))
  local _opt_=${3%/}
  local _dig_="yes"
  local _old_=$(date '+%Y-%m-%d %H:%M:%S' -r "$1" 2>/dev/null)
  local _now_=$(date '+%Y-%m-%d %H:%M:%S')
  local _old_sec_=`date --date="${_old_}" +%s`
  local _now_sec_=`date --date="${_now_}" +%s`
  local _sec_=$((${_now_sec_} - ${_old_sec_}))
  local _num_=${_neardays_:-60}
  local _ddf_=$((${_sec_#-}/86400 - ${_num_}))
  local _dth_=$((${_lvl_} - ${_dirdepth_}))
  local _act_=""
  local _cur_=''
  local _sub_=''
  local _new_=''
  local _ymd_=''

  if [[ ${_lvl_} == 0 ]] || [[ ${_lvl_} == -255 ]]; then
    # Do NOT echo. Should be handled by upper caller.
    _dig_="no"
  fi

  for d in "${_dir_}"/*; do
    if [[ -d "$d" ]] && [[ "${_dig_}" == "yes" ]]; then
      touchdbyfile "$d" ${_lvl_} ${_opt_}
    fi
  done

  if [[ "${_dth_#-}" -gt 1 ]]; then
    if [[ ${_act_mode_} =~ --qui(ck|et) ]]; then
      if [[ ${_ddf_} -gt 0 ]]; then
        _day_skip_="${_ddf_}"
      elif [[ ! "${_day_skip_}" == "" ]]; then
        _day_skip_=""
      fi
    fi
  fi

  if [[ ! "${_day_skip_}" == "" ]]; then
    if [[ ! ${_act_mode_} =~ --qui(ck|et) ]]; then
      echo "Skipping ${_old_} beyond ${_num_} days on ${_dir_}"
      _new_line_=""
    fi
    return ${_ddf_}
  fi

  for f in "${_dir_}"/*; do
    if [[ -d "$f" ]]; then
      _ymd_=${_sub_}
    else
      _ymd_=${_cur_}
    fi
    _new_=$(date '+%Y-%m-%d %H:%M:%S' -r "$f" 2>/dev/null)
    if [[ "${_opt_}" == "--asc-sort" ]]; then
      if [[ "${_ymd_}" == "" ]] || \
         [[ "${_new_}" < "${_ymd_}" ]] ; then
        _ymd_=${_new_}
      fi
    else
      if [[ "${_new_}" > "${_ymd_}" ]]; then
        _ymd_=${_new_}
      fi
    fi
    if [[ -d "$f" ]]; then
      _sub_=${_ymd_}
    else
      _cur_=${_ymd_}
    fi
  done

  _ymd_=${_cur_:-${_sub_:-${_new_:-${_old_}}}}

  # echo "_dir_=${_dir_}"
  # echo "_ymd_=${_ymd_}"
  # echo "_old_=${_old_} ,_new_=${_new_}"
  # echo "_cur_=${_cur_} ,_sub_=${_sub_}"
  # echo "_opt_=${_opt_}"

  if [[ "${_ymd_}" == "${_old_}" ]]; then
    _act_="Matching ${_ymd_} on ${_dir_}"
  fi
  if [[ "${_ymd_}" > "${_old_}" ]] && \
     [[ "${_opt_}" == "" ]]; then
      _act_="Reserved ${_old_} on ${_dir_}"
  fi
  if [[ "${_act_}" == "" ]]; then
    if [[ "${_new_line_}" == "" ]]; then
      echo ""
    fi
    echo Applying ${_ymd_} to ${_dir_} [${_old_}]
    touch -d "${_ymd_}" "${_dir_}"
    _new_line_="done"
    echo ""
  elif [[ ! ${_act_mode_} =~ --quiet ]]; then
    _new_line_=""
    echo ${_act_}
  fi
}

############################################################
# function: Use `touch -d` to apply all basepath recursively
# Params: $1 a source dir paths
############################################################
function touchdpath {
  if [[ ! -d "$1" ]]; then return 1; fi
  # echo "touchdpath $@"
  local _spath_="$( cd "$( echo "${1}" )" && pwd )"
  local _sbase_="$( cd "${_spath_}/.." && pwd )"
  local _upper_="$( cd "${_sbase_}/.." && pwd )"
  local _slash_=${_spath_//[!\/]}
  local _depth_=${#_slash_}
  local _order_=$2

  if [[ ${_depth_} -gt 2 ]]; then
    touchdbyfile "${_spath_}" 1 ${_order_}
    touchdpath "${_sbase_}" ${_order_}
  fi
}

############################################################
# function: Use `touch -d` on file/dir
# Params: a file/dir, or FMT "%Y-%m-%d %H:%M"
############################################################
function touchd() {
  local _awk_="awk '{print \$6,\$7}'"
  local _arg_='-l --time-style=long-iso'
  local _asc_sort_=''
  local _act_mode_='--quick'
  local _neardays_=${NEARDAYS:-60}
  local _day_skip_=''
  local _fmt_date_='%Y-%m-%d %H:%M'
  local _fmt_regx_='[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]( [0-9][0-9]:[0-9][0-9])?'
  local _iso_date_=''
  local _dir_from_=''
  local _dir_file_=''
  local _dir_path_=''
  local _dirdepth_=${DIRDEPTH:-3}
  local _grp_dirs_=()
  local _grp_file_=()
  local _new_line_=""

  # echo "---args: $@"
  for p in "$@"; do
    if [[ "$p" =~ ^${_fmt_regx_}$ ]]; then
      _iso_date_="$p"
    elif [[ -e "$p" ]]; then
      if [[ -d "$p" ]]; then
        _grp_dirs_+=("${p%/}")
        _dir_path_="${p%/}"
      elif [[ "${_iso_date_}" == "" ]] && \
         [[ "${_dir_from_}" == "" ]]; then
        _iso_date_=`date +"${_fmt_date_}" -r "$p"`
        _dir_from_="${p%/}"
      else # non-directory
        # echo "File: $p"
        _grp_file_+=("${p%/}")
        _dir_file_="${p%/}"
      fi
      # if [[ $p =~ (${_fmt_regx_}) ]] && \
      #    [[ "${_iso_date_}" == "" ]]; then
      #   _iso_date_="${BASH_REMATCH[1]} 00:00"
      # fi
    elif [[ "$p" =~ ([/-]{1,2}L?)([0-9]{1,3}) ]]; then
      _dirdepth_=${BASH_REMATCH[2]}
    elif [[ "$p" =~ ([/-]{1,2}N)([1-9][0-9]{0,3}) ]]; then
      _neardays_=${BASH_REMATCH[2]}
    elif [[ "$p" =~ [/-]{1,2}[vV] ]]; then
      _asc_sort_='--asc-sort'
    elif [[ "$p" =~ [/-]{1,2}[fF] ]]; then
      _act_mode_='--always'
    elif [[ "$p" =~ [/-]{1,2}[kK] ]]; then
      _act_mode_='--quick'
    elif [[ "$p" =~ [/-]{1,2}[qQ] ]]; then
      _act_mode_='--quiet'
    fi
  done
  if [[ "${_dirdepth_}" == "0" ]]; then
    _act_mode_='--always'
  fi

  echo ""
  # echo "Date: ${_iso_date_}"
  if [[ "${_iso_date_}" =~ ${_fmt_regx_} ]]; then
    echo "Fetched: '${_iso_date_}' from ${_dir_from_}"
    if [[ ${#_grp_dirs_[@]} -gt 1 ]]; then
      echo "Applying '${_iso_date_}' on dirs..."
      for _dir_ in "${_grp_dirs_[@]}"; do
        echo "Applying '${_iso_date_}' on ${_dir_}"
        touch -d "${_iso_date_}" "${_dir_}"
      done
    elif [[ -d "${_dir_path_}" ]]; then
      echo "Applying '${_iso_date_}' on ${_dir_path_}/*"
      touch -d "${_iso_date_}" "${_dir_path_}"/* && echo OK
    fi
    if [[ ${#_grp_file_[@]} -gt 1 ]]; then
      echo "Applying '${_iso_date_}' to files..."
      for _file_ in "${_grp_file_[@]}"; do
        echo "Applying '${_iso_date_}' to ${_file_}"
        touch -d "${_iso_date_}" "${_file_}"
      done
    elif [[ -e "${_dir_file_}" ]]; then
      echo "Applying '${_iso_date_}' to ${_dir_file_}"
      touch -d "${_iso_date_}" "${_dir_file_}" && echo OK
    fi
  elif [[ -d "${_dir_path_}" ]]; then
    if [[ ! "${_asc_sort_}" == "" ]]; then
      echo "+-----------------+"
      echo " Sorting by oldest "
      echo "+-----------------+"
    fi
    if [[ ${_act_mode_} =~ --qui(ck|et) ]]; then
      _act_mode_="${_act_mode_} [mode] near ${_neardays_} days, depth: ${_dirdepth_}"
    fi
    echo "Recuring on ${#_grp_dirs_[@]} dir(s) ... ${_act_mode_}"
    _new_line_="done"
    echo ""
    for _dir_ in "${_grp_dirs_[@]}"; do
      local _dir_base_="$( cd "${_dir_}/.." && pwd )"
      if [[ ${_dirdepth_} -ne 0 ]]; then
        touchdbyfile "${_dir_}" ${_dirdepth_} ${_asc_sort_}
      else
        touchdpath "${_dir_}" ${_asc_sort_}
      fi
    done
  else
    echo "\$ touchd $@ [args]"
    echo ""
    echo "┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓"
    echo "┃ Syntax:                                          ┃"
    echo "┃  touchd <file>|'yyyy-mm-dd HH:MM' <dir>|<file>   ┃"
    echo "┃     or:                                          ┃"
    echo "┃  touchd <dir> [-L<depth>] [-fkq] [-v]            ┃"
    echo "┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛"
    echo ""
  fi
  if [[ "${_new_line_}" == "" ]]; then
    echo ""
  fi
}



if [[ "$0" == "${BASH_SOURCE}" ]]; then
  echo ""
  echo '.........................................................................'
  echo '!! Please `source` this script in order to export envirnment variables !!'
  exit 9
else
  main $@
fi
