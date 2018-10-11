package rating_vote_config

import (
  "BigBang/internal/pkg/error_config"
  "github.com/stretchr/testify/suite"
  "BigBang/internal/platform/postgres_config/client_config"
  "testing"
)

const ProjectId1 = "ProjectId1"
const MilestoneId1 = 1
const ObjectiveId1 = 1
const Voter1 = "Voter1"
const Voter2 = "Voter2"
const Voter3 = "Voter3"
const Voter4 = "Voter4"
const Voter5 = "Voter5"


var RatingVoteRecord1 = RatingVoteRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId1,
  Voter: Voter1,
  Rating: 10,
  Weight: 10,
}

var RatingVoteRecord2 = RatingVoteRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId1,
  Voter: Voter2,
  Rating: 10,
  Weight: 10,
}

var RatingVoteRecord3 = RatingVoteRecord {
  ID: 1,
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId1,
  Voter: Voter3,
  Rating: 10,
  Weight: 10,
}

var RatingVoteRecord4 = RatingVoteRecord {
  ID: 2,
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId1,
  Voter: Voter4,
  Rating: 10,
  Weight: 10,
}

var RatingVoteRecord5 = RatingVoteRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId1,
  Voter: Voter5,
  Rating: 10,
  Weight: 10,
}

type RatingVoteTestSuite struct {
  suite.Suite
  RatingVoteExecutor RatingVoteExecutor
}

func (suite *RatingVoteTestSuite) SetupSuite() {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  suite.RatingVoteExecutor = RatingVoteExecutor{*postgresBigBangClient}
  suite.RatingVoteExecutor.DeleteRatingVoteTable()
  suite.RatingVoteExecutor.CreateRatingVoteTable()
}

func (suite *RatingVoteTestSuite) TearDownSuite() {
  suite.RatingVoteExecutor.DeleteRatingVoteTable()
  suite.RatingVoteExecutor.C.Close()
}

func (suite *RatingVoteTestSuite) SetupTest() {
  suite.RatingVoteExecutor.ClearRatingVoteTable()
}

func (suite *RatingVoteTestSuite) TestEmptyQueryForGetRatingVoteRecordByIDsAndVoter() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoRatingVoteVoterExisting, message.ErrorCode)
    }
  }()
  suite.RatingVoteExecutor.GetRatingVoteRecordByIDsAndVoter(ProjectId1, MilestoneId1, ObjectiveId1, Voter1)
}

func (suite *RatingVoteTestSuite) TestNonEmptyQueryForGetRatingVoteRecordByIDsAndVoter() {
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  ratingVoteRecord := suite.RatingVoteExecutor.GetRatingVoteRecordByIDsAndVoter(ProjectId1, MilestoneId1, ObjectiveId1, Voter1)
  suite.Equal(RatingVoteRecord1.ProjectId, ratingVoteRecord.ProjectId)
  suite.Equal(RatingVoteRecord1.MilestoneId, ratingVoteRecord.MilestoneId)
  suite.Equal(RatingVoteRecord1.ObjectiveId, ratingVoteRecord.ObjectiveId)
  suite.Equal(RatingVoteRecord1.Voter, ratingVoteRecord.Voter)
  suite.Equal(RatingVoteRecord1.Rating, ratingVoteRecord.Rating)
  suite.Equal(RatingVoteRecord1.Weight, ratingVoteRecord.Weight)
}

func (suite *RatingVoteTestSuite) TestEmptyQueryForGetRatingVoteRecordsByIDs() {
  listObjectiveUUDs := suite.RatingVoteExecutor.GetRatingVoteRecordsByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
  suite.Equal(0, len(*listObjectiveUUDs))
}

func (suite *RatingVoteTestSuite) TestUpsertRatingVoteRecord() {
  defer func() {
    errPanic := recover();
    suite.Nil(errPanic)
  }()
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord2)
}

func (suite *RatingVoteTestSuite) TestNonEmptyQueryForGetRatingVoteRecordsByIDs() {
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord2)
  expectedRatingVoteRecords := []RatingVoteRecord {RatingVoteRecord1, RatingVoteRecord2}
  ratingVoteRecords:= suite.RatingVoteExecutor.GetRatingVoteRecordsByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
  suite.Equal(len(expectedRatingVoteRecords), len(*ratingVoteRecords))
  for index, ratingVoteRecord := range *ratingVoteRecords {
    suite.Equal(expectedRatingVoteRecords[index].ProjectId, ratingVoteRecord.ProjectId)
    suite.Equal(expectedRatingVoteRecords[index].MilestoneId, ratingVoteRecord.MilestoneId)
    suite.Equal(expectedRatingVoteRecords[index].ObjectiveId, ratingVoteRecord.ObjectiveId)
    suite.Equal(expectedRatingVoteRecords[index].Voter, ratingVoteRecord.Voter)
    suite.Equal(expectedRatingVoteRecords[index].Rating, ratingVoteRecord.Rating)
    suite.Equal(expectedRatingVoteRecords[index].Weight, ratingVoteRecord.Weight)
  }
}

func (suite *RatingVoteTestSuite) TestVerifyNonExitingRatingVoteVoter() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoRatingVoteVoterExisting, message.ErrorCode)
    }
  }()
  suite.RatingVoteExecutor.VerifyRatingVoteRecordExisting(ProjectId1, MilestoneId1, ObjectiveId1, Voter1)
}

