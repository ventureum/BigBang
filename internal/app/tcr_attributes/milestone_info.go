package tcr_attributes

import (
	"encoding/json"
	"github.com/jmoiron/sqlx/types"
	"log"
)

type MilestonesInfo struct {
	CurrentMilestone       int64        `json:"currentMilestone,required"`
	NumMilestones          int64        `json:"numMilestones,required"`
	NumMilestonesCompleted int64        `json:"numMilestonesCompleted,required"`
	Milestones             *[]Milestone `json:"milestones,omitempty"`
}

func (milestonesInfo *MilestonesInfo) ToJsonText() types.JSONText {
	marshaled, err := json.Marshal(milestonesInfo)
	if err != nil {
		log.Panicf("Failed to marshal MilestonesInfo %+v with error: %+v\n", milestonesInfo, err)
	}
	return types.JSONText(string(marshaled))
}
