package lambda_feed_update_post_rewards_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
  "time"
)

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()
  getStreamClient := getstream_config.ConnectGetStreamClient()
  postgresBigBangClient.Begin()
  redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{*postgresBigBangClient}

  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
  postRewardsForUpdates := postRewardsRecordExecutor.UpdatePostRewardsRecordsByAggregationsTx()
  for _, postRewardsForUpdate := range *postRewardsForUpdates {
    getStreamClient.UpdateFeedPostRewardsByForeignIdAndTimestamp(
      postRewardsForUpdate.Object,
      postRewardsForUpdate.PostTime,
      postRewardsForUpdate.WithdrawableMPs)
  }
  currentRedeemBlock := time.Now().UTC().Unix() / (60 * 60 * 24 * 7)
  redeemBlockInfoRecordExecutor.UpdateTotalEnrolledMilestonePointsForRedeemBlockInfoRecordTx(currentRedeemBlock)
  redeemBlockInfoRecordExecutor.InitRedeemBlockInfoTx(currentRedeemBlock + 1)
  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler() (response Response, err error) {
  response.Ok = false
  ProcessRequest(&response)
  return response, nil
}
