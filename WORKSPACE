# Bazel WORKSPACE

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Contains Go repository rules
http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz",
    sha256 = "77dfd303492f2634de7a660445ee2d3de2960cbd52f97d8c0dffa9362d3ddef9",
)

# Contains Gazelle
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

# Load Go repository rules
load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_rules_dependencies",
    "go_register_toolchains",
)
go_rules_dependencies()
go_register_toolchains()

# Load Gazelle
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

# Contains Git repository rules
http_archive(
    name = "io_bazel",
    sha256 = "621d2a97899a88850a913eabf9285778331a309fd4658b225b1377f80060fa85",
    url = "https://github.com/bazelbuild/bazel/releases/download/0.24.0/bazel-0.24.0-dist.zip",
)

# Load Git repository rules
load(
    "@io_bazel//tools/build_defs/repo:git.bzl",
    "git_repository",
    "new_git_repository",
)

# Load Go Proto repository rules
load("@io_bazel_rules_go//proto:def.bzl",
    "go_proto_library",
)

# Dependencies

go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b4da8ac2ed8007337bc5eb5007998", # 2016-01-25
    importpath = "github.com/golang/glog"
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "aa810b61a9c79d51363740d207bb46cf8e620ed5", # 2017-05-12 v1.2.0 
    importpath = "github.com/golang/protobuf",
)

# This is required by gRPC.
go_repository(
    name = "org_golang_google_genproto",
    commit = "411e09b969b1170a9f0c467558eb4c4c110d9c77", # 2017-04-04
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "d2e1b51f33ff8c5e4a15560ff049d200e83726c5", # 2017-04-18
    importpath = "google.golang.org/grpc",
)

# This is required by gRPC.
go_repository(
    name = "org_golang_x_text",
    commit = "19e51611da83d6be54ddafce4a4af510cb3e9ea4", # 2017-04-21
    importpath = "golang.org/x/text",
)

go_repository(
    name = "org_golang_x_net",
    commit = "34057069f4ab13dc4433c68d368737ebeafcccdc", # 2017-05-09
    importpath = "golang.org/x/net",
)
