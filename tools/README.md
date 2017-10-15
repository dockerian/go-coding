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

  See https://www.goinggo.net/2016/05/installing-go-and-your-workspace.html

  After installed `go` and set `$GOPATH` (which typically is `$HOME/go`)

  ```
  mkdir -p $GOPATH/src/github.com/dockerian
  cd $GOPATH/src/github.com/dockerian
  git clone https://github.com/dockerian/go-coding.git
  cd go-coding
  ```

  *NOTE:* This assumes you have Git installed.  If you don’t, you can find the installation instructions here: https://git-scm.com/

  Optionally create a soft link (as shortcut) in `$HOME/projects`

  ```
  ln -s $GOPATH/src/github.com/dockerian/go-coding $HOME/projects/go-coding
  cd -P $HOME/projects/go-coding
  ```

  **IMPORTANT NOTES**:
  `$GOPATH/src/github.com` may contain both `origin` (forked by, e.g. cyberdev)
  and `upstream` (the repo forked from, e.g. dockerian)
    - $GOPATH/src/github.com/dockerian/cyberint-sng-api
    - $GOPATH/src/github.com/cyberdev/cyberint-sng-api


<a name="godep"><br /></a>
### Golang dependency management

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
      glide update
      ```

  - Sync/Downloading packages per packages list to `./vendor`

      ```
      glide install
      govendor sync
      ```

  - Downloading packages to `./vendor`

      ```
      godep get
      glide get
      govendor fetch  # govendor get: to both ./vendor and $GOPATH
                      # govendor add or update: add/update from $GOPATH to ./vendor
      ```

  - Removing a package dependency

      ```
      glide remove
      govendor remove
      dep prune       # only remove unused
      ```
