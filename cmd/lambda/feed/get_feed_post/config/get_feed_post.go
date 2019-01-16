package lambda_get_feed_post_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_votes_counters_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_config"
	"BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_counters_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	PostHash  string `json:"postHash,required"`
	Requestor string `json:"requestor,omitempty"`
}

type ResponseContent struct {
	Actor                  string                         `json:"actor"`
	Username               string                         `json:"username,required"`
	PhotoUrl               string                         `json:"photoUrl,required"`
	BoardId                string                         `json:"boardId"`
	ParentHash             string                         `json:"parentHash"`
	PostHash               string                         `json:"postHash"`
	PostType               string                         `json:"postType"`
	Content                *feed_attributes.Content       `json:"content"`
	DeltaFuel              int64                          `json:"deltaFuel"`
	DeltaReputation        int64                          `json:"deltaReputation"`
	DeltaMilestonePoints   int64                          `json:"deltaMilestonePoints"`
	WithdrawableMPs        int64                          `json:"withdrawableMPs"`
	RepliesLength          int64                          `json:"repliesLength"`
	PostVoteCountInfo      *feed_attributes.VoteCountInfo `json:"postVoteCountInfo,omitempty"`
	RequestorVoteCountInfo *feed_attributes.VoteCountInfo `json:"requestorVoteCountInfo,omitempty"`
}

type Response struct {
	Post    *ResponseContent        `json:"post,omitempty"`
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func PostRecordResultToResponseContent(result *post_config.PostRecordResult) *ResponseContent {
	return &ResponseContent{
		Actor:      result.Actor,
		BoardId:    result.BoardId,
		ParentHash: result.ParentHash,
		PostHash:   result.PostHash,
		PostType:   result.PostType,
		Content:    result.Content,
	}
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Post = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	postHash := request.Body.PostHash
	requestor := request.Body.Requestor
	auth.AuthProcess(request.PrincipalId, requestor, postgresBigBangClient)

	postExecutor := post_config.PostExecutor{*postgresBigBangClient}
	postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
	postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresBigBangClient}
	postVotesCounterRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresBigBangClient}
	actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{*postgresBigBangClient}
	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}

	postExecutor.VerifyPostRecordExistingTx(postHash)
	if requestor != "" {
		actorProfileRecordExecutor.VerifyActorExistingTx(requestor)
		actorRewardsInfoRecordExecutor.VerifyActorExistingTx(requestor)
	}

	postRecordResult := postExecutor.GetPostRecordTx(postHash).ToPostRecordResult()
	response.Post = PostRecordResultToResponseContent(postRecordResult)
	actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecordTx(postRecordResult.Actor)
	response.Post.Username = actorProfileRecord.Username
	response.Post.PhotoUrl = actorProfileRecord.PhotoUrl
	response.Post.RepliesLength = postRepliesRecordExecutor.GetPostRepliesRecordCountTx(postHash)
	postRewardsRecord := postRewardsRecordExecutor.GetPostRewardsRecordByPostHashTx(postHash)
	response.Post.DeltaFuel = postRewardsRecord.DeltaFuel
	response.Post.DeltaReputation = postRewardsRecord.DeltaReputation
	response.Post.DeltaMilestonePoints = postRewardsRecord.DeltaMilestonePoints
	response.Post.WithdrawableMPs = postRewardsRecord.WithdrawableMPs

	log.Printf("Post Content is loaded for postHash %s\n", postHash)

	postVotesCounterRecord := postVotesCounterRecordExecutor.GetPostVotesCountersRecordByPostHashTx(postHash)
	response.Post.PostVoteCountInfo = &feed_attributes.VoteCountInfo{
		DownVoteCount:  postVotesCounterRecord.DownVoteCount,
		UpVoteCount:    postVotesCounterRecord.UpVoteCount,
		TotalVoteCount: postVotesCounterRecord.TotalVoteCount,
	}

	log.Printf("PostVoteInfo is loaded for postHash %s\n", postHash)

	if requestor != "" {
		actorVotesCountersRecord := actorVotesCountersRecordExecutor.GetActorVotesCountersRecordByPostHashAndActorTx(postHash, requestor)
		response.Post.RequestorVoteCountInfo = &feed_attributes.VoteCountInfo{
			DownVoteCount:  actorVotesCountersRecord.DownVoteCount,
			UpVoteCount:    actorVotesCountersRecord.UpVoteCount,
			TotalVoteCount: actorVotesCountersRecord.TotalVoteCount,
		}
		log.Printf("RequestorVoteInfo is loaded for postHash %s\n", postHash)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
