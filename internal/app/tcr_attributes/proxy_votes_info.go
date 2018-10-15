package tcr_attributes

type ProxyVotesInfo struct {
  Actor string `json:"actor,required"`
  ProjectId  string   `json:"projectId,required"`
  ProxyVotesList *[]ProxyVotes `json:"proxyVotesList,required"`
}
