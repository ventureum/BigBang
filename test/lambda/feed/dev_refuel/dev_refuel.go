package dev_refuel_test

import (
	"BigBang/cmd/lambda/feed/dev_refuel/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_dev_refuel_config.Request
		response lambda_dev_refuel_config.Response
		err      error
	}{
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor1,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor2,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor3,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor4,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor5,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor6,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_dev_refuel_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_dev_refuel_config.RequestContent{
					Actor:           test_constants.Actor7,
					Fuel:            100,
					Reputation:      100,
					MilestonePoints: 100,
				},
			},
			response: lambda_dev_refuel_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_dev_refuel_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
