package lambda_get_recent_votes_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor string `json:"actor,required"`
	Limit int64  `json:"limit,omitempty"`
}

type Response struct {
	RecentVotes *[]post_votes_record_config.PostVotesRecord `json:"recentVotes,omitempty"`
	Ok          bool                                        `json:"ok"`
	Message     *error_config.ErrorInfo                     `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.RecentVotes = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)
	limit := request.Body.Limit

	if limit == 0 {
		limit = 20
	}

	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{
		*postgresBigBangClient}
	postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{
		*postgresBigBangClient}

	actorProfileRecordExecutor.VerifyActorExistingTx(actor)
	actorRewardsInfoRecordExecutor.VerifyActorExistingTx(actor)

	response.RecentVotes = postVotesRecordExecutor.GetRecentPostVotesRecordsByActorTx(actor, limit)

	postgresBigBangClient.Commit()

	log.Printf("RecentPostVotesRecords is loaded for actor %s\n", actor)
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
