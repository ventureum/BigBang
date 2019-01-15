package principal_proxy_votes_config

import (
	"BigBang/internal/pkg/utils"
	"fmt"
	"time"
)

type PrincipalProxyVotesRecord struct {
	ID             string    `json:"id" db:"id"`
	Actor          string    `json:"actor" db:"actor"`
	ProjectId      string    `json:"projectId" db:"project_id"`
	Proxy          string    `json:"proxy" db:"proxy"`
	BlockTimestamp int64     `json:"block_timestamp" db:"block_timestamp"`
	VotesInPercent int64     `json:"votesInPercent" db:"votes_in_percent"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

func (principalProxyVotesRecord *PrincipalProxyVotesRecord) GenerateID() {
	idStr := fmt.Sprintf("%d:%s:%s:%s", principalProxyVotesRecord.BlockTimestamp, principalProxyVotesRecord.Actor, principalProxyVotesRecord.ProjectId, principalProxyVotesRecord.Proxy)
	principalProxyVotesRecord.ID = idStr
}

func (principalProxyVotesRecord *PrincipalProxyVotesRecord) EncodeID() string {
	idStr := fmt.Sprintf("%s:%s:%s", principalProxyVotesRecord.Actor, principalProxyVotesRecord.ProjectId, principalProxyVotesRecord.Proxy)
	return utils.Base64EncodeIdByInt64AndStr(principalProxyVotesRecord.BlockTimestamp, idStr)
}

func GenerateEncodedPrincipalProxyVotesRecordID(actor string, projectId string, proxy string, blockTimestamp int64) string {
	idStr := fmt.Sprintf("%s:%s:%s", actor, projectId, proxy)
	return utils.Base64EncodeIdByInt64AndStr(blockTimestamp, idStr)
}

func GeneratePrincipalProxyVotesRecordID(actor string, projectId string, proxy string, blockTimestamp int64) string {
	idStr := fmt.Sprintf("%d:%s:%s:%s", blockTimestamp, actor, projectId, proxy)
	return idStr
}
