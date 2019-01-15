package lambda_get_project_list_config

import (
	"BigBang/cmd/lambda/TCR/common"
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
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
	Projects   *[]tcr_attributes.Project `json:"projects,omitempty"`
	NextCursor string                    `json:"nextCursor,omitempty"`
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

	var cursor string
	if cursorStr != "" {
		cursor = utils.Base64DecodeToString(cursorStr)
	}

	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

	projectRecords := projectExecutor.GetProjectRecordsByCursorTx(cursor, limit+1)

	response.ResponseData = &ResponseData{
		NextCursor: "",
		Projects:   nil,
	}

	var projects []tcr_attributes.Project
	for index, projectRecord := range *projectRecords {
		if index < int(limit) {
			project := common.ConstructProjectFromProjectRecordTx(&projectRecord, postgresBigBangClient)
			projects = append(projects, *project)
		} else {
			response.ResponseData.NextCursor = utils.Base64EncodeStr(projectRecord.ID)
		}
	}

	response.ResponseData.Projects = &projects
	if cursorStr == "" {
		log.Printf("ProjectRecords is loaded for first query with limit %d\n", limit)
	} else {
		log.Printf("ProjectRecords is loaded for query with cursor %s and limit %d\n", cursorStr, limit)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
