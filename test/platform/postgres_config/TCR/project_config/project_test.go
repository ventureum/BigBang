package project_config

import (
  "BigBang/internal/pkg/error_config"
  "github.com/stretchr/testify/suite"
  "BigBang/internal/platform/postgres_config/client_config"
  "testing"
)

const ProjectId1 = "ProjectId1"
const ProjectId2 = "ProjectId2"
const ProjectId3 = "ProjectId3"
const ProjectId4 = "ProjectId4"
const ProjectId5 = "ProjectId5"

const Admin1 = "Admin1"
const MilestoneId1 = 1
const MilestoneId2 = 2
const MilestoneId3 = 3
const ObjectiveId1 = 1
const Content1 = "Conent1"
const Voter1 = "Voter1"
const Voter2 = "Voter2"
const Voter3 = "Voter3"
const Voter4 = "Voter4"
const Voter5 = "Voter5"


var ProjectRecord1 = ProjectRecord {
  ProjectId: ProjectId1,
  Admin: Admin1,
  Content: Content1,
}

var ProjectRecord2 = ProjectRecord {
  ProjectId: ProjectId2,
  Admin: Admin1,
  Content: Content1,
}

var ProjectRecord3 = ProjectRecord {
  ProjectId: ProjectId3,
  Admin: Admin1,
  Content: Content1,
}

var ProjectRecord4 = ProjectRecord {
  ProjectId: ProjectId4,
  Admin: Admin1,
  Content: Content1,
}

var ProjectRecord5 = ProjectRecord {
  ProjectId: ProjectId5,
  Admin: Admin1,
  Content: Content1,
}

type ProjectTestSuite struct {
  suite.Suite
  ProjectExecutor ProjectExecutor
}

func (suite *ProjectTestSuite) SetupSuite() {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  suite.ProjectExecutor = ProjectExecutor{*postgresBigBangClient}
  suite.ProjectExecutor.DeleteProjectTable()
  suite.ProjectExecutor.CreateProjectTable()
}

func (suite *ProjectTestSuite) TearDownSuite() {
  suite.ProjectExecutor.DeleteProjectTable()
  suite.ProjectExecutor.C.Close()
}

func (suite *ProjectTestSuite) SetupTest() {
  suite.ProjectExecutor.ClearProjectTable()
}

func (suite *ProjectTestSuite) TestEmptyQueryForGetProjectRecord() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoProjectIdExisting, message.ErrorCode)
    }
  }()
  suite.ProjectExecutor.GetProjectRecord(ProjectId1)
}

func (suite *ProjectTestSuite) TestNonEmptyQueryForGetProjectRecord() {
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)
  projectRecord := suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(ProjectRecord1.ProjectId, projectRecord.ProjectId)
  suite.Equal(ProjectRecord1.Admin, projectRecord.Admin)
  suite.Equal(ProjectRecord1.AvgRating, projectRecord.AvgRating)
  suite.Equal(ProjectRecord1.Content, projectRecord.Content)
  suite.Equal(ProjectRecord1.TotalWeight, projectRecord.TotalWeight)
  suite.Equal(ProjectRecord1.TotalRating, projectRecord.TotalRating)
  suite.Equal(ProjectRecord1.CurrentMilestone, projectRecord.CurrentMilestone)
  suite.Equal(ProjectRecord1.NumMilestones, projectRecord.NumMilestones)
  suite.Equal(ProjectRecord1.NumMilestonesCompleted, projectRecord.NumMilestonesCompleted)
}

func (suite *ProjectTestSuite) TestUpsertProjectRecord() {
  defer func() {
    errPanic := recover();
    suite.Nil(errPanic)
  }()
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord2)
}


func (suite *ProjectTestSuite) TestVerifyNonExitingProjectVoter() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoProjectIdExisting, message.ErrorCode)
    }
  }()
  suite.ProjectExecutor.VerifyProjectRecordExisting(ProjectId1)
}

func (suite *ProjectTestSuite) TestDeleteProjectRecord() {
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      suite.Equal(error_config.NoProjectIdExisting, message.ErrorCode)
    }
  }()
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)
  suite.ProjectExecutor.DeleteProjectRecord(ProjectId1)
  suite.ProjectExecutor.VerifyProjectRecordExisting(ProjectId1)
}

func (suite *ProjectTestSuite) TestNonEmptyQueryForGetProjectRecordsByCursorFirstQuery() {
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord2)

  expectedProjectRecords := []ProjectRecord {ProjectRecord2, ProjectRecord1}
  projectRecords := suite.ProjectExecutor.GetProjectRecordsByCursor( 0, 100)

  suite.Equal(len(expectedProjectRecords), len(*projectRecords))
  for index, projectRecord:= range *projectRecords {
    suite.Equal(expectedProjectRecords[index].ProjectId, projectRecord.ProjectId)
    suite.Equal(expectedProjectRecords[index].Admin, projectRecord.Admin)
    suite.Equal(expectedProjectRecords[index].AvgRating, projectRecord.AvgRating)
    suite.Equal(expectedProjectRecords[index].Content, projectRecord.Content)
    suite.Equal(expectedProjectRecords[index].TotalWeight, projectRecord.TotalWeight)
    suite.Equal(expectedProjectRecords[index].TotalRating, projectRecord.TotalRating)
    suite.Equal(expectedProjectRecords[index].CurrentMilestone, projectRecord.CurrentMilestone)
    suite.Equal(expectedProjectRecords[index].NumMilestones, projectRecord.NumMilestones)
    suite.Equal(expectedProjectRecords[index].NumMilestonesCompleted, projectRecord.NumMilestonesCompleted)
  }
}

