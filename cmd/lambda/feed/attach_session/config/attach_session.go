package lambda_attach_session_config

import (
	"gopkg.in/GetStream/stream-go2.v1"
	"log"

	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/getstream_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/post_config"
	"BigBang/internal/platform/postgres_config/feed/session_record_config"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor              string                  `json:"actor,required"`
	PostHash           string                  `json:"postHash,required"`
	StartTime          int64                   `json:"startTime,required"`
	EndTime            int64                   `json:"endTime,required"`
	Content            feed_attributes.Content `json:"content,required"`
	GetStreamApiKey    string                  `json:"getStreamApiKey,omitempty"`
	GetStreamApiSecret string                  `json:"getStreamApiSecret,omitempty"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToSessionRecord() *session_record_config.SessionRecord {
	return &session_record_config.SessionRecord{
		Actor:     request.Body.Actor,
		PostHash:  request.Body.PostHash,
		StartTime: request.Body.StartTime,
		EndTime:   request.Body.EndTime,
		Content:   request.Body.Content.ToJsonText(),
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

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, request.Body.Actor, postgresBigBangClient)

	var err error
	var getStreamIOClient *stream.Client
	if request.Body.GetStreamApiKey != "" && request.Body.GetStreamApiSecret != "" {
		getStreamIOClient, err = stream.NewClient(request.Body.GetStreamApiKey, request.Body.GetStreamApiSecret)
	} else {
		getStreamIOClient, err = stream.NewClientFromEnv()
	}

	if err != nil {
		log.Panic(err)
	}

	getStreamClient := &getstream_config.GetStreamClient{C: getStreamIOClient}
	sessionRecord := request.ToSessionRecord()

	postExecutor := post_config.PostExecutor{*postgresBigBangClient}
	sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresBigBangClient}

	postExecutor.VerifyPostRecordExistingTx(sessionRecord.PostHash)

	postRecord := postExecutor.GetPostRecordTx(sessionRecord.PostHash)
	sessionRecordExecutor.UpsertSessionRecordTx(sessionRecord)

	activity := postRecord.ToActivity(
		feed_attributes.OFF_CHAIN, feed_attributes.BlockTimestamp(postRecord.CreatedAt.Unix()))

	// Insert Activity to GetStream
	sessionRecord.EmbedSessionRecordToActivity(activity)
	getStreamClient.UpdateFeedActivityToGetStream(activity)

	postgresBigBangClient.Commit()

	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
