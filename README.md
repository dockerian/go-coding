# go-coding

This is a project for [Golang](https://golang.org/) exercises.


## Project Structure

  - [Data structure solutions](ds)
  - [Dev interview: coding solutions](puzzle)
  - [Dev interview: examples](interview)
  - [Go API example](api)
  - [Go CLI example](cli) (TBD)
  - [Online coding examples](demo) | [Golang Notes](demo/golang-notes.md)
  - [Package solutions](pkg)
  - [Utilities](utils)
  - [How to build, test, and run](#build-test-run)
  - [Learning Go](#learning-go)


<a name="readme"><br/></a>
## Introduction


<a name="using-docker"><br/></a>
### Using Docker

  Installing `Go` may not be needed if you choose to use [Docker](#docker). With running a go-coding container, you can clone this repo at any location on your disk, for example `$HOME/projects`, without having to set ```$GOPATH```. And you can still access (e.g. for editing) the source code locally.

    ```
    # assume in your projects folder
    cd $HOME/projects
    git clone https://github.com/dockerian/go-coding.git
    cd go-coding
    ```

To build and run in docker container, see [here](tools/README.md#docker).



<a name="build-test-run"><br/></a>
### Build, test and run

The `Makefile` has included `build`, `test`, `run` targets. For example, to build, simply change to the project directory and run

  ```
  make build # or ./build.sh
  ```

or to run tests

  ```
  make test  # or ./run.sh test
  ```




<br/><a name="learning-go"></a>
### Learning Go
- [awesome-go](https://github.com/avelino/awesome-go)
- [go books](https://github.com/dariubs/GoBooks)
