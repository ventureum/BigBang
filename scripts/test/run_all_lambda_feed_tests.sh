#!/usr/bin/env bash
bazel clean
bazel run //:gazelle
./run_unit_test.sh  //test/lambda/feed/feed_events_table_creator:go_default_test
./run_unit_test.sh  //test/lambda/feed/profile:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_profile:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_batch_profiles:go_default_test

./run_unit_test.sh  //test/lambda/feed/set_actor_private_key:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_actor_private_key:go_default_test

./run_unit_test.sh  //test/lambda/feed/get_actor_uuid_from_public_key:go_default_test
./run_unit_test.sh  //test/lambda/feed/feed_post:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_feed_post:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_batch_posts:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_recent_posts:go_default_test
./run_unit_test.sh  //test/lambda/feed/attach_session:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_session:go_default_test
./run_unit_test.sh  //test/lambda/feed/refuel:go_default_test
./run_unit_test.sh  //test/lambda/feed/dev_refuel:go_default_test
./run_unit_test.sh  //test/lambda/feed/feed_update_post_rewards:go_default_test
./run_unit_test.sh  //test/lambda/feed/feed_upvote:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_recent_votes:go_default_test
./run_unit_test.sh  //test/lambda/feed/feed_token_generator:go_default_test
./run_unit_test.sh  //test/lambda/feed/reset_actor_fuel:go_default_test
./run_unit_test.sh  //test/lambda/feed/add_tracked_wallet_addresses:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_tracked_wallet_addresses:go_default_test
./run_unit_test.sh  //test/lambda/feed/delete_tracked_wallet_addresses:go_default_test
./run_unit_test.sh  //test/lambda/feed/set_next_redeem:go_default_test
./run_unit_test.sh  //test/lambda/feed/get_next_redeem:go_default_test

./run_unit_test.sh  //test/lambda/feed/deactivate_actor:go_default_test
