package tcr_attributes

type ProxyVoting struct {
	Proxy          string `json:"proxy,required" db:"proxy"`
	BlockTimestamp int64  `json:"block_timestamp,required" db:"block_timestamp"`
	VotesInPercent int64  `json:"votesInPercent,required" db:"votes_in_percent"`
}
