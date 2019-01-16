package feed_attributes

import (
	"log"
	"strings"
)

type ObjectType string

type Object struct {
	ObjType ObjectType `json:"objType"`
	ObjId   string     `json:"objId"`
}

const (
	ReplyObjectType ObjectType = "reply"
	PostObjectType  ObjectType = "post"
)

func CreateObject(objType string, objId string) Object {
	return Object{
		ObjType: ObjectType(objType),
		ObjId:   objId,
	}
}

func CreateObjectFromValue(value string) Object {
	s := strings.Split(value, ":")
	if len(s) != 2 {
		log.Fatal("value is not valid when creating Object: ", value)
	}
	return CreateObject(s[0], s[1])
}

func (obj Object) Value() string {
	return string(obj.ObjType) + ":" + obj.ObjId
}
