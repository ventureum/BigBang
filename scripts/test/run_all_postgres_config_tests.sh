#!/usr/bin/env bash
bazel clean
bazel run //:gazelle
./run_unit_test.sh  //internal/platform/postgres_config/TCR/proxy_config:go_default_test
