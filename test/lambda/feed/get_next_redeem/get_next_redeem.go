package get_next_redeem_test

import (
	"BigBang/cmd/lambda/feed/get_next_redeem/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
	"BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	nextRedeemBlock := feed_attributes.MoveToNextNRedeemBlock(1)
	executedAt := nextRedeemBlock.ConvertToTime()
	postgresBigBangClient.Begin()
	redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{
		*postgresBigBangClient}
	redeemBlockInfoRecordExecutor.UpdateTotalEnrolledMilestonePointsForRedeemBlockInfoRecordTx(nextRedeemBlock)
	postgresBigBangClient.Commit()
	postgresBigBangClient.Begin()
	milestonePointsRedeemRequestRecordExecutor := milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{
		*postgresBigBangClient}

	tests := []struct {
		request  lambda_get_next_redeem_config.Request
		response lambda_get_next_redeem_config.Response
		err      error
	}{
		{
			request: lambda_get_next_redeem_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_next_redeem_config.RequestContent{
					Actor: test_constants.Actor1,
				},
			},
			response: lambda_get_next_redeem_config.Response{
				NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
					Actor:                   test_constants.Actor1,
					TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
					ActualMilestonePoints:   100,
					EstimatedTokens:         2500,
					RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
						RedeemBlock:                  nextRedeemBlock,
						TotalEnrolledMilestonePoints: 400,
						TokenPool:                    10000,
						ExecutedAt:                   executedAt,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_next_redeem_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_next_redeem_config.RequestContent{
					Actor: test_constants.Actor2,
				},
			},
			response: lambda_get_next_redeem_config.Response{
				NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
					Actor:                   test_constants.Actor2,
					TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular2,
					ActualMilestonePoints:   100,
					EstimatedTokens:         2500,
					RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
						RedeemBlock:                  nextRedeemBlock,
						TotalEnrolledMilestonePoints: 400,
						TokenPool:                    10000,
						ExecutedAt:                   executedAt,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_next_redeem_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_next_redeem_config.RequestContent{
					Actor: test_constants.Actor3,
				},
			},
			response: lambda_get_next_redeem_config.Response{
				NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
					Actor:                   test_constants.Actor3,
					TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular3,
					ActualMilestonePoints:   100,
					EstimatedTokens:         1250,
					RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
						RedeemBlock:                  nextRedeemBlock,
						TotalEnrolledMilestonePoints: 400,
						TokenPool:                    10000,
						ExecutedAt:                   executedAt,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_next_redeem_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_next_redeem_config.RequestContent{
					Actor: test_constants.Actor4,
				},
			},
			response: lambda_get_next_redeem_config.Response{
				NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
					Actor:                   test_constants.Actor4,
					TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular4,
					ActualMilestonePoints:   100,
					EstimatedTokens:         1250,
					RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
						RedeemBlock:                  nextRedeemBlock,
						TotalEnrolledMilestonePoints: 400,
						TokenPool:                    10000,
						ExecutedAt:                   executedAt,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_next_redeem_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_next_redeem_config.RequestContent{
					Actor: test_constants.Actor5,
				},
			},
			response: lambda_get_next_redeem_config.Response{
				NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
					Actor:                   test_constants.Actor5,
					TargetedMilestonePoints: test_constants.RedeemMiletonePointsMax,
					ActualMilestonePoints:   100,
					EstimatedTokens:         2500,
					RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
						RedeemBlock:                  nextRedeemBlock,
						TotalEnrolledMilestonePoints: 400,
						TokenPool:                    10000,
						ExecutedAt:                   executedAt,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_next_redeem_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_next_redeem_config.RequestContent{
					Actor: test_constants.Actor6,
				},
			},
			response: lambda_get_next_redeem_config.Response{
				NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
					Actor:                   test_constants.Actor6,
					TargetedMilestonePoints: test_constants.RedeemMiletonePointsZero,
					ActualMilestonePoints:   100,
					EstimatedTokens:         0,
					RedeemBlockInfo: &feed_attributes.RedeemBlockInfo{
						RedeemBlock:                  nextRedeemBlock,
						TotalEnrolledMilestonePoints: 400,
						TokenPool:                    10000,
						ExecutedAt:                   executedAt,
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		result, err := lambda_get_next_redeem_config.Handler(test.request)
		submittedAt := milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequestTx(test.request.Body.Actor).SubmittedAt
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.NextRedeem.Actor, result.NextRedeem.Actor)
		assert.Equal(t, test.response.NextRedeem.TargetedMilestonePoints, result.NextRedeem.TargetedMilestonePoints)
		assert.Equal(t, test.response.NextRedeem.ActualMilestonePoints, result.NextRedeem.ActualMilestonePoints)
		assert.Equal(t, test.response.NextRedeem.EstimatedTokens, result.NextRedeem.EstimatedTokens)
		assert.Equal(t, submittedAt.Unix(), result.NextRedeem.SubmittedAt.Unix())
		assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.TokenPool, result.NextRedeem.RedeemBlockInfo.TokenPool)
		assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.RedeemBlock, result.NextRedeem.RedeemBlockInfo.RedeemBlock)
		assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.TotalEnrolledMilestonePoints, result.NextRedeem.RedeemBlockInfo.TotalEnrolledMilestonePoints)
		assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.ExecutedAt.Unix(), result.NextRedeem.RedeemBlockInfo.ExecutedAt.Unix())
	}

	postgresBigBangClient.Commit()
}
