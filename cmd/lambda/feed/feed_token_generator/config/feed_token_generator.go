package lambda_feed_token_generator_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"gopkg.in/GetStream/stream-go2.v1"
	"log"
	"os"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	FeedSlug           string `json:"feedSlug,required"`
	UserId             string `json:"userId,required"`
	GetStreamApiKey    string `json:"getStreamApiKey,omitEmpty"`
	GetStreamApiSecret string `json:"getStreamApiSecret,omitEmpty"`
}

type Response struct {
	FeedToken string                  `json:"feedToken,omitempty"`
	Ok        bool                    `json:"ok"`
	Message   *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.FeedToken = ""
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	var client *stream.Client
	var err error
	if request.Body.GetStreamApiKey != "" && request.Body.GetStreamApiSecret != "" {
		client, err = stream.NewClient(request.Body.GetStreamApiKey, request.Body.GetStreamApiSecret)
	} else {
		client, err = stream.NewClientFromEnv()
		request.Body.GetStreamApiSecret = os.Getenv("STREAM_API_SECRET")
		request.Body.GetStreamApiKey = os.Getenv("STREAM_API_KEY")
	}

	if err != nil {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.GetStreamClientConnectionError,
		}
		log.Printf("Failed to connect GetStream client with error: %+v\n", err)
		log.Panicln(errorInfo.Marshal())
	}

	client.FlatFeed(request.Body.FeedSlug, request.Body.UserId)
	feedID := feed_attributes.CreateFeedId(request.Body.FeedSlug, request.Body.UserId)
	response.FeedToken = feedID.FeedToken(request.Body.GetStreamApiSecret)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
