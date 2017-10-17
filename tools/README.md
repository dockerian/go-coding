# Tools


This is a set of tools and scripts for `go-coding` project.

<br/><a name="contents"></a>
## Contents

  - [Dockerization](#docker)
  - [Installing Go](#install-go)
  - [Golang Dependency Management](#godep)


<a name="docker"><br /></a>
## Docker Notes

### Build and run in Docker container

**Install Docker Toolbox**  

  - [Install Docker Platform](https://www.docker.com/products/overview#/install_the_platform)
  - [Docker Toolbox](https://www.docker.com/products/docker-toolbox)

**Clean up go-coding container and image**

  ```
  docker rm -f $(docker ps -a | grep go-coding | awk '{print $1}')
  docker rmi -f $(docker images -a | grep go-coding | awk '{print $1}')
  ```

**Build Docker container**

  ```
  make build
  ```

  or

  ```
  # current path is the source root where Dockerfile exists
  docker build -t dockerian/go-coding .
  ```

**Start Docker container**

  Recommend to run inside the docker container, simply by

    ```
    make  # or `make cmd`, which starts a bash shell in the container
    ```

  or

    ```
    docker run -it --rm --name go-coding \
      -v "$PWD":/go/src/github.com/dockerian/go-coding \
      dockerian/go-coding

    ```

  Now `golang` environment is available (in the container);
  otherwise, using the hybrid script `run.sh` to call any `Makefile` target,
  default is `test` :

    ```
    ./run.sh  # inside or outside of the container
    ```


### Dockerfile ENTRYPOINT vs CMD

  - No `ENTRYPOINT`

    | CMD form             | Actual calling       |
    |:---------------------|:---------------------|
    | No `CMD`             | *error, not allowed* |
    | `CMD cmd arg`        | /bin/sh -c cmd arg   |
    | `CMD ["cmd", "arg"]` | cmd arg              |

  - Shell form `ENTRYPOINT exec param`

    | CMD form             | Actual calling                           |
    |:---------------------|:-----------------------------------------|
    | No `CMD`             | /bin/sh -c exec param                    |
    | `CMD cmd arg`        | /bin/sh -c exec param /bin/sh -c cmd arg |
    | `CMD ["cmd", "arg"]` | /bin/sh -c exec param cmd arg            |

  - Exec form: `ENTRYPOINT ["exec", "param"]`

    | CMD form             | Actual calling                |
    |:---------------------|:------------------------------|
    | No `CMD`             | exec param                    |
    | `CMD cmd arg`        | exec param /bin/sh -c cmd arg |
    | `CMD ["cmd", "arg"]` | exec param cmd arg            |



<a name="gotool"><br /></a>
## Golang Tools

<a name="install-go"></a>
### Installing Go

#### Local Installation

  - See https://www.goinggo.net/2016/05/installing-go-and-your-workspace.html
  - See https://git-scm.com/ for git installation instructions

  After installed `go`, set `$GOPATH` (which typically is `$HOME/go`)
  and optionally `$GOROOT` (usually `/usr/local/go`) in `$PATH`

    ```
    export GOPATH=$HOME/go
    export GOROOT=/usr/local/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

    mkdir -p $GOPATH/src/github.com/dockerian
    cd $GOPATH/src/github.com/dockerian
    git clone git@github.com:dockerian/go-coding.git
    cd go-coding
    ```

  For a developer with GitHub account, e.g. "cyberdev", working on fork -

    ```
    mkdir -p $GOPATH/src/github.com/dockerian
    cd $GOPATH/src/github.com/dockerian
    git clone git@github.com:cyberdev/go-coding.git
    cd go-coding
    git remote add upstream git@github.com:dockerian/go-coding.git
    git fetch --all -v
    ```

  Optionally create a soft link (as shortcut) in a project folder, e.g. `$HOME/gh`

    ```
    ln -s $GOPATH/src/github.com/dockerian/go-coding $HOME/gh/go-coding
    cd $HOME/gh/go-coding
    cd -P .  # on docker host
    ```

  **IMPORTANT NOTES**:

  * `$GOPATH/src/github.com` may contain both `origin` (forked by, e.g. cyberdev)
  and `upstream` (the repo forked from, e.g. dockerian)
    - $GOPATH/src/github.com/dockerian/go-coding
    - $GOPATH/src/github.com/cyberdev/go-coding

  However, in order to use [godep](#godep) (and some other go packages manager),
  any github fork should be cloned/checked-out to its upstream, e.g. dockerian,
  path under `$GOPATH/src/github.com`.

  * Running go test and/or build requires current path (`$PWD`) under `$GOPATH`.
    - Use `cd -P .` if the project is a soft link to the repository path
    - Add the following script, e.g. to `./.bashrc`, as a helper function

    ```
    function goto() {
      cd $(find $GOPATH/src -type d -name "$1" 2>/dev/null | head -n 1); pwd
    }
    ```


<a name="godep"></a>
## Golang Dependency Management

  Since `vendor` support in go 1.6, there are many existing 3rd-party tools
  can help to manage package dependencies. Here are some references:
  - [Package management tools](https://github.com/golang/go/wiki/PackageManagementTools)
  - [Go package manager comparison](https://github.com/Masterminds/glide/wiki/Go-Package-Manager-Comparison)

  Before [dep](https://github.com/golang/dep) becomes official, here are
  a few top options:
  - [godep](https://github.com/tools/godep)
  - [glide](https://github.com/Masterminds/glide), see [also](https://www.meta.sc/tech/govendoring/)
  - [govendor](https://github.com/kardianos/govendor)
  - [gpm](https://github.com/pote/gpm)

  Using package management in development can have following situations:

  - Creating `./vendor` to save packages (by parsing project `import`'s) with
    a packages list (metadata info)

      ```
      dep init      # creating ./Gopkg.toml, ./Gopkg.lock, and ./vendor
      godep save    # creating ./Godeps/Godeps.json and ./vendor
      glide init    # creating ./glide.yaml and ./vendor
      govendor init # creating ./vendor/vendor.json
      govendor add +external
      ```

  - Restoring packages from packages list info and `./vendor` to `$GOPATH` (not recommended)

      ```
      godep restore
      ```

  - Checking/comparing package status and versions between `./vendor` and `$GOPATH`

      ```
      dep status  # -dot requiring http://www.graphviz.org/
      glide list
      glide tree  # deprecated but nice to show dependencies in a tree view
      ```

  - Updating `./vendor` package(s) as well as packages list to the latest

      ```
      dep ensure -update
      godep update
      glide update  # update to glide.lock
      govendor update +external
      ```

  - Sync/Downloading packages per packages list to `./vendor`

      ```
      dep ensure
      glide install
      govendor sync +external
      ```

  - Downloading packages to `./vendor`

      ```
      godep get
      glide get
      govendor fetch  # +external: updating to ./vendor only
                      # govendor get: to both ./vendor and $GOPATH to ./vendor
                      # govendor add or update: add/update from $GOPATH
      ```

  - Removing a package dependency

      ```
      glide remove
      govendor remove
      dep prune       # only remove unused
      ```
