# go-coding

This is a project for [Golang](https://golang.org/) exercises.

[![Code Coverage](https://codecov.io/gh/dockerian/go-coding/branch/master/graph/badge.svg)](https://codecov.io/gh/dockerian/go-coding)
[![GoDoc](https://godoc.org/github.com/dockerian/go-coding?status.svg)](http://godoc.org/github.com/dockerian/go-coding)
[![Build Status](https://travis-ci.org/dockerian/go-coding.svg?branch=master)](https://travis-ci.org/dockerian/go-coding)
[![Go ReportCard](https://goreportcard.com/badge/github.com/dockerian/go-coding)](https://goreportcard.com/report/github.com/dockerian/go-coding)


## Project Structure

  - [Data structure solutions](ds/README.md)
  - [Dev interview: coding solutions](puzzle/README.md)
  - [Dev interview: examples](interview/README.md)
  - [Go API example](api/README.md)
  - [Go CLI example](cli) (TBD)
  - [Online coding examples](demo/README.md) |
    [Golang Notes](demo/golang-notes.md) |
    [Gotchas](demo/gotchas.md)
  - [Package solutions](pkg/README.md)
  - [Utilities](utils/README.md)
  * [About this project](#readme)
    - [How to build, test, and run](#build-test-run)
    - [End-to-end tests](#e2e)
    - [Documentation](#godoc)
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


<a name="e2e"><br/></a>
### End-to-end tests

  * MessageSender test
    - In order to run the test, inject environment variables to
      set user, pass, recipient, e.g.

      ```
      MAIL_TO='a@hotmail.com' \
      MAIL_USER='a@hotmail.com' MAIL_PASS='pass!wd' \
      make e2e
      ```
      Default mail server (`MAIL_HOST`): `smtp.office365.com`.<br/>
      Default auth: `MAIL_TLS=true` (otherwise using just `LoginAuth`)<br/>
      Default sender (`MAIL_USER`): `jason_zhuyx@hotmail.com`.<br/>
      And default recipient (`MAIL_TO`): `jason_zhuyx@hotmail.com`.<br/>


<a name="godoc"><br/></a>
### Documentation

  This project uses [godocdown](https://github.com/robertkrimen/godocdown)
  and `$(DOC_PACKAGES)` in `Makefile` to generate documentations for some library packages

  ```
  make doc
  ```



<br/><a name="learning-go"></a>
## Learning Go
- [awesome-go](https://github.com/avelino/awesome-go)
- [go books](https://github.com/dariubs/GoBooks)


<p><br/></p>

[![Code Coverage](https://codecov.io/gh/dockerian/go-coding/branch/master/graph/badge.svg)](https://codecov.io/gh/dockerian/go-coding)
[![GoDoc](https://godoc.org/github.com/dockerian/go-coding?status.svg)](http://godoc.org/github.com/dockerian/go-coding)
[![Build Status](https://travis-ci.org/dockerian/go-coding.svg?branch=master)](https://travis-ci.org/dockerian/go-coding)
[![Go ReportCard](https://goreportcard.com/badge/github.com/dockerian/go-coding)](https://goreportcard.com/report/github.com/dockerian/go-coding)
