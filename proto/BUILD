load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_library")

package(default_visibility = [
    "//visibility:private",
])

go_proto_library(
    name = "library_service_proto",
    srcs = [
        "library_service.proto",
    ],
    has_services = 1,
    rules_go_repo_only_for_internal_use = "@",
    visibility = [
        "//client:__subpackages__",
        "//server:__subpackages__",
    ],
    deps = [],
)
