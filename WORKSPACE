# Bazel WORKSPACE
# Dependencies last updated on 2019-03-29

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
    commit = "b5d812f8a3706043e23a9cd5babf2e5423744d30", # 2019-02-28 v1.3.1 
    importpath = "github.com/golang/protobuf",
)

# This is required by gRPC.
go_repository(
    name = "org_golang_google_genproto",
    commit = "d831d65fe17df2e52bcc4316d4a9f7a418701f43", # 2019-03-27
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "3507fb8e1a5ad030303c106fef3a47c9fdad16ad", # 2019-03-20 v1.19.1 
    importpath = "google.golang.org/grpc",
)

# This is required by gRPC.
go_repository(
    name = "org_golang_x_text",
    commit = "e3703dcdd614d2d7488fff034c75c551ea25da95", # 2018-12-15
    importpath = "golang.org/x/text",
)

go_repository(
    name = "org_golang_x_net",
    commit = "74de082e2cca95839e88aa0aeee5aadf6ce7710f", # 2019-02-21
    importpath = "golang.org/x/net",
)

go_repository(
    name = "com_github_google_go_cmp",
    commit = "3af367b6b30c263d47e8895973edcca9a49cf029", # 2018-02-19 v0.2.0
    importpath = "github.com/google/go-cmp",
)