func (suite *ProjectTestSuite) TestNonEmptyQueryForGetProjectRecordsByCursorInterQuery() {
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord2)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord3)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord4)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord5)

  expectedProjectRecords := []ProjectRecord {ProjectRecord4, ProjectRecord3}
  projectRecords := suite.ProjectExecutor.GetProjectRecordsByCursor(4, 2)

  suite.Equal(len(expectedProjectRecords), len(*projectRecords))
  for index, projectRecord:= range *projectRecords {
    suite.Equal(expectedProjectRecords[index].ProjectId, projectRecord.ProjectId)
    suite.Equal(expectedProjectRecords[index].Admin, projectRecord.Admin)
    suite.Equal(expectedProjectRecords[index].AvgRating, projectRecord.AvgRating)
    suite.Equal(expectedProjectRecords[index].Content, projectRecord.Content)
    suite.Equal(ProjectRecord1.TotalWeight, projectRecord.TotalWeight)
    suite.Equal(ProjectRecord1.TotalRating, projectRecord.TotalRating)
    suite.Equal(expectedProjectRecords[index].CurrentMilestone, projectRecord.CurrentMilestone)
    suite.Equal(expectedProjectRecords[index].NumMilestones, projectRecord.NumMilestones)
    suite.Equal(expectedProjectRecords[index].NumMilestonesCompleted, projectRecord.NumMilestonesCompleted)
  }
}

func (suite *ProjectTestSuite) TestNonEmptyQueryForGetProjectRecordsByCursorFinalQuery() {
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord2)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord3)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord4)
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord5)

  expectedProjectRecords := []ProjectRecord {ProjectRecord3, ProjectRecord2, ProjectRecord1}
  projectRecords := suite.ProjectExecutor.GetProjectRecordsByCursor(3, 6)

  suite.Equal(len(expectedProjectRecords), len(*projectRecords))
  for index, projectRecord:= range *projectRecords {
    suite.Equal(expectedProjectRecords[index].ProjectId, projectRecord.ProjectId)
    suite.Equal(expectedProjectRecords[index].Admin, projectRecord.Admin)
    suite.Equal(expectedProjectRecords[index].AvgRating, projectRecord.AvgRating)
    suite.Equal(expectedProjectRecords[index].Content, projectRecord.Content)
    suite.Equal(expectedProjectRecords[index].TotalWeight, projectRecord.TotalWeight)
    suite.Equal(expectedProjectRecords[index].TotalRating, projectRecord.TotalRating)
    suite.Equal(expectedProjectRecords[index].CurrentMilestone, projectRecord.CurrentMilestone)
    suite.Equal(expectedProjectRecords[index].NumMilestones, projectRecord.NumMilestones)
    suite.Equal(expectedProjectRecords[index].NumMilestonesCompleted, projectRecord.NumMilestonesCompleted)
  }
}

func (suite *ProjectTestSuite) TestMilestoneInfo() {
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)

  projectRecord := suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(0) , projectRecord.NumMilestones)
  suite.Equal(int64(0), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(0), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.IncreaseNumMilestones(ProjectId1)
  suite.ProjectExecutor.IncreaseNumMilestones(ProjectId1)
  suite.ProjectExecutor.IncreaseNumMilestones(ProjectId1)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(0), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(0), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.SetCurrentMilestone(ProjectId1, MilestoneId1)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(0), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(1), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.IncreaseNumMilestonesCompleted(ProjectId1)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(1), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(0), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.SetCurrentMilestone(ProjectId1, MilestoneId2)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(1), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(2), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.IncreaseNumMilestonesCompleted(ProjectId1)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(2), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(0), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.SetCurrentMilestone(ProjectId1, MilestoneId3)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(2), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(3), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.IncreaseNumMilestonesCompleted(ProjectId1)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(3), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(0), projectRecord.CurrentMilestone)

  suite.ProjectExecutor.IncreaseNumMilestonesCompleted(ProjectId1)
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(3), projectRecord.NumMilestones)
  suite.Equal(int64(3), projectRecord.NumMilestonesCompleted)
  suite.Equal(int64(0), projectRecord.CurrentMilestone)
}

func (suite *ProjectTestSuite) TestAddRatingAndWeight() {
  suite.ProjectExecutor.UpsertProjectRecord(&ProjectRecord1)

  projectRecord := suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(0) , projectRecord.AvgRating)
  suite.Equal(int64(0), projectRecord.TotalRating)
  suite.Equal(int64(0), projectRecord.TotalWeight)

  delatRating := 30
  deltaWeight := 20
  suite.ProjectExecutor.AddRatingAndWeight(ProjectId1, int64(delatRating), int64(deltaWeight))
  projectRecord = suite.ProjectExecutor.GetProjectRecord(ProjectId1)
  suite.Equal(int64(delatRating/deltaWeight) , projectRecord.AvgRating)
  suite.Equal(int64(delatRating) , projectRecord.TotalRating)
  suite.Equal(int64(deltaWeight), projectRecord.TotalWeight)
}

func TestProjectTestSuite(t *testing.T) {
  suite.Run(t, new(ProjectTestSuite))
}
