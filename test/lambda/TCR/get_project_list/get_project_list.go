package get_project_list_test

import (
	"BigBang/cmd/lambda/TCR/get_project_list/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/utils"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_project_list_config.Request
		response lambda_get_project_list_config.Response
		err      error
	}{
		{
			request: lambda_get_project_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_list_config.RequestContent{
					Limit: 0,
				},
			},
			response: lambda_get_project_list_config.Response{
				ResponseData: &lambda_get_project_list_config.ResponseData{
					Projects:   &[]tcr_attributes.Project{},
					NextCursor: utils.Base64EncodeIdByInt64AndStr(test_constants.ProjectBlockTimestamp5, test_constants.ProjectId6),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_project_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_list_config.RequestContent{
					Limit: 2,
				},
			},
			response: lambda_get_project_list_config.Response{
				ResponseData: &lambda_get_project_list_config.ResponseData{
					Projects: &[]tcr_attributes.Project{
						{
							ProjectId:      test_constants.ProjectId6,
							Admin:          test_constants.ProjectAdmin6,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp5,
						},
						{
							ProjectId:      test_constants.ProjectId5,
							Admin:          test_constants.ProjectAdmin5,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp5,
						},
					},
					NextCursor: utils.Base64EncodeIdByInt64AndStr(test_constants.ProjectBlockTimestamp4, test_constants.ProjectId4),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_project_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_list_config.RequestContent{
					Limit:  2,
					Cursor: utils.Base64EncodeIdByInt64AndStr(test_constants.ProjectBlockTimestamp4, test_constants.ProjectId4),
				},
			},
			response: lambda_get_project_list_config.Response{
				ResponseData: &lambda_get_project_list_config.ResponseData{
					Projects: &[]tcr_attributes.Project{
						{
							ProjectId:      test_constants.ProjectId4,
							Admin:          test_constants.ProjectAdmin4,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp4,
						},
						{
							ProjectId:      test_constants.ProjectId3,
							Admin:          test_constants.ProjectAdmin3,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp3,
						},
					},
					NextCursor: utils.Base64EncodeIdByInt64AndStr(test_constants.ProjectBlockTimestamp2, test_constants.ProjectId2),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_project_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_list_config.RequestContent{
					Limit:  5,
					Cursor: utils.Base64EncodeIdByInt64AndStr(test_constants.ProjectBlockTimestamp4, test_constants.ProjectId4),
				},
			},
			response: lambda_get_project_list_config.Response{
				ResponseData: &lambda_get_project_list_config.ResponseData{
					Projects: &[]tcr_attributes.Project{
						{
							ProjectId:      test_constants.ProjectId4,
							Admin:          test_constants.ProjectAdmin4,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp4,
						},
						{
							ProjectId:      test_constants.ProjectId3,
							Admin:          test_constants.ProjectAdmin3,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp3,
						},
						{
							ProjectId:      test_constants.ProjectId2,
							Admin:          test_constants.ProjectAdmin2,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp2,
						},
						{
							ProjectId:      test_constants.ProjectId1,
							Admin:          test_constants.ProjectAdmin1,
							Content:        test_constants.ProjectContent1,
							BlockTimestamp: test_constants.ProjectBlockTimestamp1,
						},
					},
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_project_list_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.ResponseData.NextCursor, result.ResponseData.NextCursor)
		resultProjects := *result.ResponseData.Projects
		responseProjects := *test.response.ResponseData.Projects
		assert.Equal(t, len(resultProjects), len(responseProjects))
		for index, responseProject := range responseProjects {
			assert.Equal(t, resultProjects[index].ProjectId, responseProject.ProjectId)
			assert.Equal(t, resultProjects[index].Content, responseProject.Content)
			assert.Equal(t, resultProjects[index].Admin, responseProject.Admin)
			assert.Equal(t, resultProjects[index].BlockTimestamp, responseProject.BlockTimestamp)
		}
	}
}
