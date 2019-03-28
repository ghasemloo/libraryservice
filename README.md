# A simple LibraryService to demonstrate how to use Bazel+Go+gRPC

A simple library service (server and client).

* Service defined in [Protocol Buffer v3](https://developers.google.com/protocol-buffers/docs/proto3)
(proto3).
* Programming language [The Go Programming Langauge](https://golang.org/) (golang).
* RPC framework [gRPC](https://www.grpc.io/).
* Build tool [Bazel](https://bazel.build).

## Pre-work

The following instructions are tested on Ubuntu 16.04 (LTS).

1.  Install Bazel

    [Bazel](https://bazel.build) is our build tool. It is written in [Java](https://www.java.com/). Install Java and then Bazel.

    ```shell
    $ sudo apt-get install bazel
    ```

    [Complete instructions for intalling Bazel on Ubuntu](https://bazel.build/versions/master/docs/install-ubuntu.html).

1. Install Git

    [Git](https://git-scm.com) is our version control system. It is used by Bazel to download some of our external dependencies (like [Protocol Buffers](https://github.com/google/protobuf), [gRPC](https://github.com/grpc/grpc-go), and Bazel's [Go rules](https://github.com/bazelbuild/rules_go).

    ```shell
    $ sudo apt-get install git
    ```

1. (Optional) Install Go

    [Go](https://golang.org) is a programming language and is very good at writting large-scale distributed systems. Our service and client are written in Go.

    ```shell
    $ sudo apt-get install golang-go
    ```

    [Compelte instructions for installing Go on Ubuntu](https://github.com/golang/go/wiki/Ubuntu)

1. (Optional) Install Gazelle

    [Gazelle](https://github.com/bazelbuild/bazel-gazelle) is a helper tool that automatically generates Bazel BUILD files for Go and makes your life easier.

    ```shell
    $ go get -u github.com/bazelbuild/bazel-gazelle/cmd/gazelle
    ```

1. (Optional) Install Buildifier

    [Bazel Buildtools](https://github.com/bazelbuild/buildtools) are helpers tools for Bazel. [Buildifier](https://github.com/bazelbuild/buildtools/blob/master/buildifier/README.md) cleans up and formats Bazel build files.

    ```shell
    $ go get github.com/bazelbuild/buildifier/buildifier
    ```

## Clone, Build, Test

```shell
$ git clone git@github.com:ghasemloo/libraryservice.git
$ cd libraryservice
$ bazel build ...
$ bazel test ...
```

## Run the server and the clinet

In one terminal run the server:
```shell
$ bazel run server/server -- --logtostderr
```
In another terminal run the client to send requests to the server:
```shell
$ bazel run client/client -- --logtostderr
```

## References

* Bazel: http://bazel.build
* Go: https://golang.org
* Google API Style Guide: https://cloud.google.com/apis/design/
* Proto3: https://developers.google.com/protocol-buffers/
* gRPC: https://www.grpc.io/
* Google RPC [Status](https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto) and 
[Canonical Error Codes](https://github.com/googleapis/googleapis/blob/master/google/rpc/codes.proto)
