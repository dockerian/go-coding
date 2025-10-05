#!/usr/bin/env bash
#######################################################################
# Example of ~/.bashrc for Mac OS X - usage: `source bashrc.sh`
#######################################################################
shopt -s promptvars
echo ""

########################################
# remove the longest match of *. from the beginning
# file_type="${file##*.}"
# remove the shortest match of .* from the end
# file_name="${file%.*}"
script_path="$( readlink -f "${BASH_SOURCE[0]}" 2>/dev/null || echo ${BASH_SOURCE[0]} )"
script_base="$( cd "${script_path%/*}" 2>/dev/null && pwd )"
script_file="${script_path##*/}" # aka `basename`
script_type="${script_file##*.}"
script_name="${script_file%.*}"
echo "script path: ${script_path}"
echo "script base: ${script_base}"
echo "script file: ${script_file}"
echo "script name: ${script_name}"
echo "script type: ${script_type}"
echo ""
########################################

# see:
#   - http://omar.io/ps1gen/
#   - https://www.kirsle.net/wizards/ps1.html
#   - https://www.askapache.com/linux/bash-power-prompt/
PS1='\n\[\033[0;36m\]\h\[\033[0m\]:\[\033[0;35m\]\u\[\033[0m\] \W [\#]:\n\$ '
PS1='$(printf "%$((COLUMNS-1))s\r")'$PS1

# disable warning message "The default interactive shell is now zsh"
export BASH_SILENCE_DEPRECATION_WARNING=1
export HISTCONTROL=erasedups:ignoredups:ignorespace
export PROMPT_COMMAND='echo -ne "\033]0;${PWD/#$HOME/~}\007"'
export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin


# Java
 unset JAVA_HOME
export JAVA_HOME="${JAVA_HOME:-$(java -XshowSettings:properties -version 2>&1 > /dev/null | grep 'java.home'|awk -F='[ .]' '{print $2}')}"
export JAVA_HOME="${JAVA_HOME:-$(/usr/libexec/java_home 2>/dev/null)}"
export JAVA_HOME="${JAVA_HOME:-$(type -p javac|xargs greadlink -f 2>/dev/null|xargs dirname|xargs dirname)}"
export JAVA_HOME="${JAVA_HOME:-$(type -p javac|xargs dirname|xargs dirname)}"

# Go
 unset GOROOT
export GOROOT="${GOROOT:-$(type -p go|xargs readlink -n 2>/dev/null|xargs dirname|xargs dirname)}"
export GOPATH="${HOME}/go"

export HOMEBREW_GITHUB_API_TOKEN="d430484ccbfc32c58135b5a3e8e1bc92a5c5a1d8"

# MYSQL
# export PATH="$MYSQL_HOME/bin:$MAVEN_HOME/bin"
# export PATH="/Library/PostgreSQL/10/bin:$PATH"
# export PATH="/usr/local/bin:$PATH"

# For compilers to find libpq
export LDFLAGS="-L/usr/local/opt/libpq/lib"
export CPPFLAGS="-I/usr/local/opt/libpq/include"
export PATH="/usr/local/opt/libpq/bin:$PATH"

# For compilers to find mysql-client
export LDFLAGS="-L/usr/local/opt/mysql-client/lib"
export CPPFLAGS="-I/usr/local/opt/mysql-client/include"
export PATH="/usr/local/opt/mysql-client/bin:$PATH"

# For compilers to find binutils
export LDFLAGS="-L/usr/local/opt/binutils/lib"
export CPPFLAGS="-I/usr/local/opt/binutils/include"

# For using GNU commands without prefix "g"
export PATH="$(brew --prefix coreutils)/libexec/gnubin:$PATH"
export PATH="$(brew --prefix findutils)/libexec/gnubin:$PATH"
export PATH="$(brew --prefix  binutils)/libexec/gnubin:$PATH"

export PATH="/usr/local/sbin:$PATH"
export PATH="/usr/local/bin:$PATH" # Add python and aws
export PATH="${JAVA_HOME}/bin:$PATH" # Add java
export PATH="${GOPATH}/bin:$PATH" # Add golang


