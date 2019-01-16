package client_config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PostgresBigBangClient struct {
	C  *sqlx.DB
	Tx *sqlx.Tx
}

func ConnectPostgresClient(dbInfo *DBInfo) *PostgresBigBangClient {
	if dbInfo == nil {
		dbInfo = CreateDefaultDBInfo()
	}
	dbInfoStr := dbInfo.ToString()
	db, err := sqlx.Connect("postgres", dbInfoStr)
	if err != nil {
		log.Panicf("Failed to connect postgres with error: %+v\n", err)
	}
	log.Println("Connected to Postgres Client")
	return &PostgresBigBangClient{C: db}
}

func (postgresBigBangClient *PostgresBigBangClient) Begin() {
	tx, err := postgresBigBangClient.C.Beginx()
	if err != nil {
		log.Panicf("Failed to Begin TX with error: %+v\n", err)
	}
	postgresBigBangClient.Tx = tx
}

func (postgresBigBangClient *PostgresBigBangClient) Commit() {
	err := postgresBigBangClient.Tx.Commit()
	if err != nil {
		log.Panicf("Failed to Commit with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) RollBack() {
	err := postgresBigBangClient.Tx.Rollback()
	if err != nil {
		log.Panicf("Failed to Rollback with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) Close() {
	err := postgresBigBangClient.C.Close()
	if err != nil {
		log.Panicf("Failed to Close with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) CreateTable(schema string, tableName string) {
	tx := postgresBigBangClient.C.MustBegin()
	_, err := tx.Exec(schema)
	if err != nil {
		log.Panicf("Failed to execute creating Table %s with error: %+v\n", tableName, err)
	}
	tx.Commit()
	log.Printf("Table %s has been created\n", tableName)
}

func (postgresBigBangClient *PostgresBigBangClient) DeleteTable(tableName string) {
	command := fmt.Sprintf("DROP TABLE IF EXISTS %s cascade;", tableName)
	tx := postgresBigBangClient.C.MustBegin()
	res, err := tx.Exec(command)
	if err != nil {
		log.Panicf("Failed to execute deleting Table %s with error: %+v\n", tableName, err)
	}
	tx.Commit()
	affected, _ := res.RowsAffected()
	log.Printf("Table %s has been deleted with %v rows affected\n", tableName, affected)
}

func (postgresBigBangClient *PostgresBigBangClient) ClearTable(tableName string) {
	command := fmt.Sprintf("TRUNCATE TABLE %s RESTART identity;", tableName)
	tx := postgresBigBangClient.C.MustBegin()
	res, err := tx.Exec(command)
	if err != nil {
		log.Panicf("Failed to execute clear Table %s with error: %+v\n", tableName, err)
	}
	tx.Commit()
	affected, _ := res.RowsAffected()
	log.Printf("Table %s has been cleared with %v rows affected\n", tableName, affected)
}

func (postgresBigBangClient *PostgresBigBangClient) CreateTimestampTrigger() {
	_, err := postgresBigBangClient.C.Exec(TRIGGER_SET_TIMESTAMP_COMMAND)
	if err != nil {
		log.Panicf("Failed to create timestamp trigger with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) RegisterTimestampTrigger(tableName string) {
	command := fmt.Sprintf(REGISTER_TIMESTAMP_TRIGGER_COMMAND, tableName)
	_, err := postgresBigBangClient.C.Exec(command)
	if err != nil {
		log.Panicf("Failed to register timestamp trigger for Table %s with error: %+v\n", tableName, err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) LoadUuidExtension() {
	_, err := postgresBigBangClient.C.Exec(LOAD_UUID_EXTENSION)
	if err != nil {
		log.Panicf("Failed to load uuid extension with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) LoadVoteTypeEnum() {
	_, err := postgresBigBangClient.C.Exec(LOAD_VOTE_TYPE_ENUM)
	if err != nil {
		log.Panicf("Failed to load vote type enum with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) LoadActorTypeEnum() {
	_, err := postgresBigBangClient.C.Exec(LOAD_ACTOR_TYPE_ENUM)
	if err != nil {
		log.Panicf("Failed to load actor type enum with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) LoadActorProfileStatusEnum() {
	_, err := postgresBigBangClient.C.Exec(LOAD_ACTOR_PROFILE_STATUS_ENUM)
	if err != nil {
		log.Panicf("Failed to load actor profile status enum with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) SetIdleInTransactionSessionTimeout(ms int64) {
	command := fmt.Sprintf(SET_IDLE_IN_TX_SESSION_TIMEOUT, ms)
	_, err := postgresBigBangClient.C.Exec(command)
	if err != nil {
		log.Panicf("Failed to Set Idle In Transaction Session Timeout: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) LoadMilestoneStateEnum() {
	_, err := postgresBigBangClient.C.Exec(LOAD_MILESTONE_STATE_ENUM)
	if err != nil {
		log.Panicf("Failed to load milestone state enum with error: %+v\n", err)
	}
}

func (postgresBigBangClient *PostgresBigBangClient) DropMilestoneStateEnum() {
	_, err := postgresBigBangClient.C.Exec(DROP_MILESTONE_STATE_ENUM)
	if err != nil {
		log.Panicf("Failed to drop milestone state enum with error: %+v\n", err)
	}
}
