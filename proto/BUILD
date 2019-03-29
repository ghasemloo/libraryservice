load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:private"])

proto_library(
    name = "library_service_proto",
    srcs = ["library_service.proto"],
    visibility = ["//visibility:private"],
)

go_proto_library(
    name = "library_service_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/ghasemloo/libraryservice/proto",
    proto = ":library_service_proto",
    visibility = ["//visibility:private"],
)

go_library(
    name = "go_default_library",
    embed = [":library_service_go_proto"],
    importpath = "github.com/ghasemloo/libraryservice/proto",
    visibility = [
        "//client:__subpackages__",
        "//server:__subpackages__",
        "//tests:__subpackages__",
    ],
)
