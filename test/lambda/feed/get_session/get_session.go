package get_session_test

import (
	"BigBang/cmd/lambda/feed/get_session/config"
	"BigBang/internal/platform/postgres_config/feed/session_record_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_session_config.Request
		response lambda_get_session_config.Response
		err      error
	}{
		{
			request: lambda_get_session_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_session_config.RequestContent{
					PostHash: test_constants.PostHash2,
				},
			},
			response: lambda_get_session_config.Response{
				Session: &session_record_config.SessionRecordResult{
					Actor:     test_constants.Actor1,
					PostHash:  test_constants.PostHash2,
					StartTime: test_constants.SessionStartTime2,
					EndTime:   test_constants.SessionEndTime2,
					Content:   &test_constants.SessionContent2,
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_session_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Session.Actor, result.Session.Actor)
		assert.Equal(t, test.response.Session.PostHash, result.Session.PostHash)
		assert.Equal(t, test.response.Session.StartTime, result.Session.StartTime)
		assert.Equal(t, test.response.Session.EndTime, result.Session.EndTime)
	}
}
