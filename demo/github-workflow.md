# GitHub workflow


<br/><a name="contents"></a>
## Contents

  * [Introduction](#intro)
  * [Github Workflow In Operation](#operation)
    - Configure SSH and git
    - Fork from upstream (one-time setup)
    - Clone and set upstream (one-time setup)
    - Pull Request (PR) and merging
  * [Advanced operations](#advanced)
    - Amend historical commits
    - Creating tags
  * [References](#reference)
    - Bash aliases for git
    - Git/GitHub GUI client options
    - Diff/Merge Tools
    - Links



<br/><a name="intro"></a>
## Introduction to the workflow

* Organization has a central GitHub account with private or public repositories. These hold the "gold copies" of the projects.

* The repository has two branches. The "**master**" branch is what in production and "development" is what deployed to the QA server.
  *Notes*:
  - Other long-lived branches could be "staging" server, or a "hackathon" (for experimental development).
  - These branches are not as same concept as in traditional cvs/svn/git workflows.
  - For initial setup, use **master** only to simplify the process.

* This "gold copy" repository (mainly **master** branch) is also used with
  - CI/CD tools (e.g. building, unit tests, functional tests, staging etc.)
  - any commit/deploy gate checking
  - deployment tools

* Developers/users fork the repository to their own personal GitHub account.
  *Note*: When the user left the organization, the fork will be automatically deleted with removed account/permission.

* Developer forks the repo, and clones from the fork (`git clone git@github.com:username/repository-name.git`), which makes "**origin**" point at the personal fork.

* Developer also adds an "**upstream**" (as naming convention) that points at the company repository (by `git add upstream git@github.com:organization-name/repository-name.git`)

* Developer uses (local) `master` and `origin/master` to sync with `upstream/master` and should never explicitly commit to any `master` branch.

* Developer always works on (local) branches, pushes to `origin` (the personal fork), and `pull`/`rebase` from `master`.

* Developer uses branches (on `origin`) to submit pull request to `upstream/master`

* Remember any dev using Github workflow manages three (3) repositories for one particular project:
  - Upstream ("upstream" as naming convention) represents the "gold copy" of the project (on github server). Should only be updated by pull request merging process rather than by committing/pushing directly.
  - Origin (named as "origin" by default) is the forked copy from "upstream" and should always sync its "`master`" branch to `upstream/master`.
  - Local copy is where checked out (by `git clone`) and should always sync its "`master`" branch to `origin/master`.

* Advantages
  - Follow the workflow with major open source community
  - Guard core/corp components as well as encourage innovation and adoption.
  - Organize cross department/team collaboration.
  - Reduce noise (on core/corp repo branches).
  - Reduce workload for a dedicated repo admin.
  - Reduce need to maintain multiple branches.
  - Streamline interaction with contractors.
  - Keep PR (pull request) conversations on github.com.
  - Easier with growing team.
  - Easier for CI/CD.


<br/><a name="operation"></a>
## In operation

### Configure SSH and git

  1. Configure SSH to use key for github.com

    - generate authentication key for github.com

      ```
      ssh-keygen -t rsa -b 4096 -C "github.com" -f ~/.ssh/github_key
      ```

    - add the key to github account. see
      - https://help.github.com/articles/adding-a-new-ssh-key-to-your-github-account/
      - https://help.github.com/articles/connecting-to-github-with-ssh/

    - add the following to `~/.ssh/config`

      ```
      host github.com
        HostName github.com
        PreferredAuthentications publickey,keyboard-interactive,password
        IdentityFile ~/.ssh/github_key
        IdentitiesOnly yes
        User git

      host github.com-foobar
        HostName github.com
        PreferredAuthentications publickey,keyboard-interactive,password
        IdentityFile ~/.ssh/github_key_foobar
        User foobar
      ```
      Note: For configuring more than one account for, e.g. `host github.com`,
            the `host` name is the identifier (for the same `HostName github.com`)
            so that later the `git remote` can set URL differently

  2. Configure git

    - Default [push](https://git-scm.com/docs/git-config/#git-config-pushdefault) option

      ```
      git config --global push.default simple
      ```

    - Cache git credential

      ```
      # store credential in mac os x key chain
      git config --global credential.helper osxkeychain

      # set git to use the credential memory cache (default 15-minute)
      git config --global credential.helper cache
      # set the cache to timeout after 1 hour (setting is in seconds)
      git config --global credential.helper 'cache --timeout=3600'

      # store credential in keyring
      sudo apt-get install libgnome-keyring-dev
      sudo make --directory=/usr/share/doc/git/contrib/credential/gnome-keyring
      git config --global credential.helper /usr/share/doc/git/contrib/credential/gnome-keyring/git-credential-gnome-keyring
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

    - Graph log

      ```
      git config --global alias.lg "log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative"
      ```
      The above is adding an alias to `~/.gitconfig`.

  * Troubleshot

    - Unset `GIT_SSH` and/or `GIT_SSH_COMMAND` or check if it matches

    - Test SSH over the HTTPS port: `ssh -T -p 443 git@ssh.github.com`

    - Test `ssh -T git@github.com`

    - Test ssh: `ssh -vvv git@github.com`


### Fork (one-time setup)

  1. Open and sign into github.com
  2. Open one of “Dockerian” repositories (e.g. “go-coding”)
  3. On the upper-left corner (under your profile icon, e.g. your-id), click on Fork
  4. On popup dialog, click on “@your-id” profile icon
  - Note: it is okay to click on Fork more than once, which won’t fork another time but shows existing forks.

### Clone and set upstream (one-time setup)

  1. Open a terminal/console
  2. Clone your forked repository:

    ```
    # Note: DO NOT use https. Use SSH instead.
    # git clone https://github.com/jasonzhuyx/go-coding.git
    git clone git@github.com:jasonzhuyx/go-coding.git
    cd go-coding
    ```

  3. Add and review upstream

    ```
    # git remote remove upstream >/dev/null 2>&1
    # Note: DO NOT use https. Use SSH instead.
    # git remote add upstream http://github.com/dockerian/go-coding.git
    git remote add upstream git@github.com:dockerian/go-coding.git

    git remote –v
    git fetch upstream
    ```

  **Note**: For using multiple github accounts in `~/.ssh/config` (e.g. `host github.com-foorbar` for the same `HostName github.com`), the URL must be set per the `HostName` associated with granted user (e.g. `foobar`). For example:

    ```
    git remote set-url upstream git@github.com-foobar:dockerian/go-coding.git
    ```


### Pull Request (PR)

  * <a name="sync-upstream"></a>Sync with upstream

    ```
    # git remote –v
    git fetch upstream
    git checkout master

    # reset to upstream master
    git fetch -v --all --prune
    git reset --hard upstream/master

    # DON'T forget to commit to origin (your fork)
    git push -f origin master
    git status
    ```

  * Start working on a new branch (usually a new feature or fix)

    ```
    git checkout -b dev
    # optionally with name convention, e.g.
    # git checkout –b FEA-1234  # a task or fix associated with ticket in tracking system (e.g. JIRA)
    # git checkout –b feature/FEA-1234  # for multitasking changes
    # git checkout –b fix/BUG-4321  # for multitasking changes

    # undo any individual change
    git checkout -- changed_file_name  # undo a change
    ```
    Note: Unless for multitasking changes, it's recommended to use a simple working branch name (e.g. just `dev`) rather than being specific.

  * Commit on branch (locally)

    ```
    # doing your work and commit
    git add *
    git commit –m "fixed BUG-4321 a bug in code"

    # optionally amend more changes after the commit
    # git add *
    # git commit --amend
    ```

  * Sync with upstream again (repeat this [step](#sync-upstream)) and merge with conflict

    ```
    # now sync on your origin/master
    git fetch -v --all --prune
    git checkout master
    git reset --hard upstream/master
    ```

    - Merging option (1) rebase:

      ```
      git checkout dev  # or `git checkout FEA-1234`
      git rebase master

      # fix conflict locally, and commit (again) or
      # git commit –amend
      ```

    - Merging option (2) pull:

      ```
      git checkout dev  # or `git checkout FEA-1234`
      git pull upstream master

      # fix conflict locally, and commit (again) or
      # git commit –amend

      ```

  * Push changes to origin branch (on your fork)

    ```
    git push  # optionally with `--force origin/branch`
    ```

  * Submit PR on github.com
    - open upstream repo on github.com
    - click on "Pull Request", then "New Pull Request" (green button)
    - click on "compare cross forks" link
    - provide brief description including (e.g. JIRA) ticket number
    - add comment of testing result
    - add reviewers and assignees
    - submit

  * Merge Pull Request
    - *note*: amending and/or adding commits to the same branch on `origin` are allowed before PR is merged.
    - review code with approval (e.g. `:+1`) or comments.
    - upon all reviewers approved the PR, one of assignees can merge the PR.
    - sync up local and "origin" with "upstream" after the PR is merged.
    - close the ticket/story after the PR is merged.


<br/><a name="advanced"></a>
## Advanced operations

### Amend historical commits

  * Listing previous logs

    ```
    git log --oneline -3  # or n, here n means last n commits
    ```
    assuming this produces

    ```
    268bb1f the last commit
    57688c5 the previous -1 commit
    e4b7303 the previous -2 commit
    ```
    and we want to edit message for `57688c5 ` - "the previous -1 commit"

  * Do an interactive rebase

    ```
    git rebase -i HEAD~3
    ```
    and this will bring up your editor with commits in reverse order:

    ```
    pick 268bb1f the last commit
    pick 57688c5 the previous -1 commit
    pick e4b7303 the previous -2 commit
    pick e4b7303 the previous -3 commit
    ```

  * Change the commend of `pick`, in the first column, to `e`

    ```
    pick e4b7303 the previous -3 commit
    pick e4b7303 the previous -2 commit
    e 57688c5 the previous -1 commit
    pick 268bb1f the last commit
    ```
    then save and quit (e.g. pressing ESC, `:wq` likely in `vi`)

  * Now amending the message by

    ```
    git commit --amend  # --author="Author Name <email@address.com>"
    ```
    this will bring up the editor to allow you editing the message, and then save and quit

  * Continue to complete

    ```
    git rebase --continue
    ```

  * Repeat last 2 steps if there are more than one commit to ammend


### Creating tags

  ```
  git tag -a <tag-name> <commit-sha> -m "Message for the tag"
  git push --tags origin master
  ```


<br/><a name="reference"></a>
## Reference

### Bash aliases for git (likely in your ~/.bashrc)

```
alias a="alias|cut -d' ' -f 2- "
alias gbc='git symbolic-ref --short -q HEAD'
alias gbd='git branch -d '  # delete branch locally
alias gbdo='git push origin --delete '  # delete branch on origin
alias gbv="git branch -v "
alias gco="git checkout "
alias gfv="git fetch -v --all --prune ; git branch -v"
alias glg="git log --graph --pretty=format:'%C(magenta)%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative"
alias gpum='git checkout master && git pull upstream master'
alias gpumgp='git checkout master && git pull upstream master && git push'
alias gpumgpf='git checkout master && git pull upstream master && git push -f'
alias grm='git rebase master'
alias grmgpf='git rebase master; git push --force'
alias grv='git remote -v'
alias gst='git status'

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
```

### Git/GitHub GUI client options

  * Mac OS X
    - SourceTree (https://www.sourcetreeapp.com/)
    - GitHub Desktop (https://desktop.github.com/)
    - Addons/plugins for Atom or IDE

  * Ubuntu
    - Giggle (https://wiki.gnome.org/Apps/giggle/)
    - Gitg (`sudo apt-get install gitg`)
    - GitKraken (https://www.gitkraken.com/)
    - Git-Cola (https://git-cola.github.io)

### Diff/Merge tools

  * from GUI (e.g. SourceTree)

  * opendiff

    ```
    which opendiff
    git config --global merge.tool opendiff
    git config --global diff.tool opendiff
    git config --global difftool.prompt false
    ```

  * KDiff3 (http://kdiff3.sourceforge.net/)
    - Mac OS X: `brew install kdiff3`
    - Ubuntu: from Software Center


### Links

Quick tip

  - https://www.sitepoint.com/quick-tip-synch-a-github-fork-via-the-command-line/

About pull reuest (PR)

  - https://help.github.com/articles/about-pull-requests/

Step-by-step

  - https://gist.github.com/colinsurprenant/9b081958b50cfecc210c
  - http://blog.scottlowe.org/2015/01/27/using-fork-branch-git-workflow/
  - https://github.com/sevntu-checkstyle/sevntu.checkstyle/wiki/Development-workflow-with-Git:-Fork,-Branching,-Commits,-and-Pull-Request
  - https://gist.github.com/Chaser324/ce0505fbed06b947d962

Comparing workflows

  - https://www.atlassian.com/git/tutorials/comparing-workflows
  - http://blogs.atlassian.com/2013/05/git-branching-and-forking-in-the-enterprise-why-fork/
  - https://github.community/t5/How-to-use-Git-and-GitHub/Branch-VS-Fork/td-p/10619
  - https://stackoverflow.com/questions/3611256/forking-vs-branching-in-github

Forks with feature branches

  - https://x-team.com/blog/our-git-workflow-forks-with-feature-branches/

More about Github workflow

  - https://github.com/asmeurer/git-workflow
  - http://hugogiraudel.com/2015/08/13/github-as-a-workflow/
  - https://github.com/servo/servo/wiki/Github-workflow

Triangle workflow

  - https://github.com/blog/2042-git-2-5-including-multiple-worktrees-and-triangular-workflows
