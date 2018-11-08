package lambda_get_next_redeem_config

import (
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
  "time"
  "BigBang/internal/app/feed_attributes"
)


type Request struct {
  Actor string `json:"actor,required"`
}

type ResponseContent struct {
  Actor string `json:"actor,required"`
  MilestonePoints feed_attributes.MilestonePoint `json:"milestonePoints,required"`
  EstimatedTokens int64 `json:"estimatedTokens,required"`
  SubmittedAt time.Time `json:"submittedAt,required"`
  RedeemBlockInfo *feed_attributes.RedeemBlockInfo `json:"redeemBlockInfo,omitempty"`
}

type Response struct {
  NextRedeem *ResponseContent `json:"nextRedeem,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()

  actor := request.Actor
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(actor)
  milestonePointsRedeemRequestRecordExecutor := milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{*postgresBigBangClient}

  //nextRedeemBlock := time.Now().UTC().Unix() / (60 * 60 * 24 * 7) + 1
  milestonePointsRedeemRequestRecordExecutor.VerifyMilestonePointsRedeemRequestExisting(actor)

  milestonePointsRedeemRequest := milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(actor)
  response.NextRedeem = &ResponseContent{
    Actor: actor,
    MilestonePoints: milestonePointsRedeemRequest.TargetedMilestonePoints,
    SubmittedAt: milestonePointsRedeemRequest.SubmittedAt,
  }

  log.Printf("Sucessfully loaded content for get_next_redeem for actor %s\n", actor)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