echo "Loading bash aliases ..."
alias a="alias|cut -d' ' -f 2- "
alias airdrop='mdfind $HOME com.apple.AirDrop'
alias bashrc='source ~/.bash_profile; title ${PWD##*/};'
alias brewery='sudo chown -R "$USER":admin /usr/local/* && brew upgrade; brew list --cask | xargs brew upgrade 2>/dev/null; brew update --debug --verbose && brew cleanup; brew doctor'
alias bu='brew update && brew upgrade && brew cleanup'
alias buc='brew list --cask | xargs brew upgrade && brew upgrade && brew cleanup'
alias cdp='cd -P .'
alias clean='find . -type f \( -name *.DS_Store -o -name Thumbs.db \) -delete 2>/dev/null'
alias cls='clear && printf "\e[3J"'
alias conv='iconv -f windows-1252 -t utf-8'
alias convgbk='iconv -f gbk -t utf-8'
alias dater='date +"%Y-%m-%d %H:%M:%S" -r'
alias dated='date +"%Y-%m-%d %H:%M:%S" -d'
alias dh='du -hs'
alias dir='ls -al '
alias demons='launchctl list'
alias dns='scutil --dns # or see `cat /etc/resolv.conf`'
alias dsclean='sudo find . -type f \( -name *.DS_Store -o -name Thumbs.db -o -name *.pyc \) -delete'
alias dsf1='diskutil secureErase freespace 1'
alias dswake='wakeonlan -i 192.168.1.218 00:11:32:aa:e3:5d'
alias envi='env | grep -i '
alias envs='env | sort'
alias fixcr='perl -i -pe '"'"'s/\r//g'"'" # remove carriage return ('\r')
alias fixgrayedout='xattr -d com.apple.FinderInfo'
alias fixmod='for f in *; do if [[ -d "$f" ]] || [[ "${f##*.}" == "sh" ]]; then chmod 755 "$f"; else chmod 644 "$f"; fi; done'
alias fixrar='/Applications/rar/rar r'
alias fixunzip='ditto -V -x -k --sequesterRsrc ' # $1.zip $2/dir'
alias hide='chflags hidden'
alias hs='history | grep'
alias ip='echo $(ipconfig getifaddr en0) $(dig +short myip.opendns.com @resolver1.opendns.com)'
alias ipwan='dig +short myip.opendns.com @resolver1.opendns.com # or `curl -s icanhazip.com`'
alias javaenv='java -version && echo JAVA_HOME='`echo ${JAVA_HOME}`
alias ll='ls -al'
alias lll='ls -al -T | sort -f -k9,9'  # --time-style=full-iso
alias lln='ls -al | sort -f -k9,9'
alias llo='ls -l --time-style=long-iso'
alias llt='ls -al -rt'
alias lg='dscl . list /groups | grep -v "_"'
alias lgv='dscacheutil -q group' # -a name staff
alias lsofi='lsof -i -n -P'
alias lports='netstat -vanp tcp|grep -e pid '
alias lu='dscl . list /users | grep -v "_"'
alias luv='dscacheutil -q user' # -a name $USER
alias ml="make -qp|awk -F':' '/^[a-zA-Z0-9][^\$#\/\t=]*:([^=]|\$)/ {split(\$1,A,/ /);for(i in A)print A[i]}'|sort"
alias path='echo $PATH|tr ":" "\n"'

alias rarx='unrar x -kb'
alias setp='(set -o posix; set|grep -v _xspec)'
alias showhidden='defaults write com.apple.finder AppleShowAllFiles YES; killall Finder /System/Library/CoreServices/Finder.app'
alias si='echo -e $(for k in ~/.ssh/*.pub;do echo -e "\\\n$(ssh-keygen -E md5 -lf $k) - $k";done)|sort -k 3; echo;echo "--- Added identities ---"; ssh-add -E md5 -l|sort -k 3'
alias sshv='ssh -v -o HostKeyAlgorithms=ssh-dss -o KexAlgorithms=diffie-hellman-group14-sha1'
alias ver='echo -e "$(uname -a)"; echo ""; echo -e "$(bash --version)"'
alias vlc='/Applications/VLC.app/Contents/MacOS/VLC --width 800 --height 600 --aspect-ratio 16x9 &'
alias ydl='yt-dlp -f bestvideo[ext=mp4]+bestaudio[ext=m4a]/mp4' # -o '%(playlist_index)s.%(ext)s'
alias t='title ${PWD##*/}'


