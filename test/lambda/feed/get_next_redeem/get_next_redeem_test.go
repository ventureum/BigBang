package get_next_redeem

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/get_next_redeem/config"
  "BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/app/feed_attributes"
  "time"
  "BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
)

var postgresBigBangClient = client_config.ConnectPostgresClient()
var redeemBlockInfoRecordExecutor = redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{*postgresBigBangClient}

var milestonePointsRedeemRequestRecordExecutor = milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{*postgresBigBangClient}
var nextRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(1)
var executedAt = nextRedeemBlock.ConvertToTime()

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_next_redeem_config.Request
    response lambda_get_next_redeem_config.Response
    err    error
  }{
    {
      request: lambda_get_next_redeem_config.Request {
        Actor: test_constants.Actor1,
      },
      response: lambda_get_next_redeem_config.Response {
          NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
            Actor:                   test_constants.Actor1,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular2,
            ActualMilestonePoints:   100,
            EstimatedTokens:         10000,
            SubmittedAt:             milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            RedeemBlockInfo: &feed_attributes.RedeemBlockInfo {
              RedeemBlock:                  nextRedeemBlock,
              TotalEnrolledMilestonePoints: 100,
              TokenPool:                    10000,
              ExecutedAt:                   executedAt,
            },
          },
          Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_get_next_redeem_config.Request {
        Actor: test_constants.Actor2,
      },
      response: lambda_get_next_redeem_config.Response {
        NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
          Actor:                   test_constants.Actor2,
          TargetedMilestonePoints: test_constants.RedeemMiletonePointsMax,
          ActualMilestonePoints:   0,
          EstimatedTokens:         0,
          SubmittedAt:             milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
          RedeemBlockInfo: &feed_attributes.RedeemBlockInfo {
            RedeemBlock:                  nextRedeemBlock,
            TotalEnrolledMilestonePoints: 100,
            TokenPool:                    10000,
            ExecutedAt:                   executedAt,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_next_redeem_config.Request {
        Actor: test_constants.Actor3,
      },
      response: lambda_get_next_redeem_config.Response {
        NextRedeem: &lambda_get_next_redeem_config.ResponseContent{
          Actor:                   test_constants.Actor3,
          TargetedMilestonePoints: test_constants.RedeemMiletonePointsZero,
          ActualMilestonePoints:   0,
          EstimatedTokens:         0,
          SubmittedAt:             milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
          RedeemBlockInfo: &feed_attributes.RedeemBlockInfo {
            RedeemBlock:                  nextRedeemBlock,
            TotalEnrolledMilestonePoints: 100,
            TokenPool:                    10000,
            ExecutedAt:                   executedAt,
          },
        },
        Ok: true,
      },
      err: nil,
    },
  }
  redeemBlockInfoRecordExecutor.UpdateTotalEnrolledMilestonePointsForRedeemBlockInfoRecord(nextRedeemBlock)
  executedAt = executedAt.In(time.UTC)
  for _, test := range tests {
    result, err := lambda_get_next_redeem_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.Equal(t, test.response.NextRedeem.Actor, result.NextRedeem.Actor)
    assert.Equal(t, test.response.NextRedeem.TargetedMilestonePoints, result.NextRedeem.TargetedMilestonePoints)
    assert.Equal(t, test.response.NextRedeem.ActualMilestonePoints, result.NextRedeem.ActualMilestonePoints)
    assert.Equal(t, test.response.NextRedeem.EstimatedTokens, result.NextRedeem.EstimatedTokens)
    assert.Equal(t, test.response.NextRedeem.SubmittedAt, result.NextRedeem.SubmittedAt)
    assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.TokenPool, result.NextRedeem.RedeemBlockInfo.TokenPool)
    assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.RedeemBlock, result.NextRedeem.RedeemBlockInfo.RedeemBlock)
    assert.Equal(t, test.response.NextRedeem.RedeemBlockInfo.TotalEnrolledMilestonePoints, result.NextRedeem.RedeemBlockInfo.TotalEnrolledMilestonePoints)
  }
}
