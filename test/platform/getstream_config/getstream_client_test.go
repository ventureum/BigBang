package main

import (
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/app/feed_attributes"
  "testing"
  "github.com/stretchr/testify/suite"
  "BigBang/test/constants"
  "time"
)

type GetstreamConfigTestSuite struct {
  suite.Suite
  GetStreamClient *getstream_config.GetStreamClient
  ActivityOne *feed_attributes.Activity
  ActivityTwo *feed_attributes.Activity
  ActivityThree *feed_attributes.Activity
}

func (suite *GetstreamConfigTestSuite) SetupSuite() {
  suite.GetStreamClient = getstream_config.ConnectGetStreamClient()
  suite.ActivityOne = test_constants.Activity1
  suite.ActivityTwo = test_constants.Activity2
  suite.ActivityThree = test_constants.Activity3
}


func (suite *GetstreamConfigTestSuite) TestAddFeedActivity() {
 suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityOne)
 activity := suite.GetStreamClient.GetFeedActivityByForeignIdAndTimestamp(
   suite.ActivityOne.ForeignId, time.Unix(int64(suite.ActivityOne.Time), 0).UTC())

 suite.Equal(suite.ActivityOne, activity)
}

func (suite *GetstreamConfigTestSuite) TestRemoveFeedActivityByForeignIdAndActor() {
 suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityTwo)

 activities := suite.GetStreamClient.GetAllFeedActivitiesByActor(string(suite.ActivityTwo.Actor))

 suite.Equal(1, len(*activities))
 suite.Equal(*suite.ActivityTwo, (*activities)[0])

 suite.GetStreamClient.RemoveFeedActivityByForeignIdAndActor(suite.ActivityTwo.ForeignId, string(suite.ActivityTwo.Actor))

 activities = suite.GetStreamClient.GetAllFeedActivitiesByActor(string(suite.ActivityTwo.Actor))

 suite.Equal(0, len(*activities))
}

func (suite *GetstreamConfigTestSuite) TestUpdateFeedPostRewardsByForeignIdAndTimestamp() {
 suite.ActivityTwo.Extra["rewards"] = int64(20)
 suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityTwo)

 activities := suite.GetStreamClient.GetAllFeedActivitiesByActor(string(suite.ActivityTwo.Actor))

 suite.Equal(1, len(*activities))
 suite.Equal(*suite.ActivityTwo, (*activities)[0])

 newRewards := int64(30)
 suite.ActivityTwo.Extra["rewards"] = newRewards
 timestamp := time.Unix(int64(suite.ActivityTwo.Time), 0)
 suite.GetStreamClient.UpdateFeedPostRewardsByForeignIdAndTimestamp(
   suite.ActivityTwo.ForeignId,
   timestamp,
   newRewards)
 activities = suite.GetStreamClient.GetAllFeedActivitiesByActor(string(suite.ActivityTwo.Actor))
 suite.Equal(1, len(*activities))
 suite.Equal(*suite.ActivityTwo, (*activities)[0])
}


func (suite *GetstreamConfigTestSuite) TestUpdateFeedPostRewardsByForeignIdAndTimestampWithNoRewardsDeclared() {
 suite.GetStreamClient.AddFeedActivityToGetStream(suite.ActivityThree)

 activities := suite.GetStreamClient.GetAllFeedActivitiesByActor(string(suite.ActivityThree.Actor))

 suite.Equal(1, len(*activities))
 suite.Equal(*suite.ActivityThree, (*activities)[0])

 newRewards := int64(30)
 suite.ActivityThree.Extra["rewards"] = newRewards
 timestamp := time.Unix(int64(suite.ActivityThree.Time), 0)
 suite.GetStreamClient.UpdateFeedPostRewardsByForeignIdAndTimestamp(
   suite.ActivityThree.ForeignId,
   timestamp,
   newRewards)
 activities = suite.GetStreamClient.GetAllFeedActivitiesByActor(string(suite.ActivityThree.Actor))
 suite.Equal(1, len(*activities))
 suite.Equal(*suite.ActivityThree, (*activities)[0])
}


func TestGetstreamConfigTestSuite(t *testing.T) {
  suite.Run(t, new(GetstreamConfigTestSuite))
}