# docker-machine
export DOCKER_EMAIL=''
export DOCKER_USERNAME=''
export DOCKER_PASSWORD=''
export GITHUB_USERNAME=''
export GITHUB_PASSWORD=''
export GITHUB_EMAIL=""
export GITHUB_USER=""

echo "Loading bash aliases for docker cli ..."
alias dm="docker-machine "
alias dme="docker-machine env default"
alias dme_disable='eval $(docker-machine env -u)'
alias dme_enable='eval $(docker-machine env)'
alias dmip="docker-machine ip default"
alias dclean='docker kill $(docker ps -aq); docker rm -f -v $(docker ps -aq); docker rmi -f $(docker images -aq)'
alias denv='env|sort|grep DOCKER'
alias di="docker images|sort|grep -v none"
alias dia="docker images -a"
alias didangling="docker images -a --filter dangling=true"
alias diq="docker images -q"
alias dlogs="docker logs -ft "
alias dps="docker ps -a"
alias dpsq="docker ps -q"
alias drm="docker rm -f -v"
alias drma='docker rm -f -v $(docker ps -aq)'
alias drme='docker rm -f -v $(docker ps -aq --filter "status=exited")'
alias drmi='docker rmi -f '
alias dvrm='docker volume rm -f $(docker volume ls -q -f dangling=true)'
alias kc="kubectl"
alias mkenv='eval $(minikube docker-env)'
alias mkenv_disable='eval $(minikube docker-env -u)'
alias mkdb='echo "# run inside debug: apk add curl --no-cache" && kubectl run -i --tty --rm debug --image=alpine --restart=Never -- sh'
alias mk="minikube"

echo "Loading bash aliases for git ..."
alias gfork='/Applications/Fork.app/Contents/MacOS/Fork . &'
alias mygit='GIT_SSH_COMMAND="ssh -i ~/.ssh/github_jasonzhuyx_2048 -F ~/.ssh/config" git '
alias gbc='git symbolic-ref --short -q HEAD'
alias gbd='git branch -d '  # delete branch locally
alias gbdo='git push origin --delete '  # delete branch on origin
alias gbv="git branch -v "
alias gco="git checkout "
alias gcp='git cherry-pick '
alias gcdev='git checkout master && git pull upstream master && git push && git checkout dev && git rebase master && git push --force && git fetch --v --all --prune ; git branch -v'
alias gcr='git clone --recurse-submodules'
alias gfv="git fetch -v --all --prune ; git branch -v; git prune"
# git log --pretty=format
# * committer date:
#   - %cr (relative)
#   - %cd (respects --date= option)
#   - %cD (RFC2822 style)
#   - %cI (strict ISO 8601 format)
#   - %ci (ISO 8601-like format)
#   - %ct (UNIX timestamp)
alias glg="git log --graph --pretty=format:'%C(magenta)%h%Creset -%C(yellow)%d%Creset %s %Cgreen[%cd] %C(bold blue)<%an>%Creset' --abbrev-commit --date=iso-strict"
alias gpom='git checkout master && git pull origin master'
alias gpum='git checkout master && git pull upstream master'
alias gpumgp='git checkout master && git pull upstream master && git push'
alias gfom='git checkout master && git fetch --all --prune && git reset --hard origin/master; git prune'
alias grm='git rebase master'
alias grmgp='git rebase master && git push'
alias grmgpf='git rebase master && git push --force'
alias grom='git rebase origin/master'
alias grum='git checkout master && git featch --all --prune && git reset --hard upstream/master'
alias grv='git remote -v'
alias gst='git status'

