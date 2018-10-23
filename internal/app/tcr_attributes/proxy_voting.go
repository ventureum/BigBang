package tcr_attributes

type ProxyVoting struct {
  Proxy          string `json:"proxy,required"`
  BlockTimestamp int64  `json:"blockTimestamp,required"`
  VotesInPercent int64  `json:"votesInPercent,required"`
}
