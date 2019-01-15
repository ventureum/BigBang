package lambda_update_received_delegate_votes_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor                      string `json:"actor,required"`
	ProjectId                  string `json:"projectId,required"`
	ReceivedDelegateVotesDelta int64  `json:"receivedDelegateVotesDelta,required"`
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

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(actor)
	projectExecutor.VerifyProjectRecordExistingTx(projectId)

	existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(actor, projectId)

	if !existing {
		actorDelegateVotesAccountExecutor.UpsertActorDelegateVotesAccountRecordTx(&actor_delegate_votes_account_config.ActorDelegateVotesAccountRecord{
			Actor:                  actor,
			ProjectId:              projectId,
			AvailableDelegateVotes: 0,
			ReceivedDelegateVotes:  request.Body.ReceivedDelegateVotesDelta,
		})
	} else {
		actorDelegateVotesAccountExecutor.UpdateReceivedDelegateVotesTx(actor, projectId, request.Body.ReceivedDelegateVotesDelta)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
