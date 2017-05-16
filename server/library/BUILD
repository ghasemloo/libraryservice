load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

package(default_visibility = [
    "//visibility:private",
])

go_library(
    name = "go_default_library",
    library = ":library",
    visibility = [
        "//server:__subpackages__",
    ],
)

go_library(
    name = "library",
    srcs = [
        "library.go",
    ],
    deps = [
        "//proto:library_service_proto",
        "//proto/internal:library_proto",
        "@com_github_golang_glog//:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)

go_test(
    name = "library_test",
    size = "small",
    srcs = [
        "library_test.go",
    ],
    library = ":library",
    deps = [
        "//proto:library_service_proto",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)
