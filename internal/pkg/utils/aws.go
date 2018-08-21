package utils

import (
  "github.com/aws/aws-sdk-go/aws/session"
  "log"
)


func CreateAwsSession()  *session.Session{
  sess, err := session.NewSession()
  if err != nil {
    log.Fatal("Failed to create aws new session")
  }
  return sess
}
