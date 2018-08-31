#!/usr/bin/env bash

echo "generating attach_session zip"
./generate_attach_session_zip.sh

echo "generating deactivate_actor zip"
./generate_deactivate_actor_zip.sh

echo "generating feed_events_table_creator zip"
./generate_feed_events_table_creator_zip.sh

echo "generating feed_post zip"
./generate_feed_post_zip.sh

echo "generating feed_token_generator zip"
./generate_feed_token_generator_zip.sh

echo "generating feed_update_post_rewards zip"
./generate_feed_update_post_rewards_zip.sh

echo "generating feed_upvote zip"
./generate_feed_upvote_zip.sh

echo "generating get_batch_posts zip"
./generate_get_batch_posts_zip.sh

echo "generating get_feed_post zip"
./generate_get_feed_post_zip.sh

echo "generating get_profile zip"
./generate_get_profile_zip.sh

echo "generating get_recent_posts zip"
./generate_get_recent_posts_zip.sh

echo "generating get_recent_votes zip"
./generate_get_recent_posts_zip.sh


echo "generating get_session zip"
./generate_get_session_zip.sh

echo "generating profile zip"
./generate_profile_zip.sh

echo "generating refuel zip"
./generate_refuel_zip.sh
