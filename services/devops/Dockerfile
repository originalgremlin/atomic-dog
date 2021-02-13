FROM golang:1.15.6

# os packages
RUN set -eu \
    && apt-get -y update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y install --no-install-recommends \
        less \
        libzmq3-dev \
        vim \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# yaegi: go repl
RUN set -eu \
    && go get -u \
        github.com/traefik/yaegi/cmd/yaegi

# go packages
RUN set -eu \
    && go get -u \
        github.com/vmihailenco/msgpack \
        github.com/pebbe/zmq4

# static files
COPY root /

ENTRYPOINT [ "/usr/local/go/bin/go" ]
