package post_rewards_record_config

import (
  "time"
)


type PostRewardsRecord struct {
  PostHash string  `json:"postHash,required" db:"post_hash"`
  Actor string `json:"actor,required" db:"actor"`
  PostType    string         `json:"postType,required" db:"post_type"`
  DeltaFuel        int64                    `json:"deltaFuel,required" db:"delta_fuel"`
  DeltaReputation  int64                    `json:"deltaReputation,required" db:"delta_reputation"`
  DeltaMilestonePoints int64                `json:"deltaMilestonePoints,required" db:"delta_milestone_points"`
  WithdrawableMPs int64 `json:"withdrawableMPs,required" db:"withdrawable_mps"`
  CreatedAt time.Time `json:"createdAt,required" db:"created_at"`
  UpdatedAt time.Time `json:"updatedAt,required" db:"updated_at"`
}
