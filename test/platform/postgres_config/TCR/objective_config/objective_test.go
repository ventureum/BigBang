package objective_config

import (
  "testing"
  "BigBang/internal/platform/postgres_config/client_config"
  "github.com/stretchr/testify/suite"
  "BigBang/internal/pkg/error_config"
)

const ProjectId1 = "ProjectId1"
const MilestoneId1 = 1
const ObjectiveId1 = 1
const ObjectiveId2 = 2

var ObjectRecord1 = ObjectiveRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId1,
  Content: "123",
  TotalRating: 10,
  TotalWeight: 10,
}

var ObjectRecord2 = ObjectiveRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  ObjectiveId: ObjectiveId2,
  Content: "123",
  TotalRating: 20,
  TotalWeight: 20,
}

type ObjectiveTestSuite struct {
  suite.Suite
  ObjectiveExecutor ObjectiveExecutor
}

func (suite *ObjectiveTestSuite) SetupSuite() {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  suite.ObjectiveExecutor = ObjectiveExecutor{*postgresBigBangClient}
  suite.ObjectiveExecutor.DeleteObjectiveTable()
  suite.ObjectiveExecutor.CreateObjectiveTable()
}

func (suite *ObjectiveTestSuite) TearDownSuite() {
  suite.ObjectiveExecutor.DeleteObjectiveTable()
  suite.ObjectiveExecutor.C.Close()
}

func (suite *ObjectiveTestSuite) SetupTest() {
  suite.ObjectiveExecutor.ClearObjectiveTable()
}

func (suite *ObjectiveTestSuite) TestEmptyQueryForGetObjectiveRecordByIDs() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoObjectiveIdExisting, message.ErrorCode)
    }
  }()
  suite.ObjectiveExecutor.GetObjectiveRecordByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
}

func (suite *ObjectiveTestSuite) TestNonEmptyQueryForGetObjectiveRecordByIDs() {
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord1)
  objectiveRecord := suite.ObjectiveExecutor.GetObjectiveRecordByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
  suite.Equal(ObjectRecord1.ProjectId, objectiveRecord.ProjectId)
  suite.Equal(ObjectRecord1.MilestoneId, objectiveRecord.MilestoneId)
  suite.Equal(ObjectRecord1.ObjectiveId, objectiveRecord.ObjectiveId)
  suite.Equal(ObjectRecord1.Content, objectiveRecord.Content)
  suite.Equal(ObjectRecord1.TotalRating, objectiveRecord.TotalRating)
  suite.Equal(ObjectRecord1.TotalWeight, objectiveRecord.TotalWeight)
}

func (suite *ObjectiveTestSuite) TestEmptyQueryForGetObjectiveRecordsByProjectIdAndMilestoneId() {
  objectiveRecords := suite.ObjectiveExecutor.GetObjectiveRecordsByProjectIdAndMilestoneId(ProjectId1, MilestoneId1)
  suite.Equal(0, len(*objectiveRecords))
}

func (suite *ObjectiveTestSuite) TestUpsertObjectiveRecord() {
  defer func() {
    errPanic := recover();
    suite.Nil(errPanic)
  }()
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord1)
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord2)
}

func (suite *ObjectiveTestSuite) TestNonEmptyQueryForGetObjectiveRecordsByProjectIdAndMilestoneId() {
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord1)
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord2)
  expectedObjectiveRecords := []ObjectiveRecord {ObjectRecord1, ObjectRecord2}
  objectiveRecords:= suite.ObjectiveExecutor.GetObjectiveRecordsByProjectIdAndMilestoneId(ProjectId1, MilestoneId1)
  suite.Equal(len(expectedObjectiveRecords), len(*objectiveRecords))
  for index, objectiveRecord := range *objectiveRecords {
    suite.Equal(expectedObjectiveRecords[index].ProjectId, objectiveRecord.ProjectId)
    suite.Equal(expectedObjectiveRecords[index].MilestoneId, objectiveRecord.MilestoneId)
    suite.Equal(expectedObjectiveRecords[index].ObjectiveId, objectiveRecord.ObjectiveId)
    suite.Equal(expectedObjectiveRecords[index].Content, objectiveRecord.Content)
    suite.Equal(expectedObjectiveRecords[index].TotalRating, objectiveRecord.TotalRating)
    suite.Equal(expectedObjectiveRecords[index].TotalWeight, objectiveRecord.TotalWeight)
  }
}

func (suite *ObjectiveTestSuite) TestVerifyNonExitingObjectiveId() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
       message := error_config.CreatedErrorInfoFromString(errPanic)
       suite.Equal(error_config.NoObjectiveIdExisting, message.ErrorCode)
    }
  }()
  suite.ObjectiveExecutor.VerifyObjectiveRecordExisting(ProjectId1, MilestoneId1, ObjectiveId1)
}

func (suite *ObjectiveTestSuite) TestDeleteObjectiveRecordByIDs() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoObjectiveIdExisting, message.ErrorCode)
    }
  }()
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord1)
  suite.ObjectiveExecutor.DeleteObjectiveRecordByIDs(ProjectId1, MilestoneId1, ObjectiveId1)
  suite.ObjectiveExecutor.VerifyObjectiveRecordExisting(ProjectId1, MilestoneId1, ObjectiveId1)
}


func (suite *ObjectiveTestSuite) TestDeleteObjectiveRecordByProjectIdAndMilestoneId() {
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord1)
  suite.ObjectiveExecutor.UpsertObjectiveRecord(&ObjectRecord2)
  suite.ObjectiveExecutor.DeleteObjectiveRecordsByProjectIdAndMilestoneId(ProjectId1, MilestoneId1)
  objectivesRecords := suite.ObjectiveExecutor.GetObjectiveRecordsByProjectIdAndMilestoneId(ProjectId1, MilestoneId1)
  suite.Equal(0, len(*objectivesRecords))
}

func TestObjectiveTestSuite(t *testing.T) {
  suite.Run(t, new(ObjectiveTestSuite))
}
