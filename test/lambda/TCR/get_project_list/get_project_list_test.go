package get_project_list

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/pkg/utils"
  "BigBang/cmd/lambda/TCR/get_project_list/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_project_list_config.Request
    response lambda_get_project_list_config.Response
    err    error
  }{
    {
      request: lambda_get_project_list_config.Request {
        Limit: 0,
      },
      response: lambda_get_project_list_config.Response {
        Projects: &[]tcr_attributes.Project{},
        NextCursor: utils.Base64EncodeIdByInt64AndStr(test_constants.BlockTimestamp5, test_constants.ProjectId6),
        Ok: true,
      },
      err: nil,
    },
    {
     request: lambda_get_project_list_config.Request {
       Limit: 2,
     },
     response: lambda_get_project_list_config.Response {
       Projects: &[]tcr_attributes.Project{
          {
            ProjectId: test_constants.ProjectId6,
            Admin: test_constants.Admin1,
            Content:  test_constants.ProjectContent1,
            BlockTimestamp: test_constants.BlockTimestamp5,
          },
          {
            ProjectId: test_constants.ProjectId5,
            Admin: test_constants.Admin1,
            Content:  test_constants.ProjectContent1,
            BlockTimestamp: test_constants.BlockTimestamp5,
         },
       },
       NextCursor: utils.Base64EncodeIdByInt64AndStr(test_constants.BlockTimestamp4, test_constants.ProjectId4),
       Ok: true,
     },
     err: nil,
    },
    {
     request: lambda_get_project_list_config.Request {
       Limit: 2,
       Cursor: utils.Base64EncodeIdByInt64AndStr(test_constants.BlockTimestamp4, test_constants.ProjectId4),
     },
     response: lambda_get_project_list_config.Response {
       Projects: &[]tcr_attributes.Project{
         {
           ProjectId: test_constants.ProjectId4,
           Admin: test_constants.Admin1,
           Content:  test_constants.ProjectContent1,
           BlockTimestamp: test_constants.BlockTimestamp4,
         },
         {
           ProjectId: test_constants.ProjectId3,
           Admin: test_constants.Admin1,
           Content:  test_constants.ProjectContent1,
           BlockTimestamp: test_constants.BlockTimestamp3,
         },
       },
       NextCursor: utils.Base64EncodeIdByInt64AndStr(test_constants.BlockTimestamp2, test_constants.ProjectId2),
       Ok: true,
     },
     err: nil,
    },
    {
     request: lambda_get_project_list_config.Request {
       Limit: 5,
       Cursor: utils.Base64EncodeIdByInt64AndStr(test_constants.BlockTimestamp4, test_constants.ProjectId4),
     },
     response: lambda_get_project_list_config.Response {
       Projects: &[]tcr_attributes.Project{
         {
           ProjectId: test_constants.ProjectId4,
           Admin: test_constants.Admin1,
           Content:  test_constants.ProjectContent1,
           BlockTimestamp: test_constants.BlockTimestamp4,
         },
         {
           ProjectId: test_constants.ProjectId3,
           Admin: test_constants.Admin1,
           Content:  test_constants.ProjectContent1,
           BlockTimestamp: test_constants.BlockTimestamp3,
         },
         {
           ProjectId: test_constants.ProjectId2,
           Admin: test_constants.Admin1,
           Content:  test_constants.ProjectContent1,
           BlockTimestamp: test_constants.BlockTimestamp2,
         },
         {
           ProjectId: test_constants.ProjectId1,
           Admin: test_constants.Admin1,
           Content:  test_constants.ProjectContent1,
           BlockTimestamp: test_constants.BlockTimestamp1,
         },
       },
       NextCursor: "",
       Ok: true,
     },
     err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_project_list_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.Equal(t, test.response.NextCursor, result.NextCursor)
    resultProjects := *result.Projects
    responseProjects := *test.response.Projects
    assert.Equal(t, len(resultProjects), len(responseProjects))
    for index, responseProject := range responseProjects {
      assert.Equal(t, resultProjects[index].ProjectId, responseProject.ProjectId)
      assert.Equal(t, resultProjects[index].Content, responseProject.Content)
      assert.Equal(t, resultProjects[index].Admin, responseProject.Admin)
      assert.Equal(t, resultProjects[index].BlockTimestamp, responseProject.BlockTimestamp)
    }
  }
}
