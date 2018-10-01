package config

import (
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
  "BigBang/internal/app/feed_attributes"
  "math"
)


type Request struct {
  Actor string `json:"actor,required"`
}

type ResponseContent struct {
  Actor string `json:"actor,required"`
  ActorType string `json:"actorType,required"`
  Username string `json:"username,required"`
  PhotoUrl string `json:"photoUrl,required"`
  TelegramId string `json:"telegramId,required"`
  PhoneNumber string `json:"phoneNumber,required"`
  Level int64 `json:"level,required"`
  RewardsInfo *feed_attributes.RewardsInfo `json:"rewardsInfo,required"`
}

type Response struct {
  Profile *ResponseContent `json:"profile,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProfileRecordResultToResponseContent(actorProfileRecord *actor_profile_record_config.ActorProfileRecord) *ResponseContent {
  return &ResponseContent{
    Actor: actorProfileRecord.Actor,
    ActorType: string(actorProfileRecord.ActorType),
    Username: actorProfileRecord.Username,
    PhotoUrl: actorProfileRecord.PhotoUrl,
    TelegramId: actorProfileRecord.TelegramId,
    PhoneNumber: actorProfileRecord.PhoneNumber,
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Profile = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()


  actor := request.Actor

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*postgresBigBangClient}

  actorProfileRecordExecutor.VerifyActorExisting(actor)
  actorRewardsInfoRecordExecutor.VerifyActorExisting(actor)

  actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecord(actor)
  response.Profile = ProfileRecordResultToResponseContent(actorProfileRecord)
  log.Printf("Loaded Profile content for actor %s\n", actor)
  rewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actor)
  log.Printf("Loaded Rewards info for actor %s\n", actor)
  response.Profile.RewardsInfo = rewardsInfo
  response.Profile.Level = int64(math.Floor(math.Log10(1 + math.Max(float64(rewardsInfo.Reputation), 0))))
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
