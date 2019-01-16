package feed_attributes

import (
	"gopkg.in/GetStream/stream-go2.v1"
	"sort"
)

type Activity struct {
	Actor     Actor                  `json:"actor"`
	Verb      Verb                   `json:"verb"`
	Object    Object                 `json:"object"`
	ForeignId string                 `json:"foreignId"`
	Time      BlockTimestamp         `json:"time"`
	PostType  PostType               `json:"postType"`
	To        []FeedId               `json:"to"`
	Extra     map[string]interface{} `json:"extra"`
}

func CreateNewActivity(
	actor Actor,
	verb Verb,
	obj Object,
	time BlockTimestamp,
	postType PostType,
	to []FeedId,
	extraParam map[string]interface{}) *Activity {

	return &Activity{
		Actor:     actor,
		Verb:      verb,
		Object:    obj,
		ForeignId: obj.Value(),
		Time:      time,
		PostType:  postType,
		To:        to,
		Extra:     extraParam,
	}
}

func ConvertStreamActivityToActivity(streamActivity *stream.Activity) *Activity {
	if streamActivity == nil {
		return nil
	}
	obj := CreateObjectFromValue(streamActivity.Object)
	postType := streamActivity.Extra["postType"].(string)

	extra := map[string]interface{}{}
	for k, v := range streamActivity.Extra {
		if k == "rewards" {
			extra[k] = int64(v.(float64))
		} else if k != "postType" {
			extra[k] = v
		}
	}
	sort.Sort(sort.StringSlice(streamActivity.To))
	return &Activity{
		Actor:     Actor(streamActivity.Actor),
		Verb:      Verb(streamActivity.Verb),
		Object:    obj,
		ForeignId: obj.Value(),
		PostType:  PostType(postType),
		Time:      BlockTimestamp(streamActivity.Time.Unix()),
		To:        ConvertFromStringArrayToFeedIds(streamActivity.To),
		Extra:     extra,
	}
}
