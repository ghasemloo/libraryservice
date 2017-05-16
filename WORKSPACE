# Bazel WORKSPACE

# Contains Git repository rules
http_archive(
    name = "io_bazel",
    sha256 = "2b737be42678900470ae9e48c975ac5b2296d9ae23c007bf118350dbe7c0552b",
    urls = ["https://github.com/bazelbuild/bazel/releases/download/0.4.5/bazel-0.4.5-dist.zip"],
)

# Load Git repository rules
load(
  "@io_bazel//tools/build_defs/repo:git.bzl",
  "git_repository",
  "new_git_repository",
)

# Cotnains Go repository rules
git_repository(
    name = "io_bazel_rules_go",
    commit = "4c9a52aba0b59511c5646af88d2f93a9c0193647", # 2017-05-05 release 0.4.4
    remote = "https://github.com/bazelbuild/rules_go.git",
)

# Load Go repository rules
load(
  "@io_bazel_rules_go//go:def.bzl",
  "go_repositories",
  "new_go_repository",
)
go_repositories()

# Load Go Proto repository rules
load(
  "@io_bazel_rules_go//proto:go_proto_library.bzl",
  "go_proto_repositories",
)
go_proto_repositories()

# Dependencies

new_go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b4da8ac2ed8007337bc5eb5007998", # 2016-01-25
    importpath = "github.com/golang/glog"
)

new_go_repository(
    name = "com_github_golang_protobuf",
    commit = "fec3b39b059c0f88fa6b20f5ed012b1aa203a8b4", # 2017-05-12
    importpath = "github.com/golang/protobuf",
)


# This is required by gRPC.
new_go_repository(
    name = "org_golang_google_genproto",
    commit = "411e09b969b1170a9f0c467558eb4c4c110d9c77", # 2017-04-04
    importpath = "google.golang.org/genproto",
)

new_go_repository(
    name = "org_golang_google_grpc",
    commit = "d2e1b51f33ff8c5e4a15560ff049d200e83726c5", # 2017-04-18
    importpath = "google.golang.org/grpc",
)

# This is required by gRPC.
new_go_repository(
    name = "org_golang_x_text",
    commit = "19e51611da83d6be54ddafce4a4af510cb3e9ea4", # 2017-04-21
    importpath = "golang.org/x/text",
)

new_go_repository(
    name = "org_golang_x_net",
    commit = "34057069f4ab13dc4433c68d368737ebeafcccdc", # 2017-05-09
    importpath = "golang.org/x/net",
)
