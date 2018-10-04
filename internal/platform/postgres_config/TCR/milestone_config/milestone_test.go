package milestone_config

import (
  "testing"
  "BigBang/internal/platform/postgres_config/client_config"
  "github.com/stretchr/testify/suite"
  "BigBang/internal/pkg/error_config"
)

const ProjectId1 = "ProjectId1"
const MilestoneId1 = 1
const MilestoneId2 = 2


var MilestoneRecord1 = MilestoneRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId1,
  Content: "123",
  StartTime: 1000,
  EndTime: 2000,
  NumObjs: 5,
  AvgRating: 10,
}

var MilestoneRecord2 = MilestoneRecord {
  ProjectId: ProjectId1,
  MilestoneId: MilestoneId2,
  Content: "456",
  StartTime: 2000,
  EndTime: 3000,
  NumObjs: 10,
  AvgRating: 20,
}

type MilestoneTestSuite struct {
  suite.Suite
  MilestoneExecutor MilestoneExecutor
}

func (suite *MilestoneTestSuite) SetupSuite() {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  suite.MilestoneExecutor = MilestoneExecutor{*postgresBigBangClient}
  suite.MilestoneExecutor.DeleteMilestoneTable()
  suite.MilestoneExecutor.CreateMilestoneTable()
}

func (suite *MilestoneTestSuite) TearDownSuite() {
  suite.MilestoneExecutor.DeleteMilestoneTable()
  suite.MilestoneExecutor.C.Close()
}

func (suite *MilestoneTestSuite) SetupTest() {
  suite.MilestoneExecutor.ClearMilestoneTable()
}

func (suite *MilestoneTestSuite) TestEmptyQueryForGetMilestoneRecordByIDs() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoMilestoneIdExisting, message.ErrorCode)
    }
  }()
  suite.MilestoneExecutor.GetMilestoneRecordByIDs(ProjectId1, MilestoneId1)
}

func (suite *MilestoneTestSuite) TestNonEmptyQueryForGetMilestoneRecordByIDs() {
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord1)
  objectiveRecord := suite.MilestoneExecutor.GetMilestoneRecordByIDs(ProjectId1, MilestoneId1)
  suite.Equal(MilestoneRecord1.ProjectId, objectiveRecord.ProjectId)
  suite.Equal(MilestoneRecord1.MilestoneId, objectiveRecord.MilestoneId)
  suite.Equal(MilestoneRecord1.MilestoneId, objectiveRecord.MilestoneId)
  suite.Equal(MilestoneRecord1.Content, objectiveRecord.Content)
  suite.Equal(MilestoneRecord1.StartTime, objectiveRecord.StartTime)
  suite.Equal(MilestoneRecord1.EndTime, objectiveRecord.EndTime)
  suite.Equal(MilestoneRecord1.NumObjs, objectiveRecord.NumObjs)
  suite.Equal(MilestoneRecord1.AvgRating, objectiveRecord.AvgRating)
}

func (suite *MilestoneTestSuite) TestEmptyQueryForGetMilestoneRecordsByProjectId() {
  listMilestoneUUDs := suite.MilestoneExecutor.GetMilestonesRecordsByProjectId(ProjectId1)
  suite.Equal(0, len(*listMilestoneUUDs))
}

func (suite *MilestoneTestSuite) TestUpsertMilestoneRecord() {
  defer func() {
    errPanic := recover();
    suite.Nil(errPanic)
  }()
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord1)
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord2)
}

func (suite *MilestoneTestSuite) TestNonEmptyQueryForGetMilestoneRecordsByProjectId() {
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord1)
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord2)
  expectedMilestoneRecords := []MilestoneRecord {MilestoneRecord1, MilestoneRecord2}
  objectiveRecords:= suite.MilestoneExecutor.GetMilestonesRecordsByProjectId(ProjectId1)
  suite.Equal(len(expectedMilestoneRecords), len(*objectiveRecords))
  for index, objectiveRecord := range *objectiveRecords {
    suite.Equal(expectedMilestoneRecords[index].ProjectId, objectiveRecord.ProjectId)
    suite.Equal(expectedMilestoneRecords[index].MilestoneId, objectiveRecord.MilestoneId)
    suite.Equal(expectedMilestoneRecords[index].MilestoneId, objectiveRecord.MilestoneId)
    suite.Equal(expectedMilestoneRecords[index].Content, objectiveRecord.Content)
    suite.Equal(expectedMilestoneRecords[index].StartTime, objectiveRecord.StartTime)
    suite.Equal(expectedMilestoneRecords[index].EndTime, objectiveRecord.EndTime)
    suite.Equal(expectedMilestoneRecords[index].NumObjs, objectiveRecord.NumObjs)
    suite.Equal(expectedMilestoneRecords[index].AvgRating, objectiveRecord.AvgRating)
  }
}

func (suite *MilestoneTestSuite) TestVerifyNonExitingMilestoneUUID() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
       message := error_config.CreatedErrorInfoFromString(errPanic)
       suite.Equal(error_config.NoMilestoneIdExisting, message.ErrorCode)
    }
  }()
  suite.MilestoneExecutor.VerifyMilestoneRecordExisting(ProjectId1, MilestoneId1)
}

func (suite *MilestoneTestSuite) TestDeleteMilestoneRecordByIDs() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoMilestoneIdExisting, message.ErrorCode)
    }
  }()
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord1)
  suite.MilestoneExecutor.DeleteMilestoneRecordByIDs(ProjectId1, MilestoneId1)
  suite.MilestoneExecutor.VerifyMilestoneRecordExisting(ProjectId1, MilestoneId1)
}


func (suite *MilestoneTestSuite) TestDeleteMilestoneRecordByProjectId() {
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord1)
  suite.MilestoneExecutor.UpsertMilestoneRecord(&MilestoneRecord2)
  suite.MilestoneExecutor.DeleteMilestoneRecordsByProjectId(ProjectId1)
  objectivesRecords := suite.MilestoneExecutor.GetMilestonesRecordsByProjectId(ProjectId1)
  suite.Equal(0, len(*objectivesRecords))
}

func TestMilestoneTestSuite(t *testing.T) {
  suite.Run(t, new(MilestoneTestSuite))
}
