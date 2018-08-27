package error_config

import (
  "regexp"
  "database/sql"
)

const RegexpForActorFkeyViolation = `violates foreign key constraint .*_actor_fkey`
const RegexpForPostHashFkeyViolation = `violates foreign key constraint .*_post_hash_fkey`

func MatchErrorString(regExpStr string, str string) (bool) {
  re := regexp.MustCompile(regExpStr)
  return re.MatchString(str)
}

func MatchError(err error, fieldName string, val interface{}, location ErrorLocation) *ErrorInfo {
  var errorInfo ErrorInfo
  errStr := err.Error()
  if (err == sql.ErrNoRows && fieldName == "actor") || MatchErrorString(RegexpForActorFkeyViolation, errStr) {
    errorInfo.ErrorCode = NoActorExisting
    errorInfo.ErrorData = map[string]interface{}{
      "actor": val,
    }
    errorInfo.ErrorLocation = location
  } else if (err == sql.ErrNoRows && fieldName == "postHash") || MatchErrorString(RegexpForPostHashFkeyViolation, errStr) {
    errorInfo.ErrorCode = NoPostHashExisting
    errorInfo.ErrorData = map[string]interface{}{
      "postHash": val,
    }
    errorInfo.ErrorLocation = location
  } else {
    errorInfo.ErrorCode = General
    errorInfo.ErrorLocation = location
    errorInfo.ErrorMessage = ErrorMessage(errStr)
  }
  return &errorInfo
}
