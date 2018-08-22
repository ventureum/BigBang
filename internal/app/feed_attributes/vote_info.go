package feed_attributes

type VoteInfo struct {
  PostHash string `json:"postHash,omitempty"`
  Actor string `json:"actor,omitempty"`
  Reputations Reputation `json:"reputations,omitempty"`
  Cost Reputation  `json:"cost,omitempty"`
  DownVoteCount int64 `json:"downvoteCount,omitempty"`
  UpVoteCount int64 `json:"upvoteCount,omitempty"`
}
