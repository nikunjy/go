load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
def docker_repos():
    http_archive(
        name = "io_bazel_rules_docker",
        sha256 = "dc97fccceacd4c6be14e800b2a00693d5e8d07f69ee187babfd04a80a9f8e250",
        strip_prefix = "rules_docker-0.14.1",
        urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.1/rules_docker-v0.14.1.tar.gz"],
    )
