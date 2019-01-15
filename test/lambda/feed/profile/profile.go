package profile_test

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/cmd/lambda/feed/profile/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_profile_config.Request
		response lambda_profile_config.Response
		err      error
	}{
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor1,
					UserType:       "ADMIN",
					Username:       test_constants.UserName1,
					PhotoUrl:       "http://123.com",
					TelegramId:     test_constants.TelegramId1,
					PhoneNumber:    "5197290001",
					PublicKey:      test_constants.PublicKey1,
					ProfileContent: test_constants.ProfileContent1,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor2,
					UserType:       "KOL",
					Username:       test_constants.UserName2,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId2,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey2,
					ProfileContent: test_constants.ProfileContent2,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor3,
					UserType:       "KOL",
					Username:       test_constants.UserName3,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId3,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey3,
					ProfileContent: test_constants.ProfileContent3,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor4,
					UserType:       "KOL",
					Username:       test_constants.UserName4,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId4,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey4,
					ProfileContent: test_constants.ProfileContent4,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor5,
					UserType:       "KOL",
					Username:       test_constants.UserName5,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId5,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey5,
					ProfileContent: test_constants.ProfileContent5,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor6,
					UserType:       "KOL",
					Username:       test_constants.UserName6,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId6,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey6,
					ProfileContent: test_constants.ProfileContent6,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor7,
					UserType:       "KOL",
					Username:       test_constants.UserName7,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId7,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey7,
					ProfileContent: test_constants.ProfileContent7,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:          test_constants.Actor8,
					UserType:       "InvalidUserType",
					Username:       test_constants.UserName8,
					PhotoUrl:       "http://567.com",
					TelegramId:     test_constants.TelegramId8,
					PhoneNumber:    "5197290002",
					PublicKey:      test_constants.PublicKey8,
					ProfileContent: test_constants.ProfileContent8,
				},
			},
			response: lambda_profile_config.Response{
				Ok: false,
				Message: &error_config.ErrorInfo{
					ErrorCode: error_config.InvalidActorType,
					ErrorData: error_config.ErrorData{
						"actorType": "InvalidUserType",
					},
					ErrorLocation: error_config.ActorTypeLocation,
				},
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:       test_constants.ProjectAdmin1,
					UserType:    "KOL",
					Username:    test_constants.ProjectAdmin1,
					PhotoUrl:    "http://567.com",
					TelegramId:  test_constants.ProjectAdmin1,
					PhoneNumber: "5197290002",
					PublicKey:   test_constants.ProjectAdmin1,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:       test_constants.ProjectAdmin2,
					UserType:    "KOL",
					Username:    test_constants.ProjectAdmin2,
					PhotoUrl:    "http://567.com",
					TelegramId:  test_constants.ProjectAdmin2,
					PhoneNumber: "5197290002",
					PublicKey:   test_constants.ProjectAdmin2,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:       test_constants.ProjectAdmin3,
					UserType:    "KOL",
					Username:    test_constants.ProjectAdmin3,
					PhotoUrl:    "http://567.com",
					TelegramId:  test_constants.ProjectAdmin3,
					PhoneNumber: "5197290002",
					PublicKey:   test_constants.ProjectAdmin3,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:       test_constants.ProjectAdmin4,
					UserType:    "KOL",
					Username:    test_constants.ProjectAdmin4,
					PhotoUrl:    "http://567.com",
					TelegramId:  test_constants.ProjectAdmin4,
					PhoneNumber: "5197290002",
					PublicKey:   test_constants.ProjectAdmin4,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:       test_constants.ProjectAdmin5,
					UserType:    "KOL",
					Username:    test_constants.ProjectAdmin5,
					PhotoUrl:    "http://567.com",
					TelegramId:  test_constants.ProjectAdmin5,
					PhoneNumber: "5197290002",
					PublicKey:   test_constants.ProjectAdmin5,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_profile_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_profile_config.RequestContent{
					Actor:       test_constants.ProjectAdmin6,
					UserType:    "KOL",
					Username:    test_constants.ProjectAdmin6,
					PhotoUrl:    "http://567.com",
					TelegramId:  test_constants.ProjectAdmin6,
					PhoneNumber: "5197290002",
					PublicKey:   test_constants.ProjectAdmin6,
				},
			},
			response: lambda_profile_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		if test.request.Body.UserType == string(feed_attributes.ADMIN_ACTOR_TYPE) {
			os.Setenv("AUTH_LEVEL", string(auth.NoAuth))
		}
		result, err := lambda_profile_config.Handler(test.request)
		if test.request.Body.UserType == string(feed_attributes.ADMIN_ACTOR_TYPE) {
			os.Setenv("AUTH_LEVEL", string(auth.UserAuth))
		}
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
