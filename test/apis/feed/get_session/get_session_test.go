package get_session

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/feed/session_record_config"
  "BigBang/cmd/lambda/feed/get_session/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_session_config.Request
    response lambda_get_session_config.Response
    err    error
  }{
    {
      request: lambda_get_session_config.Request {
        PostHash: test_constants.PostHash2,
      },
      response: lambda_get_session_config.Response {
        Session: &session_record_config.SessionRecordResult{
          Actor: test_constants.Actor1,
          PostHash: test_constants.PostHash2,
          StartTime: test_constants.SessionStartTime2,
          EndTime: test_constants.SessionEndTime2,
          Content: &test_constants.SessionContent2,
        },
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    responseMessageGetSession := api.SendPost(test.request, api.GetSessionAlphaEndingPoint)
    var responseGetSession lambda_get_session_config.Response
    mapstructure.Decode(*responseMessageGetSession, &responseGetSession)
    responseGetSession.Session = &session_record_config.SessionRecordResult{}
    mapstructure.Decode((*responseMessageGetSession).(map[string]interface{})["session"], responseGetSession.Session)
    assert.Equal(t, test.response.Ok, responseGetSession.Ok)
    assert.Equal(t, test.response.Session.Actor, responseGetSession.Session.Actor)
    assert.Equal(t, test.response.Session.PostHash, responseGetSession.Session.PostHash)
    assert.Equal(t, test.response.Session.StartTime, responseGetSession.Session.StartTime)
    assert.Equal(t, test.response.Session.EndTime, responseGetSession.Session.EndTime)
  }
}
