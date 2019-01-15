package lambda_profile_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
	"BigBang/internal/platform/postgres_config/feed/refuel_record_config"
	"log"
	"strings"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor          string `json:"actor,required"`
	UserType       string `json:"userType,required"`
	Username       string `json:"username,required"`
	PhotoUrl       string `json:"photoUrl,omitempty"`
	TelegramId     string `json:"telegramId,omitempty"`
	PhoneNumber    string `json:"phoneNumber,omitempty"`
	PublicKey      string `json:"publicKey,omitempty"`
	ProfileContent string `json:"profileContent,omitempty"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToActorProfileRecord() *actor_profile_record_config.ActorProfileRecord {
	return &actor_profile_record_config.ActorProfileRecord{
		Actor:          request.Body.Actor,
		ActorType:      auth.ValidateAndCreateActorTypeWithAuthLevel(request.Body.UserType),
		Username:       request.Body.Username,
		PhotoUrl:       request.Body.PhotoUrl,
		TelegramId:     request.Body.TelegramId,
		PhoneNumber:    request.Body.PhoneNumber,
		PublicKey:      strings.ToLower(request.Body.PublicKey),
		ProfileContent: request.Body.ProfileContent,
	}
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	actor := request.Body.Actor

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	inserted := actorProfileRecordExecutor.UpsertActorProfileRecordTx(request.ToActorProfileRecord())

	if inserted {
		refuelRecordExecutor := refuel_record_config.RefuelRecordExecutor{*postgresBigBangClient}
		actorReputationsRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
			*postgresBigBangClient}
		milestonePointsRedeemRequestRecordExecutor := milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{*postgresBigBangClient}

		initFuel := feed_attributes.MaxFuelForFuelUpdateInterval
		initReputation := feed_attributes.Reputation(initFuel)

		actorReputationsRecord := actor_rewards_info_record_config.ActorRewardsInfoRecord{
			Actor:                     actor,
			Reputation:                initReputation,
			Fuel:                      initFuel,
			MilestonePointsFromVotes:  0,
			MilestonePointsFromPosts:  0,
			MilestonePointsFromOthers: 0,
			MilestonePoints:           0,
		}
		actorReputationsRecordExecutor.UpsertActorRewardsInfoRecordTx(&actorReputationsRecord)
		refuelRecordExecutor.UpsertRefuelRecordTx(&refuel_record_config.RefuelRecord{
			Actor:           actor,
			Fuel:            initFuel,
			Reputation:      initReputation,
			MilestonePoints: 0,
		})
		log.Printf("Created Actor Fuel Account for actor %s", actor)

		milestonePointsRedeemRequestRecordExecutor.UpsertMilestonePointsRedeemRequestRecordTx(
			&milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecord{
				Actor:                   actor,
				NextRedeemBlock:         0,
				TargetedMilestonePoints: 0,
			})
	}

	postgresBigBangClient.Commit()

	if inserted {
		log.Printf("Created Profile for actor %s", actor)
	} else {
		log.Printf("Updated Profile for actor %s", actor)
	}

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
