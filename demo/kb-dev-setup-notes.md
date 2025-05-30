# Dev Setup Notes

> Developer machine setup with MacOS, Ubuntu, and Windows


<br/><a name="contents"></a>
## Contents

  * [Mac OS X Restore](#mac-os-x-restore)
  * [Setup on Mac OS X](#setup-mac-os)
    - [Keyboard Shortcuts](#keys)
    - [Dev Env Example](#devenv)
  * [Software Install on Mac OS X](#software-install-mac-os)
  * [Software Install on Ubuntu](#software-install-ubuntu)
    - [Command lines](#command-lines)
    - [Keyboard shortcuts](#keyboard-shortcuts)
    - [Install Apps](#install-apps)
  * [Software and Tools on Windows](#software-windows)
  * [VirtualBox Setup and Network](#vbox-setup)
  * [Online Tools](#online-tools)


<br/><a name="mac-os-x-restore"></a>
## Mac OS X Restore
  * Recovery by holding [Option] key (https://support.apple.com/kb/DL1433)
  * Re-install OS X by holding [`Command+R`] during restart from [macOS Recovery](https://support.apple.com/en-us/102518)
  * System Integrity protection
    - Boot with holding [`Command+R`]
    - Launch Terminal from Utilities menu
    - Disable System Integrity Protection (partially)

      ```bash
      csrutil enable --without debug
      ```

    - Revert System Integrity Protection

      ```bash
      csrutil clear
      ```


<br/><a name="setup-mac-os"></a>
## Setup (Mac OS)

  * Apple ID issue
    - Previously installed app may use a different Apple ID
    - App Store prompts "Update unavailable with this Apple ID"
    - Fix: rename or remove "Contents/\_MASRecipt" folder in app package
    - Associate current user with a valid Apple ID
    - And retry update in App Store

  * Apple Startup key combinations for Mac OS
    - See [Apple Support](https://support.apple.com/en-us/102603)
    - Recovery mode (start up from MacOS Recovery over the Internet)
      * `Command-R`: reinstalling the latest MacOS installed on Mac.
      * `Option-Command-R`: upgrading to the latest MacOS that is compatible.
      * `Shift-Option-Command-R`: reinstalling the macOS that came with Mac, or the version closest to it that is still available.
    - `Shift(⇧)`: Safe mode.
    - `Command-S`: Single-user mode.
    - `Command-V`: Start up in verbose mode.
    - `C`: Start up from an available CD, DVD, or USB drive
    - `D`: Start [Apple Diagnostics](https://support.apple.com/en-us/102550), or [`Option-D`] over the internet.
    - `N`: Start up from a [compatible](https://support.apple.com/en-us/101676) NetBoot server. To use the default boot image on the NetBoot server, hold down [`Option-N`] instead.
    - `Option (⌥)` or [`Alt`]: [Startup Manager](https://support.apple.com/guide/mac-help/change-your-mac-startup-disk-mchlp1034/mac), to choose other startup disk.
    - `Option-Command (⌘)-P-R`:	Reset NVRAM or PRAM.
    - `T`: Start up in [target disk mode](https://support.apple.com/guide/mac-help/transfer-files-mac-computers-target-disk-mode-mchlp1443/mac).
    - `X`: Start up from startup disk (e.g. a Windows partition), or use Startup Manager.
    - `Eject (⏏)` or [`F12`]: Eject removable media, such as an optical disc.

  * Check kernal

    ```
    uname -v # /usr/local/opt/coreutils/libexec/gnubin/uname [-vipsorm]
             # /usr/bin/uname [-amnprsv]
    ```

  * Get basic info

    ```bash
    sw_vers
    uname -a
    system_profiler SPSoftwareDataType
    system_profiler -detailLevel full # -xml output to XML
    /usr/sbin/sysctl -a machdep.cpu
    ```

  * Arrange windows / menu bar / finder / mission control
    - http://apple.stackexchange.com/questions/9659/what-window-management-options-exist-for-os-x
    - https://computers.tutsplus.com/tutorials/customizing-the-os-x-menu-bar--mac-49391

  * Mac OS X key shortcuts <a name="keys"></a>
    * block/column selection
      - Atom: `Control+Shift+`[`up`|`down`].
      - Atom with Sublime Column Selection package: `Shift+`[`left`|`right`].
      - TextMate: `Option (⌥)` or [`Alt`] to toggle for a selection.
    * force quit: `Cmd+Alt+Esc` (Force Quit) or `Cmd+Alt+Esc` (for active window)
    * moving cursor between words (iTerm)
      - `⌥+←` or `Alt/Option+Left`: (Send Escape Sequence) `^[B` (ESC+B)
      - `⌥+→` or `Alt/Option+Right` (Send Escape Sequence) `^[F` (ESC+F)
    * moving cursor to beginning and end of line (Terminal)
      - `Fn+←`: `Home`, Send Text: `\001` (no modifier)
      - `Fn+→`: `End`, Send Text: `\005` (no modifier)
    * clear current command line during input
      - `ESC` or `Cmd+L`: Send Hex `0x05 0x15` (moving to end and delete line)
    * screenshot:
      - `Cmd+Shift+3`: copy screen to file;
      - `Cmd+Shift+Alt+3`: copy screen to clipboard.
      - `Cmd+Shift+4`: start rectangle selection;
      - `Cmd+Shift+4`, `Space`: to select a window or dialog;
      - `Cmd+Shift+Alt+4` to clipboard (original `Control+Cmd+Shift+4`).
      - `Cmd+Shift+5`: screenshot or recording options.
    * show hidden files in open file dialog: `Cmd+Shift+.`(Period).
    * switch between app windows: Cmd+\` (back-quote above Tab key)
    * switch between Finder history: `Cmd+[`, `Cmd+]`
    * see
      - https://support.apple.com/en-us/HT201236
      - https://www.danrodney.com/mac/

  <br/><a name="devenv"></a>
  * Dev Env [example on macOS 15 Sequoia](https://blog.driftingruby.com/my-development-environment-on-macos-15-sequoia/)

    - Installing Homebrew

      ```bash
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      ```
      and then be prompted to run some additional commands

      ```
      (echo; echo 'eval "$(/usr/local/bin/brew shellenv)"') >> /Users/$(whoami)/.zprofile
      eval "$(/usr/local/bin/brew shellenv)"
      ```

    - Installing from previous `brew list`

      ```
      # brew_list.txt was from `brew list` standard output
      xargs brew install < brew_list.txt
      ```

    - Installing core packages

      ```
      brew install \
        autoconf bash ed flex \
        binutils coreutils diffutils findutils moreutils \
        gnu-indent gnu-sed gnu-tar gnu-which \
        grep gawk gpatch gzip less \
        gcc git ffmpeg jq perl ruby rar rsync \
        m4 make nano openssh rar screen tmux \
        wakeonlan watch wdiff wget
      ```

    - Installing extendable version manager (with support for Ruby, Node.js, Erlang & more)

      ```
      brew install asdf
      ```

    - Installing Apps

      ```
      brew install \
        iterm2 \
        orbstack \
        1password \
        cleanshot \
        cyberduck \
        obsidian \
        google-chrome \
        rectangle \
        magic-wormhole \
        ollama --cask \
        tableplus \
        tree

      # For other Rails development with both MySQL and PostgreSQL.
      brew install libpq mysql-client

      # Install VS Code
      brew install --cask visual-studio-code

      # Intellij IDEA Community Edition
      brew install --cask intellij-idea-ce

      # Android Studio
      brew install --cask android-studio
      ```

    - Installing Python ([the right way](https://medium.com/marvelous-mlops/the-rightway-to-install-python-on-a-mac-f3146d9d9a32))

      ```
      xcode-select --install
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      brew update; brew upgrade
      brew install pyenv pyenv-virtualenv
      ```
      and the following in `.bash_profile`:
      ```
      eval "$(pyenv init -)"
      if which pyenv-virtualenv-init > /dev/null; then eval "$(pyenv virtualenv-init -)"; fi
      ```

    - Fixing lock issue

      ```
      # ls $(brew --prefix)/var/homebrew/locks
      sudo chown -R "$USER":admin /usr/local/*
      brew cleanup
      brew doctor
      ```

    - Cron job

      ```
      brew update && brew upgrade && brew cleanup
      ```
      **Note** `brew cleanup` only removes files (under `~/Library/Caches/Homebrew`) more than 120 days old; to force removing all, use

      ```
      brew cleanup --prune=all ; or =<days>
      ```

    - [Oh-my-zsh](https://ohmyz.sh/?ref=blog.driftingruby.com) | [comparison](https://www.educba.com/zsh-vs-bash/)

      ```
      sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

      # and separate a `.devrc` from the main `.zshrc` file
      echo 'source ~/.devrc' >> ~/.zshrc
      ```

  * Finder
    - Preview `*.webp`
      ```
      brew install --cask WebPQuickLook
      # available with Preview after restart
      # and licensing requirements for MacOS Catalina
      xattr -d com.apple.quarantine ~/Library/QuickLook/WebpQuickLook.qlgenerator
      ```
    - Show hidden files in dialog: pressing `CMD + Shift + '.'`

    - For hidden dot-underscore (`._*`) files not shown in Finder on the drive that does not support [AppleDouble](https://en.wikipedia.org/wiki/AppleSingle_and_AppleDouble_formats) format, use [`dot_clean`](https://ss64.com/mac/dot_clean.html) command.

    - Show hidden files in Finder

      ```
      defaults write com.apple.finder AppleShowAllFiles YES
      killall Finder
      ```
      or using the following AppleScript (in Script Editor):

      ```
      set vis to do shell script "defaults read com.apple.Finder AppleShowAllFiles"

      if vis = "0" then
      	do shell script "defaults write com.apple.Finder AppleShowAllFiles 1"
      else
      	do shell script "defaults write com.apple.Finder AppleShowAllFiles 0"
      end if

      tell application "Finder" to quit
      delay 1
      tell application "Finder" to activate
      ```
      or change file/flder hidden flag

      ```
      chflags hidden|nohidden folder_or_file
      ```

  * Bash
    - Default Shell:
      * Go to `System Settings`|`Users & Groups`
      * Hold `Control` key and click on the user
      * Select `Advanced Options`
      * Change `Login Shell`
    - Change Shell:

      ```
      chsh -s /bin/bash # or chsh -s /bin/zsh
      ```

    - Behind [OS X's Unix tools](http://robservatory.com/behind-os-xs-modern-face-lies-an-aging-collection-of-unix-tools/)
    - Shell [Parameter Expansion](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html) | [ss64](https://ss64.com/mac/syntax-expand.html)
    - Advanced [Bash-Scripting Guide](http://tldp.org/LDP/abs/html/index.html)
    - File [test operators](http://tldp.org/LDP/abs/html/fto.html)
    - Example: Rename

      ```
      mv $dir/{oldname,newname}
      ```

  * Bash update (using `#!/usr/bin/env bash` instead of `#!/bin/bash`)

      ```
      brew update
      brew install bash bash-completion
      echo $BASH_VERSION
      # changing login shell to new version bash
      # chsh -s $(brew --prefix)/bin/bash
      ```

  * Change default opener ("Always Open With")
    - select a file with some extension
    - open context menu and click "`Get Info`"
    - select "`Open With`"
    - click "`Change All`"

  * Command history

    ```bash
    # ignore duplicate commands, ignore commands starting with a space
    export HISTCONTROL=erasedups:ignorespace
    # keep the last 5000 entries
    export HISTSIZE=5000
    # append to the history instead of overwriting (good for multiple connections)
    shopt -s histappend
    ```

  * Delete all `*.pyc` and "`.DS_Store`" recursively

    ```bash
    find . -type f \( -name *.DS_Store -o -name Thumbs.db \) -delete 2>/dev/null
    find . -name *.pyc -delete
    rm -rf **/*.pyc
    ```

  * Disable `.DS_Store`

    ```
    defaults write com.apple.desktopservices DSDontWriteNetworkStores -bool TRUE
    defaults write com.apple.desktopservices DSDontWriteUSBStores -bool TRUE
    ```

  * Disable Airplay

    ```bash
    sudo chmod 000 /System/Library/CoreServices/AirPlayUIAgent.app/Contents/MacOS/AirPlayUIAgent
    ```

  * Disable/Enable startup chime/sound

    ```bash
    cd /Applications/Utilities/
    sudo nvram SystemAudioVolume=%80 # disable the sound
    sudo nvram SystemAudioVolume=%01 # some mac may require different syntax
    sudo nvram -d SystemAudioVolume  # enable the sound
    ```
    **Note** in new OSX: `System Settings` => `Sound`, uncheck "`Play sound on startup`".

  * DNS:
    - using `dig` or `host` over `nslookup`
    - dig [web](http://www.digwebinterface.com/)
    - dns [web](https://www.whatsmydns.net)

  * Find WAN IP

    ```
    dig +short myip.opendns.com @resolver1.opendns.com
    wget -O - -q http://whatismyip.org/
    curl -s https://4.ifcfg.me/
    curl -s icanhazip.com
    ```

  * Get IP addresses

    ```bash
    ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1'
    ```

  * List DNS

    ```
    scutil --dns # or using `cat /etc/resolv.conf`
    ```

  * List demons/agents

    ```
    launchctl list
    ```

  * List all users

    ```
      dscl . list /Users | grep -v "^_"
    ```

  * List all open ports

    ```
    netstat
    netstat -nr # routing table (route print)
    netstat -atp tcp | grep -i "listen"
    netstat -Watnlv | grep LISTEN | \
      awk '{"ps -o comm= -p " $9 | getline procname;colred="\033[01;31m";colclr="\033[0m"; print cred "proto: " colclr $1 colred " | addr.port: " colclr $4 colred " | pid: " colclr $9 colred " | name: " colclr procname;  }' | column -t -s "|"
    sudo lsof -i -n -P | grep "listen"
    sudo lsof -PiTCP -sTCP:LISTEN
    ```
    see
    - https://wilsonmar.github.io/ports-open/
    - https://stackoverflow.com/questions/4421633/who-is-listening-on-a-given-tcp-port-on-mac-os-x

  * Mount NTFS drive
    - see https://www.howtogeek.com/236055/how-to-write-to-ntfs-drives-on-a-mac/
    - For Seagate/Samsung/Maxtor, use Paragon Driver for Mac OS (10.9. and above)
      - [free](http://www.seagate.com/support/downloads/item/ntfs-driver-for-mac-os-master-dl/) or
      - [buy](https://www.paragon-software.com/ufsdhome/ntfs-mac/)

    * Using [Mounty](https://mounty.app/)
      - download: https://mounty.app/releases/Mounty.dmg
      - brew: `brew cask install mounty`

    - Using OSX Fuse
      - download [OSX Fuse](https://github.com/osxfuse/osxfuse/releases)
      - or install osxfuse from cmd (```brew install Caskroom/cask/osxfuse```)
      - and Apple Developer Tools: ```xcode-select --install```
      - use `brew` to install ntfs-3g (```brew install homebrew/fuse/ntfs-3g```)
      - or simply ```brew install ntfs-3g```
      - create `/Volumes/NTFS` folder (`sudo mkdir -p /Volumes/NTFS`)
      - get disk identifier (```diskutil list```)
      - unmount the disk

        ```
        sudo umount $(diskutil list|awk '{if($2=="Windows_NTFS") print $NF}')
        ```
      - mount the disk

        ```
        sudo /usr/local/bin/ntfs-3g $(diskutil list|awk '{if($2=="Windows_NTFS") print $NF}') /Volumes/NTFS -olocal -oallow_other
        ```
      - review disk list (```diskutil list```)

    - For automatically mount with OSX Fuse (**not recommended**)
      - reboot Mac with holding `Command+R` to enter into recovery mode
      - disable SIP (```csrutil disable```) in Terminal
      - reboot normally
      - use new ```mount_ntfs```

        ```
        sudo mv /sbin/mount_ntfs /sbin/mount_ntfs.orig
        sudo ln -s /usr/local/Cellar/ntfs-3g/2015.3.14/sbin/mount_ntfs /sbin/mount_ntfs
        # or
        sudo ln -s /usr/local/sbin/mount_ntfs /sbin/mount_ntfs
        ```
      - reboot to recovery mode and re-enable System Integrity Protection
        ```
        csrutil enable
        ```
      - undo change

        ```
        sudo rm /sbin/mount_ntfs
        sudo mv /sbin/mount_ntfs.original /sbin/mount_ntfs
        brew uninstall ntfs-3g
        ```
    - For native support (**highly unrecommended**)
      - add `LABEL=VOLUME_NAME none ntfs rw,auto,nobrowse` to `/etc/fstab`
      - add difference name for each

  * MySQL
    - install by HomeBrew
      ```
      brew install mysql maven sonar-scanner
      ```
      configured to only allow connections from localhost
      ```
      mysql -u root
      ```
      start without a background service
      ```
      /usr/local/opt/mysql/bin/mysqld_safe --datadir\=/usr/local/var/mysql
      ```
      start as a background service
      ```
      brew services start mysql
      ```

    - start mysql.server

      ```
      sudo launchctl load -F /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist
      # or
      sudo /usr/local/mysql/support-files/mysql.server start  # or restart
      # or start in safe mode (without password)
      mysqld_safe --skip-grant-tables &
      ```

    - stop mysql.server

      ```
      sudo launchctl unload -F /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist
      # or
      sudo /usr/local/mysql/support-files/mysql.server stop
      ```

    - change root password: (ref)[https://dev.mysql.com/doc/refman/5.7/en/resetting-permissions.html]

      - creat a init file (e.g `~/mysql-init.sql`)

        ```
        UPDATE mysql.user
            SET authentication_string = PASSWORD('password'), password_expired = 'N'
            WHERE User = 'root' AND Host = 'localhost';
        FLUSH PRIVILEGES;
        ```

      - start mysqld with the init file

        ```
        mysqld --init-file=~/mysql-init.sql
        ```

      - or with existing password

        ```
        mysqladmin -u root --password='password' new_password
        # or
        mysql -u root --password='pass' -h 127.0.0.1 mysql  # using default db
        mysql> use mysql;
        mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'NewPassword';
        mysql> UPDATE user SET password=PASSWORD("NEWPASS") WHERE User='root';
        mysql> UPDATE user SET authentication_string=PASSWORD("NEWPASS") WHERE User='root';
        mysql> select user, password_expired, password_last_changed, password_lifetime, account_locked, authentication_string from user where user = 'root';
        mysql> FLUSH PRIVILEGES;
        mysql> quit
        ```

  * Rar
    - Official website: https://www.rarlab.com/download.htm
    - WinRar for Mac: https://www.techspot.com/downloads/5169-winrar-for-mac/similar/
    - See https://best-mac-tips.com/2013/02/01/install-free-command-line-unrar-mac/
    - Command "`unrar`"

      ```bash
      brew install rar  # no unrar formulae from brew
      ```

  * Ruby

    ```bash
    curl -L https://get.rvm.io | bash -s stable
    ```

  * Shell prompt
    - My default prompt

    ```
    PS1='\[\e[0;36m\]\h\[\e[0m\]:\[\e[0;35m\]\u\[\e[0m\] \W [\#]:\n\$ '
    ```

    - Prompt Generator
      - http://bashrcgenerator.com/
      - https://www.kirsle.net/wizards/ps1.html
      - https://xta.github.io/HalloweenBash/

  * Tail and watch

    ```bash
    tail -f tests.log | while read LOGLINE
    do [[ "$LOGLINE" =~ "keyword" ]] && echo "$LOGLINE [$$]" && pkill -P $$ && break;
    done
    ```

  * Verify/Repair Disk Permissions

    ```
    sudo /usr/libexec/repair_packages --verify --standard-pkgs /
    ```

  * User (or system) profile

    ```bash
    set -o posix; set
    /etc/paths.d
    /etc/bash_profile  # source /etc/bashrc
    /etc/bashrc
    ~/.bash_profile  # source ~/.bashrc
    ~/.bash_login
    ~/.profile
    ~/.bashrc
    ```
    see https://github.com/dockerian/go-coding/blob/master/tools/bashrc.sh

  * Widgets
    - [Countdown](https://www.apple.com/downloads/dashboard/status/countdownx.html)
    - [Currency Converter](https://www.apple.com/downloads/dashboard/calculate_convert/currencyconverter_palplesoftware.html)
    - [Starry Night](https://www.apple.com/downloads/dashboard/information/starrynightwidget.html)
    - [Symbol Candy](https://www.apple.com/downloads/dashboard/developer/symbolcaddy.html)
    - [Time Scroller](https://www.apple.com/downloads/dashboard/business/timescroller.html)
    - [What's Different](https://www.apple.com/downloads/dashboard/games/whatsdifferentbygeorge.html)

  * Xcode

    ```
    xcode-select --install
    ```

  * VLC
    - cannot start/launch

    ```
    rm ~/Library/Preferences/org.videolan.vlc/vlcrc  # or rename to vlcrc.bak
    # optionally to reset defaults if clearing preference does not work
    defaults delete org.videolan.vlc
    ```

  * Git/GitHub
    - SSH (see [SSH keys](#keys-ssh))

      ```
      # ssh-keygen -t rsa -b 4096 -C "dockeria@gmail.com"
      chmod 600 ~/.ssh/id*
      chmod 644 ~/.ssh/id*.pub
      # using `-key` option to specify key file
      # git clone <repo> -key
      ```

    - In `~/.ssh/config`, add:

      ```
      host github.com
       HostName github.com
       PreferredAuthentications publickey,keyboard-interactive,password
       IdentityFile ~/.ssh/id_rsa_github
       IdentitiesOnly yes
       User git
      ```
      Note `IdentitiesOnly` is to prevent from sending the default identity file for each protocol

    - Cache credential

      ```
      # set git to use the credential memory cache (default 15-minute)
      git config --global credential.helper cache
      # set the cache to timeout after 1 hour (setting is in seconds)
      git config --global credential.helper 'cache --timeout=3600'
      # store credential in mac os x key chain
      git config --global credential.helper osxkeychain
      ```

    - Clear credential in keychain

      ```
      git credential-osxkeychain erase
      host=github.com
      protocol=https
      [Press Return]
      ```

    - Diff/Merge Tool

      ```
      which opendiff
      git config --global merge.tool opendiff
      git config --global diff.tool opendiff
      git config --global difftool.prompt false
      ```
      should generate following in `~/.gitconfig`:

      ```
      [diff]
      	tool = opendiff
      [merge]
      	tool = opendiff
      ```

      for kdiff3 on Mac OS X, download from
      [sourceforge](https://sourceforge.net/projects/kdiff3/files/), or

      ```
      brew install kdiff3
      git config --global merge.tool kdiff3
      git config --global mergetool.kdiff3.path /Applications/kdiff3.app/Contents/MacOS/kdiff3
      ```
      should generate following in `~/.gitconfig`:

      ```
      [diff]
          tool = kdiff3
      [difftool "kdiff3"]
          cmd = /Applications/kdiff3.app/Contents/MacOS/kdiff3
          args = $base $local $other -o $output
          trustExitCode = false
      [merge]
          tool = kdiff3
      [mergetool "kdiff3"]
          cmd = /Applications/kdiff3.app/Contents/MacOS/kdiff3
          args = $base $local $other -o $output
          trustExitCode = false
      ```

    - Graph log

      ```
      git config --global alias.lg "log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative"
      ```
      The above is adding an alias to `~/.gitconfig`.

  * Fixing brew and ruby:

    ```bash
    cd $(brew --repo); git add . ; git fetch; git reset --hard origin/master
    brew update-reset

    sudo gem update --system
    ```

  * Python and pyenv

    ```bash
    brew install openssl readline sqlite3 xz zlib

    curl https://pyenv.run | bash

    # Add to ~/.bashrc
    export PATH="$HOME/.pyenv/bin:$PATH"
    eval "$(pyenv init -)"
    ```

  * Others
    - Install homebrew

      ```bash
      ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
      export PATH="$(brew --prefix coreutils)/libexec/gnubin:/usr/local/bin:$PATH"
      brew update # update homebrew
      brew upgrade # upgrade Packages
      brew doctor # check issue
      brew install # install package to /usr/local/Cellar
      brew list # list packages
      ```

    - Install coreutils

      ```bash
      brew install coreutils
      ln -s /usr/local/opt/coreutils/bin/greadlink /usr/local/bin/readlink
      # or
      alias readlink='/usr/local/opt/coreutils/bin/greadlink'
      ```

    - Install commonly-used commends

      ```bash
      brew tap homebrew/dupes # run only once
      # The --default-names option will prevent Homebrew from
      # prepending gs to the newly installed commands,
      # thus we could use these commands as default ones
      # over the ones shipped by OS X.
      brew install binutils
      brew install diffutils
      brew install ed --default-names
      brew install fcrackzip
      brew install findutils
      brew install gawk
      brew install gnu-indent
      brew install gnu-sed
      brew install gnu-tar
      brew install gnu-which
      brew install gnutls
      brew install grep
      brew install gzip
      brew install screen
      brew install watch
      brew install wdiff --with-gettext
      brew install wget
      ```

    - Install newer versions

      ```bash
      brew install bash
      brew install emacs
      brew install gdb  # gdb requires further actions to make it work. See `brew info gdb`.
      brew install gpatch
      brew install m4
      brew install make
      brew install nano
      ```

    - Install extra

      ```bash
      brew install file-formula
      brew install git
      brew install git-lfs
      brew install less
      brew install openssh
      brew install perl518   # must run "brew tap homebrew/versions" first!
      brew install python
      brew install rsync
      brew install svn
      brew install unzip
      brew install macvim
      brew install vim
      ```


<br/><a name="software-install-mac-os"></a>
## Software Install (Mac OS)

  * Admin
    - Add Administrator as Admin user: @dmini$tr@t0r
    - Enable root user: R00tU$er4M@cB00kPr0
    - Run <code>`sudo su`</code>

  * Bible Study (App store)

  * Clipboard managers:
    - [1Clipboard](http://1clipboard.io/) ★★★
    - AppStore [PasteBar](https://www.pastebar.app/) ★★★★
    - [ClipMenu](http://www.clipmenu.com/)
    - [ClipTools](https://macmost.com/cliptools) ★★★★★
    - [ClipUp](https://clickup.com/)
    - [CopyClip](https://itunes.apple.com/us/app/copyclip-clipboard-history/id595191960)
    - [Flycut - app store](https://itunes.apple.com/in/app/flycut-clipboard-manager/id442160987)
    - [Jumpcut](http://jumpcut.sourceforge.net/)

  * Developer editors
    - [CotEditor](https://coteditor.com/)
    - [TextMate](https://macromates.com/)
    - [MacVim](https://macvim.org/)

      ```
      # global settings in ~/.vimrc
      set guifont=Menlo\ Regular:h18
      ```

  * Download managers:
    - [FDM](https://www.freedownloadmanager.org/download-fdm-for-mac.htm)
    - [Folx](https://www.mac-downloader.com/) or from [eltima.com](https://mac.eltima.com/download-manager.html) ★★
    - [iGetter](http://www.igetter.net/downloads.html)
    - [Motrix](https://motrix.app/) ★★★★
    - [Progressive Downloader](https://www.macpsd.net/)
    - [qBittorrent](https://www.qbittorrent.org/) ★★★★★
    - [res-downloader](https://github.com/putyy/res-downloader) ★★★★
      * 🖥️ 支持Windows、Mac、Linux
      * 🌐 支持视频、音频、图片、m3u8、直播流等常见网络资源
      * 💪 支持微信视频号、小程序、抖音、快手、小红书、酷狗音乐、QQ音乐等网络资源下载
      * 👼 支持设置代理以获取特殊网络下的资源
    - [SEO Spider Tool](https://www.screamingfrog.co.uk/)
    - [SiteSucker](https://ricks-apps.com/osx/sitesucker/index.html) ★★
    - [Xtreme](https://sourceforge.net/projects/xdman/)

  * Free VPN with kill switch
    - [ProtonVPN](https://protonvpn.com/)
    - [Hotspot Shield](https://apps.apple.com/us/app/vpn-hotspot-shield-super-vpn/id443369807)
    - Free [TunnelBear](https://www.tunnelbear.com/)
    - [Windscribe](https://windscribe.com/)

  * Homebrew [brew.sh](http://brew.sh/)

    ```
    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
    brew update
    brew upgrade
    brew install tree
    ```

  * Java
    - [Basic guide](https://www.freecodecamp.org/news/how-to-set-up-java-development-environment-a-comprehensive-guide/)
      - see [versions history](java-history.png) | [wiki](https://en.wikipedia.org/wiki/Java_version_history)
    - [Jabba](https://github.com/shyiko/jabba) - a lightweight version manager

      ```
      brew install jabba
      jabba ls-remote
      jabba install <version> # ~/.jabba/jdk/
      jabba use <version>
      ```

    - [SDKMAN](https://sdkman.io/) - a dev kit manager

      ```
      curl -s "https://get.sdkman.io" | bash
      source "$HOME/.sdkman/bin/sdkman-init.sh"
      sdk list java
      sdk install java 17.0.1-open # sdk install java <version> at ~/.sdkman/candidates/java/
      sdk use java 17.0.1-open
      java -version
      ```

  * MongoDB

    - install

      ```
      brew install mongodb
      ```

  * Node.js

    - install [node](https://nodejs.org/en/download/package-manager/#macos)

      ```
      brew install node
      # alternatives:
      # port install nodejs7  # using MacPorts
      # pkgin -y install nodejs  # using pkgsrc
      # cd pkgsrc/lang/nodejs && bmake install  # build manually from pkgsrc
      # or
      # curl "https://nodejs.org/dist/latest/node-${VERSION:-$(wget -qO- https://nodejs.org/dist/latest/ | sed -nE 's|.*>node-(.*)\.pkg</a>.*|\1|p')}.pkg" > "$HOME/Downloads/node-latest.pkg" && sudo installer -store -pkg "$HOME/Downloads/node-latest.pkg" -target "/"
      ```

    - install [yarn](https://yarnpkg.com/lang/en/docs/install/)

      ```
      brew install yarn  # without node
      brew upgrade yarn
      # or from MacPorts
      sudo port install yarn
      # or by node
      node install -g yarn
      ```

  * Python

    - install python

      ```
      brew upgrade
      brew install python # to replace built-in python installation
      pip install --upgrade setuptools
      pip install ipython
      sudo easy_install pip
      sudo pip install ipython[all]
      sudo pip uninstall pyzmq
      sudo pip install pyzmq
      ```

    - unit test ?

      ```
      pip install -U pytest
      # install nose2
      pip install nose2
      # developer version
      pip install -e git+git://github.com/nose-devs/nose2.git#egg=nose2
      ```

  * Snort (NDIS - network intrusion )

    ```
    # install Xcode and MacPorts
    # install PCRE
    sudo port install pcre
    # or from [PCRE](http://pcre.org/) and https://ftp.pcre.org/pub/pcre/)
    ./configure & make & sudo make install
    # install wget
    sudo port install wget # or `brew install wget`
    # download snort (see https://snort.org/downloads)
    SNORT_VER=2.9.9.0 wget –no-check-certificate -O snort-${SNORT_VER}.tar.gz http://www.snort.org/ports/snort-current/snort-${SNORT_VER}.tar.gz
    # unpack
    SNORT_VER=2.9.9.0 tar zxvf snort-${SNORT_VER}.tar.gz && cd snort-${SNORT_VER}
    # make and install
    ./configure && make
    sudo make install
    ```

  * System tools
    - [94FBR](https://www.google.com/search?q=94FBR)
    - [Better365](https://www.better365.cn/apps.html) tools
    - [GPG Tools](https://gpgtools.org)
    - [AWS](http://docs.aws.amazon.com/cli/latest/userguide/installing.html)
    - [Adobe Flash Player](http://labs.adobe.com/downloads/flashplayer.html)
    - [Baidu](http://srf.baidu.com/input/mac.html)
    * Disk tools
      - [CC Cleaner](https://www.ccleaner.com/)
      - [Clean-Me](https://github.com/Kevin-De-Koninck/Clean-Me)
      - [Disk Drill 3](https://www.cleverfiles.com)
      - [DiskInventoryX](http://www.derlien.com/)
      - [Daisy Disk](https://daisydiskapp.com/) (free trial)
      - [Grand Perspective](http://grandperspectiv.sourceforge.net/)
      - [OmniDiskSweeper](https://omnidisksweeper.en.softonic.com/)
    - [Docker Toolbox](https://www.docker.com/products/docker-toolbox)
    - [Docker Version Manager](https://howtowhale.github.io/dvm/install.html)
    - [BetterTouchTool](https://www.boastr.net/)
    - [Charles](https://www.charlesproxy.com/)
    - [Chinese Lunar Calendar/WanNianLi](http://calendar.zfdang.com/) | [顺历](https://www.zhihu.com/question/35251600)
    - [GeekTool](https://www.tynsoe.org/v2/geektool/)
    - [f.lux](https://justgetflux.com/news/pages/macquickstart/) - brightness tool
    - [Karabiner](https://pqrs.org/osx/karabiner/)
    - [MacPorts](https://guide.macports.org/#installing)
    - [Omni* Apps](https://www.omnigroup.com/more)
    - [Onyx](http://www.titanium.free.fr/onyx.html) - Titanium system unitils
    - [OSQuery](https://osquery.io/) |
      [schema](https://osquery.io/schema/current) |
      [github](https://github.com/osquery/osquery)
    - [QuickLook for Webp](https://github.com/emin/WebPQuickLook)
    - [RealVNC](https://www.realvnc.com/)
    - [Screen Mirroring](https://letsview.com/)
    - [Slack](https://slack.com/) and ~[HipChat](http://hipchat.com)~
    - [Snip](http://snip.qq.com/)
    - [Snort](https://michaelok.tumblr.com/post/1095392081/how-to-install-snort-on-mac-os-x)
    - [Spectacle](https://www.spectacleapp.com/) - keyboard shortcuts
    - [Syncios Manager](https://www.syncios.com/download.html)
    - [tmate](https://tmate.io/): `brew install tmate`
    - [Tunnelblick](https://tunnelblick.net/)
    - [Unarchiver](http://wakaba.c3.cx/s/apps/unarchiver.html)
    - Oracle [VirtualBox](https://www.virtualbox.org/)
    - [VMware Fusion for Mac](https://www.vmware.com/products/fusion.html)
    - [Wine](https://www.winehq.org/)
      - [PlayOnMac](https://www.playonmac.com)
      - [Porting Kit](http://portingkit.com/)
      - [WineBottler](http://winebottler.kronenberg.org/)
      - [WineSkin](http://wineskin.urgesoftware.com/)
    - [WireShark](https://www.wireshark.org/download.html)
    - Xcode (from App Store)
      - Xcode Command Line Developer Tools
      - see http://railsapps.github.io/xcode-command-line-tools.html
    - [XtraFinder](https://www.trankynam.com/xtrafinder/)
      - hold `Cmd+R` on boot up then open Terminal
      - run `csrutil enable --without debug`
    - See https://github.com/alebcay/awesome-shell

  * Browsers
    - Chrome: Cast, Currently, DHC, Dictionary, Firebug, LastPass, Markdown, ScrollMaps, Vimium, Exif
    - Firefox: Firebug, Google Code Wiki Viewer, Dictionary, Flash Video Downloader, LastPass, Markdown, Poster
    - Opera: Dictionary, LastPass

  * Dictionary
    - 翻译软件[Bob](https://github.com/ripperhe/Bob)
    - http://diary.taskinghouse.com/posts/383137-mac-built-in-dictionary-install-traditional-chinese-dictionary
    - http://blogger.gtwang.org/2013/03/mac-os-x-dicttionary-add-chinese.html
    - Rhyme [github](https://github.com/shaunplee/homebrew-rhyme)

      ```
      brew tap shaunplee/rhyme
      brew install rhyme
      rhyme test
      ```

  * Diff tools
    - [DiffMerge](https://sourcegear.com/diffmerge/)
    - [Kdiff3](http://kdiff3.sourceforge.net/)
  * Developer tools
    * Atom and plugins
      **NOTE**: copy `~/.atom` (except `~/.atom/.gitignore`) for migration between computers
      - atom-beautify, linter, sort-lines, tabs-to-spaces
      - convert-to-utf8, file-icons, file-type-icons
      - find-and-replace, fuzzy-finder
      - grammar-selector, markdown-preview
      - highlight-selected, minimap,
      - open-recent, simple-drag-drop-text
      - git-diff, git-plus, git-blame
      - go plugins, concourse-vis
      - nuclide, node-debugger
      - status-bar, whitespace, wrap-guide
      - tree-view
    * [Eclipse](https://eclipse.org/downloads/)
      - [Pydev and Extension](http://pydev.org/updates)
    * [Docker](https://www.docker.com/products/docker-toolbox)
    * [Nuclide.io](http://nuclide.io/docs/quick-start/getting-started/)
    * Microsoft Visual Studio
    * Microsoft VS Code [VSCode](https://www.visualstudio.com/vs/visual-studio-mac/) + go + react
    * Concourse/VirtualBox/Vagrant
    * [GitHub Desktop](https://desktop.github.com/)
    * [Graphviz](http://www.graphviz.org/download/)
    * ipython
    * [Java](http://www.oracle.com/technetwork/java)
    * Java decompilers
      - http://jd.benow.ca/
      - http://www.brouhaha.com/~eric/software/mocha/
      - https://github.com/Storyyeller/Krakatau
      - http://www.benf.org/other/cfr/
      - http://www.neshkov.com/
    * [jq](https://stedolan.github.io/jq/download/)
    * [SourceTree](https://www.sourcetreeapp.com/)
    * [GitHub Desktop](https://desktop.github.com/)
    * [IntelliJ IDEA](https://www.jetbrains.com/idea/)
    * SQL Schema
      - [DbSchema](http://www.dbschema.com)
      - [MySQLWorkbenchm](https://www.mysql.com/products/workbench/)
      - [Oracle Developer Data Modeler](http://www.oracle.com/technetwork/developer-tools/datamodeler/)
      - [SQL Power Architect](http://software.sqlpower.ca/page/architect_download_os)
      - [SQL DBM](https://sqldbm.com)
    * [Xcode](https://developer.apple.com/xcode/)
    * [MySQL](http://dev.mysql.com/downloads/mysql/)
    * [MySQLWorkbench](https://dev.mysql.com/downloads/workbench/)
    * [PostgreSQL](https://www.postgresql.org/download/)
      - For [Homebrew on Mac OS](https://www.moncefbelyamani.com/how-to-install-postgresql-on-a-mac-with-homebrew-and-lunchy/)

        ```
        brew update
        brew doctor
        brew install postgresql
        ```
        Note: add, e.g. `/Library/PostgreSQL/10/bin` to `$PATH` after installation.
      - For [Ubuntu](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04)

        ```
        sudo apt-get update
        sudo apt-get install postgresql postgresql-contrib
        ```
      * PostgreSQL design/modeling tools comparison
        - http://www.databaseanswers.org/modelling_tools.htm
        - https://wiki.postgresql.org/wiki/GUI_Database_Design_Tools
        - https://wiki.postgresql.org/wiki/Design_Tools

      * Cross-platform tools selection
        - Free: [DBeaver](https://dbeaver.io/)
        - Free: [PSequel](http://psequel.com) - similar to [Sequel Pro](http://www.sequelpro.com)
        - Licensed: [DBSchema](https://www.dbschema.com/index.html)
        - Browser-based: [pgAdmin](https://www.pgadmin.org/download/)
        - IDE: [DataGrip](https://www.jetbrains.com/datagrip/)
    * [Sequel Pro](http://www.sequelpro.com/)
  * Hex Editors
    - [010 Editor](http://www.sweetscape.com/010editor/)
    - [Hex Fiend](http://ridiculousfish.com/hexfiend/)
      - iHex on App store
      - see https://en.wikipedia.org/wiki/Comparison_of_hex_editors
    - [HxD](https://sourceforge.net/projects/osxhxd/)
    - [wxHexEditor](http://www.wxhexeditor.org/download.php)
    - [wxMEditor](https://wxmedit.github.io/)
    - Xcode: Open file with Xcode, press [`Command + Shift + J`], right click file name in left pane, `Open as` Hex
    - vim:
      ```
      :%!xxd # transform buffer to hex
      :%!xxd -r # reverse transformation
      ```

  * Kubernetes
    - Install a Hypervisor
      - macOS: [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
        or [VMware Fusion](https://www.vmware.com/products/fusion),
        or [HyperKit](https://github.com/moby/hyperkit).
      - Linux: [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
        or [KVM](http://www.linux-kvm.org/).
      - **Note**: Minikube also supports a `--vm-driver=none` option that runs the Kubernetes components on the host and not in a VM. Using this driver requires Docker, but not a hypervisor.

    - Install Docker
    - Install [Docker Version Manager](https://howtowhale.github.io/dvm/)

      ```
      brew install kubernetes-cli
      kubectl version
      # using `brew` on Mac OS X
      brew update && brew install dvm
      # or directly without Homebrew
      curl -sL https://howtowhale.github.io/dvm/downloads/latest/install.sh | sh
      ```

    - Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

      ```
      sudo apt-get update && sudo apt-get install -y apt-transport-https
      curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
      echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | \
      sudo tee -a /etc/apt/sources.list.d/kubernetes.list
      sudo apt-get update
      sudo apt-get install -y kubectl
      brew install kubernetes-cli
      kubectl version
      ```

    - Uninstall minikube (Mac OS X)

      ```
      minikube stop
      minikube delete
      minikube delete && rm -rf ~/.minikube  # delete old minikube
      brew cask uninstall --force minikube
      ```

    - Install [MiniKube](https://github.com/kubernetes/minikube/releases/latest)

      ```
      brew cask install minikube  # on Mac OS X
      brew update && brew cask reinstall minikube  # upgrade to the latest

      (
      export MINIKUBE_VER=v0.28.2  # or latest
      export MINIKUBE_BIN="minikube-$(uname|tr [:upper:] [:lower:])-amd64"
      export MINIKUBE_URL=https://storage.googleapis.com/minikube/releases
      curl -Lo minikube ${MINIKUBE_URL}/${MINIKUBE_VER}/${MINIKUBE_BIN} \
        && chmod +x minikube && sudo cp minikube /usr/local/bin/ && rm minikube
      )
      ```

    - Start

      ```
      # minikube config set WantUpdateNotification false
      minikube version
      minikube ip  # this should return an IP address
      minikube start
      minikube dashboard
      minikube ssh 'docker ps -a'  # check current docker process
      minikube logs -f
      ```

  * Microsoft Office
    - See [msoffice-for-mac](./msoffice-for-mac.md)

  * PDF Editor/Reader
    - [Foxit PDF Reader](https://www.foxit.com/pdf-reader/)

          ```
          brew install --cask foxitreader
          ```

    - [LibreOffice](https://www.libreoffice.org/download/download-libreoffice/?type=mac-x86_64)
    - [PDF Reader Pro: Convert & Edit](https://apps.apple.com/us/app/pdf-reader-pro-convert-edit/id919472673?mt=12)

      ```
      brew install --cask pdf-reader-pro
      ```

    - [PDFescape](https://www.pdfescape.com/)
    - [PDFgear](https://www.pdfgear.com/)

  * Photo Tools
    - [DigiKam](https://www.digikam.org/download)

      ```
      # Fixing "Finder does not have permission" issue
      for s in digikam showfoto; do \
      sudo codesign --remove-signature /Applications/digiKam.org/$s.app; \
      sudo codesign --force --deep --sign - /Applications/digiKam.org/$s.app; \
      done
      ```

    - [DxO OpticsPro 10](http://www.dxo.com)
    - [ExifTool](http://www.sno.phy.queensu.ca/~phil/exiftool/install.html) ★★★★
    - [GIMP](https://www.gimp.org/downloads/) ★★★
    - [Pixlr](https://pixlr.com/desktop)
    - [QuickLook for Webp](https://github.com/emin/WebPQuickLook)
    - Raw editors: [darktable](https://www.darktable.org/) | [RawTherapee](https://rawtherapee.com/)
    - [XnViewMP](http://www.xnview.com/) ★★★★
  * Movie Editors
    - http://filmora.wondershare.com/video-editor/free-video-editing-software-mac.html
    - http://www.makeuseof.com/tag/top-6-free-video-editors-mac-os/
    - [DaVinci Resolve](https://www.blackmagicdesign.com/products/davinciresolve)
    - [Kdenlive](https://kdenlive.org)
    - [OpenShot](https://www.openshot.org/)
    - [Shotcut](https://www.shotcut.org/)
  * Mutlimedia/Media Players
    - [5K Player](http://www.5kplayer.com/)
    - [4K Video Downloader 4.2](https://www.4kdownload.com/products/product-videodownloader)
      - Serial Keys: XS3W4CD4VF5GB8HU / NJKM9JN8UHYB7GTV / F5RCD4X3XAQSW4CD / E5FRV6GTBH6YN8JU
      - License Keys: MI0KMPOK0I9NJUBHV / G5FC5DX4ESW3AXWS
      - Activation Keys: DCF5VG68B9HP0KMI / N9B7Y6VT5CR4X3HS
    - [AviDemux](http://avidemux.sourceforge.net/) ★★★★
    - [Audacity](https://www.audacityteam.org/download/)
    - [foobar2000](https://www.foobar2000.org/)
    - [HandBrake](https://handbrake.fr/downloads.php) - video transcoder ★★★
    - [Shutter Encoder](https://www.shutterencoder.com/)
    - [Kid3](https://kid3.sourceforge.io/) ★★★★★
    - [XLD](https://sourceforge.net/projects/xld/) ★★★★★
    - [Calibre](https://calibre-ebook.com/) ★★★★★
    - [Sigil](https://github.com/Sigil-Ebook/Sigil) ★★★★

  <br/><a name="netdrive"></a>
  * 雲盘、网盘应用
    - [阿里](https://www.aliyundrive.com)
    - [百度](https://pan.baidu.com)
    - [城通](https://ctfile.com/get_mac)
    - [苹果iCloud](https://www.icloud.com)
    - [谷歌](https://drive.google.com) | [Google Drive](https://workspace.google.com/products/drive)
    - [微软](https://onedrive.live.com)
    - [MEGA cloud storage](https://mega.io)
    - [FlowUs](https://flowus.cn/product)
    - [坚果雲](https://www.jianguoyun.com/)
    - [巴别鸟](http://www.babel.cc)
    - [迅雷](https://pan.xunlei.com/)
    - [夸克](https://pan.quark.cn/)
    - [天翼](https://cloud.189.cn)

  * Weather
    - [Meteorologist](http://macappstore.org/meteorologist/)
    - WeatherBug (in app store)


<br/><a name="software-install-ubuntu"></a>
## Software Install (Ubuntu)

<br/><a name="command-lines"></a>
### Command lines

  - basic info

    ```
    cat /proc/cpuinfo # for processor info
    cat /proc/meminfo # for RAM status
    ```

  - check installed packages

    ```
    dpkg -l
    ```

  - check opening port and pid:

    ```
    netstat -lnp  # run `kill -9` with the pid to close the port
    ```

  - check server mem: ```free -mt```
  - check disk usage: ```df -h```

  - check system services

    ```
    service --status-all
    sudo systemctl list-unit-files
    ```

  - check system monitor

    ```
    gnome-system-monitor
    ```

  - convert to utf-8:

    ```bash
    for file in *; do
      iconv -f gbk -t utf-8 "$file" -o "utf8/${file}"
    done
    iconv -f iso-8859-1 -t utf-8 < ${file} > ${file_new}
    ```
    see also [ConvertFilesToUtf8.js](https://www.ultraedit.com/resources/scripts/ConvertFilesToUtf8.js)

  - docker post-insatll
    - see https://docs.docker.com/engine/installation/linux/linux-postinstall/

      ```
      unset DOCKER_HOST
      ls -l /var/run/docker.sock
      sudo systemctl enable docker
      sudo gpasswd -a ${USER} docker
      # check if ${USER} is added to docker group
      cat /etc/group | grep ^docker
      groups # should contain `docker`
      # restart docker or have to restart the system
      sudo service docker restart  
      ```

  - download website:

    ```bash
    wget --mirror –w 2 –p --HTML-extension –-convert-links –P ${website}
    ```
    other options
    * [Pavuk](http://pavuk.sourceforge.net/)
    * [SiteSucker](https://ricks-apps.com/osx/sitesucker/index.html) (free before v3)
    * [A1](https://www.microsystools.com/products/website-download/)

  - git

    ```
    # push only to current remote/origin branch
    git config --global push.default simple

    # set git to use the credential memory cache (default 15-minute)
    git config --global credential.helper cache
    # set the cache to timeout after 1 hour (setting is in seconds)
    git config --global credential.helper 'cache --timeout=3600'
    git config --global credential.https://github.com.username jzhuyx

    # store credential in keyring
    sudo apt-get install libgnome-keyring-dev
    sudo make --directory=/usr/share/doc/git/contrib/credential/gnome-keyring
    git config --global credential.helper /usr/share/doc/git/contrib/credential/gnome-keyring/git-credential-gnome-keyring

    # push by a different credential
    GIT_SSH_COMMAND="ssh -i ~/.ssh/github_private_key" git push

    # delete remote branch
    git push origin --delete branch-x  # delete both remote and origin/branch-x
    git branch -d branch-x  # delete the local branch-x
    git fetch --all --prune

    # rename branch
    git branch -m old-name new-name
    git push origin :old-name new-name
    git push origin -u new-name
    ```

  - json tool [jq](https://stedolan.github.io/jq/download/)

  - node from [NodeSource](https://nodejs.org/en/download/package-manager/#debian-and-ubuntu-based-linux-distributions)

    ```
    curl -sL https://deb.nodesource.com/setup_9.x | sudo -E bash -
    sudo apt-get install -y nodejs
    # then install yarn
    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list

    # alternative
    sudo apt-get install -y build-essential
    # then install yarn
    sudo apt-get update && sudo apt-get install yarn
    ```

  - networking

    ```
    dhclient -r  # release the current DHCP lease
    dhclient
    ```

  - power commands

    ```
    sudo add-apt-repository ppa:atareao/atareao
    sudo apt-get update && sudo apt-get install power-commands
    ```

    - shutdown:

      ```
      dbus-send --system --print-reply --dest="org.freedesktop.ConsoleKit" /org/freedesktop/ConsoleKit/Manager org.freedesktop.ConsoleKit.Manager.Stop;;
      ```

    - hibernate:

      ```
      dbus-send --system --print-reply --dest="org.freedesktop.UPower" /org/freedesktop/UPower org.freedesktop.UPower.Hibernate;;
      ```

    - suspend:

      ```
      dbus-send --system --print-reply --dest="org.freedesktop.UPower" /org/freedesktop/UPower org.freedesktop.UPower.Suspend;;
      ```

    - restart:

      ```
      dbus-send --system --print-reply --dest="org.freedesktop.ConsoleKit" /org/freedesktop/ConsoleKit/Manager org.freedesktop.ConsoleKit.Manager.Restart;;
      ```

    - logout:

      ```
      gnome-session-quit --logout --no-prompt
      ```

    - lock:

      ```
      gnome-screensaver-command --lock
      ```

  - remember sessions from last login: dconf Editor >> gnome >> gnome-session
  - restore desktop:

    ```
    sudo apt-get remove unity-control-center
    sudo apt-get install unity-control-center
    sudo apt-get install ubuntu-desktop
    ```

  - set screen lock in 60 seconds

    ```
    gsettings set org.gnome.desktop.session idle-delay 60
    gsettings set org.gnome.desktop.screensaver lock-enabled true
    # disable screen lock
    gsettings set org.gnome.desktop.screensaver ubuntu-lock-on-suspend false
    gsettings set org.gnome.desktop.lockdown disable-lock-screen true
    xset s off # disable screen saver ?

    ```

  - kvm

    ```
    sudo apt install qemu-kvm libvirt-bin
    sudo virsh edit ubuntu16.04
    # find video section and change vgamem according to vram
    # review vm details in "Video QXL" section

    # for copy/paste, on guest vm
    sudo apt install spice-vdagent
    ```

  - python

    ```
    sudo apt-get install python-pip python-dev build-essential
    sudo pip install upgrade pip
    sudo pip install --upgrade virtualenv
    ```

  - vi
    - copy (yank): `yy` or `2yy` (yank 2 lines)
    - cut: `dd` (cut current line), `dw` (cut current word)
    - paste: `p` (paste after cursor), `P` (capital P, paste before cursor)
    - search and replace: `:%s/old/new/g`
    - toggle Hex mode (stream vi buffer thru external program `xxd`)

      ```
      :%!xxd  # to turn off :%!xxd -r
      ```

    - see
      - http://ryanstutorials.net/linuxtutorial/cheatsheetvi.php
      - http://www.lagmonster.org/docs/vi.html
      - http://www.openvim.com/


<br/><a name="keyboard-shortcuts"></a>
### Keyboard Shortcuts

  - copy screen to clipboard: Ctrl+PrtScr
  - copy screenshot of an area to clipboard: Ctrl+Shift+PrtScr
  - copy screenshot of a window to clipboard: Ctrl+Alt+PrtScr
  - save screen to picture: PrtScr
  - save screenshot of an area to picture: Shift+PrtScr
  - save screenshot of a window to picture: Alt+PrtScr


<br/><a name="install-apps"></a>
### Install Apps

  * Atom

    ```
    sudo add-apt-repository ppa:webupd8team/atom
    sudo apt-get update
    sudo apt-get install atom
    ```

  * AWS CLI

    ```
    # brew install awscli
    pip install awscli --upgrade --user
    pip install awscli-plugin-endpoint --user # if `aws` has `ImportError`
    aws --version
    ```

  * Diff tools
    - command line: diff, colordiff, wdiff,
    - [vimdiff](http://vimdoc.sourceforge.net/htmldoc/diff.html)
    - [DiffMerge](https://sourcegear.com/diffmerge/)
    - [Diffuse](http://diffuse.sourceforge.net/)
    - [Kdiff3](http://kdiff3.sourceforge.net/)
    - [Kompare](https://www.kde.org/applications/development/kompare/)
    - [TkDiff](https://sourceforge.net/projects/tkdiff/)
    - [Meld](http://meldmerge.org/)


  * DigiKam

    ```
    sudo add-apt-repository ppa:philip5/extra
    sudo apt update
    sudo apt install digikam5
    ```

  * Docker, from package
    - download `.deb` [package](https://apt.dockerproject.org/repo/pool/main/d/docker-engine/)
    - cd to download folder to run `dpkg`:

      ```
      sudo dpkg -i /path/to/package.deb
      ```

  * Docker - see [doc](https://docs.docker.com/engine/installation/linux/)

    - prerequisites, with extra packages:

    ```
    sudo apt-get update
    sudo apt-get install curl \
      linux-image-extra-$(uname -r) \
      linux-image-extra-virtual
    ```
    - install packages to allow apt to use a repository over https:

    ```
    sudo apt-get install apt-transport-https ca-certificates
    ```
    - add Docker official GPG key:

    ```
    curl -fsSL https://yum.dockerproject.org/gpg | sudo apt-key add -

    apt-key fingerprint 58118E89F3A912897C070ADBF76221572C52609D
    ```
    verify key id:

    ```
    pub   4096R/2C52609D 2015-07-14
          Key fingerprint = 5811 8E89 F3A9 1289 7C07  0ADB F762 2157 2C52 609D
    uid        Docker Release Tool (releasedocker) <docker@docker.com>
    ```

    - setup stable repository (with `testing` after `main` to enable testing/unstable/non-production repository):

    ```
    sudo add-apt-repository \
       "deb https://apt.dockerproject.org/repo/ \
       ubuntu-$(lsb_release -cs) \
       main"
    ```

    - check docker versions:

    ```
    apt-cache madison docker-engine
    ```

    - install the latest version docker:

    ```
    sudo apt-get update
    sudo apt-get -y install docker-engine

    # or with specific version:
    sudo apt-get -y install docker-engine=$(apt-cache madison docker-engine|head -n 1|awk '{print $3}')
    ```

    - verify the install:

    ```
    sudo docker run hello-world
    ```

  * ExifTool

    ```
    # see http://www.sno.phy.queensu.ca/~phil/exiftool/install.html#Unix
    # download from http://www.sno.phy.queensu.ca/~phil/exiftool/index.html
    # cd to download folder
    gzip -dc Image-ExifTool-10.40.tar.gz | tar -xf -
    cd Image-ExifTool-10.40
    perl Makefile.PL
    # optionally to run `make test`
    sudo make install
    # now `exiftool` should be in `/usr/local/bin`
    # using `perldoc` or `man` to consult ExifTool documentation
    man exiftool # same as `perldoc exiftool`
    man Image::ExifTool
    man Image::ExifTool::TagNames
    ```

  * Eclipse

  * Git GUI Tools
    - [Giggle](https://wiki.gnome.org/Apps/giggle/)
      ```
      apt-get install giggle
      ```

    - Gitg

      ```
      sudo apt-get install gitg
      ```

    - [GitEye](http://www.collab.net/products/giteye) free registered account

    - [GitKraken](https://www.gitkraken.com/) - personal free

    - [Git-Cola](https://git-cola.github.io)

    - [RabbitVCS](http://rabbitvcs.org)

    - [SmartGit](https://www.syntevo.com/smartgit/)

    - Diff/Merge: KDiff3 (in Software Center)

  * Go

    ```
    sudo apt-get update
    sudo apt-get -y upgrade
    sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
    sudo apt install golang-go
    sudo apt-get install golang-1.7-go
    # find download link at https://golang.org/dl/
    wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
    sudo tar -xvf go$VERSION*.$OS*-$ARCH*.tar.gz
    sudo mv go /usr/local
    # add the following to ~/.bashrc
    export GOROOT="${GOROOT:-$(type -p go|xargs greadlink -f 2>/dev/null|xargs dirname|xargs dirname)}"
    export GOPATH=${HOME}/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
    # test go
    go version
    go env
    ```

  * Java 8

    ```
    sudo apt-get install default-jre
    sudo apt-get install default-jdk
    ```
    or

    ```
    sudo add-apt-repository ppa:webupd8team/Java
    sudo apt-get update
    sudo apt-get install oracle-java8-installer
    sudo update-alternatives --config java
    ```

  * JetBrains IntelliJ IDEA

    ```
    sudo apt install ubuntu-make
    sudo add-apt-repository ppa:ubuntu-desktop/ubuntu-make  
    sudo apt-get update
    sudo apt-get install ubuntu-make
    umake ide idea
    ```
    or

    ```
    # download IDEA from www.jetbrains.com/idea/download/
    # extract ideaIC-*.tar.gz using
    tar -zxvf ideaIC-*.tar.gz
    # run idea.sh in bin directory inside the extracted folder
    # to create command-line runner, Tools > Create Command-line Launcher
    # to create a desktop entry, Tools > Create Desktop Entry
    ```
    or (with more manual steps)

    ```
    # extract ideaIC-*.tar.gz using
    tar -zxvf ideaIC-*.tar.gz
    sudo -i
    # move the extracted folder to /opt/idea
    mv ideaIC-* /opt/idea
    # create a desktop file:
    cat <<EOF >idea.desktop
    [Desktop Entry]
    Name=IntelliJ IDEA
    Type=Application
    Exec=idea.sh
    Terminal=false
    Icon=idea
    Comment=Integrated Development Environment
    NoDisplay=false
    Categories=Development;IDE;
    Name[en]=IntelliJ IDEA
    EOF
    # install the desktop file in the unity:
    desktop-file-install idea.desktop
    # create a symlink in /usr/local/bin using
    cd /usr/local/bin
    ln -s /opt/idea/bin/idea.sh
    # add idea icon in dash:
    cp /opt/idea/bin/idea.png /usr/share/pixmaps/idea.png
    ```
    For license see: http://idea.lanyus.com/

  * Kid3 (audio file metadata editor)

    ```
    sudo add-apt-repository ppa:ufleisch/kid3
    sudo apt update
    sudo apt install kid3
    sudo apt install kid3-qt
    sudo apt install kid3-cli
    sudo apt remove kid3 kid3-qt kid3-cli && sudo apt autoremove
    ```

  * Snort (https://www.snort.org/documents)

  * tmate (https://tmate.io/)

    ```
    sudo apt-get install software-properties-common && \
    sudo add-apt-repository ppa:tmate.io/archive    && \
    sudo apt-get update                             && \
    sudo apt-get install tmate
    ```

  * Ubuntu Software Center

    ```
    sudo add-apt-repository ppa:ubuntu-desktop/gnome-software
    sudo apt-get update
    sudo apt-get install gnome-software packagekit
    sudo apt remove software-center
    sudo apt install ubuntu-software
    ```

  * WireShark

    ```
    sudo apt-get install -y wireshark
    sudo groupadd wireshark
    sudo usermod -a -G wireshark $USER
    sudo chgrp wireshark /usr/bin/dumpcap
    sudo chmod 750 /usr/bin/dumpcap
    sudo setcap cap_net_raw,cap_net_admin=eip /usr/bin/dumpcap
    sudo getcap /usr/bin/dumpcap

    # it is required to log out and restart after the install

    sudo dpkg-reconfigure wireshark-common
    ```

  * Thunderbird

    ```
    apt-cache search thunderbird
    apt-cache policy thunderbird
    sudo apt-get upgrade
    sudo apt-get install thunderbird
    # add ExQuilla in Addons
    ```

  * System
    - BleachBit
    - CompizConfig
    - dconf Editor
    - Tweek Tool / Unity Tweek Tool
    - System Monitor
    - Settings



<br/><a name="software-windows"></a>
## Software and Tools (Windows)

### System Information

  * [CCleaner](https://www.ccleaner.com/ccleaner/download)
  * [CCleaner/Piriform Speccy](https://www.ccleaner.com/speccy)
  * [CPU-Z](https://www.cpuid.com/softwares/cpu-z.html)
    - [HWMonitor](https://www.cpuid.com/softwares/hwmonitor.html)
    - [Perfmonitor](https://www.cpuid.com/softwares/perfmonitor-2.html)
    - [PowerMax](https://www.cpuid.com/softwares/powermax.html)
  * [HWiNFO](https://www.hwinfo.com/download/) (Windows 32/64-bit)
  * [MsInfo32](https://support.microsoft.com/en-us/help/184075/description-of-microsoft-system-information-msinfo32-exe-tool)
  * [NirSoft](https://www.nirsoft.net/utils/index.html) Utilities
  * [Others](https://www.majorgeeks.com/mg/sortdate/system_info.html)

### Dictionaries

  * [WordWeb for Windows](https://wordweb.info/free/) - support systemwide right-click
  * [TheSage English Dictionary and Thesaurus](https://www.sequencepublishing.com/1/)
  * [Perfect Dictionary](https://www.microsoft.com/en-us/p/perfect-dictionary/9wzdncrd2ljv)
  * [Dictionary](https://www.microsoft.com/en-in/p/dictionary/9wzdncrfj3zp) by Flow Simulation Ltd.
  * [Bible Dictionary](https://www.microsoft.com/en-us/p/bible-dictionary-app/9nblggh09lr5)
  * [Advanced English Dictionary](https://www.microsoft.com/en-in/p/advanced-english-dictionary/9wzdncrfhmvb)
  * [The Free Dictionary](https://www.microsoft.com/en-us/p/dictionary/9wzdncrfj1z0) from Microsoft Store
  * (Native) Search Box: e.g. "fox mean"



<br/><a name="vbox-setup"></a>
## VirtualBox Setup and Network

### Install and Setup

  ```
  brew install virtualbox
  ```

### Network

VBox's implementation of NAT only supports TCP & UDP among all transport layer protocols [[1]](https://www.virtualbox.org/manual/ch06.html#nat-limitations), so VPN via PPTP won't work. As support for bridged networking with wireless interfaces is also limited [[2]](https://www.virtualbox.org/manual/ch06.html#network_bridged). Consider the followings alternative.

**Host-only network with Internet** can be achieved with IP routing by host. Good with Windows/Ubuntu host and guests. It should work for MacOS too.

In OSX host (e.g. 192.168.8.109), enable ip routing by

	```
	sudo sysctl -w net.inet.ip.forwarding=1
	```

In Win10 guest with host-only adapter (e.g. 192.168.56.10), set default gateway to OSX (e.g. 192.168.56.1), set DNS to router.

In router (e.g. 192.168.8.1), add static route via web interface or SSH

	```
	route add -net 192.168.56.0/24 gw 192.168.8.109
	```

Read this [lab manual](http://www.cs.dartmouth.edu/~sergey/cs60/lab3/vm-networking.pdf) (courtesy of CS@Dartmouth) for setting up internet in VBox guest with Host-only Adapter and native NAT by MacOS host.



<br/><a name="online-tools"></a>
## Online tools

  * Chrome extensions:
    - Better Go Playground
    - DHC (Dynamic HTTP Client)
    - Exif Meta Viewer
    - JSON Formatter | JSON Viewer
    - LastPass (password vault management)
    - World Clocks

  * Coding interview tools, see [interview](../interview/README.md#resources)
    - [C++ Shell](http://cpp.sh/)
    - [CodeBunk](https://codebunk.com/) (online interview tool)
    - [CodeChef](http://codechef.com) (judge code by other engineers)
    - [CodePad](https://codepad.remoteinterview.io) (remote interview supports golang)
    - [CodeShare](https://codeshare.io/) (sharing code in real time with others)
    - [CodeTable](https://code.hackerearth.com) (online code compiler and runner)
    - [Coding Ground](https://www.tutorialspoint.com/codingground.htm) (dev environments)
    - [Codvia](https://www.codiva.io) (online compiler and IDE for C/C++, Java)
    - [CollabEdit](http://collabedit.com/) (online coding interview, free version of https://codinghire.com)
    - [Compiler Explorer](https://godbolt.org/#)
    - [IDE One](https://ideone.com/) (online compiler and debugging tool)
    - [JDoodle](https://www.jdoodle.com/) (online compilers and editors)
    - [Js.do](https://js.do/) (online javascript editor)
    - [Go Playground](https://play.golang.org/)
    - [Oracle Live SQL](https://livesql.oracle.com/apex/livesql/file/index.html)
    - [OnlineGDB](https://www.onlinegdb.com/) (online C++ compiler)
    - [Paiza](https://paiza.io) (online PHP editor)
    - [PHP Fiddle](http://phpfiddle.org/)
    - [Pramp](https://www.pramp.com) (practicing interviews)
    - [RemoteInterview](https://www.remoteinterview.io/)
    - [RexTester](http://rextester.com/) (online C# compiler)
    - [Repl.it](https://repl.it)

  * Developer tools
    - https://www.browserling.com/tools
    - https://code.tutsplus.com/articles/20-tools-to-make-the-life-of-a-web-developer-easier--net-5338
    - http://www.commandlinefu.com/ (all about shell commands)
    - http://www.schrockguide.net/online-tools.html
    - https://www.tools4noobs.com/online_tools/
    - http://www.tutorialspoint.com/online_dev_tools.htm
    - http://tutorialzine.com/2014/09/50-awesome-tools-and-resources-for-web-developers/
    - https://vim.rtorr.com/ (vim cheatsheet)

  * Decode/Encode Base64
    - https://www.base64decode.org/
    - http://www.url-encode-decode.com/base64-encode-decode/
    - http://base64decode.net/
    - http://www.freeformatter.com/base64-encoder.html
    - http://www.httputility.net/base64-encode-decode.aspx
    - https://opinionatedgeek.com/Codecs/Base64Decoder
    - http://ostermiller.org/calc/encode.html
  * Decode/Encode URL
    - http://meyerweb.com/eric/tools/dencoder/

  * Download icons: [flaticon.com](http://www.flaticon.com/)

  * GIF/PNG animation tools: [ezgif.com](https://ezgif.com/)

  * JavaScript
    - [Codepen](https://codepen.io/)
    - [CodeSandbox](https://codesandbox.io/)
    - [ES6 Katas](http://es6katas.org/)
    - [JS.do - Online JavaScript Editor](https://js.do/)
    - [JSFiddle](https://jsfiddle.net/)
    - [Plunker](https://plnkr.co/) - Online Snippet Previewer
    - [JSBin](https://jsbin.com/) - Collaborative JS Editor/Debugger
    - [JavaScript Minifier](http://www.danstools.com/javascript-minify/)
    - [Javascript Obfuscator](http://www.danstools.com/javascript-obfuscate/)
    - [JsBeautifier](http://jsbeautifier.org/)
    - [JsPerf](http://jsperf.com/) - JavaScript performance playground
    - [JsLint](http://www.jslint.com/) - JavaScript Code Quality Tool
    - [JsHint](http://jshint.com/) - JavaScript Error Detector
    - [Stackblitz - Online VS Code IDE for Angular & React](https://stackblitz.com/)
    - [Vueditor - WYSIWYG Editor For Vue.js](http://www.vuescript.com/wysiwyg-editor-vue-js-vueditor/)
    - [VueJS Editor](https://www.tutorialspoint.com/online_vuejs_editor.php)
    - [more ...](http://blog.liveedu.tv/10-best-online-javascript-editors/)

  * JSON Formatters
    - http://codebeautify.org/jsonviewer
    - https://jsonformatter.curiousconcept.com/
    - http://www.jsoneditoronline.org/
    - http://jsonlint.com/
    - http://www.freeformatter.com/json-formatter.html
    - http://www.webtoolkitonline.com/json-formatter.html
    - http://jsonviewer.stack.hu/
    - http://jsonformatter.org/
    - http://jsonformat.com/

  * FireFox add-ons: Adblock, Firebug, Flash Downloader, ImTranslator, Markdown

  * [Google Draw](https://www.draw.io/)
  * [Go Playground](https://play.golang.org/)
  * [Morse code](https://morsecode.scphillips.com/translator.html)

  * Markdown editors
    - [Dillinger](http://dillinger.io/)
    - [GitHub-Flavored](https://jbt.github.io/markdown-editor/)
    - [Haroop](http://pad.haroopress.com/user.html) - Linux/Mac/Windows
    - [MacDown](http://macdown.uranusjr.com/) - Mac OS X only
    - [MarkMyWords](https://github.com/voldyman/MarkMyWords) - Linux only
    - [Markable](https://markable.in/)
    - [Markdown converter](https://daringfireball.net/projects/markdown/dingus)
    - [MEditor](https://pandao.github.io/editor.md/en.html)
    - [Remarkable](https://remarkableapp.github.io/) - Linux/Windows
    - [StackEdit](https://stackedit.io/editor)
    - [Mou](http://25.io/mou/)

  * Online resources for web developers
    - [Argento Theme](http://argentotheme.com/)
    - [ExpressionEngine](https://ellislab.com/) (content management tool by EllisLab)
    - [Free Formatters](https://www.freeformatter.com)
    - [GoJs](https://gojs.net) (interactive JavaScript diagrams in HTML)
    - [Gridgum](https://gridgum.com/)
    - [HTML5 Maker](http://html5maker.com/)
    - [Icon Finder](https://www.iconfinder.com/)
    - [Icon Maker](https://iconsflow.com/)
    - [Kukook](https://www.kukook.com/) (professional resume template)
    - [MotoCMS](https://www.motocms.com/) (website builder)
    - [Template Master](https://templates-master.com/)
    - [Themify](http://themify.me/)

  * Online SVG editors
    - [Boxy SVG](https://boxy-svg.com/app)
    - [DrawSVG](http://www.drawsvg.org/#tools)
    - [Janvas](https://www.janvas.com/) tools
    - [rollApp](https://www.rollapp.com/apps) tools collection
    - [SVG Editor](http://www.clker.com/inc/svgedit/svg-editor.html) | [https://github.com/SVG-Edit/svgedit](source)

  * Online syntax highlighting tools
    - [Markup Highligher](http://markup.su/highlighter/)
    - [Hohli Highligher](http://highlight.hohli.com/)
    - [Code Beautify](http://codebeautify.org/)
    - [ToHtml syntax highlighting for masses](https://tohtml.com/)
    - [Hilite](http://hilite.me/)

  * Regular expression
    - Regex 101 [regex101.com](https://regex101.com)
    - Regex Pal [regular expression tester](http://www.regexpal.com/)
    - Regex Tester [regexstorm](http://regexstorm.net/tester)
    - Regexr [regexr](http://regexr.com/)

  * [Swagger](http://editor.swagger.io/#/)


<br/>
&raquo; Back to <a href="#contents">Contents</a> &laquo;
