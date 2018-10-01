package config

import (
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/feed/refuel_record_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  UserType string `json:"userType,required"`
  Username string `json:"username,required"`
  PhotoUrl string `json:"photoUrl,omitempty"`
  TelegramId string `json:"telegramId,omitempty"`
  PhoneNumber string `json:"phoneNumber,omitempty"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToActorProfileRecord() (*actor_profile_record_config.ActorProfileRecord) {
  return &actor_profile_record_config.ActorProfileRecord{
    Actor:      request.Actor,
    ActorType: feed_attributes.ActorType(request.UserType),
    Username: request.Username,
    PhotoUrl: request.PhotoUrl,
    TelegramId: request.TelegramId,
    PhoneNumber: request.PhoneNumber,
  }
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

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}

  inserted := actorProfileRecordExecutor.UpsertActorProfileRecordTx(request.ToActorProfileRecord())

  if inserted {
    refuelRecordExecutor := refuel_record_config.RefuelRecordExecutor{*postgresBigBangClient}
    actorReputationsRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
      *postgresBigBangClient}
    actorReputationsRecord := actor_rewards_info_record_config.ActorRewardsInfoRecord{
      Actor:           request.Actor,
      Reputation:      feed_attributes.Reputation(feed_attributes.MuMinFuel.Value()),
      Fuel:            feed_attributes.MuMinFuel,
      MilestonePoints: 0,
    }
    actorReputationsRecordExecutor.UpsertActorRewardsInfoRecordTx(&actorReputationsRecord)
    refuelRecordExecutor.UpsertRefuelRecordTx(&refuel_record_config.RefuelRecord{
      Actor: request.Actor,
      Fuel: feed_attributes.MuMinFuel,
      Reputation: feed_attributes.Reputation(feed_attributes.MuMinFuel),
      MilestonePoints: 0,
    })

    log.Printf("Created Actor Fuel Account for actor %s", request.Actor)
  }

  postgresBigBangClient.Commit()

  if inserted {
    log.Printf("Created Profile for actor %s", request.Actor)
  } else {
    log.Printf("Updated Profile for actor %s", request.Actor)
  }

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
