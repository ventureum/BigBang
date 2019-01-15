package refuel_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type RefuelRecord struct {
	UUID            string                         `json:"uuid,required" db:"uuid"`
	Actor           string                         `json:"actor,required" db:"actor"`
	Fuel            feed_attributes.Fuel           `json:"fuel,required" db:"fuel"`
	Reputation      feed_attributes.Reputation     `json:"reputation,required" db:"reputation"`
	MilestonePoints feed_attributes.MilestonePoint `json:"milestonePoints,required" db:"milestone_points"`
	CreatedAt       time.Time                      `db:"created_at"`
	UpdatedAt       time.Time                      `db:"updated_at"`
}
