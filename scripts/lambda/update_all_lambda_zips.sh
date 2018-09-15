#!/usr/bin/env bash

echo "updating attach_session zip"
.//update_lambda_zip.sh attach_session $1 $2

echo "updating deactivate_actor zip"
.//update_lambda_zip.sh deactivate_actor $1 $2

echo "updating feed_events_table_creator zip"
.//update_lambda_zip.sh feed_events_table_creator $1 $2

echo "updating feed_post zip"
.//update_lambda_zip.sh feed_post $1 $2

echo "updating feed_token_generator zip"
.//update_lambda_zip.sh feed_token_generator $1 $2

echo "updating feed_update_post_rewards zip"
.//update_lambda_zip.sh feed_update_post_rewards $1 $2

echo "updating feed_upvote zip"
.//update_lambda_zip.sh feed_upvote $1 $2

echo "updating get_batch_posts zip"
.//update_lambda_zip.sh get_batch_posts $1 $2

echo "updating get_feed_post zip"
.//update_lambda_zip.sh get_feed_post $1 $2

echo "updating get_profile zip"
.//update_lambda_zip.sh get_profile $1 $2

echo "updating get_recent_posts zip"
.//update_lambda_zip.sh get_recent_posts $1 $2

echo "updating get_recent_votes zip"
.//update_lambda_zip.sh get_recent_votes $1 $2

echo "updating get_session zip"
.//update_lambda_zip.sh get_session $1 $2

echo "updating profile zip"
.//update_lambda_zip.sh profile $1 $2

echo "updating refuel zip"
.//update_lambda_zip.sh refuel $1 $2

echo "updating dev_refuel zip"
.//update_lambda_zip.sh dev_refuel $1 $2
