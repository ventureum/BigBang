package tcr_attributes

type ProxyVotingInfoKey struct {
	Actor     string `json:"actor,required"`
	ProjectId string `json:"projectId,required"`
}

type ProxyVotingInfo struct {
	Actor                  string         `json:"actor,required"`
	ProjectId              string         `json:"projectId,required"`
	AvailableDelegateVotes int64          `json:"availableDelegateVotes,required"`
	ReceivedDelegateVotes  int64          `json:"receivedDelegateVotes,required"`
	ProxyVotingList        *[]ProxyVoting `json:"proxyVotingList,required"`
}
