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
    $ go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
    ```

1. (Optional) Install Buildifier

    [Bazel Buildtools](https://github.com/bazelbuild/buildtools) are helpers tools for Bazel. [Buildifier](https://github.com/bazelbuild/buildtools/blob/master/buildifier/README.md) cleans up and formats Bazel build files.

    ```shell
    $ go install github.com/bazelbuild/buildtools/buildifier@latest
    ```

## Clone

```shell
$ git clone git@github.com:ghasemloo/libraryservice.git
$ cd libraryservice
```

## Bazel

### Build and Test

```shell
$ bazel build ...
$ bazel test ...
```

### Run the server and the client

In one terminal run the server:
```shell
$ bazel run server/server -- --logtostderr
```
In another terminal run the client to send requests to the server:
```shell
$ bazel run client/client -- --logtostderr
```

## Go (without Bazel)

### Go Module

```shell
$ go mod init github.com/ghasemloo/libraryservice
$ go mod tidy
```

### Proto Compiler for Go

Install Protobuf Compiler to generate Go packages for proto files:

```shell
$ sudo apt-get install protobuf-compiler
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Generated Go Proto Packages

```shell
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/api/api.proto
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/storage/storage.proto
```

### Build and Test

```shell
$ go build ./...
$ go test ./...
```

### Run the server and the client

In two different terminals run:

```shell
$ go run ./server/main.go
```

```shell
$ go run ./client/main.go
```

## Gazelle

Updating `WORKSPACE.bazel` from `go.mod`:

```shell
$ gazelle update-repos --proto_import_prefix="github.com/ghasemloo/libraryservice" --from_file=./go.mod --prune
```

Creating/updating `BUILD.bazel` files:

```shell
$ gazelle fix -go_prefix github.com/ghasemloo/libraryservice
```

You can also use Bazel (`gazelle` commands and paramters are configured in the `BUILD.bazel`):

```shell
$ bazel run //:gazelle-update-repos
$ bazel run //:gazelle
```

## References

* Bazel: http://bazel.build
* Go: https://golang.org
* Google API Style Guide: https://cloud.google.com/apis/design/
* Proto3: https://developers.google.com/protocol-buffers/
* gRPC: https://www.grpc.io/
* gRPC Example: https://grpc.io/docs/languages/go/quickstart/
* Google RPC [Status](https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto) and 
[Canonical Error Codes](https://github.com/googleapis/googleapis/blob/master/google/rpc/codes.proto)
