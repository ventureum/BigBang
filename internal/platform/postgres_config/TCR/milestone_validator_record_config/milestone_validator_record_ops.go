package milestone_validator_record_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"database/sql"
	"log"
)

type MilestoneValidatorRecordExecutor struct {
	client_config.PostgresBigBangClient
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) CreateMilestoneValidatorRecordTable() {
	milestoneValidatorRecordExecutor.CreateTimestampTrigger()
	milestoneValidatorRecordExecutor.CreateTable(TABLE_SCHEMA_FOR_MILESTONE_VALIDATOR_RECORD, TABLE_NAME_FOR_MILESTONE_VALIDATOR_RECORD)
	milestoneValidatorRecordExecutor.RegisterTimestampTrigger(TABLE_NAME_FOR_MILESTONE_VALIDATOR_RECORD)
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) DeleteMilestoneValidatorRecordTable() {
	milestoneValidatorRecordExecutor.DeleteTable(TABLE_NAME_FOR_MILESTONE_VALIDATOR_RECORD)
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) UpsertMilestoneValidatorRecordTx(milestoneValidatorRecord *MilestoneValidatorRecord) {
	_, err := milestoneValidatorRecordExecutor.Tx.NamedExec(UPSERT_MILESTONE_VALIDATOR_RECORD_COMMAND, milestoneValidatorRecord)
	if err != nil {
		log.Panicf("Failed to upsert milestone validator record: %+v with error: %+v\n", milestoneValidatorRecord, err)
	}
	log.Printf("Sucessfully upserted milestone validator record for projectId %s and milestoneId %s\n",
		milestoneValidatorRecord.ProjectId, milestoneValidatorRecord.MilestoneId)
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) DeleteMilestoneValidatorRecordsByProjectIdAndMilestoneIdTx(
	projectId string, milestoneId int64) {
	_, err := milestoneValidatorRecordExecutor.Tx.Exec(
		DELETE_ALL_MILESTONE_VALIDATOR_RECORDS_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND, projectId, milestoneId)
	if err != nil {
		log.Panicf("Failed to delete milestone validator records for projectId %s and milestoneId %d with error: %+v\n", projectId, milestoneId, err)
	}
	log.Printf("Sucessfully deleted milestone validator records for projectId %s and milestoneId %d\n", projectId, milestoneId)
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) DeleteMilestoneValidatorRecordByIDsAndValidatorTx(
	projectId string, milestoneId int64, validator string) {
	_, err := milestoneValidatorRecordExecutor.Tx.Exec(
		DELETE_MILESTONE_VALIDATOR_RECORD_BY_IDS_AND_VALIDATOR_COMMAND, projectId, milestoneId, validator)
	if err != nil {
		log.Panicf("Failed to delete milestone validator record for projectId %s, milestoneId %d and validator %s with error: %+v\n",
			projectId, milestoneId, validator, err)
	}
	log.Printf("Sucessfully deleted milestone validator record for projectId %s, milestoneId %d and validator %s\n",
		projectId, milestoneId, validator)
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) GetMilestoneValidatorListByProjectIdAndMilestoneIdTx(
	projectId string, milestoneId int64) *[]string {
	var milestoneValidatorList []string
	err := milestoneValidatorRecordExecutor.Tx.Select(
		&milestoneValidatorList, QUERY_MILESTONE_VALIDATOR_LIST_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND, projectId, milestoneId)
	if err != nil && err != sql.ErrNoRows {
		log.Panicf(
			"Failed to get milestone validator list for projectId %s and milestoneId %d with error: %+v\n", projectId, milestoneId, err)
	}
	return &milestoneValidatorList
}

func (milestoneValidatorRecordExecutor *MilestoneValidatorRecordExecutor) CheckMilestoneValidatorRecordExistingTx(
	projectId string, milestoneId int64, validator string) bool {
	var existing bool
	err := milestoneValidatorRecordExecutor.Tx.Get(&existing, VERIFY_MILESTONE_VALIDATOR_EXISTING_COMMAND, projectId, milestoneId, validator)
	if err != nil {
		errInfo := error_config.MatchError(err, "milestoneId", milestoneId, error_config.MilestoneRecordLocation)
		log.Printf("Failed to verify milestone validator record existing for projectId %s , milestoneId %d and validator %s with error: %+v\n",
			projectId, milestoneId, validator, err)
		errInfo.ErrorData["projectId"] = projectId
		errInfo.ErrorData["validator"] = validator
		log.Panicln(errInfo.Marshal())
	}
	return existing
}
