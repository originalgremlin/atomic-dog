FROM golang:1.15.6

# os packages
RUN set -eu \
    && apt-get -y update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install --no-install-recommends \
        libczmq-dev \
        libsodium-dev \
        libzmq3-dev \
        vim \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# gore: go repl
RUN set -eu \
    && go get -u \
        github.com/motemen/gore/cmd/gore \
        github.com/mdempsky/gocode \
        github.com/k0kubun/pp

# go packages
RUN set -eu \
    && go get -u \
        github.com/zeromq/goczmq

# static files
COPY root /

ENTRYPOINT [ "/usr/local/go/bin/go" ]