echo "Loading bash aliases for dev ..."
alias apache='httpd -v; sudo apachectl '
alias exif='exiftool -sort -s'
alias ipy='ipython -i --ext=autoreload -c "%autoreload 2"'
alias ipy2='python2 -m IPython -i --ext=autoreload -c "%autoreload 2"'
alias ipy3='python3 -m ipython -i --ext=autoreload -c "%autoreload 2"'
alias goback='cd ${GOPATH}/$(cut -d/ -f2,3,4 <<< "${PWD/$GOPATH/}")'
alias gopath='cd -P ${GOPATH} && pwd'
alias pipf='pip freeze'
alias pipi='pip install'
alias pipiu='pip install --upgrade'
alias pipl='pip list'
alias pylib='pip show pip | grep Location | awk '\''{print substr($0, index($0,$2))}'\'''
alias pyserver='python -m SimpleHTTPServer'
alias sitedl='wget --mirror –w 2 –p --HTML-extension –-convert-links –P '
alias dvenv='deactivate'
alias venv='source .venv/bin/activate'
alias el='echo ".exit: Close the I/O stream, causing the REPL to exit."; echo ".help: Show this list of special commands."; echo; node_modules/.bin/electron --interactive'
alias nr="npm run "

# WineHQ app
echo "Loading bash aliases for wine ..."
alias wine='/Applications/WineHQ.app/Contents/Resources/wine/bin/wine'
alias wine64='/Applications/WineHQ.app/Contents/Resources/wine/bin/wine64'
alias winecfg='/Applications/WineHQ.app/Contents/Resources/wine/bin/winecfg'
alias geosetter='/Applications/WineHQ.app/Contents/Resources/wine/bin/wine ~/.wine/drive_c/App/geosetter/GeoSetter.exe'
alias iview='/Applications/WineHQ.app/Contents/Resources/wine/bin/wine64 /Users/jasonzhu/.wine/drive_c/App/iview/i_view64.exe'
alias npp='/Applications/WineHQ.app/Contents/Resources/wine/bin/wine ~/.wine/drive_c/App/npp/notepad++.exe'

echo ""



