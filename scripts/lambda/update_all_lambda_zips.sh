#!/usr/bin/env bash

echo "generating attach_session zip"
.//update_lambda_zip attach_session $1 $2

echo "generating deactivate_actor zip"
.//update_lambda_zip deactivate_actor $1 $2

echo "generating feed_events_table_creator zip"
.//update_lambda_zip feed_events_table_creator $1 $2

echo "generating feed_post zip"
.//update_lambda_zip feed_post $1 $2

echo "generating feed_token_generator zip"
.//update_lambda_zip feed_token_generator $1 $2

echo "generating feed_update_post_rewards zip"
.//update_lambda_zip feed_update_post_rewards $1 $2

echo "generating feed_upvote zip"
.//update_lambda_zip feed_upvote $1 $2

echo "generating get_batch_posts zip"
.//update_lambda_zip get_batch_posts $1 $2

echo "generating get_feed_post zip"
.//update_lambda_zip get_feed_post $1 $2

echo "generating get_profile zip"
.//update_lambda_zip get_profile $1 $2

echo "generating get_recent_posts zip"
.//update_lambda_zip get_recent_posts $1 $2

echo "generating get_recent_votes zip"
.//update_lambda_zip get_recent_votes $1 $2


echo "generating get_session zip"
.//update_lambda_zip get_session $1 $2

echo "generating profile zip"
.//update_lambda_zip profile $1 $2

echo "generating refuel zip"
.//update_lambda_zip refuel $1 $2
