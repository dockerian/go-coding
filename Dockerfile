FROM golang:latest

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
 && echo "export PS1='\n\u@\h \w [\#]:\n\$ ' " >> ~/.bashrc \
 && echo "alias ll='ls -al'" >> ~/.bashrc \
 && echo "" >> ~/.bashrc

RUN go get -d github.com/tools/godep && \
    go install github.com/tools/godep && \
    go get -u github.com/golang/lint/golint && \
    go get -u github.com/sanbornm/go-selfupdate && \
    go install github.com/sanbornm/go-selfupdate

#
# downloading the latest go-coding source code so that it allows to
# run the container without mapping to any local go-coding copy
# e.g.
#       docker build -t dockerian/go-coding .
#       docker run --rm -it dockerian/go-coding
#
ENV GOPATH=/go \
    PROJECT_DIR=/go/src/github.com/dockerian/go-coding \
    SHELL=/bin/bash
RUN mkdir -p /go/src/github.com/dockerian \
 && git clone \
    https://github.com/dockerian/go-coding \
    /go/src/github.com/dockerian/go-coding \
 && mkdir -p "$PROJECT_DIR"

WORKDIR $PROJECT_DIR

# ENTRYPOINT ["/bin/bash", "-c"]

CMD ["/bin/bash"]
