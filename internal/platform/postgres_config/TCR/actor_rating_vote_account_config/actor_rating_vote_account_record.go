package actor_rating_vote_account_config
import (
  "time"
)

type ActorRatingVoteAccountRecord struct {
  Actor string  `json:"actor" db:"actor"`
  AvailableRatingVotes int64 `json:"availableRatingVotes" db:"available_rating_votes"`
  ReceivedRatingVotes int64 `json:"receivedRatingVotes" db:"received_rating_votes"`
  CreatedAt time.Time `json:"createdAt" db:"created_at"`
  UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
