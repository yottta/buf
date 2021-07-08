# Generated by "buf config migrate-v1beta1". Edit as necessary, and
# remove this comment when you're finished.
#
# This module represents the "." root found in
# the previous configuration file for the
# "buf.build/test/test" module.
version: v1
name: buf.build/test/test
deps:
  - buf.build/beta/googleapis
  - buf.build/beta/envoy
build:
  excludes:
    - dir1
breaking:
  use:
    - FILE
  ignore:
    - dir2/file.proto
  ignore_only:
    FIELD_SAME_JSON_NAME:
      - dir3/file.proto
lint:
  use:
    - DEFAULT
  ignore:
    - dir2/file.proto
  ignore_only:
    ENUM_PASCAL_CASE:
      - dir3/file.proto