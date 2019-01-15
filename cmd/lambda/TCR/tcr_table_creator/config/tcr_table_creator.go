package lambda_tcr_table_creator_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_validator_record_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
	"BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/TCR/proxy_config"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/internal/platform/postgres_config/client_config"
)

type Request struct {
	DBInfo *client_config.DBInfo `json:"dbInfo,omitempty"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(request.DBInfo)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()

	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
	proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
	ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
	actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
	principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}
	milestoneValidatorRecordExecutor := milestone_validator_record_config.MilestoneValidatorRecordExecutor{*postgresBigBangClient}

	principalProxyVotesExecutor.DeletePrincipalProxyVotesTable()
	actorDelegateVotesAccountExecutor.DeleteActorRatingVoteAccountTable()
	objectiveExecutor.DeleteObjectiveTable()
	milestoneExecutor.DeleteMilestoneTable()
	projectExecutor.DeleteProjectTable()
	proxyExecutor.DeleteProxyTable()
	ratingVoteExecutor.DeleteRatingVoteTable()
	milestoneValidatorRecordExecutor.DeleteMilestoneValidatorRecordTable()

	projectExecutor.CreateProjectTable()
	milestoneExecutor.CreateMilestoneTable()
	objectiveExecutor.CreateObjectiveTable()
	proxyExecutor.CreateProxyTable()
	ratingVoteExecutor.CreateRatingVoteTable()
	actorDelegateVotesAccountExecutor.CreateActorRatingVoteAccountTable()
	principalProxyVotesExecutor.CreatePrincipalProxyVotesTable()
	milestoneValidatorRecordExecutor.CreateMilestoneValidatorRecordTable()

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
