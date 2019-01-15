package lambda_reset_actor_fuel_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
)

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

	postgresBigBangClient.Begin()
	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*postgresBigBangClient}
	actorRewardsInfoRecordExecutor.ResetAllActorFuelTx(feed_attributes.MaxFuelForFuelUpdateInterval)
	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler() (response Response, err error) {
	response.Ok = false
	ProcessRequest(&response)
	return response, nil
}
