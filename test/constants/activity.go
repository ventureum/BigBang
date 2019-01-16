package test_constants

import "BigBang/internal/app/feed_attributes"

const ActivityPostHash1 = "0xActivityPostHash1"
const ActivityPostBlockTimestamp1 = 1539108020
const ActivityPostHash2 = "0xActivityPostHash2"
const ActivityPostBlockTimestamp2 = 1539107020
const ActivityPostHash3 = "0xActivityPostHash3"
const ActivityPostBlockTimestamp3 = 1539107120
const ActivityPostHash4 = "0xActivityPostHash4"
const ActivityPostBlockTimestamp4 = 1539107124
const ActivityPostHash5 = "0xActivityPostHash5"
const ActivityPostBlockTimestamp5 = 1539107128

var Activity1 = feed_attributes.CreateNewActivity(
	feed_attributes.Actor(Actor1),
	feed_attributes.SubmitVerb,
	feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash1),
	ActivityPostBlockTimestamp1,
	feed_attributes.PostPostType,
	[]feed_attributes.FeedId{},
	map[string]interface{}{
		"rewards": int64(0),
	},
)

var Activity2 = feed_attributes.CreateNewActivity(
	feed_attributes.Actor(Actor2),
	feed_attributes.SubmitVerb,
	feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash2),
	ActivityPostBlockTimestamp2,
	feed_attributes.PostPostType,
	[]feed_attributes.FeedId{},
	map[string]interface{}{
		"rewards": int64(0),
	},
)

var Activity3 = feed_attributes.CreateNewActivity(
	feed_attributes.Actor(Actor3),
	feed_attributes.SubmitVerb,
	feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash3),
	ActivityPostBlockTimestamp3,
	feed_attributes.PostPostType,
	[]feed_attributes.FeedId{},
	map[string]interface{}{
		"rewards": int64(0),
	},
)

var Activity4 = feed_attributes.CreateNewActivity(
	feed_attributes.Actor(Actor4),
	feed_attributes.SubmitVerb,
	feed_attributes.CreateObject(string(feed_attributes.PostObjectType), ActivityPostHash4),
	ActivityPostBlockTimestamp4,
	feed_attributes.PostPostType,
	[]feed_attributes.FeedId{},
	map[string]interface{}{
		"rewards": int64(0),
	},
)

var Activity5 = feed_attributes.CreateNewActivity(
	feed_attributes.Actor(Actor5),
	feed_attributes.ReplyVerb,
	feed_attributes.CreateObject(string(feed_attributes.ReplyPostType), ActivityPostHash5),
	ActivityPostBlockTimestamp5,
	feed_attributes.ReplyPostType,
	[]feed_attributes.FeedId{},
	map[string]interface{}{
		"rewards": int64(0),
	},
)
