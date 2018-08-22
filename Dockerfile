FROM golang:1.10.3
RUN apt-get update -qqy --fix-missing
RUN apt-get install -qqy build-essential \
    git \
    curl \
    libcurl4-openssl-dev \
    libjansson-dev \
    gdb \
    sudo \
    clang
RUN go get -u golang.org/x/vgo/...
ADD . /go/src/github.com/STNS/STNS
WORKDIR /go/src/github.com/STNS/STNS
RUN make depsdev