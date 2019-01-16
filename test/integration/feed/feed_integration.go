package feed_integration_test

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/test/lambda/feed/add_tracked_wallet_addresses"
	"BigBang/test/lambda/feed/attach_session"
	"BigBang/test/lambda/feed/deactivate_actor"
	"BigBang/test/lambda/feed/delete_batch_actors"
	"BigBang/test/lambda/feed/delete_tracked_wallet_addresses"
	"BigBang/test/lambda/feed/dev_refuel"
	"BigBang/test/lambda/feed/feed_events_table_creator"
	"BigBang/test/lambda/feed/feed_post"
	"BigBang/test/lambda/feed/feed_redeem_milestone_points"
	"BigBang/test/lambda/feed/feed_token_generator"
	"BigBang/test/lambda/feed/feed_update_post_rewards"
	"BigBang/test/lambda/feed/feed_upvote"
	"BigBang/test/lambda/feed/get_actor_private_key"
	"BigBang/test/lambda/feed/get_actor_uuid_from_public_key"
	"BigBang/test/lambda/feed/get_batch_posts"
	"BigBang/test/lambda/feed/get_batch_profiles"
	"BigBang/test/lambda/feed/get_feed_post"
	"BigBang/test/lambda/feed/get_next_redeem"
	"BigBang/test/lambda/feed/get_profile"
	"BigBang/test/lambda/feed/get_recent_posts"
	"BigBang/test/lambda/feed/get_recent_votes"
	"BigBang/test/lambda/feed/get_redeem_block_info"
	"BigBang/test/lambda/feed/get_redeem_history"
	"BigBang/test/lambda/feed/get_session"
	"BigBang/test/lambda/feed/get_tracked_wallet_addresses"
	"BigBang/test/lambda/feed/profile"
	"BigBang/test/lambda/feed/refuel"
	"BigBang/test/lambda/feed/set_actor_private_key"
	"BigBang/test/lambda/feed/set_next_redeem"
	"BigBang/test/lambda/feed/set_token_pool"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("feed_events_table_creator", feed_events_table_creator_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("profile", profile_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_profile", get_profile_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_batch_profiles", get_batch_profiles_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("set_actor_private_key", set_actor_private_key_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_actor_private_key", get_actor_private_key_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_actor_uuid_from_public_key", get_actor_uuid_from_public_key_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("feed_post", feed_post_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_feed_post", get_feed_post_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_recent_posts", get_recent_posts_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("get_batch_posts", get_batch_posts_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("attach_session", attach_session_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_session", get_session_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("refuel", refuel_test.TestHandlerWithDebugMode)
	t.Run("refuel", refuel_test.TestHandlerWithoutDebugMode)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("dev_refuel", dev_refuel_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("feed_upvote", feed_upvote_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_recent_votes_test", get_recent_votes_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("feed_token_generator", feed_token_generator_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("reset_actor_fuel", dev_refuel_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("add_tracked_wallet_addresses", add_tracked_wallet_addresses_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("add_tracked_wallet_addresses", get_tracked_wallet_addresses_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("delete_tracked_wallet_addresses", delete_tracked_wallet_addresses_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("set_next_redeem", set_next_redeem_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_next_redeem", get_next_redeem_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("feed_update_post_rewards", feed_update_post_rewards_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("feed_redeem_milestone_points", feed_redeem_milestone_points_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
	t.Run("get_redeem_history", get_redeem_history_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("set_token_pool", set_token_pool_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_redeem_block_info", get_redeem_block_info_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("deactivate_actor", deactivate_actor_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("delete_batch_actors", delete_batch_actors.TestHandler)
}
