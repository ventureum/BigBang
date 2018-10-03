#!/usr/bin/env bash
bazel clean
bazel run //:gazelle
./run_unit_test.sh  //cmd/lambda/TCR/tcr_table_creator/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/new_project/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/get_project/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/get_project_list/config:go_default_test
