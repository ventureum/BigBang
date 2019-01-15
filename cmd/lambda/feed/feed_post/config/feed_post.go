package lambda_feed_post_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/getstream_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_config"
	"BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
	"gopkg.in/GetStream/stream-go2.v1"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor              string                  `json:"actor,required"`
	BoardId            string                  `json:"boardId,required"`
	ParentHash         string                  `json:"parentHash,required"`
	PostHash           string                  `json:"postHash,required"`
	TypeHash           string                  `json:"typeHash,required"`
	Content            feed_attributes.Content `json:"content,required"`
	GetStreamApiKey    string                  `json:"getStreamApiKey,omitempty"`
	GetStreamApiSecret string                  `json:"getStreamApiSecret,omitempty"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToPostRecord() *post_config.PostRecord {
	return &post_config.PostRecord{
		Actor:      request.Body.Actor,
		BoardId:    request.Body.BoardId,
		ParentHash: request.Body.ParentHash,
		PostHash:   request.Body.PostHash,
		PostType:   feed_attributes.CreatePostTypeFromHashStr(request.Body.TypeHash).Value(),
		Content:    request.Body.Content.ToJsonText(),
	}
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	actor := request.Body.Actor

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	var err error
	var getStreamIOClient *stream.Client
	if request.Body.GetStreamApiKey != "" && request.Body.GetStreamApiSecret != "" {
		getStreamIOClient, err = stream.NewClient(request.Body.GetStreamApiKey, request.Body.GetStreamApiSecret)
	} else {
		getStreamIOClient, err = stream.NewClientFromEnv()
	}

	if err != nil {
		log.Panic(err.Error())
	}

	getStreamClient := &getstream_config.GetStreamClient{C: getStreamIOClient}

	postRecord := request.ToPostRecord()
	ProcessPostRecord(postRecord, getStreamClient, postgresBigBangClient, feed_attributes.OFF_CHAIN)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func ProcessPostRecord(
	postRecord *post_config.PostRecord,
	getStreamClient *getstream_config.GetStreamClient,
	postgresBigBangClient *client_config.PostgresBigBangClient,
	source feed_attributes.Source) {
	postExecutor := post_config.PostExecutor{*postgresBigBangClient}
	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(postRecord.Actor)
	actorRewardsInfoRecordExecutor.VerifyActorExistingTx(postRecord.Actor)

	updateCount := postExecutor.GetPostUpdateCountTx(postRecord.PostHash)
	fuelsPenalty := feed_attributes.FuelsPenaltyForPostType(
		feed_attributes.PostType(postRecord.PostType), updateCount)

	log.Printf("UpdateCount for PostHash %s: %d", postRecord.PostHash, updateCount)
	log.Printf("Fuel Penalty for PostHash %s: %d", postRecord.PostHash, fuelsPenalty)

	// Update Actor Fuel
	actorRewardsInfoRecordExecutor.SubActorFuelTx(postRecord.Actor, fuelsPenalty)

	// Insert Post Record
	createdTimestamp := postExecutor.UpsertPostRecordTx(postRecord)
	activity := postRecord.ToActivity(source, feed_attributes.BlockTimestamp(createdTimestamp.Unix()))

	postRewardsRecordExecutor.UpsertPostRewardsRecordTx(&post_rewards_record_config.PostRewardsRecord{
		PostHash:  postRecord.PostHash,
		Actor:     postRecord.Actor,
		PostType:  postRecord.PostType,
		Object:    activity.Object.Value(),
		PostTime:  createdTimestamp,
		DeltaFuel: int64(fuelsPenalty.Neg()),
	})

	// Insert Activity to GetStream
	if updateCount == 0 {
		getStreamClient.AddFeedActivityToGetStream(activity)
	} else {
		getStreamClient.UpdateFeedActivityToGetStream(activity)
	}

	// Update Post Replies Record
	if activity.Verb == feed_attributes.ReplyVerb {
		postRepliesRecord := post_replies_record_config.PostRepliesRecord{
			PostHash:  postRecord.ParentHash,
			ReplyHash: postRecord.PostHash,
		}
		postRepliesRecordExecutor.UpsertPostRepliesRecordTx(&postRepliesRecord)
	}
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
