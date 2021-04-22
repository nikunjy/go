load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

local_repository(
    name = "bazel_rules_go",
    path = "bazel/go/",
)

load("@bazel_rules_go//:repos.bzl", "add_go_repos")
add_go_repos()
load("@bazel_rules_go//:def.bzl", "go_rules_deps")
go_rules_deps()

load("//bazel/docker:repos.bzl", "docker_repos")
docker_repos()

load("//bazel/docker:def.bzl", "docker_deps")
docker_deps()

http_archive(
    name = "golink",
    urls = ["https://github.com/nikunjy/golink/archive/v1.1.0.tar.gz"],
    sha256 = "c505a82b7180d4315bbaf05848e9b7d2683e80f1b16159af51a0ecae6fb2d54d",
    strip_prefix = "golink-1.1.0"
)


load("//:go_third_party.bzl", "go_deps")
# gazelle:repository_macro go_third_party.bzl%go_deps
go_deps()
