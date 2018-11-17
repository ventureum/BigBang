package attach_session

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/attach_session/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_attach_session_config.Request
    response lambda_attach_session_config.Response
    err    error
  }{
    {
      request: lambda_attach_session_config.Request {
        Actor: test_constants.Actor1,
        PostHash: test_constants.PostHash1,
        StartTime: test_constants.SessionStartTime1,
        EndTime: test_constants.SessionEndTime1,
        Content: test_constants.SessionContent1,
      },
      response: lambda_attach_session_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_attach_session_config.Request {
        Actor: test_constants.Actor1,
        PostHash: test_constants.PostHash2,
        StartTime: test_constants.SessionStartTime2,
        EndTime: test_constants.SessionEndTime2,
        Content: test_constants.SessionContent2,
      },
      response: lambda_attach_session_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    responseMessageAttachSession := api.SendPost(test.request, api.AttachSessionAlphaEndingPoint)
    var responseAttachSession lambda_attach_session_config.Response
    mapstructure.Decode(*responseMessageAttachSession, &responseAttachSession)
    assert.Equal(t, test.response, responseAttachSession)
  }
}
