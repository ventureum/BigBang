package actor_profile_record_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type ActorProfileRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) CreateActorProfileRecordTable() {
	actorProfileRecordExecutor.CreateTimestampTrigger()
	actorProfileRecordExecutor.CreateTable(
		TABLE_SCHEMA_FOR_ACTOR_PROFILE_RECORD, TABLE_NAME_FOR_ACTOR_PROFILE_RECORD)
	actorProfileRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_ACTOR_PROFILE_RECORD)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecordTable() {
	actorProfileRecordExecutor.DeleteTable(TABLE_NAME_FOR_ACTOR_PROFILE_RECORD)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) UpsertActorProfileRecordTx(
	actorProfileRecord *ActorProfileRecord) bool {
	var inserted sql.NullBool
	rows, err := actorProfileRecordExecutor.Tx.NamedQuery(
		UPSERT_ACTOR_PROFILE_RECORD_COMMAND, actorProfileRecord)
	for rows.Next() {
		err = rows.Scan(&inserted)
	}
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actorProfileRecord.Actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to upsert profile record: %+v with error: %+v\n", actorProfileRecord, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully upserted profile record for actor %s\n", actorProfileRecord.Actor)
	return inserted.Bool
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeleteActorProfileRecordsTx(actor string) {
	_, err := actorProfileRecordExecutor.Tx.Exec(DELETE_ACTOR_PROFILE_RECORD_COMMAND, actor)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to delete profile records for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deleted profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) DeactivateActorProfileRecordsTx(actor string) {
	_, err := actorProfileRecordExecutor.Tx.Exec(DEACTIVATE_ACTOR_PROFILE_RECORD_COMMAND, actor)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to deactivate profile records for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully deactivated profile records for actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorProfileRecordTx(actor string) *ActorProfileRecord {
	var actorProfileRecord ActorProfileRecord
	err := actorProfileRecordExecutor.Tx.Get(&actorProfileRecord, QUERY_ACTOR_PROFILE_RECORD_COMMAND, actor)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to get actor profile record for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}
	return &actorProfileRecord
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) SetActorPrivateKeyTx(actor string, privateKey string) {
	_, err := actorProfileRecordExecutor.Tx.Exec(UPDATE_ACTOR_PRIVATE_KEY_COMMAND, actor, privateKey)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to update private key for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}
	log.Printf("Sucessfully updated private keyfor actor %s\n", actor)
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorUuidFromPublicKeyTx(publicKey string) string {
	var actor string
	err := actorProfileRecordExecutor.Tx.Get(&actor, QUERY_ACTOR_UUID_FROM_PRIVATE_KEY_COMMAND, publicKey)

	if err == sql.ErrNoRows {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoActorExistingForPublicKey,
			ErrorData: map[string]interface{}{
				"publicKey": publicKey,
			},
			ErrorLocation: error_config.ProfileAccountLocation,
		}
		log.Printf("No actor for publicKey %s", publicKey)
		log.Panicln(errorInfo.Marshal())
	}

	if err != nil {
		errInfo := error_config.MatchError(err, "publicKey", publicKey, error_config.ProfileAccountLocation)
		log.Printf("Failed to get actor uuid for publicKey %s with error: %+v\n", publicKey, err)
		log.Panicln(errInfo.Marshal())
	}

	return actor
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) CheckActorExistingTx(actor string) bool {
	var existing bool
	err := actorProfileRecordExecutor.Tx.Get(&existing, VERIFY_ACTOR_EXISTING_COMMAND, actor)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to check actor existing for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}
	return existing
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) VerifyActorExistingTx(actor string) {
	existing := actorProfileRecordExecutor.CheckActorExistingTx(actor)
	if !existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoActorExisting,
			ErrorData: map[string]interface{}{
				"actor": actor,
			},
			ErrorLocation: error_config.ProfileAccountLocation,
		}
		log.Printf("No Actor Profile Acount for actor %s", actor)
		log.Panicln(errorInfo.Marshal())
	}
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) CheckActorTypeTx(actor string, actorType feed_attributes.ActorType) bool {
	var match bool
	err := actorProfileRecordExecutor.Tx.Get(&match, VERIFY_ACTOR_TYPE_COMMAND, actor, actorType)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		errInfo.ErrorData["actorType"] = actorType
		log.Printf("Failed to verify ActorType %s for actor %s with error: %+v\n", actorType, actor, err)
		log.Panicln(errInfo.Marshal())
	}
	return match
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorTypeTx(actor string) feed_attributes.ActorType {
	var actorType feed_attributes.ActorType
	err := actorProfileRecordExecutor.Tx.Get(&actorType, QUERY_ACTOR_TYPE_COMMAND, actor)
	if err == sql.ErrNoRows {
		return feed_attributes.ActorType("")
	}
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		errInfo.ErrorData["actorType"] = actorType
		log.Printf("Failed to get ActorType for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}
	return actorType
}

func (actorProfileRecordExecutor *ActorProfileRecordExecutor) GetActorPrivateKeyTx(actor string) string {
	var privateKey string
	err := actorProfileRecordExecutor.C.Get(&privateKey, QUERY_ACTOR_PRIVATE_KEY_COMMAND, actor)
	if err != nil {
		errInfo := error_config.MatchError(err, "actor", actor, error_config.ProfileAccountLocation)
		log.Printf("Failed to get actor private key for actor %s with error: %+v\n", actor, err)
		log.Panicln(errInfo.Marshal())
	}

	if privateKey == "" {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoPrivateKeyExistingForActor,
			ErrorData: map[string]interface{}{
				"actor": actor,
			},
			ErrorLocation: error_config.ProfileAccountLocation,
		}
		log.Printf("No private key existing for actor %s", actor)
		log.Panicln(errorInfo.Marshal())
	}
	return privateKey
}
