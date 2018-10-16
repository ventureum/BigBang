package test_constants

import "BigBang/internal/app/feed_attributes"


const ActivityPostHash1 = "0xActivityPostHash1"
const ActivityPostBlockTimestamp1 = 1539108020
const ActivityPostHash2 = "0xActivityPostHash2"
const ActivityPostBlockTimestamp2 = 1539107020
const ActivityPostHash3 = "0xActivityPostHash3"
const ActivityPostBlockTimestamp3 = 1539107120

var Activity1 = feed_attributes.CreateNewActivity(
  feed_attributes.Actor(Actor1),
  feed_attributes.SubmitVerb,
  feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash1),
  ActivityPostBlockTimestamp1,
  feed_attributes.PostPostType,
  []feed_attributes.FeedId{},
  map[string]interface{}{},
)

var Activity2 = feed_attributes.CreateNewActivity(
  feed_attributes.Actor(Actor2),
  feed_attributes.SubmitVerb,
  feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash2),
  ActivityPostBlockTimestamp2,
  feed_attributes.PostPostType,
  []feed_attributes.FeedId{},
  map[string]interface{}{},
)

var Activity3 = feed_attributes.CreateNewActivity(
  feed_attributes.Actor(Actor3),
  feed_attributes.SubmitVerb,
  feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash3),
  ActivityPostBlockTimestamp3,
  feed_attributes.PostPostType,
  []feed_attributes.FeedId{},
  map[string]interface{}{},
)
