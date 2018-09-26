#!/usr/bin/env bash

bazel clean
bazel run //:gazelle
./run_unit_test.sh  //test/apis:go_default_test