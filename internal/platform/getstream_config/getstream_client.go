package getstream_config

import (
	"BigBang/internal/app/feed_attributes"
	"gopkg.in/GetStream/stream-go2.v1"
	"log"
	"time"
)

type GetStreamClient struct {
	C *stream.Client
}

func ConnectGetStreamClient() *GetStreamClient {
	client, err := stream.NewClientFromEnv()
	if err != nil {
		log.Panicf("Failed to connect to GetStreamClient with error: %+v", err)
	}
	log.Println("Connected to GetStream Client")
	return &GetStreamClient{C: client}
}

func (getStreamClient *GetStreamClient) CreateFlatFeed(feedSlug string, userId string) *stream.FlatFeed {
	return getStreamClient.C.FlatFeed(feedSlug, userId)
}

func (getStreamClient *GetStreamClient) CreateFlatFeedFromFeedId(feedId feed_attributes.FeedId) *stream.FlatFeed {
	return getStreamClient.C.FlatFeed(string(feedId.FeedSlug), string(feedId.UserId))
}

func (getStreamClient *GetStreamClient) AddFeedActivityToGetStream(activity *feed_attributes.Activity) {
	actor := string(activity.Actor)
	verb := string(activity.Verb)
	obj := activity.Object.Value()
	timestamp := activity.Time
	extra := map[string]interface{}{
		"postType": activity.PostType.Value(),
		"rewards":  0,
	}
	for k, v := range activity.Extra {
		extra[k] = v
	}

	streamActivity := stream.Activity{
		Actor:  actor,
		Verb:   verb,
		Object: obj,
		Time: stream.Time{
			Time: time.Unix(int64(timestamp), 0).UTC(),
		},
		ForeignID: obj,
		To:        feed_attributes.ConvertToStringArray(activity.To),
		Extra:     extra,
	}
	var flatFeed *stream.FlatFeed
	if activity.Verb == feed_attributes.ReplyVerb {
		flatFeed = getStreamClient.CreateFlatFeed(string(feed_attributes.UserCommentFeedSlug), string(activity.Actor))
	} else {
		flatFeed = getStreamClient.CreateFlatFeed(string(feed_attributes.UserPostFeedSlug), string(activity.Actor))
	}

	flatFeed.AddActivities(streamActivity)
	log.Printf("Added feed activity to GetStream with object: %s by user %s with stream activty: %v+\n",
		obj, actor, streamActivity)
}

func (getStreamClient *GetStreamClient) GetAllFeedActivitiesByFeedSlugAndUserId(
	feedSlug string, userId string) *[]feed_attributes.Activity {

	flatFeed := getStreamClient.CreateFlatFeed(feedSlug, userId)

	flatFeedResponse, err := flatFeed.GetActivities()

	log.Printf("flatFeedResponse %+v", flatFeedResponse)
	if err != nil {
		log.Panicf("Failed to get activities for feedId %s with error: %+v\n", flatFeed.ID(), err)
	}

	var activities []feed_attributes.Activity
	for _, activity := range flatFeedResponse.Results {
		activities = append(activities, *feed_attributes.ConvertStreamActivityToActivity(&activity))
	}
	return &activities
}

func (getStreamClient *GetStreamClient) GetFeedActivityByForeignIdAndTimestamp(foreignId string, timestamp time.Time) *feed_attributes.Activity {
	response, err := getStreamClient.C.GetActivitiesByForeignID(
		stream.NewForeignIDTimePair(foreignId, ToStreamTime(timestamp)))
	if err != nil {
		log.Panicf("Failed to get activity for foreignId %s and timestamp %s with error: %+v\n", foreignId, timestamp, err)
	}

	if len(response.Results) == 0 {
		return nil
	}

	return feed_attributes.ConvertStreamActivityToActivity(&response.Results[0])
}

func (getStreamClient *GetStreamClient) UpdateFeedActivityToGetStream(activity *feed_attributes.Activity) {
	actor := string(activity.Actor)
	verb := string(activity.Verb)
	obj := activity.Object.Value()
	timestamp := activity.Time
	extra := map[string]interface{}{
		"postType": activity.PostType.Value(),
	}
	for k, v := range activity.Extra {
		extra[k] = v
	}
	streamActivity := stream.Activity{
		Actor:  actor,
		Verb:   verb,
		Object: obj,
		Time: stream.Time{
			Time: time.Unix(int64(timestamp), 0).UTC(),
		},
		ForeignID: obj,
		To:        feed_attributes.ConvertToStringArray(activity.To),
		Extra:     extra,
	}

	getStreamClient.C.UpdateActivities(streamActivity)
	log.Printf("Updated feed activity to GetStream with object: %s by user %s with stream activty: %v+\n",
		obj, actor, streamActivity)
}

func (getStreamClient *GetStreamClient) UpdateFeedActivityToGetStreamByForeignIdAndTimestamp(
	foreignId string, timestamp time.Time, set map[string]interface{}, unset []string) {
	getStreamClient.C.UpdateActivityByForeignID(
		foreignId,
		ToStreamTime(timestamp),
		set,
		unset)
	log.Printf("Successfully updated feed activity to GetStream by Foreign Id %s and Timestamp %s with set %+v and unset %+v",
		foreignId, timestamp, set, unset)
}

func (getStreamClient *GetStreamClient) RemoveFeedActivityByFeedSlugAndUserIdAndForeignId(feedSlug string, userId string, foreignId string) {
	flatFeed := getStreamClient.CreateFlatFeed(feedSlug, userId)
	flatFeed.RemoveActivityByForeignID(foreignId)
	log.Printf("Successfully deleted feed activity by Foreign Id %s for FeedSlug %s and UserId %s", foreignId, feedSlug, userId)
}

func (getStreamClient *GetStreamClient) UpdateFeedPostRewardsByForeignIdAndTimestamp(
	foreignId string, timestamp time.Time, newRewards int64) {
	set := map[string]interface{}{
		"rewards": newRewards,
	}
	var unset []string
	getStreamClient.UpdateFeedActivityToGetStreamByForeignIdAndTimestamp(
		foreignId,
		timestamp,
		set,
		unset)
	log.Printf("Successfully updated feed activity by Foreign Id %s and timestamp %s", foreignId, timestamp)
}

func ToStreamTime(timestamp time.Time) stream.Time {
	return stream.Time{
		Time: time.Unix(int64(timestamp.Unix()), 0).UTC(),
	}
}
