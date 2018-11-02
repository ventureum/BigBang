package lambda_get_batch_profiles_config

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
  Actors []string `json:"actors,required"`
}

type ResponseContent struct {
  Actor string `json:"actor,required"`
  ActorType string `json:"actorType,required"`
  Username string `json:"username,required"`
  PhotoUrl string `json:"photoUrl,required"`
  TelegramId string `json:"telegramId,required"`
  PhoneNumber string `json:"phoneNumber,required"`
  PublicKey string `json:"publicKey,required"`
  Level int64 `json:"level,required"`
  RewardsInfo *feed_attributes.RewardsInfo `json:"rewardsInfo,required"`
}

type Response struct {
  Profiles *[]ResponseContent `json:"profiles,omitempty"`
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
    PublicKey: actorProfileRecord.PublicKey,
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Profiles = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*postgresBigBangClient}
  actors := request.Actors

  for _ , actor := range actors {
    actorProfileRecordExecutor.VerifyActorExisting(actor)
    actorRewardsInfoRecordExecutor.VerifyActorExisting(actor)
  }

  var profiles []ResponseContent
  for _ , actor := range actors {
    actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecord(actor)
    profile := ProfileRecordResultToResponseContent(actorProfileRecord)
    log.Printf("Loaded Profile content for actor %s\n", actor)
    rewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actor)
    log.Printf("Loaded Rewards info for actor %s\n", actor)
    profile.RewardsInfo = rewardsInfo
    profile.Level = int64(math.Floor(math.Log10(1 + math.Max(float64(rewardsInfo.Reputation), 0))))
    profiles = append(profiles, *profile)
  }

  response.Profiles = &profiles
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
