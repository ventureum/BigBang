package tcr_attributes

type RatingVote struct {
	Voter          string `json:"voter,required" db:"voter"`
	BlockTimestamp int64  `json:"blockTimestamp,required" db:"block_timestamp"`
	Rating         int64  `json:"rating,required" db:"rating"`
	Weight         int64  `json:"weight,required" db:"weight"`
}
