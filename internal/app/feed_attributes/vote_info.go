package feed_attributes

type VoteInfo struct {
  PostHash string `json:"postHash"`
  Actor string `json:"actor"`
  Reputations Reputation `json:"reputations"`
  Cost Reputation  `json:"cost"`
  DownVoteCount int64 `json:"downvoteCount"`
  UpVoteCount int64 `json:"upvoteCount"`
}
