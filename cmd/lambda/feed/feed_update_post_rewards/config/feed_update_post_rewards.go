package lambda_feed_update_post_rewards_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/getstream_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
)

// Request only used for test
type Request struct {
	IncreasedRedeemBlockNum int64 `json:"increasedRedeemBlockNum,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()
	getStreamClient := getstream_config.ConnectGetStreamClient()
	postgresBigBangClient.Begin()
	postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
	postRewardsForUpdates := postRewardsRecordExecutor.UpdatePostRewardsRecordsByAggregationsTx()
	for _, postRewardsForUpdate := range *postRewardsForUpdates {
		getStreamClient.UpdateFeedPostRewardsByForeignIdAndTimestamp(
			postRewardsForUpdate.Object,
			postRewardsForUpdate.PostTime,
			postRewardsForUpdate.WithdrawableMPs)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler() (response Response, err error) {
	response.Ok = false
	ProcessRequest(&response)
	return response, nil
}
