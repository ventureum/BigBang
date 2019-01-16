package redeem_block_info_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"time"
)

type RedeemBlockInfoRecord struct {
	RedeemBlock                  int64                          `db:"redeem_block"`
	TotalEnrolledMilestonePoints feed_attributes.MilestonePoint `db:"total_enrolled_milestone_points"`
	TokenPool                    feed_attributes.MilestonePoint `db:"token_pool"`
	CreatedAt                    time.Time                      `db:"created_at"`
	UpdatedAt                    time.Time                      `db:"updated_at"`
}
