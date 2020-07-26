load("@bazel_gazelle//:def.bzl", "gazelle")
load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle_binary")

gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES + ["@golink//gazelle/go_link:go_default_library"],
    visibility = ["//visibility:public"],
)

# gazelle:prefix github.com/nikunjy/go
gazelle(
    name = "gazelle",
    gazelle = "//:gazelle_binary",
)
