package milestone_validator_record_config

import (
	"time"
)

type MilestoneValidatorRecord struct {
	UUID        string    `json:"uuid,required" db:"uuid"`
	ProjectId   string    `json:"projectId,required" db:"project_id"`
	MilestoneId int64     `json:"milestoneId,required" db:"milestone_id"`
	Validator   string    `json:"validator,required" db:"validator"`
	CreatedAt   time.Time `json:"createdAt,required" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt,required" db:"updated_at"`
}
