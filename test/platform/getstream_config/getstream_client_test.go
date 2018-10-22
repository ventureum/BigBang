package main

import (
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/app/feed_attributes"
  "testing"
  "github.com/stretchr/testify/suite"
  "BigBang/test/constants"
  "time"
  "sort"
)

type GetstreamConfigTestSuite struct {
  suite.Suite
  GetStreamClient *getstream_config.GetStreamClient
  ActivityOne *feed_attributes.Activity
  ActivityTwo *feed_attributes.Activity
  ActivityThree *feed_attributes.Activity
  ActivityFour *feed_attributes.Activity
  ActivityFive *feed_attributes.Activity
}

func (suite *GetstreamConfigTestSuite) SetupSuite() {
  suite.GetStreamClient = getstream_config.ConnectGetStreamClient()
  suite.ActivityOne = test_constants.Activity1
  suite.ActivityTwo = test_constants.Activity2
  suite.ActivityThree = test_constants.Activity3
  suite.ActivityFour = test_constants.Activity4
  suite.ActivityFive = test_constants.Activity5
}


func (suite *GetstreamConfigTestSuite) TestAddFeedActivity() {
suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityOne)
activity := suite.GetStreamClient.GetFeedActivityByForeignIdAndTimestamp(
  suite.ActivityOne.ForeignId, time.Unix(int64(suite.ActivityOne.Time), 0).UTC())

suite.Equal(suite.ActivityOne, activity)
}

func (suite *GetstreamConfigTestSuite) TestRemoveFeedActivityByForeignIdAndActor() {
suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityTwo)

activities := suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
  string(feed_attributes.UserPostFeedSlug), string(suite.ActivityTwo.Actor))

suite.Equal(1, len(*activities))
suite.Equal(*suite.ActivityTwo, (*activities)[0])

suite.GetStreamClient.RemoveFeedActivityByFeedSlugAndUserIdAndForeignId(
  string(feed_attributes.UserPostFeedSlug), string(suite.ActivityTwo.Actor), suite.ActivityTwo.ForeignId)

activities = suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
  string(feed_attributes.UserPostFeedSlug), string(suite.ActivityTwo.Actor))

suite.Equal(0, len(*activities))
}

func (suite *GetstreamConfigTestSuite) TestUpdateFeedPostRewardsByForeignIdAndTimestamp() {
suite.ActivityTwo.Extra["rewards"] = int64(20)
suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityTwo)

activities := suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
  string(feed_attributes.UserPostFeedSlug), string(suite.ActivityTwo.Actor))

suite.Equal(1, len(*activities))
suite.Equal(*suite.ActivityTwo, (*activities)[0])

newRewards := int64(30)
suite.ActivityTwo.Extra["rewards"] = newRewards
timestamp := time.Unix(int64(suite.ActivityTwo.Time), 0)
suite.GetStreamClient.UpdateFeedPostRewardsByForeignIdAndTimestamp(
  suite.ActivityTwo.ForeignId,
  timestamp,
  newRewards)
activities = suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
  string(feed_attributes.UserPostFeedSlug), string(suite.ActivityTwo.Actor))
suite.Equal(1, len(*activities))
suite.Equal(*suite.ActivityTwo, (*activities)[0])
}


func (suite *GetstreamConfigTestSuite) TestUpdateFeedPostRewardsByForeignIdAndTimestampWithNoRewardsDeclared() {
suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityThree)

activities := suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
  string(feed_attributes.UserFeedSlug), string(suite.ActivityThree.Actor))

suite.Equal(1, len(*activities))
suite.Equal(*suite.ActivityThree, (*activities)[0])

newRewards := int64(30)
suite.ActivityThree.Extra["rewards"] = newRewards
timestamp := time.Unix(int64(suite.ActivityThree.Time), 0)
suite.GetStreamClient.UpdateFeedPostRewardsByForeignIdAndTimestamp(
  suite.ActivityThree.ForeignId,
  timestamp,
  newRewards)
activities = suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
  string(feed_attributes.UserPostFeedSlug), string(suite.ActivityThree.Actor))
suite.Equal(1, len(*activities))
suite.Equal(*suite.ActivityThree, (*activities)[0])
}


func (suite *GetstreamConfigTestSuite) TestFeedActivityPostTypeWithTo() {
  suite.ActivityFour.To = []feed_attributes.FeedId{
    {
      FeedSlug: feed_attributes.BoardFeedSlug,
      UserId: feed_attributes.AllBoardId,
    },
    {
      FeedSlug: feed_attributes.BoardFeedSlug,
      UserId: feed_attributes.UserId(test_constants.BoardId1),
    },
  }

  sort.Slice(suite.ActivityFour.To, func(i, j int) bool {
    return  suite.ActivityFour.To[i].Value() <  suite.ActivityFour.To[j].Value()
  })
  suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityFour)


  activities := suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
    string(feed_attributes.UserPostFeedSlug), string(suite.ActivityFour.Actor))

  suite.ActivityFour.To = []feed_attributes.FeedId{}
  actualActivity := (*activities)[0]
  actualActivity.To = []feed_attributes.FeedId{}

  suite.Equal(1, len(*activities))
  suite.Equal(*suite.ActivityFour, actualActivity)

  activity := suite.GetStreamClient.GetFeedActivityByForeignIdAndTimestamp(
    suite.ActivityFour.ForeignId, time.Unix(int64(suite.ActivityFour.Time), 0).UTC())

  activity.To = []feed_attributes.FeedId{}
  suite.Equal(*suite.ActivityFour, *activity)
}

func (suite *GetstreamConfigTestSuite) TestFeedActivityCommentTypeWithTo() {

  sort.Slice(suite.ActivityFive.To, func(i, j int) bool {
    return  suite.ActivityFive.To[i].Value() <  suite.ActivityFive.To[j].Value()
  })
  suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityFive)


  activities := suite.GetStreamClient.GetAllFeedActivitiesByFeedSlugAndUserId(
    string(feed_attributes.UserCommentFeedSlug), string(feed_attributes.UserId(suite.ActivityFive.Actor)))

  suite.ActivityFive.To = []feed_attributes.FeedId{}
  actualActivity := (*activities)[0]
  actualActivity.To = []feed_attributes.FeedId{}

  suite.Equal(1, len(*activities))
  suite.Equal(*suite.ActivityFive, actualActivity)

  activity := suite.GetStreamClient.GetFeedActivityByForeignIdAndTimestamp(
    suite.ActivityFive.ForeignId, time.Unix(int64(suite.ActivityFive.Time), 0).UTC())

  suite.ActivityFive.To = []feed_attributes.FeedId{}
  suite.Equal(*suite.ActivityFive, *activity)
}

func TestGetstreamConfigTestSuite(t *testing.T) {
  suite.Run(t, new(GetstreamConfigTestSuite))
}
