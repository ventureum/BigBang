package get_project_test

import (
	"BigBang/cmd/lambda/TCR/get_project/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

var EmptyObjectives []tcr_attributes.Objective

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_project_config.Request
		response lambda_get_project_config.Response
		err      error
	}{
		{
			request: lambda_get_project_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_project_config.RequestContent{
					ProjectId: test_constants.ProjectId1,
				},
			},
			response: lambda_get_project_config.Response{
				Project: &tcr_attributes.Project{
					ProjectId:      test_constants.ProjectId1,
					Admin:          test_constants.ProjectAdmin1,
					Content:        test_constants.ProjectContent1,
					BlockTimestamp: test_constants.ProjectBlockTimestamp1,
					MilestonesInfo: &tcr_attributes.MilestonesInfo{
						CurrentMilestone:       0,
						NumMilestonesCompleted: 4,
						NumMilestones:          4,
						Milestones: &[]tcr_attributes.Milestone{
							{
								ProjectId:      test_constants.ProjectId1,
								MilestoneId:    test_constants.MilestoneId1,
								Content:        test_constants.MilestoneContent1,
								BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
								StartTime:      test_constants.MilestoneStartTime1,
								EndTime:        test_constants.MilestoneEndTime1,
								State:          tcr_attributes.CompleteMilestoneState,
								NumObjectives:  2,
								Objectives: &[]tcr_attributes.Objective{
									{
										ProjectId:      test_constants.ProjectId1,
										MilestoneId:    test_constants.MilestoneId1,
										ObjectiveId:    test_constants.ObjectiveId1,
										Content:        test_constants.ObjectiveContent1,
										BlockTimestamp: test_constants.ObjectiveBlockTimestamp1,
									},
									{
										ProjectId:      test_constants.ProjectId1,
										MilestoneId:    test_constants.MilestoneId1,
										ObjectiveId:    test_constants.ObjectiveId2,
										Content:        test_constants.ObjectiveContent1,
										BlockTimestamp: test_constants.ObjectiveBlockTimestamp2,
									},
								},
							},
							{
								ProjectId:      test_constants.ProjectId1,
								MilestoneId:    test_constants.MilestoneId2,
								Content:        test_constants.MilestoneContent2,
								BlockTimestamp: test_constants.MilestoneBlockTimestamp2,
								StartTime:      test_constants.MilestoneStartTime2,
								EndTime:        test_constants.MilestoneEndTime2,
								State:          tcr_attributes.CompleteMilestoneState,
								NumObjectives:  3,
								Objectives: &[]tcr_attributes.Objective{
									{
										ProjectId:      test_constants.ProjectId1,
										MilestoneId:    test_constants.MilestoneId2,
										ObjectiveId:    test_constants.ObjectiveId1,
										Content:        test_constants.ObjectiveContent1,
										BlockTimestamp: test_constants.ObjectiveBlockTimestamp3,
									},
									{
										ProjectId:      test_constants.ProjectId1,
										MilestoneId:    test_constants.MilestoneId2,
										ObjectiveId:    test_constants.ObjectiveId2,
										Content:        test_constants.ObjectiveContent1,
										BlockTimestamp: test_constants.ObjectiveBlockTimestamp4,
									},
									{
										ProjectId:      test_constants.ProjectId1,
										MilestoneId:    test_constants.MilestoneId2,
										ObjectiveId:    test_constants.ObjectiveId3,
										Content:        test_constants.ObjectiveContent1,
										BlockTimestamp: test_constants.ObjectiveBlockTimestamp5,
									},
								},
							},
							{
								ProjectId:      test_constants.ProjectId1,
								MilestoneId:    test_constants.MilestoneId3,
								Content:        test_constants.MilestoneContent1,
								BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
								StartTime:      test_constants.MilestoneStartTime3,
								EndTime:        test_constants.MilestoneEndTime3,
								State:          tcr_attributes.CompleteMilestoneState,
								NumObjectives:  0,
								Objectives:     &EmptyObjectives,
							},
							{
								ProjectId:      test_constants.ProjectId1,
								MilestoneId:    test_constants.MilestoneId4,
								Content:        test_constants.MilestoneContent1,
								BlockTimestamp: test_constants.MilestoneBlockTimestamp3,
								StartTime:      test_constants.MilestoneStartTime4,
								EndTime:        test_constants.MilestoneEndTime4,
								State:          tcr_attributes.CompleteMilestoneState,
								NumObjectives:  0,
								Objectives:     &EmptyObjectives,
							},
						},
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_project_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.Project.ProjectId, result.Project.ProjectId)
		assert.Equal(t, test.response.Project.Admin, result.Project.Admin)
		assert.Equal(t, test.response.Project.Content, result.Project.Content)
		assert.Equal(t, test.response.Project.BlockTimestamp, result.Project.BlockTimestamp)
		assert.Equal(t, test.response.Project.MilestonesInfo, result.Project.MilestonesInfo)

	}
}