func (suite *RatingVoteTestSuite) TestDeleteRatingVoteRecordByIDsAndVoter() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoRatingVoteVoterExisting, message.ErrorCode)
    }
  }()
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.DeleteRatingVoteRecordByIDsAndVoter(ProjectId1, MilestoneId1, ObjectiveId1, Voter1)
  suite.RatingVoteExecutor.VerifyRatingVoteRecordExisting(ProjectId1, MilestoneId1, ObjectiveId1, Voter1)
}


func (suite *RatingVoteTestSuite) TestDeleteRatingVoteRecordByIDs() {
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord2)
  suite.RatingVoteExecutor.DeleteRatingVoteRecordsByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
  objectivesRecords := suite.RatingVoteExecutor.GetRatingVoteRecordsByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
  suite.Equal(0, len(*objectivesRecords))
}

func (suite *RatingVoteTestSuite) TestNonEmptyQueryForGetRatingVoteRecordsByCursorFirstQuery() {
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord2)

  expectedRatingVoteRecords := []RatingVoteRecord {RatingVoteRecord2, RatingVoteRecord1}
  ratingVoteRecords := suite.RatingVoteExecutor.GetRatingVoteRecordsByCursor(
    ProjectId1, MilestoneId1, ObjectiveId1,0, 100)

  suite.Equal(len(expectedRatingVoteRecords), len(*ratingVoteRecords))
  for index, ratingVoteRecord:= range *ratingVoteRecords {
    suite.Equal(expectedRatingVoteRecords[index].ProjectId, ratingVoteRecord.ProjectId)
    suite.Equal(expectedRatingVoteRecords[index].MilestoneId, ratingVoteRecord.MilestoneId)
    suite.Equal(expectedRatingVoteRecords[index].ObjectiveId, ratingVoteRecord.ObjectiveId)
    suite.Equal(expectedRatingVoteRecords[index].Voter, ratingVoteRecord.Voter)
    suite.Equal(expectedRatingVoteRecords[index].Rating, ratingVoteRecord.Rating)
    suite.Equal(expectedRatingVoteRecords[index].Weight, ratingVoteRecord.Weight)
  }
}

func (suite *RatingVoteTestSuite) TestNonEmptyQueryForGetRatingVoteRecordsByCursorInterQuery() {
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord2)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord3)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord4)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord5)

  expectedRatingVoteRecords := []RatingVoteRecord {RatingVoteRecord4, RatingVoteRecord3}
  ratingVoteRecords := suite.RatingVoteExecutor.GetRatingVoteRecordsByCursor(
    ProjectId1, MilestoneId1, ObjectiveId1,4, 2)

  suite.Equal(len(expectedRatingVoteRecords), len(*ratingVoteRecords))
  for index, ratingVoteRecord:= range *ratingVoteRecords {
    suite.Equal(expectedRatingVoteRecords[index].ProjectId, ratingVoteRecord.ProjectId)
    suite.Equal(expectedRatingVoteRecords[index].MilestoneId, ratingVoteRecord.MilestoneId)
    suite.Equal(expectedRatingVoteRecords[index].ObjectiveId, ratingVoteRecord.ObjectiveId)
    suite.Equal(expectedRatingVoteRecords[index].Voter, ratingVoteRecord.Voter)
    suite.Equal(expectedRatingVoteRecords[index].Rating, ratingVoteRecord.Rating)
    suite.Equal(expectedRatingVoteRecords[index].Weight, ratingVoteRecord.Weight)
  }
}

func (suite *RatingVoteTestSuite) TestNonEmptyQueryForGetRatingVoteRecordsByCursorFinalQuery() {
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord1)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord2)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord3)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord4)
  suite.RatingVoteExecutor.UpsertRatingVoteRecord(&RatingVoteRecord5)

  expectedRatingVoteRecords := []RatingVoteRecord {RatingVoteRecord3, RatingVoteRecord2, RatingVoteRecord1}
  ratingVoteRecords := suite.RatingVoteExecutor.GetRatingVoteRecordsByCursor(
    ProjectId1, MilestoneId1, ObjectiveId1,3, 6)

  suite.Equal(len(expectedRatingVoteRecords), len(*ratingVoteRecords))
  for index, ratingVoteRecord := range *ratingVoteRecords {
   suite.Equal(expectedRatingVoteRecords[index].ProjectId, ratingVoteRecord.ProjectId)
   suite.Equal(expectedRatingVoteRecords[index].MilestoneId, ratingVoteRecord.MilestoneId)
   suite.Equal(expectedRatingVoteRecords[index].ObjectiveId, ratingVoteRecord.ObjectiveId)
   suite.Equal(expectedRatingVoteRecords[index].Voter, ratingVoteRecord.Voter)
   suite.Equal(expectedRatingVoteRecords[index].Rating, ratingVoteRecord.Rating)
   suite.Equal(expectedRatingVoteRecords[index].Weight, ratingVoteRecord.Weight)
  }
}


func TestRatingVoteTestSuite(t *testing.T) {
  suite.Run(t, new(RatingVoteTestSuite))
}
