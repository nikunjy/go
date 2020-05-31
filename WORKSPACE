load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("//bazel/go:repos.bzl", "add_go_repos")
add_go_repos()
load("//bazel/go:def.bzl", "go_rules_deps")
go_rules_deps()

# Download the rules_docker repository at release v0.14.1
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "dc97fccceacd4c6be14e800b2a00693d5e8d07f69ee187babfd04a80a9f8e250",
    strip_prefix = "rules_docker-0.14.1",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.1/rules_docker-v0.14.1.tar.gz"],
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")
container_repositories()
load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")
container_deps()
load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
_go_image_repos()



load("//:go_third_party.bzl", "go_deps")

# gazelle:repository_macro go_third_party.bzl%go_deps
go_deps()