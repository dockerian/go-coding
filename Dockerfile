FROM golang:latest

ARG DOCKER_IMAG=dockerian/go-coding
ARG GITHUB_REPO=github.com/dockerian/go-coding
MAINTAINER Jason Zhu <jason.zhuyx@gmail.com>
LABEL maintainer="jason.zhuyx@gmail.com"
LABEL organization="Dockerian"
LABEL project="go-coding"

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    bash \
    make \
    jq \
    tree \
    tar \
    zip \
 && rm -rf /var/lib/apt/lists/* \
 && rm /bin/sh && ln -sf /bin/bash /bin/sh \
 && mkdir ~/.ssh \
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
# COPY tools/entrypoint.sh /usr/local/bin/entrypoint.sh

RUN go get -u github.com/golang/lint/golint \
 && go get -u github.com/sanbornm/go-selfupdate \
 && go get -u github.com/ory/go-acc \
 && go get -u github.com/golang/dep/cmd/dep \
 && go get -u github.com/Masterminds/glide \
 && go get -u github.com/kardianos/govendor \
 && go get -u github.com/tools/godep

#
# downloading the latest go-coding source code so that it allows to
# run the container without mapping to any local go-coding copy
# e.g.
#       docker build -t dockerian/go-coding .
#       docker run --rm -it dockerian/go-coding
#
ENV GOPATH=/go \
    PROJECT=go-coding \
    PROJECT_DIR=/go/src/github.com/dockerian/go-coding \
    SHELL=/bin/bash

# creating "$PROJECT_DIR" and adding Godeps
ADD Godeps "$PROJECT_DIR/Godeps"

RUN cd -P "$PROJECT_DIR" && godep restore

WORKDIR $PROJECT_DIR

# this ENTRYPOINT requires gosu
# ENTRYPOINT $PROJECT_DIR/tools/entrypoint.sh
# ENTRYPOINT ["/bin/bash", "-c"]

CMD ["/bin/bash"]
