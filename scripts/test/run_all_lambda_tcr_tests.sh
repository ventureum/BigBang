#!/usr/bin/env bash
bazel clean
bazel run //:gazelle
./run_unit_test.sh  //cmd/lambda/TCR/tcr_table_creator/config:go_default_test
./run_unit_test.sh  //test/lambda/TCR/project:go_default_test
./run_unit_test.sh  //test/lambda/TCR/milestone:go_default_test
./run_unit_test.sh  //test/lambda/TCR/objective:go_default_test

./run_unit_test.sh  //test/lambda/TCR/get_project:go_default_test
./run_unit_test.sh  //test/lambda/TCR/get_project_list:go_default_test
./run_unit_test.sh  //test/lambda/TCR/get_milestone:go_default_test
./run_unit_test.sh  //test/lambda/TCR/get_objective:go_default_test

./run_unit_test.sh  //test/lambda/TCR/delete_objective:go_default_test
./run_unit_test.sh  //test/lambda/TCR/delete_milestone:go_default_test
./run_unit_test.sh  //test/lambda/TCR/delete_project:go_default_test

./run_unit_test.sh  //cmd/lambda/TCR/add_proxy/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/get_proxy_list/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/delete_proxy/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/rating_vote/config:go_default_test
./run_unit_test.sh  //cmd/lambda/TCR/get_rating_vote_list/config:go_default_test
