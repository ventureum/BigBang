package error_config

import (
  "fmt"
  "strings"
)

type ErrorCode string

// For feed
const GetStreamClientConnectionError ErrorCode = "GetStreamClientConnectionError"
const InsufficientWaitingTimeToRefuel ErrorCode = "InsufficientWaitingTimeToRefuel"
const InsufficientReputaionsAmount ErrorCode = "InsufficientReputaionsAmount"
const InsufficientFuelsAmount ErrorCode = "InsufficientFuelsAmount"
const NoActorExisting ErrorCode = "NoActorExisting"
const NoPostHashExisting ErrorCode = "NoPostHashExisting"
const ExceedingUpvoteLimit ErrorCode = "ExceedingUpvoteLimit"
const ExceedingDownvoteLimit ErrorCode = "ExceedingDownvoteLimit"
const General ErrorCode = "General"

// For TCR
const NoProjectIdExisting ErrorCode = "NoProjectIdExisting"
const NoObjectiveIdExisting ErrorCode = "NoObjectiveIdExisting"
const NoMilestoneIdExisting ErrorCode = "NoMilestoneIdExisting"
const NoRatingVoteVoterExisting ErrorCode = "NoRatingVoteVoterExisting"
const NoProxyUUIDExisting ErrorCode = "NoProxyUUIDExisting"
const ProxyUUIDAlreadyExisting ErrorCode = "ProxyUUIDAlreadyExisting"

func CreateNoExistingErrorCode(tag string) ErrorCode {
  errCodeStr := fmt.Sprintf("No%sExisting", strings.Title(tag))
  return ErrorCode(errCodeStr)
}