# The go packages in this repository will have this string as a prefox in import statements.
load("@io_bazel_rules_go//go:def.bzl", "go_prefix")

go_prefix("github.com/ghasemloo/libraryservice")

# The gazelle target to fix BUILD files.
load("@io_bazel_rules_go//go:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/example/project",
)

#load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_google_protobuf")
#go_google_protobuf()
