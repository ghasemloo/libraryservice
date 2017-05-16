load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

package(default_visibility = [
    "//visibility:private",
])

go_binary(
    name = "server",
    srcs = [
        "main.go",
    ],
    visibility = [
        "//visibility:public",
    ],
    deps = [
        "//proto:library_service_proto",
        "//server/library:go_default_library",
        "@com_github_golang_glog//:go_default_library",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)
