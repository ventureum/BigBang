package client_config

import (
  "log"
  _ "github.com/lib/pq"
  "fmt"
  "github.com/jmoiron/sqlx"
  "os"
)

type PostgresFeedClient struct {
  C *sqlx.DB
  Tx *sqlx.Tx
}

func ConnectPostgresClient() *PostgresFeedClient {
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbName := os.Getenv("DB_NAME")
  dbHost := os.Getenv("DB_HOST")
  dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
    dbHost, dbUser, dbPassword, dbName)
  db, err := sqlx.Connect("postgres", dbinfo)
  if err != nil {
    log.Panicf("Failed to connect postgres with error: %+v\n", err)
  }
  log.Println("Connected to Postgres Client")
  return &PostgresFeedClient{C: db}
}

func (postgresFeedClient *PostgresFeedClient) Begin() {
  tx, err := postgresFeedClient.C.Beginx()
  if err != nil {
    log.Panicf("Failed to Begin TX with error: %+v\n", err)
  }
  postgresFeedClient.Tx = tx
}

func (postgresFeedClient *PostgresFeedClient) Commit() {
  err := postgresFeedClient.Tx.Commit()
  if err != nil {
    log.Panicf("Failed to Commit with error: %+v\n", err)
  }
}

func (postgresFeedClient *PostgresFeedClient) Close() {
  err := postgresFeedClient.C.Close()
  if err != nil {
    log.Panicf("Failed to Close with error: %+v\n", err)
  }
}

func (postgresFeedClient *PostgresFeedClient) CreateTable(schema string, tableName string) {
  tx := postgresFeedClient.C.MustBegin()
  _ , err := tx.Exec(schema)
  if err != nil {
    log.Panicf("Failed to execute creating Table %s with error: %+v\n", tableName, err)
  }
  tx.Commit()
  log.Printf("Table %s has been created\n", tableName)
}

func (postgresFeedClient *PostgresFeedClient) DeleteTable(tableName string) {
  command := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
  tx := postgresFeedClient.C.MustBegin()
  res, err := tx.Exec(command)
  if err != nil {
    log.Panicf("Failed to execute deleting Table %s with error: %+v\n", tableName, err)
  }
  tx.Commit()
  affected, _ := res.RowsAffected()
  log.Printf("Table %s has been deleted with %v rows affected\n", tableName,  affected)
}

func (postgresFeedClient *PostgresFeedClient) CreateTimestampTrigger() {
  _, err := postgresFeedClient.C.Exec(TRIGGER_SET_TIMESTAMP_COMMAND)
  if err != nil {
    log.Panicf("Failed to create timestamp trigger with error: %+v\n", err)
  }
}

func (postgresFeedClient *PostgresFeedClient) RegisterTimestampTrigger(tableName string) {
  command := fmt.Sprintf(REGISTER_TIMESTAMP_TRIGGER_COMMAND, tableName)
  _, err := postgresFeedClient.C.Exec(command)
  if err != nil {
    log.Panicf("Failed to register timestamp trigger for Table %s with error: %+v\n", tableName, err)
  }
}

func (postgresFeedClient *PostgresFeedClient) LoadUuidExtension() {
  _, err := postgresFeedClient.C.Exec(LOAD_UUID_EXTENSION)
  if err != nil {
    log.Panicf("Failed to load uuid extension with error: %+v\n", err)
  }
}

func (postgresFeedClient *PostgresFeedClient) LoadVoteTypeEnum() {
  _, err := postgresFeedClient.C.Exec(LOAD_VOTE_TYPE_ENUM)
  if err != nil {
    log.Fatalf("Failed to load vote type enum with error: %+v\n", err)
  }
}

func (postgresFeedClient *PostgresFeedClient) LoadActorTypeEnum() {
  _, err := postgresFeedClient.C.Exec(LOAD_ACTOR_TYPE_ENUM)
  if err != nil {
    log.Fatalf("Failed to load actor type enum with error: %+v\n", err)
  }
}
