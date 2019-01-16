package lambda_get_recent_posts_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor    string `json:"actor,required"`
	TypeHash string `json:"typeHash,required"`
	Limit    int64  `json:"limit,omitempty"`
}

type ResponseContent struct {
	PostHash             string `json:"postHash"`
	Actor                string `json:"actor"`
	PostType             string `json:"postType"`
	DeltaFuel            int64  `json:"deltaFuel"`
	DeltaReputation      int64  `json:"deltaReputation"`
	DeltaMilestonePoints int64  `json:"deltaMilestonePoints"`
	WithdrawableMPs      int64  `json:"withdrawableMPs"`
}

type Response struct {
	RecentPosts *[]post_rewards_record_config.PostRewardsRecord `json:"recentPosts,omitempty"`
	Ok          bool                                            `json:"ok"`
	Message     *error_config.ErrorInfo                         `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.RecentPosts = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	postType := feed_attributes.CreatePostTypeFromHashStr(request.Body.TypeHash)
	limit := request.Body.Limit

	if limit == 0 {
		limit = 20
	}

	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}

	actorProfileRecordExecutor.VerifyActorExistingTx(actor)
	actorRewardsInfoRecordExecutor.VerifyActorExistingTx(actor)

	response.RecentPosts = postRewardsRecordExecutor.GetRecentPostRewardsRecordsByActorTx(actor, postType, limit)

	postgresBigBangClient.Commit()

	log.Printf("RecentPostRewardsRecords is loaded for actor %s\n", actor)
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
