package tcr_attributes

type ProxyVotes struct {
  Proxy          string `json:"proxy,required"`
  BlockTimestamp int64  `json:"blockTimestamp,required"`
  VotesInPercent int64  `json:"votes,required"`
}
