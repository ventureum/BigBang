package tcr_attributes

type RatingVote struct {
  Voter string `json:"voter,required"`
  Rating int64 `json:"rating,required"`
  Weight int64 `json:"weight,required"`
  BlockTimestamp  int64  `json:"blockTimestamp,required"`
}
