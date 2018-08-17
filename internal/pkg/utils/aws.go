package utils

import (
  "github.com/aws/aws-sdk-go/aws/session"
  "log"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)


func CreateAwsSession()  *session.Session{
  sess, err := session.NewSession()
  if err != nil {
    log.Fatal("Failed to create aws new session")
  }
  return sess
}

func CreateAwsDynamoDBClient() *dynamodb.DynamoDB {
  sess := CreateAwsSession()
  return dynamodb.New(sess)
}
