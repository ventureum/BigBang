package main

import (
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/app/feed_attributes"
  "gopkg.in/GetStream/stream-go2.v1"
  "time"
  "log"
)

func main() {
  getStreamClient := getstream_config.ConnectGetStreamClient()

  actor := feed_attributes.Actor("0xdavid007")
  verb  := feed_attributes.SubmitVerb
  obj := feed_attributes.CreateObjectFromValue("post:0x001")
  timestamp := feed_attributes.BlockTimestamp(1535493347)
  postType := feed_attributes.PostPostType
  to := []feed_attributes.FeedId{}
  activity := feed_attributes.CreateNewActivity(actor, verb, obj, timestamp, postType, to, map[string]interface{}{})
  getStreamClient.AddFeedActivityToGetStream(activity)
  flatFeed := getStreamClient.CreateFlatFeed(string(feed_attributes.UserFeedSlug), string(actor))
  getStreamClient.GetAllFeedActivitiesByFlatFeed(flatFeed)
  reponse, _ := getStreamClient.C.GetActivitiesByForeignID(
    stream.NewForeignIDTimePair(activity.Object.Value(), stream.Time{time.Unix(int64(timestamp), 0).UTC()}))
  log.Printf("%+v", reponse)

  activity.Extra["test"] = 1234
  getStreamClient.UpdateFeedActivityToGetStream(activity)
  reponse, _ = getStreamClient.C.GetActivitiesByForeignID(
    stream.NewForeignIDTimePair(activity.Object.Value(), stream.Time{time.Unix(int64(timestamp), 0).UTC()}))
  log.Printf("%+v", reponse)
}
