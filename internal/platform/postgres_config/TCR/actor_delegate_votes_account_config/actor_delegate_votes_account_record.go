package actor_delegate_votes_account_config

import (
	"time"
)

type ActorDelegateVotesAccountRecord struct {
	Actor                  string    `json:"actor" db:"actor"`
	ProjectId              string    `json:"projectId" db:"project_id"`
	AvailableDelegateVotes int64     `json:"availableDelegateVotes" db:"available_delegate_votes"`
	ReceivedDelegateVotes  int64     `json:"receivedDelegateVotes" db:"received_delegate_votes"`
	CreatedAt              time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt              time.Time `json:"updatedAt" db:"updated_at"`
}
