package config

import (
  "log"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/refuel_record_config"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
  "strconv"
  "os"
  "time"
  "math"
)

type Request struct {
  Actor string `json:"actor,required"`
}

type Response struct {
  Ok      bool   `json:"ok"`
  RefuelAmount int64 `json:"refuelAmount,omitempty"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.RefuelAmount = 0
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()

  actor := request.Actor

  debugMode, _ := strconv.ParseInt(os.Getenv("DEBUG_MODE"), 10, 64)
  refuelInterval, _ := strconv.ParseInt(os.Getenv("REFUEL_INTERVAL"), 10, 64)
  refuelReplenishmentHourly, _ := strconv.ParseInt(os.Getenv("FUEL_REPLENISHMENT_HOURLY"), 10, 64)

  postgresFeedClient.Begin()

  refuelRecordExecutor := refuel_record_config.RefuelRecordExecutor{
    *postgresFeedClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresFeedClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}

  actorProfileRecordExecutor.VerifyActorExistingTx(actor)
  actorRewardsInfoRecordExecutor.VerifyActorExistingTx(actor)

  lastRefuelTime := refuelRecordExecutor.GetLastRefuelTimeTx(actor)
  deltaTime := time.Now().UTC().Unix() - lastRefuelTime.Unix()

  log.Printf("lastRefuelTime %+v", lastRefuelTime)
  log.Printf("deltaTime %+v", deltaTime)

  if deltaTime < refuelInterval * 3600 && debugMode != 1 {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.InsufficientWaitingTimeToRefuel,
      ErrorData: map[string]interface{} {
        "lastRefuelTimestamp": lastRefuelTime.Unix(),
      },
    }
    log.Printf("Insufficient Waiting Time To Refuel for actor %s", actor)
    log.Panicln(errorInfo.Marshal())
  }

  hoursSinceLastReplenishment := deltaTime / 3600

  newFuelIncremental := feed_attributes.Fuel(math.Min(
    float64(feed_attributes.MuMaxFuel),
    float64(hoursSinceLastReplenishment * refuelReplenishmentHourly)))
  actorRewardsInfoRecordExecutor.AddActorFuelTx(actor, newFuelIncremental)
  refuelRecordExecutor.UpsertRefuelRecordTx(&refuel_record_config.RefuelRecord{
    Actor: actor,
    Fuel: newFuelIncremental,
    Reputation: 0,
    MilestonePoints: 0,
  })
  postgresFeedClient.Commit()

  log.Printf("Refuel %d fuel to actor %s", newFuelIncremental, actor)
  response.RefuelAmount = int64(newFuelIncremental)
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
