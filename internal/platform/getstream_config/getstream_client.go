package getstream_config

import (
  "gopkg.in/GetStream/stream-go2.v1"
  "log"
  "time"
  "BigBang/internal/app/feed_attributes"
)

type GetStreamClient struct {
  C *stream.Client
}

func ConnectGetStreamClient() (*GetStreamClient) {
  client, err := stream.NewClientFromEnv()
  if err != nil {
    log.Panicf("Failed to connect to GetStreamClient with error: %+v", err)
  }
  log.Println("Connected to GetStream Client")
  return &GetStreamClient{C:  client}
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
  extra := map[string]interface{} {
    "postType": activity.PostType.Value(),
  }
  for k, v := range activity.Extra {
    extra[k] = v
  }
  streamActivity := stream.Activity{
    Actor: actor,
    Verb: verb,
    Object: obj,
    Time: stream.Time{
      Time: time.Unix(int64(timestamp), 0).UTC(),
    },
    ForeignID: obj,
    To: feed_attributes.ConvertToStringArray(activity.To),
    Extra: extra,
  }

  flatFeed := getStreamClient.CreateFlatFeed(string(feed_attributes.UserFeedSlug), actor)
  flatFeed.AddActivities(streamActivity)
  log.Printf("Added feed activity to GetStream with object: %s by user %s with stream activty: %v+\n",
    obj, actor, streamActivity)
}

func (getStreamClient *GetStreamClient) GetAllFeedActivitiesByFeedId(feedId feed_attributes.FeedId) {
  flatFeed := getStreamClient.C.FlatFeed(string(feedId.FeedSlug), string(feedId.UserId))
  flatFeedResponse, err := flatFeed.GetActivities(stream.WithActivitiesLimit(10))
  if err != nil {
    log.Panicf("Failed to get activities for feedId %s\n with error: %+v\n", feedId.Value(), err)
  }
  log.Printf("%+v\n",flatFeedResponse.Results)
}

func (getStreamClient *GetStreamClient) GetAllFeedActivitiesByFlatFeed(flatFeed *stream.FlatFeed) {
  flatFeedResponse, err := flatFeed.GetActivities()
  if err != nil {
    log.Panicf("Failed to get activities for feedId %s with error: %+v\n", flatFeed.ID(), err)
  }
  log.Printf("Activities for Feed Id %s\n %+v\n",flatFeed.ID(), flatFeedResponse.Results)
}

func (getStreamClient *GetStreamClient) UpdateFeedActivityToGetStream(activity *feed_attributes.Activity) {
  actor := string(activity.Actor)
  verb := string(activity.Verb)
  obj := activity.Object.Value()
  timestamp := activity.Time
  extra := map[string]interface{} {
    "postType": activity.PostType.Value(),
  }
  for k, v := range activity.Extra {
    extra[k] = v
  }
  streamActivity := stream.Activity{
    Actor: actor,
    Verb: verb,
    Object: obj,
    Time: stream.Time{
      Time: time.Unix(int64(timestamp), 0).UTC(),
    },
    ForeignID: obj,
    To: feed_attributes.ConvertToStringArray(activity.To),
    Extra: extra,
  }

  getStreamClient.C.UpdateActivities(streamActivity)
  log.Printf("Updated feed activity to GetStream with object: %s by user %s with stream activty: %v+\n",
    obj, actor, streamActivity)
}
