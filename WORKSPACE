# Bazel WORKSPACE
# Dependencies last updated on 2023-08-10

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# -----------------------------------------------------------------------------
# Bazel rules.
# -----------------------------------------------------------------------------

# Contains Git repository rules.
http_archive(
    name = "io_bazel",
    sha256 = "8cd7feac58193be2bcba451ba6688a46824d37ca6359ff58e0d44eb98f042948",  # 2023-08-08
    url = "https://github.com/bazelbuild/bazel/releases/download/6.3.2/bazel-6.3.2-dist.zip",
)

# Contains Go repository rules.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",  # 2023-07-10
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

# Contains Gazelle.
http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",  # 2023-07-11
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

# Contains Protocol Buffer rules.
http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",  # 2022-12-27
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

# Contains Protocol Buffer rules.
http_archive(
    name = "com_google_protobuf",
    sha256 = "850357336189c470e429e9bdffca92229d8cd5b7f84aa2f3b4c5fdb80ce8351b",  # 2023-08-08
    strip_prefix = "protobuf-24.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v24.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/releases/download/v24.0/protobuf-24.0.tar.gz",
    ],
)

# Contains Google APIs proto files.
http_archive(
    name = "googleapis",
    sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",  # 2022-10-20
    strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
    urls = [
        "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
    ],
)

# Contains Bazel build tools.
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "977a0bd4593c8d4c8f45e056d181c35e48aa01ad4f8090bdb84f78dca42f47dc",  # 2023-04-27 v6.1.2
    strip_prefix = "buildtools-6.1.2",
    urls = [
        "https://github.com/bazelbuild/buildtools/archive/refs/tags/v6.1.2.tar.gz",
    ],
)

# -----------------------------------------------------------------------------
# Load Bazel rules.
# -----------------------------------------------------------------------------

# Load Git rules.
load(
    "@io_bazel//tools/build_defs/repo:git.bzl",
    "git_repository",  # @unused
    "new_git_repository",  # @unused
)

# Load Go rules.
load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

# Load Gazelle rules.
load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
)

# Load Protocol Buffers rules.
load(
    "@rules_proto//proto:repositories.bzl",
    "rules_proto_dependencies",
    "rules_proto_toolchains",
)

# Load Protocol Buffers dependencies.
load(
    "@com_google_protobuf//:protobuf_deps.bzl",
    "protobuf_deps",
)

# Load Google APIs rules.
load(
    "@googleapis//:repository_rules.bzl",
    "switched_rules_by_language",
)

switched_rules_by_language(
    name = "com_google_googleapis_imports",
)

# -----------------------------------------------------------------------------
# Go dependencies.
# -----------------------------------------------------------------------------

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

# -----------------------------------------------------------------------------
# Register Bazel rules.
# -----------------------------------------------------------------------------

rules_proto_dependencies()

rules_proto_toolchains()

go_rules_dependencies()

go_register_toolchains(version = "1.21.0")  # 2023-08-08

gazelle_dependencies()  #  (go_repository_default_config = "@//:WORKSPACE.bazel")

protobuf_deps()
