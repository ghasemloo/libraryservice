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

    _Bazel_ is our build tool.

    [Instructions for intalling Bazel on Ubuntu](https://bazel.build/versions/master/docs/install-ubuntu.html).

1. Install Git

    _Git_ is used by Bazel to download some external dependencies.

    ```shell
    $ sudo apt-get install git
    ```

1. (Optional) Install Go

    The service and client are written in _Go_.

    [Instructions for installing Go on Ubuntu](https://github.com/golang/go/wiki/Ubuntu)

1. (Optional) Install Gazelle

    _Gazelle_ automatically generated Bazel build files for Go.

    ```shell
    $ go install github.com/bazelbuild/rules_go/go/tools/gazelle/gazelle
    ```

1. (Optional) Install Buildifier

    _Buildifier_ cleans up and formats Bazel build files.

    ```shell
    $ go get github.com/bazelbuild/buildifier/buildifier
    ```

## Clone, Build, Test

```shell
$ git clone github.com/ghasemloo/libraryservice
$ cd libraryservice
$ bazel build ...
$ bazel test ...
```

## Run the server and the clinet

In one terminal run the server:
```shell
$ ./bazel-bin/server/server --logtostderr
```
In another terminal run the client to send requests to the server:
```shell
$ ./bazel-bin/client/client --logtostderr
```

## References

* Bazel: http://bazel.build
* Go: https://golang.org
* Google API Style Guide: https://cloud.google.com/apis/design/
* Proto3: https://developers.google.com/protocol-buffers/
* gRPC: https://www.grpc.io/
* Google RPC [Status](https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto) and 
[Canonical Error Codes](https://github.com/googleapis/googleapis/blob/master/google/rpc/codes.proto)
