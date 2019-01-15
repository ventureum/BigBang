package attach_session_test

import (
	"BigBang/cmd/lambda/feed/attach_session/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_attach_session_config.Request
		response lambda_attach_session_config.Response
		err      error
	}{
		{
			request: lambda_attach_session_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_attach_session_config.RequestContent{
					Actor:     test_constants.Actor1,
					PostHash:  test_constants.PostHash1,
					StartTime: test_constants.SessionStartTime1,
					EndTime:   test_constants.SessionEndTime1,
					Content:   test_constants.SessionContent1,
				},
			},
			response: lambda_attach_session_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_attach_session_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_attach_session_config.RequestContent{
					Actor:     test_constants.Actor1,
					PostHash:  test_constants.PostHash2,
					StartTime: test_constants.SessionStartTime2,
					EndTime:   test_constants.SessionEndTime2,
					Content:   test_constants.SessionContent2,
				},
			},
			response: lambda_attach_session_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_attach_session_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
