package tcr_integration_test

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/test/lambda/TCR/activate_milestone"
	"BigBang/test/lambda/TCR/add_milestone"
	"BigBang/test/lambda/TCR/add_proxy"
	"BigBang/test/lambda/TCR/add_proxy_voting_for_principal"
	"BigBang/test/lambda/TCR/batch_objectives"
	"BigBang/test/lambda/TCR/delete_batch_objectives"
	"BigBang/test/lambda/TCR/delete_milestone"
	"BigBang/test/lambda/TCR/delete_objective"
	"BigBang/test/lambda/TCR/delete_project"
	"BigBang/test/lambda/TCR/delete_proxy"
	"BigBang/test/lambda/TCR/finalize_milestone"
	"BigBang/test/lambda/TCR/finalize_validators"
	"BigBang/test/lambda/TCR/get_batch_finalized_validators"
	"BigBang/test/lambda/TCR/get_batch_proxy_voting_info"
	"BigBang/test/lambda/TCR/get_batch_rating_vote_list"
	"BigBang/test/lambda/TCR/get_finalized_validators"
	"BigBang/test/lambda/TCR/get_milestone"
	"BigBang/test/lambda/TCR/get_objective"
	"BigBang/test/lambda/TCR/get_project"
	"BigBang/test/lambda/TCR/get_project_id_by_admin"
	"BigBang/test/lambda/TCR/get_project_list"
	"BigBang/test/lambda/TCR/get_proxy_list"
	"BigBang/test/lambda/TCR/get_proxy_voting_info"
	"BigBang/test/lambda/TCR/get_rating_vote_list"
	"BigBang/test/lambda/TCR/get_validator_recent_rating_vote_activities"
	"BigBang/test/lambda/TCR/objective"
	"BigBang/test/lambda/TCR/project"
	"BigBang/test/lambda/TCR/rating_vote"
	"BigBang/test/lambda/TCR/tcr_table_creator"
	"BigBang/test/lambda/TCR/update_available_delegate_votes"
	"BigBang/test/lambda/TCR/update_batch_available_delegate_votes"
	"BigBang/test/lambda/TCR/update_batch_received_delegate_votes"
	"BigBang/test/lambda/TCR/update_received_delegate_votes"
	"BigBang/test/lambda/feed/feed_events_table_creator"
	"BigBang/test/lambda/feed/profile"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("feed_events_table_creator", feed_events_table_creator_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("tcr_table_creator", tcr_table_creator_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("profile", profile_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("project", project_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("add_milestone", add_milestone_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("activate_milestone", activate_milestone_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("finalize_milestone", finalize_milestone_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("objective", objective_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("batch_objectives", batch_objectives_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_project", get_project_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_project_list", get_project_list_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_milestone", get_milestone_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_objective", get_objective_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("rating_vote", rating_vote_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("add_proxy", add_proxy_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("update_available_delegate_votes", update_available_delegate_votes_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("update_received_delegate_votes", update_received_delegate_votes_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("update_batch_available_delegate_votes", update_batch_available_delegate_votes_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("update_batch_received_delegate_votes", update_batch_received_delegate_votes_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("add_proxy_voting_for_principal", add_proxy_voting_for_principal_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_rating_vote_list", get_rating_vote_list_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_batch_rating_vote_list", get_batch_rating_vote_list_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_validator_recent_rating_vote_activities",
		get_validator_recent_rating_vote_activities_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_proxy_voting_info", get_proxy_voting_info_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_batch_proxy_voting_info", get_batch_proxy_voting_info_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("finalize_validators", finalize_validators_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_finalized_validators", get_finalized_validators_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_batch_finalized_validators", get_batch_finalized_validators_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("delete_objective", delete_objective_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("delete_batch_objectives", delete_batch_objectives_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("delete_milestone", delete_milestone_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("delete_project", delete_project_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_proxy_list", get_proxy_list_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
	t.Run("get_project_id_by_admin", get_project_id_by_admin_test.TestHandler)

	os.Setenv("AUTH_LEVEL", string(auth.AdminAuth))
	t.Run("delete_proxy", delete_proxy_test.TestHandler)
}
