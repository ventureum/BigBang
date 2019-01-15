package auth

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"log"
	"os"
)

func RegisterAuth(principalId string, actor string, postgresBigBangClient *client_config.PostgresBigBangClient) {
	auth := os.Getenv("AUTH_LEVEL")

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}

	if principalId != "" && AuthLevel(auth) == AdminAuth {
		match := actorProfileRecordExecutor.CheckActorTypeTx(principalId, feed_attributes.ADMIN_ACTOR_TYPE)
		if match {
			return
		}
	} else if principalId != "" && AuthLevel(auth) == UserAuth {
		match := actorProfileRecordExecutor.CheckActorTypeTx(principalId, feed_attributes.ADMIN_ACTOR_TYPE)
		if match || (actor != "" && principalId == actor) {
			return
		}
	} else if principalId != "" && AuthLevel(auth) == NoAuth {
		return
	}
	errorInfo := error_config.ErrorInfo{
		ErrorCode: error_config.InvalidAuthRegister,
		ErrorData: error_config.ErrorData{
			"principalId": principalId,
			"actor":       actor,
		},
		ErrorLocation: error_config.Auth,
	}
	log.Printf("Invalid Auth Register for principalId %s and actor %s", principalId, actor)
	log.Panicln(errorInfo.Marshal())
}

func AuthProcess(principalId string, actor string, postgresBigBangClient *client_config.PostgresBigBangClient) {
	auth := os.Getenv("AUTH_LEVEL")

	if principalId == "" {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.InvalidPrincipalId,
			ErrorData: error_config.ErrorData{
				"principalId": principalId,
			},
			ErrorLocation: error_config.Auth,
		}
		log.Printf("Invalid PrincipalId: %s", principalId)
		log.Panicln(errorInfo.Marshal())
	}

	if AuthLevel(auth) == NoAuth {
		return
	}

	if postgresBigBangClient == nil {
		postgresBigBangClient = client_config.ConnectPostgresClient(nil)
	}

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	actorType := actorProfileRecordExecutor.GetActorTypeTx(principalId)

	if actorType == "" {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.NoPrincipalIdExisting,
			ErrorData: error_config.ErrorData{
				"principalId": principalId,
			},
			ErrorLocation: error_config.Auth,
		}
		log.Printf("No PrincipalId  %s exists", principalId)
		log.Panicln(errorInfo.Marshal())
	}

	if AuthLevel(auth) == AdminAuth {
		if actorType == feed_attributes.ADMIN_ACTOR_TYPE {
			return
		}
	} else if AuthLevel(auth) == UserAuth {
		if actorType != feed_attributes.ActorType("") &&
			(actor == "" || (actor != "" &&
				((actorType == feed_attributes.ADMIN_ACTOR_TYPE) ||
					(actorType != feed_attributes.ADMIN_ACTOR_TYPE && principalId == actor)))) {
			return
		}
	}

	errorInfo := error_config.ErrorInfo{
		ErrorCode: error_config.InvalidAuthAccess,
		ErrorData: error_config.ErrorData{
			"principalId": principalId,
		},
		ErrorLocation: error_config.Auth,
	}
	log.Printf("Invalid Auth Access for principalId %s", actor)
	log.Panicln(errorInfo.Marshal())

}

func ValidateAndCreateActorTypeWithAuthLevel(actorTypeStr string) feed_attributes.ActorType {
	var actorType feed_attributes.ActorType
	auth := os.Getenv("AUTH_LEVEL")
	if AuthLevel(auth) == NoAuth && feed_attributes.ActorType(actorTypeStr) == feed_attributes.ADMIN_ACTOR_TYPE {
		return feed_attributes.ADMIN_ACTOR_TYPE
	}
	switch feed_attributes.ActorType(actorTypeStr) {
	case feed_attributes.USER_ACTOR_TYPE, feed_attributes.KOL_ACTOR_TYPE, feed_attributes.PF_ACTOR_TYPE:
		actorType = feed_attributes.ActorType(actorTypeStr)
	default:
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.InvalidActorType,
			ErrorData: error_config.ErrorData{
				"actorType": actorTypeStr,
			},
			ErrorLocation: error_config.ActorTypeLocation,
		}
		log.Printf("Invalid actorType: %s", actorTypeStr)
		log.Panicln(errorInfo.Marshal())
	}
	return actorType
}
