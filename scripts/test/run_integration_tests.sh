#!/usr/bin/env bash

bazel clean
bazel run //:gazelle
bazel test --config=ci  //test/integration/...
