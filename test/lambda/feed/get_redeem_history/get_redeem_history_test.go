package get_redeem_history

import (
  "BigBang/test/constants"
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/get_redeem_history/config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/utils"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_milestone_points_redeem_history_record_config"
  "log"
)

var postgresBigBangClient = client_config.ConnectPostgresClient()
var milestonePointsRedeemRequestRecordExecutor = milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{*postgresBigBangClient}
var actorMilestonePointsRedeemHistoryRecordExecutor = actor_milestone_points_redeem_history_record_config.ActorMilestonePointsRedeemHistoryRecordExecutor{*postgresBigBangClient}

var nextRedeemBlock = feed_attributes.MoveToNextNRedeemBlock(1)
var executedAt = nextRedeemBlock.ConvertToTime()

var redeemBlock = feed_attributes.MoveToNextNRedeemBlock(1)
func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_redeem_history_config.Request
    response lambda_get_redeem_history_config.Response
    err    error
  }{
    {
      request: lambda_get_redeem_history_config.Request {
        Actor: test_constants.Actor1,
        Limit: 0,
      },
      response: lambda_get_redeem_history_config.Response {
        Redeems: &[]feed_attributes.MilestonePointsRedeemHistory{},
        NextCursor: utils.Base64EncodeIdByInt64AndStr(int64(feed_attributes.MoveToNextNRedeemBlock(6)), test_constants.Actor1),
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_redeem_history_config.Request {
        Actor: test_constants.Actor1,
        Limit: 2,
      },
      response: lambda_get_redeem_history_config.Response {
        Redeems: &[]feed_attributes.MilestonePointsRedeemHistory{
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(6),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(5),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
        },
        NextCursor: utils.Base64EncodeIdByInt64AndStr(int64(feed_attributes.MoveToNextNRedeemBlock(4)), test_constants.Actor1),
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_redeem_history_config.Request {
        Actor: test_constants.Actor1,
        Cursor: utils.Base64EncodeIdByInt64AndStr(int64(feed_attributes.MoveToNextNRedeemBlock(4)), test_constants.Actor1),
        Limit: 2,
      },
      response: lambda_get_redeem_history_config.Response {
        Redeems: &[]feed_attributes.MilestonePointsRedeemHistory{
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(4),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(3),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
        },
        NextCursor: utils.Base64EncodeIdByInt64AndStr(int64(feed_attributes.MoveToNextNRedeemBlock(2)), test_constants.Actor1),
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_get_redeem_history_config.Request {
        Actor: test_constants.Actor1,
        Cursor: utils.Base64EncodeIdByInt64AndStr(int64(feed_attributes.MoveToNextNRedeemBlock(4)), test_constants.Actor1),
        Limit: 5,
      },
      response: lambda_get_redeem_history_config.Response {
        Redeems: &[]feed_attributes.MilestonePointsRedeemHistory{
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(4),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(3),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(2),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
          {
            Actor: test_constants.Actor1,
            RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(1),
            TokenPool: 10000,
            TotalEnrolledMilestonePoints: 400,
            TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
            ActualMilestonePoints: 100,
            ConsumedMilestonePoints: 100,
            RedeemedTokens: 2500,
            SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
            ExecutedAt: executedAt,
          },
        },
        NextCursor: "",
        Ok: true,
      },
      err: nil,
    },
  }

  actorMilestonePointsRedeemHistoryRecord := &actor_milestone_points_redeem_history_record_config.ActorMilestonePointsRedeemHistoryRecord{
    Actor: test_constants.Actor1,
    RedeemBlock:  feed_attributes.MoveToNextNRedeemBlock(2),
    TokenPool: 10000,
    TotalEnrolledMilestonePoints: 400,
    TargetedMilestonePoints: test_constants.RedeemMiletonePointsRegular1,
    ActualMilestonePoints: 100,
    ConsumedMilestonePoints: 100,
    RedeemedTokens: 2500,
    SubmittedAt: milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequest(test_constants.Actor1).SubmittedAt,
    ExecutedAt: executedAt,
  }

  actorMilestonePointsRedeemHistoryRecord.GenerateID()

  log.Printf("%+v\n", actorMilestonePointsRedeemHistoryRecord)
  actorMilestonePointsRedeemHistoryRecordExecutor.UpsertActorMilestonePointsRedeemHistoryRecord(actorMilestonePointsRedeemHistoryRecord)

  actorMilestonePointsRedeemHistoryRecord.RedeemBlock = feed_attributes.MoveToNextNRedeemBlock(3)
  actorMilestonePointsRedeemHistoryRecord.GenerateID()
  actorMilestonePointsRedeemHistoryRecordExecutor.UpsertActorMilestonePointsRedeemHistoryRecord(actorMilestonePointsRedeemHistoryRecord)

  actorMilestonePointsRedeemHistoryRecord.RedeemBlock = feed_attributes.MoveToNextNRedeemBlock(4)
  actorMilestonePointsRedeemHistoryRecord.GenerateID()
  actorMilestonePointsRedeemHistoryRecordExecutor.UpsertActorMilestonePointsRedeemHistoryRecord(actorMilestonePointsRedeemHistoryRecord)

  actorMilestonePointsRedeemHistoryRecord.RedeemBlock = feed_attributes.MoveToNextNRedeemBlock(5)
  actorMilestonePointsRedeemHistoryRecord.GenerateID()
  actorMilestonePointsRedeemHistoryRecordExecutor.UpsertActorMilestonePointsRedeemHistoryRecord(actorMilestonePointsRedeemHistoryRecord)

  actorMilestonePointsRedeemHistoryRecord.RedeemBlock = feed_attributes.MoveToNextNRedeemBlock(6)
  actorMilestonePointsRedeemHistoryRecord.GenerateID()
  actorMilestonePointsRedeemHistoryRecordExecutor.UpsertActorMilestonePointsRedeemHistoryRecord(actorMilestonePointsRedeemHistoryRecord)

  for _, test := range tests {
    result, err := lambda_get_redeem_history_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
    assert.Equal(t, test.response.NextCursor, result.NextCursor)
    assert.Equal(t, len(*test.response.Redeems), len(*result.Redeems))
    expectedRedeems := *test.response.Redeems
    redeems := *result.Redeems
    for index, redeem := range redeems {
       assert.Equal(t, expectedRedeems[index].Actor, redeem.Actor)
       assert.Equal(t, expectedRedeems[index].RedeemBlock, redeem.RedeemBlock)
       assert.Equal(t, expectedRedeems[index].TokenPool, redeem.TokenPool)
       assert.Equal(t, expectedRedeems[index].TotalEnrolledMilestonePoints, redeem.TotalEnrolledMilestonePoints)
       assert.Equal(t, expectedRedeems[index].ActualMilestonePoints, redeem.ActualMilestonePoints)
       assert.Equal(t, expectedRedeems[index].ConsumedMilestonePoints, redeem.ConsumedMilestonePoints)
       assert.Equal(t, expectedRedeems[index].SubmittedAt, redeem.SubmittedAt)
       assert.Equal(t, expectedRedeems[index].ExecutedAt, redeem.ExecutedAt)
    }
  }
}