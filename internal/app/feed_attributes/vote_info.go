package feed_attributes

type VoteInfoType string

type VoteCountInfo struct {
	DownVoteCount  int64 `json:"downvoteCount"`
	UpVoteCount    int64 `json:"upvoteCount"`
	TotalVoteCount int64 `json:"totalVoteCount"`
}

type VoteInfo struct {
	Actor                  string         `json:"actor"`
	PostHash               string         `json:"postHash"`
	RewardsInfo            *RewardsInfo   `json:"rewardsInfo"`
	FuelCost               Fuel           `json:"fuelCost"`
	PostVoteCountInfo      *VoteCountInfo `json:"postVoteCountInfo"`
	RequestorVoteCountInfo *VoteCountInfo `json:"requestorVoteCountInfo"`
}
