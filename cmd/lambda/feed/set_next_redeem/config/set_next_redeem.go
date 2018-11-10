package lambda_set_next_redeem_config

import (
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
  "BigBang/internal/app/feed_attributes"
)


type Request struct {
  Actor string `json:"actor,required"`
  MilestonePoints int64 `json:"milestonePoints,required"`
}

type Response struct {
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

  postgresBigBangClient.Begin()

  actor := request.Actor
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(actor)

  nextRedeemBlock := feed_attributes.MoveToNextNRedeemBlock(1)
  milestonePoints := request.MilestonePoints

  if milestonePoints < 0 {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.InvalidMilestonePoints,
      ErrorData: map[string]interface{} {
        "milestonePoints": milestonePoints,
        "actor": actor,
      },
      ErrorLocation: error_config.MilestonePointsRedeemRequestRecordLocation,
    }
    log.Printf("Invalid milestonePoints for Milestone Points Redeem Request for actor %s: %d", actor, milestonePoints)
    log.Panicln(errorInfo.Marshal())
  } else if milestonePoints == 0  {
    nextRedeemBlock = 0
  }

  milestonePointsRedeemRequestRecordExecutor := milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{*postgresBigBangClient}

  milestonePointsRedeemRequestRecordExecutor.UpsertMilestonePointsRedeemRequestRecordTx(
    &milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecord {
      Actor: request.Actor,
      NextRedeemBlock: nextRedeemBlock,
      TargetedMilestonePoints: feed_attributes.MilestonePoint(milestonePoints),})

  postgresBigBangClient.Commit()

  log.Printf("Sucessfully set next redeem for actor %s with milestonePoints %d\n", actor, milestonePoints)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
