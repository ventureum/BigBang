package feed_attributes

import (
  "BigBang/internal/pkg/error_config"
  "log"
)

type ActorType string

const USER_ACTOR_TYPE ActorType = "USER"

const KOL_ACTOR_TYPE ActorType = "KOL"

const ADMIN_ACTOR_TYPE ActorType = "ADMIN"

// Project Founder
const PF_ACTOR_TYPE ActorType = "PF"

func ValidateAndCreateActorType (actorTypeStr string) ActorType {
  var actorType  ActorType
  switch ActorType(actorTypeStr) {
     case USER_ACTOR_TYPE, KOL_ACTOR_TYPE, ADMIN_ACTOR_TYPE, PF_ACTOR_TYPE:
         actorType = ActorType(actorTypeStr)
     default:
         errorInfo := error_config.ErrorInfo{
           ErrorCode: error_config.InvalidActorType,
           ErrorData: error_config.ErrorData {
             "actorType": actorTypeStr,
           },
           ErrorLocation: error_config.ActorTypeLocation,
         }
         log.Printf("Invalid actorType: %s", actorTypeStr)
         log.Panicln(errorInfo.Marshal())
  }
  return actorType
}

