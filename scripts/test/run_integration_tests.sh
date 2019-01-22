#!/usr/bin/env bash
bazel clean
bazel run //:gazelle

soda reset -e development -p ../../migrations

bazel test --config=ci  //test/integration/...
