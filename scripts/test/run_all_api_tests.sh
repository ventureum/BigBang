#!/usr/bin/env bash

bazel clean
bazel run //:gazelle
./run_unit_test.sh  //test/apis/feed:go_default_test
./run_unit_test.sh  //test/apis/TCR:go_default_test