############################################################
# function: Main
############################################################
function main() {

  # echo `date +"%Y-%m-%d %H:%M:%S"` "Login to docker with ${DOCKER_USERNAME}..."
  # docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}
  # eval $(minikube docker-env)
  # eval $(docker-machine env)

  # (brew update && brew upgrade && brew cleanup) &
  find . -name '.DS_Store' -type f -delete 2>/dev/null &

  ipy_autoreload
  fixpath
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
# function: Add $1 (a dir path) to $PATH, if not added yet
############################################################
function addtopath() {
    if [ -d "$1" ] && [[ ! "$PATH" =~ (^|:)"${1}"(:|$) ]]; then
      export PATH+=:$1
    fi
}

############################################################
# function: Configure AWS profile
############################################################
function aws_config() {
  local profile=${1:-default}
  echo "Loading/Setting AWS CLI profile: ${profile}"
  local aws_access_key_id="$(aws configure get aws_access_key_id --profile ${profile} 2>/dev/null)"
  local aws_secret_access_key="$(aws configure get aws_secret_access_key --profile ${profile} 2>/dev/null)"
  local aws_default_region="$(aws configure get profile.${profile}.region)"
  local s3_bucket="$(aws configure get profile.${profile}.s3_bucket)"

  if [[ "${aws_access_key_id}" == "" ]]; then
    if [[ "${AWS_ACCESS_KEY_ID}" == "" ]]; then
      echo "ERROR: Environment variable AWS_ACCESS_KEY_ID is not set."
    else
      echo "  - setting aws_access_key_id to default and profile ${profile}"
      aws configure set aws_access_key_id ${AWS_ACCESS_KEY_ID}
      aws configure set aws_access_key_id ${AWS_ACCESS_KEY_ID} --profile ${profile}
    fi
  else
    echo "  - setting environment variable AWS_ACCESS_KEY_ID"
    export AWS_ACCESS_KEY_ID="${aws_access_key_id}"
  fi
  if [[ "${aws_secret_access_key}" == "" ]]; then
    if [[ "${AWS_SECRET_ACCESS_KEY}" == "" ]]; then
      echo "ERROR: Environment variable AWS_SECRET_ACCESS_KEY is not set."
    else
      echo "  - setting aws_secret_access_key to default and profile ${profile}"
      aws configure set aws_secret_access_key ${AWS_SECRET_ACCESS_KEY}
      aws configure set aws_secret_access_key ${AWS_SECRET_ACCESS_KEY} --profile ${profile}
    fi
  else
    echo "  - setting environment variable AWS_SECRET_ACCESS_KEY"
    export AWS_SECRET_ACCESS_KEY="${aws_secret_access_key}"
  fi
  if [[ "${aws_default_region}" == "" ]]; then
    if [[ "${AWS_DEFAULT_REGION}" == "" ]]; then
      echo "ERROR: Environment variable AWS_DEFAULT_REGION is not set."
    else
      echo "  - setting default.region and profile.${profile}.region"
      aws configure set default.region ${AWS_DEFAULT_REGION}
      aws configure set profile.${profile}.region ${AWS_DEFAULT_REGION}
    fi
  else
    echo "  - setting environment variable AWS_DEFAULT_REGION"
    export AWS_DEFAULT_REGION="${aws_default_region}"
  fi
  if [[ "${s3_bucket}" == "" ]]; then
    if [[ "${S3_BUCKET}" == "" ]]; then
      echo "ERROR: Environment variable S3_BUCKET is not set."
    else
      aws configure set default.s3_bucket ${S3_BUCKET}
      aws configure set profile.${profile}.s3_bucket ${S3_BUCKET}
    fi
  else
    export S3_BUCKET="${s3_bucket}"
  fi
}

############################################################
# function: docker inspect extract config in docker image
############################################################
function dex() {
  docker history --no-trunc "$1" | \
  sed -n -e 's,.*/bin/sh -c #(nop) \(MAINTAINER .[^ ]\) 0 B,\1,p' | \
  head -1
  docker inspect --format='{{range $e := .Config.Env}}
  ENV {{$e}}
  {{end}}{{range $e,$v := .Config.ExposedPorts}}
  EXPOSE {{$e}}
  {{end}}{{range $e,$v := .Config.Volumes}}
  VOLUME {{$e}}
  {{end}}{{with .Config.User}}USER {{.}}{{end}}
  {{with .Config.WorkingDir}}WORKDIR {{.}}{{end}}
  {{with .Config.Entrypoint}}ENTRYPOINT {{json .}}{{end}}
  {{with .Config.Cmd}}CMD {{json .}}{{end}}
  {{with .Config.OnBuild}}ONBUILD {{json .}}{{end}}' "$1"
}

############################################################
# function: Extract for common file formats
# -- https://github.com/xvoland/Extract/blob/master/extract.sh
############################################################
function extract() {
  IFS_SAVED=$IFS
  IFS=$(echo -en "\n\b")

  if [ -z "$1" ]; then
    # display usage if no parameters given
    echo "Usage: extract <path/file_name>.<zip|rar|bz2|gz|tar|tbz2|tgz|Z|7z|xz|ex|tar.bz2|tar.gz|tar.xz>"
    echo "       extract <path/file_name_1.ext> [path/file_name_2.ext] [path/file_name_3.ext]"
    return 1
  else
    for n in $@
    do
      if [ -f "$n" ] ; then
          case "${n%,}" in
            *.tar.bz2|*.tar.gz|*.tar.xz|*.tbz2|*.tgz|*.txz|*.tar)
                         tar xvf "$n"       ;;
            *.lzma)      unlzma ./"$n"      ;;
            *.bz2)       bunzip2 ./"$n"     ;;
            *.rar)       unrar x -ad ./"$n" ;;
            *.gz)        gunzip ./"$n"      ;;
            *.zip)       unzip ./"$n"       ;;
            *.z)         uncompress ./"$n"  ;;
            *.7z|*.arj|*.cab|*.chm|*.deb|*.dmg|*.iso|*.lzh|*.msi|*.rpm|*.udf|*.wim|*.xar)
                         7z x ./"$n"        ;;
            *.xz)        unxz ./"$n"        ;;
            *.exe)       cabextract ./"$n"  ;;
            *)
                         echo "extract: '$n' - unknown archive method"
                         return 1
                         ;;
          esac
      else
          echo "'$n' - file does not exist"
          return 1
      fi
    done
  fi

  IFS=$IFS_SAVED
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
# function: Rename git branch name locally and on origin
############################################################
function gb-rename() {
  echo "Fetching git branches ..."
  git fetch -v --all --prune
  echo ""

  local old_name=$1
  local new_name=$2
  # get current branch, optionally using:
  #   - `git branch --no-color | grep -E '^\*' | awk '{print $2}'`
  #   - `git symbolic-ref --short -q HEAD`)
  local current_branch="$(git rev-parse --symbolic-full-name --abbrev-ref HEAD)"
  echo "Current branch: ${current_branch}"
  echo ""

  if [[ "$2" == "" ]]; then
    echo "Missing argument(s) on renaming git branch: "
    echo ""
    echo "${FUNCNAME} old_name new_name"
    echo ""
    return -2
  fi

  if [[ "$1" == "master" ]] || [[ "$2" == "master" ]]; then
    echo "Cannot rename 'master' branch."
    echo ""
    return -1
  fi

  if [[ "$1" == "${current_branch}" ]] || [[ "$2" == "${current_branch}" ]]; then
    echo "Currently on branch [${current_branch}] to be renamed: "
    echo ""
    echo "${FUNCNAME} $1 $2"
    echo ""
    return 9
  fi

  local chk_name=""
  for b in $(git branch --no-color | grep -E '^ '); do
    if [[ "${b}" == "${new_name}" ]]; then
      echo "Branch name [${new_name}] already exists."
      echo ""
      return 2
    fi
    if [[ "${b}" == "${old_name}" ]]; then
      chk_name="${old_name}"
    fi
  done

  if [[ "${chk_name}" == "" ]]; then
    echo "Cannot find branch [${old_name}]. Please fetch and sync to origin."
    echo ""
    return 1
  fi

  git branch -m ${old_name} ${new_name}
  git push origin :${old_name} ${new_name}
  git push origin -u ${new_name}

  echo ""
  echo "Done."
  echo ""
}

