load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:private"])

proto_library(
    name = "library_proto",
    srcs = ["library.proto"],
    visibility = ["//visibility:private"],
)

go_proto_library(
    name = "library_go_proto",
    importpath = "github.com/ghasemloo/libraryservice/proto/storage",
    proto = ":library_proto",
    visibility = ["//visibility:private"],
)

go_library(
    name = "go_default_library",
    embed = [":library_go_proto"],
    importpath = "github.com/ghasemloo/libraryservice/proto/storage",
    visibility = ["//server:__subpackages__"],
)
