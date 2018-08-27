package error_config

import (
  "fmt"
  "strings"
)

type ErrorCode string

const InsufficientReputaionsAmount ErrorCode = "InsufficientReputaionsAmount"
const NoActorExisting ErrorCode = "NoActorExisting"
const NoPostHashExisting ErrorCode = "NoPostHashExisting"
const ExceedingUpvoteLimit ErrorCode = "ExceedingUpvoteLimit"
const ExceedingDownvoteLimit ErrorCode = "ExceedingDownvoteLimit"
const General ErrorCode = "Gerenal"

func CreateNoExistingErrorCode(tag string) ErrorCode {
  errCodeStr := fmt.Sprintf("No%sExisting", strings.Title(tag))
  return ErrorCode(errCodeStr)
}