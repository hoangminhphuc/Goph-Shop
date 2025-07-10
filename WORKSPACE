load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# rules_go: Go support for Bazel
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/v0.44.0/rules_go-v0.44.0.zip"],
    sha256 = "c8035e8ae248b56040a65ad3f0b7434712e2037e5dfdcebfe97576e620422709",
)

# gazelle: auto-generates BUILD.bazel files from Go code (optional for later)
http_archive(
    name = "bazel_gazelle",
    sha256 = "75df288c4b31c81eb50f51e2e14f4763cb7548daae126817247064637fd9ea62",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.36.0/bazel-gazelle-v0.36.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.36.0/bazel-gazelle-v0.36.0.tar.gz",
    ],
)

# ───────────── rules_go setup ─────────────
# load rules_go setup functions
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
# go deps
go_rules_dependencies()
# use system-installed Go (avoid having Bazel download one)
go_register_toolchains(version = "host")

# ───────────── gazelle setup ─────────────
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()