############################################################
# function: Delete git branch locally and on origin/remote
############################################################
function gbd-all() {
  if [[ "$1" != "" ]] && [[ "$1" != "master" ]]; then
    git push origin --delete $1
    git branch -d $1
  else
    echo "Missing valid branch name in argument."
    echo ""
  fi
  git fetch --all --prune
  git branch -v
}

############################################################
# function: Find a directory in $GOPATH/src and change to it
############################################################
function goto() {
  cd $(find $GOPATH/src -type d -name "$1" 2>/dev/null | head -n 1); pwd
}

############################################################
# function: Add config for ipython with autoreload extension
############################################################
function ipy_autoreload() {
  echo ""
  local ipy_script="
c.InteractiveShellApp.extensions = ['autoreload']
c.InteractiveShellApp.exec_lines = ['%autoreload 2', 'print("")']
c.InteractiveShellApp.exec_lines.append('print(\"**ATTENTION**: %autoreload 2 enabled.\\n\")')
"
  local ipy_config=".ipython/profile_default/ipython_config.py"
  if [[ ! -e "$HOME/${ipy_config}" ]]; then
    echo "Saving to '~/${ipy_config}' ..."
    echo "${ipy_script}" > "$HOME/${ipy_config}"
  else
    echo "Please configure '~/${ipy_config}' for autoreload in ipython."
  fi
  echo "........................................................................"
  echo "${ipy_script}"
  echo "........................................................................"
  echo ""
}

############################################################
# function: Keep watching a log ($1) for keyword ($2)
############################################################
function logwatch () {
  if [[ ! -f "${1}" ]] || [[ "${2// /}" == "" ]]; then
    echo ""
    echo '........................................'
    echo "  Usage: ${FUNCNAME[0]} <log> '<keyword>'"
    echo ''
    return 9
  fi
  echo ""
  echo `date '+%Y-%m-%d %H:%M:%S'`" Start watching '$2' in $1"
  echo '...................'
  echo ""
  tail -f "$1" | while read LOGLINE
  do [[ "$LOGLINE" =~ "${2}" ]] && echo "$LOGLINE [$$]" && pkill -P $$ && break;
  done
}

