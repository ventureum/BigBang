package actor_delegate_votes_account_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type ActorDelegateVotesAccountExecutor struct {
	client_config.PostgresBigBangClient
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) CreateActorRatingVoteAccountTable() {
	actorDelegateVotesAccountExecutor.CreateTimestampTrigger()
	actorDelegateVotesAccountExecutor.CreateTable(TABLE_SCHEMA_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT, TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT)
	actorDelegateVotesAccountExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) DeleteActorRatingVoteAccountTable() {
	actorDelegateVotesAccountExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_DELEGATE_VOTES_ACCOUNT)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) UpsertActorDelegateVotesAccountRecordTx(
	actorDelegateVotesAccountRecord *ActorDelegateVotesAccountRecord) {
	_, err := actorDelegateVotesAccountExecutor.Tx.NamedExec(UPSERT_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actorDelegateVotesAccountRecord)

	if err != nil {
		errorInfo := error_config.MatchError(err, "actor", actorDelegateVotesAccountRecord.Actor, error_config.ActorDelegateVotesAccountRecordLocation)
		log.Printf("Failed to upsert Actor Delegate Votes Account Record for actor %s with error: %+v\n", actorDelegateVotesAccountRecord.Actor, err)
		log.Panic(errorInfo.Marshal())
	}

	log.Printf("Sucessfully upserted Actor Delegate Votes Account Record for actor %s\n", actorDelegateVotesAccountRecord.Actor)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) DeleteActorDelegateVotesAccountRecordTx(actor string, projectId string) {
	_, err := actorDelegateVotesAccountExecutor.Tx.Exec(DELETE_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actor, projectId)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
		errInfo.ErrorData["projectId"] = projectId
		log.Printf("Failed to delete Actor Delegate Votes Account Record for actor %s and projectId %s with error: %+v\n",
			actor, projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted Actor Delegate Votes Account Record for actor %s and projectId %s\n", actor, projectId)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) GetActorDelegateVotesAccountRecordTx(
	actor string, projectId string) *ActorDelegateVotesAccountRecord {
	var actorDelegateVotesAccountRecord ActorDelegateVotesAccountRecord
	err := actorDelegateVotesAccountExecutor.Tx.Get(&actorDelegateVotesAccountRecord, QUERY_ACTOR_DELEGATE_VOTES_ACCOUNT_RECORD_COMMAND, actor, projectId)
	if err != nil && err != sql.ErrNoRows {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
		errInfo.ErrorData["projectId"] = projectId
		log.Printf("Failed to get Actor Delegate Votes Account Record for actor %s and projectId %s with error: %+v\n",
			actor, projectId, err)
		log.Panicln(errInfo.Marshal())
	}

	if err == sql.ErrNoRows {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoActorExisting,
			ErrorData: map[string]interface{}{
				"actor":     actor,
				"projectId": projectId,
			},
			ErrorLocation: error_config.ActorDelegateVotesAccountRecordLocation,
		}
		log.Printf("No Actor Delegate Votes Account Record for actor %s and projectId %s\n", actor, projectId)
		log.Panicln(errorInfo.Marshal())
	}
	return &actorDelegateVotesAccountRecord
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) UpdateAvailableDelegateVotesTx(actor string, projectId string, availableDelegateVotes int64) {
	_, err := actorDelegateVotesAccountExecutor.Tx.Exec(UPDATE_AVAILABLE_DELEGATE_VOTES_COMMAND, actor, projectId, availableDelegateVotes)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
		errInfo.ErrorData["projectId"] = projectId
		log.Printf("Failed to update Actor Available Delegate Votes for actor %s and projectId %s with error: %+v\n",
			actor, projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully updated Actor Available Delegate Votes for actor %s and projectId %s\n", actor, projectId)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) UpdateReceivedDelegateVotesTx(actor string, projectId string, receivedDelegateVotesDelta int64) {
	_, err := actorDelegateVotesAccountExecutor.Tx.Exec(UPDATE_RECEIVED_DELEGATE_VOTES_COMMAND, actor, projectId, receivedDelegateVotesDelta)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
		errInfo.ErrorData["projectId"] = projectId
		log.Printf("Failed to update Actor Received Delegate Votes for actor %s and projectId %s with error: %+v\n",
			actor, projectId, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully updated Actor Received Delegate Votes for actor %s and projectId %s\n", actor, projectId)
}

func (actorDelegateVotesAccountExecutor *ActorDelegateVotesAccountExecutor) VerifyDelegateVotesAccountExistingTx(
	actor string, projectId string) bool {
	var existing bool
	err := actorDelegateVotesAccountExecutor.Tx.Get(&existing, VERIFY_DELEGATE_VOTES_ACCOUNT_EXISTING_COMMAND, actor, projectId)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ActorDelegateVotesAccountRecordLocation)
		log.Printf("Failed to verify  Delegate Votes Account existing for actor %s and projectId %s with error: %+v\n",
			actor, projectId, err)
		errInfo.ErrorData["projectId"] = projectId
		log.Panicln(errInfo.Marshal())
	}
	return existing
}
