load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")
load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")
load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")
load("@io_bazel_rules_docker//container:container.bzl", "container_pull")
def extra_container_repos():
    container_pull(
        name = "ubuntu16.04",
        registry = "gcr.io",
        repository = "tensorflow-testing/nosla-ubuntu16.04",
        digest = "sha256:b90dcf2f35f3354909f4491bdf019c110b4b4d95ef0395ebf178bc5d523a4208",
    )


def docker_deps():
    container_repositories()
    container_deps()
    _go_image_repos()
    extra_container_repos()
