load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")
load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")
load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
def docker_deps():
    container_repositories()
    container_deps()
    _go_image_repos()
