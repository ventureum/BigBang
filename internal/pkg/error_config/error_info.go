package error_config

import (
	"encoding/json"
	"github.com/jmoiron/sqlx/types"
	"log"
	"reflect"
)

type ErrorInfo struct {
	ErrorCode        ErrorCode     `json:"errorCode,omitempty"`
	ErrorData        ErrorData     `json:"errorData,omitempty"`
	ErrorMessage     ErrorMessage  `json:"errorMessage,omitempty"`
	ErrorLocation    ErrorLocation `json:"errorLocation,omitempty"`
	ErrorUnknownType interface{}   `json:"errorUnknownType,omitempty"`
}

func (errorInfo *ErrorInfo) Marshal() string {
	res, _ := json.Marshal(*errorInfo)
	return string(res)
}

func CreatedErrorInfoFromString(errObj interface{}) *ErrorInfo {
	var errorInfo ErrorInfo
	switch reflect.TypeOf(errObj) {
	case reflect.TypeOf(string("")):
		jsonText := types.JSONText(errObj.(string))
		err := jsonText.Unmarshal(&errorInfo)
		if err != nil {
			log.Printf("Failed to unmarshal errorInfo %+v with error: %+v\n", jsonText, err)
			errorInfo.ErrorCode = General
			errorInfo.ErrorMessage = ErrorMessage(errObj.(string))
		}
	default:
		errorInfo.ErrorUnknownType = errObj
	}
	return &errorInfo
}

func (errorInfo *ErrorInfo) AddErrorData(fieldName string, val interface{}) {
	errorInfo.ErrorData[fieldName] = val
}
