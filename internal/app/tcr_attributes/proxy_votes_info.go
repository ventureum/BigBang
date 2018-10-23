package tcr_attributes

type ProxyVotesInfo struct {
  Actor string `json:"actor,required"`
  ProjectId  string   `json:"projectId,required"`
  AvailableDelegateVotes int64     `json:"availableDelegateVotes,required"`
  ReceivedDelegateVotes  int64     `json:"receivedDelegateVotes,required"`
  ProxyVotesList *[]ProxyVotes `json:"proxyVotesList,required"`
}
