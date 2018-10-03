package tcr_attributes

import (
  "github.com/jmoiron/sqlx/types"
  "log"
  "encoding/json"
)

type MilestoneInfo struct {
  NumMilestones int64 `json:"numMilestones,required"`
  NumMilestonesCompleted int64 `json:"numMilestonesCompleted,required"`
}

func (milestoneInfo *MilestoneInfo) ToJsonText() types.JSONText {
  marshaled, err := json.Marshal(milestoneInfo)
  if err != nil {
    log.Panicf("Failed to marshal MilestoneInfo %+v with error: %+v\n", milestoneInfo, err)
  }
  return types.JSONText(string(marshaled))
}


func CreatedMilestoneInfoFromJsonText(jsonText types.JSONText) *MilestoneInfo {
  var milestoneInfo MilestoneInfo
  err := jsonText.Unmarshal(&milestoneInfo)
  if err != nil {
    log.Panicf("Failed to unmarshal jsonText %+v with error: %+v\n", jsonText, err)
  }
  return &milestoneInfo
}

