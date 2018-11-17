package lambda_get_redeem_block_info_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
)


type Request struct {
  RedeemBlockTimestamp int64 `json:"redeemBlockTimestamp,required"`
}

type Response struct {
  RedeemBlockInfo *feed_attributes.RedeemBlockInfo `json:"redeemBlockInfo,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.RedeemBlockInfo = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{*postgresBigBangClient}

  redeemBlockTimestamp := request.RedeemBlockTimestamp
  redeemBlock := feed_attributes.CreateRedeemBlockFromUnix(redeemBlockTimestamp)
  redeemBlockInfoRecordExecutor.VerifyRedeemBlockInfoExisting(redeemBlock)

  response.RedeemBlockInfo = redeemBlockInfoRecordExecutor.GetRedeemBlockInfo(redeemBlock)

  log.Printf("Sucessfully loaded content of RedeemBlockInfo for redeemBlock %d\n", redeemBlock)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
