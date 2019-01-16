package lambda_set_token_pool_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	RedeemBlockTimestamp int64 `json:"redeemBlockTimestamp,required"`
	TokenPool            int64 `json:"tokenPool,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
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

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{*postgresBigBangClient}

	redeemBlockTimestamp := request.Body.RedeemBlockTimestamp
	tokenPool := request.Body.TokenPool

	redeemBlock := feed_attributes.CreateRedeemBlockFromUnix(redeemBlockTimestamp)
	redeemBlockInfoRecordExecutor.VerifyRedeemBlockInfoExistingTx(redeemBlock)

	redeemBlockInfoRecordExecutor.UpdateTokenPoolForRedeemBlockInfoRecordTx(redeemBlock, tokenPool)

	postgresBigBangClient.Commit()

	log.Printf("Sucessfully set token pool for redeemBlock %d\n", redeemBlock)

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
