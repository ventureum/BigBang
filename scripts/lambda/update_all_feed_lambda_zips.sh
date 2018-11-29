#!/usr/bin/env bash

echo "updating add_tracked_wallet_addresses zip"
.//update_lambda_zip.sh feed add_tracked_wallet_addresses $1 $2 UserAuth

echo "updating attach_session zip"
.//update_lambda_zip.sh feed attach_session $1 $2 UserAuth

echo "updating deactivate_actor zip"
.//update_lambda_zip.sh feed deactivate_actor $1 $2 AdminAuth

echo "updating delete_tracked_wallet_addresses zip"
.//update_lambda_zip.sh feed delete_tracked_wallet_addresses $1 $2 UserAuth

echo "updating dev_refuel zip"
.//update_lambda_zip.sh feed dev_refuel $1 $2 AdminAuth

echo "updating feed_events_table_creator zip"
.//update_lambda_zip.sh feed feed_events_table_creator $1 $2 AdminAuth

echo "updating feed_post zip"
.//update_lambda_zip.sh feed feed_post $1 $2 UserAuth

echo "updating feed_redeem_milestone_points zip"
.//update_lambda_zip.sh feed feed_redeem_milestone_points $1 $2 AdminAuth

echo "updating feed_token_generator zip"
.//update_lambda_zip.sh feed feed_token_generator $1 $2 AdminAuth

echo "updating feed_update_post_rewards zip"
.//update_lambda_zip.sh feed feed_update_post_rewards $1 $2 AdminAuth

echo "updating feed_upvote zip"
.//update_lambda_zip.sh feed feed_upvote $1 $2 UserAuth

echo "updating get_actor_private_key zip"
.//update_lambda_zip.sh feed get_actor_private_key $1 $2 UserAuth

echo "updating get_actor_uuid_from_public_key zip"
.//update_lambda_zip.sh feed get_actor_uuid_from_public_key $1 $2 UserAuth

echo "updating get_batch_posts zip"
.//update_lambda_zip.sh feed get_batch_posts $1 $2 AdminAuth

echo "updating get_batch_profiles zip"
.//update_lambda_zip.sh feed get_batch_profiles $1 $2 AdminAuth

echo "updating get_feed_post zip"
.//update_lambda_zip.sh feed get_feed_post $1 $2 UserAuth

echo "updating get_next_redeem zip"
.//update_lambda_zip.sh feed get_next_redeem  $1 $2 UserAuth

echo "updating get_profile zip"
.//update_lambda_zip.sh feed get_profile $1 $2 UserAuth

echo "updating get_recent_posts zip"
.//update_lambda_zip.sh feed get_recent_posts $1 $2 UserAuth

echo "updating get_recent_votes zip"
.//update_lambda_zip.sh feed get_recent_votes $1 $2 UserAuth

echo "updating get_redeem_block_info zip"
.//update_lambda_zip.sh feed get_redeem_block_info  $1 $2 NoAuth

echo "updating get_redeem_history zip"
.//update_lambda_zip.sh feed get_redeem_history  $1 $2 UserAuth

echo "updating get_session zip"
.//update_lambda_zip.sh feed get_session $1 $2 UserAuth

echo "updating get_tracked_wallet_addresses zip"
.//update_lambda_zip.sh feed get_tracked_wallet_addresses $1 $2 UserAuth

echo "updating profile zip"
.//update_lambda_zip.sh feed profile $1 $2 	UserAuth

echo "updating refuel zip"
.//update_lambda_zip.sh feed refuel $1 $2 	UserAuth

echo "updating reset_actor_fuel zip"
.//update_lambda_zip.sh feed reset_actor_fuel $1 $2 AdminAuth

echo "updating set_actor_private_key zip"
.//update_lambda_zip.sh feed set_actor_private_key $1 $2 AdminAuth

echo "updating set_next_redeem zip"
.//update_lambda_zip.sh feed set_next_redeem  $1 $2 UserAuth

echo "updating set_token_pool zip"
.//update_lambda_zip.sh feed set_token_pool  $1 $2 AdminAuth
