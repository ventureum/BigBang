package lambda_get_redeem_history_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/pkg/utils"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/feed/actor_milestone_points_redeem_history_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  Limit int64 `json:"limit,required"`
  Cursor string `json:"cursor,omitempty"`
}

type Response struct {
  Redeems *[]feed_attributes.MilestonePointsRedeemHistory `json:"redeems,omitempty"`
  NextCursor string `json:"nextCursor,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Redeems = nil
      response.NextCursor = ""
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}

  actor := request.Actor
  limit := request.Limit
  cursorStr := request.Cursor

  actorProfileRecordExecutor.VerifyActorExisting(actor)

  var cursor string
  if cursorStr != "" {
    cursor = utils.Base64DecodeToString(cursorStr)
  }

  actorMilestonePointsRedeemHistoryRecordExecutor := actor_milestone_points_redeem_history_record_config.ActorMilestonePointsRedeemHistoryRecordExecutor{*postgresBigBangClient}

  actorMilestonePointsRedeemHistory := actorMilestonePointsRedeemHistoryRecordExecutor.GetActorMilestonePointsRedeemHistoryByCursor(
    actor, cursor, limit + 1)

  response.NextCursor = ""

  var redeems []feed_attributes.MilestonePointsRedeemHistory
  for index, redeem := range *actorMilestonePointsRedeemHistory {
    if index < int(limit) {
      redeems = append(redeems, redeem)
    } else {
      response.NextCursor = redeem.GenerateRecordID()
    }
  }

  response.Redeems = &redeems

  if cursorStr == "" {
    log.Printf("ActorMilestonePointsRedeemHistory is loaded for first query with actor %s and limit %d\n",
      actor, limit)
  } else {
    log.Printf("ActorMilestonePointsRedeemHistory is loaded for query with actor %s, cursor %s and limit %d\n",
      actor, cursorStr, limit)
  }
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}