package feed_attributes

type VoteInfoType string

const VoteInfoForVoter VoteInfoType = "Voter"
const VoteInfoForPost VoteInfoType = "Post"

type VoteCountInfo struct {
  DownVoteCount int64 `json:"downvoteCount"`
  UpVoteCount int64 `json:"upvoteCount"`
  TotalVoteCount int64 `json:"totalVoteCount"`
}

type VoteInfo struct {
  For VoteInfoType `json:"postHash"`
  Actor string `json:"actor"`
  PostHash string `json:"postHash"`
  Reputations Reputation `json:"reputations"`
  Cost Reputation  `json:"cost"`
  VoteCountInfo
}
