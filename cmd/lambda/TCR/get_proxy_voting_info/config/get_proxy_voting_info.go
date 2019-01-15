package lambda_get_proxy_voting_info_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
	"BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor     string `json:"actor,required"`
	ProjectId string `json:"projectId,required"`
	Limit     int64  `json:"limit,required"`
	Cursor    string `json:"cursor,omitempty"`
}

type ResponseData struct {
	ProxyVotingInfo *tcr_attributes.ProxyVotingInfo `json:"proxyVotingInfo,omitempty"`
	NextCursor      string                          `json:"nextCursor,omitempty"`
}

type Response struct {
	ResponseData *ResponseData           `json:"responseData,omitempty"`
	Ok           bool                    `json:"ok"`
	Message      *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.ResponseData = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	projectId := request.Body.ProjectId
	limit := request.Body.Limit

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}

	actorProfileRecordExecutor.VerifyActorExistingTx(actor)
	projectExecutor.VerifyProjectRecordExistingTx(projectId)

	principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}

	existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(actor, projectId)
	if !existing {
		response.ResponseData = &ResponseData{
			NextCursor: "",
			ProxyVotingInfo: &tcr_attributes.ProxyVotingInfo{
				Actor:                  actor,
				ProjectId:              projectId,
				AvailableDelegateVotes: 0,
				ReceivedDelegateVotes:  0,
				ProxyVotingList:        nil,
			},
		}
	} else {

		cursorStr := request.Body.Cursor
		var cursor string
		if cursorStr != "" {
			cursor = utils.Base64DecodeToString(cursorStr)
		}

		actorDelegateVotesAccount := actorDelegateVotesAccountExecutor.GetActorDelegateVotesAccountRecordTx(actor, projectId)
		principalProxyVotesRecordList := principalProxyVotesExecutor.GetPrincipalProxyVotesRecordListByCursorTx(
			actor, projectId, cursor, limit+1)

		response.ResponseData = &ResponseData{
			NextCursor: "",
			ProxyVotingInfo: &tcr_attributes.ProxyVotingInfo{
				Actor:                  actor,
				ProjectId:              projectId,
				AvailableDelegateVotes: actorDelegateVotesAccount.AvailableDelegateVotes,
				ReceivedDelegateVotes:  actorDelegateVotesAccount.ReceivedDelegateVotes,
			},
		}

		var proxyVotesList []tcr_attributes.ProxyVoting
		for index, principalProxyVotesRecord := range *principalProxyVotesRecordList {
			if index < int(limit) {
				ratingVote := tcr_attributes.ProxyVoting{
					Proxy:          principalProxyVotesRecord.Proxy,
					BlockTimestamp: principalProxyVotesRecord.BlockTimestamp,
					VotesInPercent: principalProxyVotesRecord.VotesInPercent,
				}
				proxyVotesList = append(proxyVotesList, ratingVote)
			} else {
				response.ResponseData.NextCursor = principalProxyVotesRecord.EncodeID()
			}
		}

		response.ResponseData.ProxyVotingInfo.ProxyVotingList = &proxyVotesList

		if cursorStr == "" {
			log.Printf("ProxyVotingInfo is loaded for first query with Actor %s, ProjectId %s and limit %d\n",
				actor, projectId, limit)
		} else {
			log.Printf("ProxyVotingInfo is loaded for query with Actor %s, ProjectId %s, cursor %s and limit %d\n",
				actor, projectId, cursorStr, limit)
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
