package(default_visibility = ["//kythe:default_visibility"])

java_test(
    name = "FileDataCacheTest",
    size = "small",
    srcs = ["FileDataCacheTest.java"],
    deps = [
        "//kythe/java/com/google/devtools/kythe/platform/shared:file_data",
        "//kythe/proto:analysis_java_proto",
        "//third_party/guava",
        "@com_google_protobuf//:protobuf_java",
        "@com_google_truth_truth//jar",
        "@junit_junit//jar",
    ],
)

java_test(
    name = "FileDataDirectoryTest",
    size = "small",
    srcs = ["FileDataDirectoryTest.java"],
    deps = [
        "//kythe/java/com/google/devtools/kythe/platform/shared:file_data",
        "//third_party/guava",
        "@com_google_truth_truth//jar",
        "@junit_junit//jar",
    ],
)
