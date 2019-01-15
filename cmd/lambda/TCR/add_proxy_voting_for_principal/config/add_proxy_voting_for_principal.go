package lambda_add_proxy_voting_for_principal_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
	"BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
	"BigBang/internal/platform/postgres_config/TCR/proxy_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor           string                       `json:"actor,required"`
	ProjectId       string                       `json:"projectId,required"`
	ProxyVotingList []tcr_attributes.ProxyVoting `json:"proxyVotingList,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	actor := request.Body.Actor
	projectId := request.Body.ProjectId
	proxyVotingList := request.Body.ProxyVotingList

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
	actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
	principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(actor)

	proxyToVotesInPercent := map[string]int64{}

	var sumPcts int64
	for _, proxyVoting := range proxyVotingList {
		proxy := proxyVoting.Proxy
		existing := proxyExecutor.VerifyProxyRecordExistingTx(proxyVoting.Proxy)
		if !existing {
			errorInfo := error_config.ErrorInfo{
				ErrorCode: error_config.NoProxyUUIDExisting,
				ErrorData: map[string]interface{}{
					"uuid": proxy,
				},
				ErrorLocation: error_config.ProxyRecordLocation,
			}
			log.Printf("No proxy exists for uuid %s", proxy)
			log.Panicln(errorInfo.Marshal())
		}

		proxyToVotesInPercent[proxy] = proxyVoting.VotesInPercent
		sumPcts += proxyVoting.VotesInPercent
	}

	principalProxyVotesRecords := principalProxyVotesExecutor.GetPrincipalProxyVotesRecordsByActorAndProjectIdTx(actor, projectId)

	for _, principalProxyVotesRecord := range *principalProxyVotesRecords {
		proxy := principalProxyVotesRecord.Proxy
		if _, exist := proxyToVotesInPercent[proxy]; !exist {
			sumPcts += principalProxyVotesRecord.VotesInPercent
		}
	}

	if sumPcts > 100 {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.TotalProxyVotingPercentageExceeding100,
			ErrorData: map[string]interface{}{
				"actor":     actor,
				"projectId": projectId,
			},
		}
		log.Printf("Total Proxy Voting Percentage Exceeds 100 for actor %s and projectId %s", actor, projectId)
		log.Panicln(errorInfo.Marshal())
	}

	existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(actor, projectId)

	if !existing {
		actorDelegateVotesAccountExecutor.UpsertActorDelegateVotesAccountRecordTx(&actor_delegate_votes_account_config.ActorDelegateVotesAccountRecord{
			Actor:                  actor,
			ProjectId:              projectId,
			AvailableDelegateVotes: 0,
			ReceivedDelegateVotes:  0,
		})
	}

	for _, proxyVoting := range proxyVotingList {
		votesInPercent := proxyVoting.VotesInPercent
		proxy := proxyVoting.Proxy

		existing = actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(proxy, projectId)

		if !existing {
			actorDelegateVotesAccountExecutor.UpsertActorDelegateVotesAccountRecordTx(&actor_delegate_votes_account_config.ActorDelegateVotesAccountRecord{
				Actor:                  proxyVoting.Proxy,
				ProjectId:              projectId,
				AvailableDelegateVotes: 0,
				ReceivedDelegateVotes:  0,
			})
		}

		if votesInPercent > 0 {
			principalProxyVotesRecord := &principal_proxy_votes_config.PrincipalProxyVotesRecord{
				Actor:          actor,
				ProjectId:      projectId,
				Proxy:          proxy,
				BlockTimestamp: proxyVoting.BlockTimestamp,
				VotesInPercent: proxyVoting.VotesInPercent,
			}
			principalProxyVotesRecord.GenerateID()
			principalProxyVotesExecutor.UpsertPrincipalProxyVotesRecordTx(principalProxyVotesRecord)
		} else {
			principalProxyVotesExecutor.DeletePrincipalProxyVotesRecordByIDsTx(actor, projectId, proxy)
		}
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
