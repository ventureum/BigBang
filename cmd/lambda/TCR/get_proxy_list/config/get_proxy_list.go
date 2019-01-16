package lambda_get_proxy_list_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"BigBang/internal/platform/postgres_config/TCR/proxy_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Limit  int64  `json:"limit,required"`
	Cursor string `json:"cursor,omitempty"`
}

type ResponseData struct {
	Proxies    *[]string `json:"proxies,omitempty"`
	NextCursor string    `json:"nextCursor,omitempty"`
}

type Response struct {
	ResponseData *ResponseData           `json:"responseData,omitempty"`
	Ok           bool                    `json:"ok"`
	Message      *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.ResponseData = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	limit := request.Body.Limit
	cursorStr := request.Body.Cursor

	var cursor int64
	if cursorStr != "" {
		cursor = utils.Base64DecodeToInt64(cursorStr)
	}

	proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}

	proxyRecords := proxyExecutor.GetListOfProxyByCursorTx(cursor, limit+1)

	response.ResponseData = &ResponseData{
		NextCursor: "",
		Proxies:    nil,
	}

	var proxyUUIDList []string
	for index, proxyRecord := range *proxyRecords {
		if index < int(limit) {
			proxyUUIDList = append(proxyUUIDList, proxyRecord.UUID)
		} else {
			response.ResponseData.NextCursor = utils.Base64EncodeInt64(proxyRecord.ID)
		}
	}

	response.ResponseData.Proxies = &proxyUUIDList

	if cursorStr == "" {
		log.Printf("proxyUUIDList  is loaded for first query with limit %d\n", limit)
	} else {
		log.Printf("ProxyUUIDList  is loaded for query with cursor %s and limit %d\n", cursorStr, limit)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
