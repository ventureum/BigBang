#!/usr/bin/env bash

echo "generating attach_session zip"
.//generate_lambda_zip.sh attach_session $1 $2

echo "generating deactivate_actor zip"
.//generate_lambda_zip.sh deactivate_actor $1 $2

echo "generating feed_events_table_creator zip"
.//generate_lambda_zip.sh feed_events_table_creator $1 $2

echo "generating feed_post zip"
.//generate_lambda_zip.sh feed_post $1 $2

echo "generating feed_token_generator zip"
.//generate_lambda_zip.sh feed_token_generator $1 $2

echo "generating feed_update_post_rewards zip"
.//generate_lambda_zip.sh feed_update_post_rewards $1 $2

echo "generating feed_upvote zip"
.//generate_lambda_zip.sh feed_upvote $1 $2

echo "generating get_batch_posts zip"
.//generate_lambda_zip.sh get_batch_posts $1 $2

echo "generating get_feed_post zip"
.//generate_lambda_zip.sh get_feed_post $1 $2

echo "generating get_profile zip"
.//generate_lambda_zip.sh get_profile $1 $2

echo "generating get_recent_posts zip"
.//generate_lambda_zip.sh get_recent_posts $1 $2

echo "generating get_recent_votes zip"
.//generate_lambda_zip.sh get_recent_votes $1 $2


echo "generating get_session zip"
.//generate_lambda_zip.sh get_session $1 $2

echo "generating profile zip"
.//generate_lambda_zip.sh profile $1 $2

echo "generating refuel zip"
.//generate_lambda_zip.sh refuel $1 $2
