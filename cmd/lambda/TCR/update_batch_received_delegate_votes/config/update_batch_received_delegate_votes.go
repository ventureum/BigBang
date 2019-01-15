package lambda_update_batch_received_delegate_votes_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
)

type Request struct {
	PrincipalId string      `json:"principalId,required"`
	Body        RequestBody `json:"body,required"`
}

type RequestContent struct {
	Actor                      string `json:"actor,required"`
	ProjectId                  string `json:"projectId,required"`
	ReceivedDelegateVotesDelta int64  `json:"receivedDelegateVotesDelta,required"`
}

type RequestBody struct {
	RequestList []RequestContent `json:"requestList,required"`
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

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	requestList := request.Body.RequestList

	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}

	for _, singleRequest := range requestList {
		actorProfileRecordExecutor.VerifyActorExistingTx(singleRequest.Actor)
		projectExecutor.VerifyProjectRecordExistingTx(singleRequest.ProjectId)
	}

	for _, singleRequest := range requestList {
		actor := singleRequest.Actor
		projectId := singleRequest.ProjectId
		receivedDelegateVotesDelta := singleRequest.ReceivedDelegateVotesDelta
		existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(actor, projectId)

		if !existing {
			actorDelegateVotesAccountExecutor.UpsertActorDelegateVotesAccountRecordTx(&actor_delegate_votes_account_config.ActorDelegateVotesAccountRecord{
				Actor:                  actor,
				ProjectId:              projectId,
				AvailableDelegateVotes: 0,
				ReceivedDelegateVotes:  receivedDelegateVotesDelta,
			})
		} else {
			actorDelegateVotesAccountExecutor.UpdateReceivedDelegateVotesTx(actor, projectId, receivedDelegateVotesDelta)
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
