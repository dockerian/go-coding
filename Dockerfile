ARG GO_VERSION=1.12
ARG ALPINE_VERSION=3.10

# NOTE: `ARG`s are reset after `FROM`
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as dev

ARG PROJECT=go-coding
ARG GITHUB_REPO=github.com/dockerian/go-coding
ARG DOCKER_NAME=go-coding
ARG DOCKER_IMAG=dockerian/$DOCKER_NAME
ARG BINARY=go-coding

MAINTAINER Jason Zhu <jason.zhuyx@gmail.com>
LABEL maintainer="jason.zhuyx@gmail.com"
LABEL organization="Dockerian Seattle"
LABEL project="Golang Practice"

RUN apk update \
 && apk upgrade \
 && apk add --no-cache --virtual .build-deps \
    bash \
    ca-certificates \
    dpkg \
    gcc \
    git \
    jq \
    make \
    musl-dev \
    net-tools \
    openssh \
    tree \
    tar \
    zip \
 && rm -rf /var/lib/apt/lists/* \
 && rm /bin/sh && ln -sf /bin/bash /bin/sh \
 && echo "export PS1='\n\u@\h \w [\#]:\n\$ ' " >> ~/.bashrc \
 && echo "alias ll='ls -al'" >> ~/.bashrc \
 && echo "" >> ~/.bashrc

# install gosu
# RUN gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys \
#     B42F6819007F00F88E364FD4036A9C25BF357DD4 \
#  && curl -o /usr/local/bin/gosu -SL \
#    "https://github.com/tianon/gosu/releases/download/1.4/gosu-$(dpkg --print-architecture)" \
#  && curl -o /usr/local/bin/gosu.asc -SL \
#    "https://github.com/tianon/gosu/releases/download/1.4/gosu-$(dpkg --print-architecture).asc" \
#  && gpg --verify /usr/local/bin/gosu.asc \
#  && chmod +x /usr/local/bin/gosu \
#  && rm /usr/local/bin/gosu.asc

# install gosu for a better su+exec command
ARG GOSU_VERSION=1.10
RUN dpkgArch="$(dpkg --print-architecture | awk -F- '{ print $NF }')" \
 && wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-$dpkgArch" \
 && chmod +x /usr/local/bin/gosu \
 && gosu nobody true

# install golang lib
RUN go get -u golang.org/x/lint/golint \
 && go get -u github.com/robertkrimen/godocdown/godocdown \
 && go get -u github.com/golang/dep/cmd/dep \
 && go get -u github.com/tools/godep

#
# downloading the latest $GITHUB_REPO source code so that it allows to
# run the container without mapping to any local $GITHUB_REPO copy
# e.g.
#       docker build -t dockerian/go-coding .
#       docker run --rm -it dockerian/go-coding
#
ENV HOME=/root \
    GOPATH=/go \
    PROJECT=$PROJECT \
    PROJECT_DIR=/go/src/$GITHUB_REPO \
    SHELL=/bin/bash

# creating "$PROJECT_DIR" and adding source code
ADD . "$PROJECT_DIR"

# env variable has content of SSH key retrieved by `ssh -v git@github.com`
ARG GITHUB_PRIVATE_KEY

# CAUTION: NOT use for production - the github key may be imported.
# see https://bit.ly/2oY3pCn
# RUN cd -P "$PROJECT_DIR" \
#  && git config --global url."git@bitbucket.org:".insteadOf https://bitbucket.org/ \
#  && git config --global url."git@github.com:".insteadOf https://github.com/ \
#  && mkdir -p $HOME/.ssh && umask 0077 \
#  && echo "${GITHUB_PRIVATE_KEY}" > $HOME/.ssh/id_rsa \
#  && ssh-keyscan bitbucket.org >> $HOME/.ssh/known_hosts \
#  && ssh-keyscan github.com >> $HOME/.ssh/known_hosts \
#  && echo "" \
#  && echo "****************************************************" \
#  && echo "CAUTION: SSH key is imported in this layer of image." \
#  && echo "****************************************************" \
#  && ls -al $HOME/.ssh && echo "$HOME/.ssh" \
#  && echo "" \
#  && tree -L 4 $GOPATH \
#  && echo "" \
#  && echo "$PROJECT_DIR" \
#  && ls -al \
#  && echo ""

RUN cd -P "$PROJECT_DIR" \
 && GO111MODULE=on go mod tidy \
#&& tree -L 4 $GOPATH \
 && ls -al

WORKDIR $PROJECT_DIR

EXPOSE 8001/TCP 8008/TCP 8080/TCP

# this ENTRYPOINT requires gosu
# ENTRYPOINT $PROJECT_DIR/tools/entrypoint.sh
# ENTRYPOINT ["/bin/bash", "-c"]

CMD ["/bin/bash"]
