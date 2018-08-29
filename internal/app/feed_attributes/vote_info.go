package feed_attributes

type VoteInfoType string

type VoteCountInfo struct {
  DownVoteCount int64 `json:"downvoteCount"`
  UpVoteCount int64 `json:"upvoteCount"`
  TotalVoteCount int64 `json:"totalVoteCount"`
}

type VoteInfo struct {
  Actor string `json:"actor"`
  PostHash string `json:"postHash"`
  Reputations Reputation `json:"reputations"`
  Cost Reputation  `json:"cost"`
  PostVoteCountInfo *VoteCountInfo `json:"postVoteCountInfo"`
  RequestorVoteCountInfo *VoteCountInfo `json:"requestorVoteCountInfo"`
}
