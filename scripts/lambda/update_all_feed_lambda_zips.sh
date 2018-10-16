#!/usr/bin/env bash

echo "updating attach_session zip"
.//update_lambda_zip.sh feed attach_session $1 $2

echo "updating deactivate_actor zip"
.//update_lambda_zip.sh feed deactivate_actor $1 $2

echo "updating dev_refuel zip"
.//update_lambda_zip.sh feed dev_refuel $1 $2

echo "updating feed_events_table_creator zip"
.//update_lambda_zip.sh feed feed_events_table_creator $1 $2

echo "updating feed_post zip"
.//update_lambda_zip.sh feed feed_post $1 $2

echo "updating feed_token_generator zip"
.//update_lambda_zip.sh feed feed_token_generator $1 $2

echo "updating feed_update_post_rewards zip"
.//update_lambda_zip.sh feed feed_update_post_rewards $1 $2

echo "updating feed_upvote zip"
.//update_lambda_zip.sh feed feed_upvote $1 $2

echo "updating get_batch_posts zip"
.//update_lambda_zip.sh feed get_batch_posts $1 $2

echo "updating get_feed_post zip"
.//update_lambda_zip.sh feed get_feed_post $1 $2

echo "updating get_profile zip"
.//update_lambda_zip.sh feed get_profile $1 $2

echo "updating get_recent_posts zip"
.//update_lambda_zip.sh feed get_recent_posts $1 $2

echo "updating get_recent_votes zip"
.//update_lambda_zip.sh feed get_recent_votes $1 $2

echo "updating get_session zip"
.//update_lambda_zip.sh feed get_session $1 $2

echo "updating profile zip"
.//update_lambda_zip.sh feed profile $1 $2

echo "updating refuel zip"
.//update_lambda_zip.sh feed refuel $1 $2

echo "updating reset_actor_fuel zip"
.//update_lambda_zip.sh feed reset_actor_fuel $1 $2
