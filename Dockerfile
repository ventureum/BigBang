FROM golang:latest

ENV GOBIN /go/bin
ENV GOPATH /go

RUN \
    # This makes add-apt-repository available.
    apt-get update && \
    apt-get -y install \
        --no-install-recommends apt-utils \
        python \
        python-pkg-resources \
        software-properties-common \
        unzip && \
    # Install bazel (https://docs.bazel.build/versions/master/install-ubuntu.html)
    apt-get -y install openjdk-8-jdk && \
    echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list && \
    curl https://bazel.build/bazel-release.pub.gpg | apt-key add - && \
    apt-get update && \
    apt-get -y install bazel && \
    apt-get -y upgrade bazel && \
    # Unpack bazel for future use.
    bazel version

# build directories
RUN mkdir  $GOPATH/src/BigBang
COPY . $GOPATH/src/BigBang

RUN rm -rf $GOPATH/src/BigBang/vendor
RUN rm -rf $GOPATH/pkg

# Go dep
RUN go get -u github.com/golang/dep/...
RUN cd $GOPATH/src/BigBang && \
    dep ensure -v &&\
    bazel run //:gazelle && \
    bazel run //:gazelle -- vendor