############################################################
# function: List members for a spcific group
############################################################
function members () {
  dscl . -list /Users | while read user; do printf "$user "; dsmemberutil checkmembership -U "$user" -G "$*"; done | grep "is a member" | cut -d " " -f 1;
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
# function: Print info
############################################################
function myinfo () {
  printf "\n"
  printf "CPU   : "
  [[ -e /proc/cpuinfo ]] && cat /proc/cpuinfo 2>/dev/null | grep "model name" | head -1 | awk '{ for (i = 4; i <= NF; i++) printf "%s ", $i }'
  [[ -x "$(which sysctl)" ]] && printf "$(sysctl -n machdep.cpu.brand_string)"
  printf "\n"

  printf "Kernel: $(uname -r) $(uname -p) $(uname -m)"
  kded4 --version 2>/dev/null | grep "KDE Development Platform" | awk '{ printf " | KDE: %s", $4 }'
  printf "\n"
  printf "OS    : $(uname -s)\n"
  # cat /etc/issue 2>/dev/null | awk '{ printf "OS    : %s %s %s %s | " , $1 , $2 , $3 , $4 }'
  printf "Host  : $(uname -n)\n"
  uptime | awk '{ printf "Uptime: %s %s %s", $3, $4, $5 }' | sed 's/,//g'
  printf "\n"
  cputemp 2>/dev/null | head -1 | awk '{ printf "%s %s %s\n", $1, $2, $3 }'
  cputemp 2>/dev/null | tail -1 | awk '{ printf "%s %s %s\n", $1, $2, $3 }'
  cputemp 2>/dev/null | awk '{ printf "%s %s", $1 $2 }'
}

############################################################
# function: List listen port on MacOS
############################################################
function netlisten () {
  netstat -Watnlv | grep LISTEN | awk '{"ps -o comm= -p " $9 | getline procname;colred="\033[01;31m";colclr="\033[0m"; print colred "proto: " colclr $1 colred " | addr.port: " colclr $4 colred " | pid: " colclr $9 colred " | name: " colclr procname;  }' | column -t -s "|"
  echo ""

  if [ $# -eq 0 ]; then
    sudo lsof -iTCP -sTCP:LISTEN -n -P
  elif [ $# -eq 1 ]; then
    sudo lsof -iTCP -sTCP:LISTEN -n -P | grep -i --color $1
  else
    echo "Usage: listening [pattern]"
  fi
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

############################################################
# function: Use youtube-dl or yt-dlp
# see
#   - https://github.com/lrvick/youtube-dl
#   - https://github.com/yt-dlp/yt-dlp
############################################################
function ydlo() {
  local _tool_=""

  if [[ -x "$(which yt-dlp)" ]]; then
    _tool_="yt-dlp"
  elif [[ -x "$(which youtube-dl)" ]]; then
    _tool_="youtube-dl"
  else
    echo ""
    echo "Cannot find yt-dlp or youtube-dl. See"
    echo "  - https://github.com/lrvick/youtube-dl"
    echo "  - https://github.com/yt-dlp/yt-dlp"
    echo ""
    return
  fi

  local _args_=""
  local _exec_=""
  local _href_=""
  local _name_=""
  local _sarg_=""
  local _earg_=""
  local _aarg_=""
  local _snum_=""
  local _user_=""
  local _pass_=""
  local _bmkv_="--merge-output-format mkv --remux-video mkv --preset-alias mkv"
  local _bmp3_="-f bestaudio -x --audio-format mp3 --audio-quality 0 --preset-alias mp3"
  local _bmp4_="-f bestvideo[ext=mp4]+bestaudio[ext=m4a]/mp4 --merge-output-format mp4 --preset-alias mp4"
  local _subt_="--write-subs --sub-format srt"
  local _enum_=""
  local _rvpl_=""
  # default sequence and extension for playlist
  local _extn_='-%(playlist_index)s.%(ext)s'
  local _ycmd_="${_tool_} ${_bmp4_}"
  # echo "---args: $@"
  for p in "$@"; do
    echo "# $p"
    if [[ "$p" =~ ^https?:// ]]; then
      _href_="$p"
    elif [[ "$p" =~ ^[0-9]+$ ]]; then
      if [[ "${_snum_}" == "" ]]; then _snum_="$p";
    elif [[ "${_enum_}" == "" ]]; then
      if [[ $p -gt ${_snum_} ]]; then _enum_="$p";
      else
        _enum_=$((${_snum_} + $p - 1)); fi; fi
    elif [[ "$p" =~ ^[/-]{1,2}mkv ]]; then
      _ycmd_="${_tool_} ${_bmkv_}"
    elif [[ "$p" =~ ^[/-]{1,2}mp3 ]]; then
      _ycmd_="${_tool_} ${_bmp3_}"
    elif [[ "$p" =~ ^[/-]{1,2}mp4 ]]; then
      _ycmd_="${_tool_} ${_bmp4_}"
    elif [[ "$p" =~ ^[/-]{1,2}best ]]; then
      _ycmd_="${_tool_}"
    elif [[ "$p" =~ ^[/-]{1,2}nosub ]]; then
      _subt_=""
    elif [[ "$p" =~ ^[/-]{1,2}[rR] ]]; then
      _extn_='-%(playlist_autonumber)02d.%(ext)s'
      _rvpl_="--playlist-reverse"
    elif [[ "$p" =~ ^[/-]{1,2}u[:=](.*)$ ]]; then
      echo "MATCH: ${BASH_REMATCH[0]}"
      _user_=${BASH_REMATCH[1]}
    elif [[ "$p" =~ ^[/-]{1,2}p[:=](.*)$ ]]; then
      _pass_=${BASH_REMATCH[1]}
    else
      _name_="$p"
    fi
  done

  if [[ "${_href_}" == "" ]]; then
    echo "┏━━━━━━━━┓"
    echo "┃ Syntax ┃"
    echo "┗━━━━━━━━┛"
    echo "  ${FUNCNAME[0]} [-r] [-best|mkv|mp4] <youtube_url> [<start#> [<end#>]]"
    echo ""
    return
  fi

  echo "----------"
  echo " name: ${_name_}"
  echo " href: ${_href_}"
  if [[ "${_href_}" =~ playlist ]]; then
    if [[ "${_name_}" =~ .*"-".* ]]; then
      _extn_='%(playlist_index)s.%(ext)s'
      if [[ ! "${_rvpl_}" == "" ]]; then
        _extn_='%(playlist_autonumber)02d.%(ext)s'
      fi
    fi
    if [[ ! "${_snum_}" == "" ]]; then
      _sarg_="--autonumber-start ${_snum_} ${_sarg_}"
      _sarg_="--playlist-start ${_snum_} ${_sarg_}"
      if [[ ! "${_enum_}" == "" ]]; then
        local _xend_=$((${_enum_} - ${_snum_} + 1))
        _sarg_="--autonumber-start ${_snum_} ${_sarg_}"
        _sarg_="--autonumber-size ${_xend_} ${_sarg_}"
        _sarg_="--playlist-end ${_xend_} ${_sarg_}"
        _sarg_="--playlist-start ${_snum_} ${_sarg_}"
      fi
      echo "start: ${_snum_}"
    fi
    if [[ ! "${_enum_}" == "" ]]; then
      if [[ "${_rvpl_}" == "" ]]; then
        _earg_="--playlist-end ${_enum_}"
        echo "  end: ${_enum_}"
      fi
    fi
    _args_=$(echo "${_args_} ${_subt_} ${_rvpl_}"|xargs)
  else # not from playlist, no need sequence
    _args_=$(echo "${_args_} ${_subt_}"|xargs)
    _extn_='.%(ext)s'
  fi
  if [[ "${_user_}" != "" && "${_pass_}" != "" ]]; then
    echo " user: ${_user_}"
    echo " pass: ${_pass_}"
    _aarg_="-u ${_user_} -p ${_pass_}"
  fi
  echo " args: ${_ycmd_}"
  echo "       ${_args_} ${_aarg_}"
  echo "       ${_sarg_} ${_earg_}"
  echo "----------"

  if [[ "${_name_}" == "" ]]; then
    _exec_="${_ycmd_} ${_sarg_} ${_earg_} ${_href_}"
    echo Downloading "${_href_}" ...
    ${_exec_}
    return
  fi

  # download with name
  echo Downloading "${_name_}""${_extn_}" ...
  echo ""
  ${_ycmd_} \
  ${_sarg_} ${_earg_} ${_args_} \
  -o "${_name_}""${_extn_}" \
  ${_href_}
}



if [[ "$0" == "${BASH_SOURCE}" ]]; then
  echo ""
  echo '.........................................................................'
  echo '!! Please `source` this script in order to export envirnment variables !!'
  exit 9
else
  main $@
fi
