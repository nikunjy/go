load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

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
    urls = ["https://github.com/nikunjy/golink/archive/v1.0.0.tar.gz"],
    sha256 = "ea728cfc9cb6e2ae024e1d5fbff185224592bbd4dad6516f3cc96d5155b69f0d",
    strip_prefix = "golink-1.0.0",
)



load("//:go_third_party.bzl", "go_deps")
# gazelle:repository_macro go_third_party.bzl%go_deps
go_deps()
