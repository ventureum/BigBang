package TCR

import (
  "testing"
  "BigBang/internal/pkg/api"
  "github.com/stretchr/testify/assert"
  projectConfig "BigBang/cmd/lambda/TCR/new_project/config"
  getProjectConfig "BigBang/cmd/lambda/TCR/get_project/config"
  "github.com/mitchellh/mapstructure"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "log"
  "time"
)

func TestProjectAndGetProjectAlpha(t *testing.T) {
  requestNewProject := projectConfig.Request{
    ProjectId: ProjectId001,
    Content: "123",
    AvgRating: 20,
    MilestoneInfo: tcr_attributes.MilestonesInfo{
      NumMilestones: 5,
      NumMilestonesCompleted: 8,
    },
  }

  expectedResponseNewProject := projectConfig.Response{
    Ok: true,
  }

  responseMessageNewProject := api.SendPost(requestNewProject, api.NewProjectAlphaEndingPoint)

  var responseProject projectConfig.Response
  mapstructure.Decode(*responseMessageNewProject, &responseProject)
  assert.Equal(t, expectedResponseNewProject, responseProject)


  requestGetProject := getProjectConfig.Request{
    ProjectId: ProjectId001,
  }

  expectedResponseGetProject := getProjectConfig.Response{
    Ok: true,
    Project: &project_config.ProjectRecordResult{
      ProjectId: ProjectId001,
      Content:   "123",
      AvgRating: 20,
      MilestoneInfo: &tcr_attributes.MilestonesInfo{
        NumMilestones:          5,
        NumMilestonesCompleted: 8,
      },
    },
  }

  responseMessageGetProject := api.SendPost(requestGetProject, api.GetProjectAlphaEndingPoint)
  var responseGetProject getProjectConfig.Response


  config := mapstructure.DecoderConfig{
    DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339Nano),
    Result:     &responseGetProject,
  }
  decoder, err := mapstructure.NewDecoder(&config)
  if err != nil {
    log.Fatalf("error to create NewDecoder: %+v", err)
  }

  err = decoder.Decode(*responseMessageGetProject)
  if err != nil {
    log.Fatalf("error to Decode: %+v", err)
  }

  assert.Equal(t, responseGetProject.Ok, expectedResponseGetProject.Ok)
  assert.Equal(t, responseGetProject.Project.ProjectId, expectedResponseGetProject.Project.ProjectId)
  assert.Equal(t, responseGetProject.Project.Content, expectedResponseGetProject.Project.Content)
  assert.Equal(t, responseGetProject.Project.AvgRating, expectedResponseGetProject.Project.AvgRating)
  assert.Equal(t, responseGetProject.Project.MilestonesInfo, expectedResponseGetProject.Project.MilestonesInfo)
}

func TestProjectAndGetProjectBeta(t *testing.T) {
  requestNewProject := projectConfig.Request{
    ProjectId: ProjectId001,
    Content: "123",
    AvgRating: 20,
    MilestoneInfo: tcr_attributes.MilestonesInfo{
      NumMilestones: 5,
      NumMilestonesCompleted: 8,
    },
  }

  expectedResponseNewProject := projectConfig.Response{
    Ok: true,
  }

  responseMessageNewProject := api.SendPost(requestNewProject, api.NewProjectBetaEndingPoint)

  var responseProject projectConfig.Response
  mapstructure.Decode(*responseMessageNewProject, &responseProject)
  assert.Equal(t, expectedResponseNewProject, responseProject)


  requestGetProject := getProjectConfig.Request{
    ProjectId: ProjectId001,
  }

  expectedResponseGetProject := getProjectConfig.Response{
    Ok: true,
    Project: &project_config.ProjectRecordResult{
      ProjectId: ProjectId001,
      Content:   "123",
      AvgRating: 20,
      MilestoneInfo: &tcr_attributes.MilestonesInfo{
        NumMilestones:          5,
        NumMilestonesCompleted: 8,
      },
    },
  }

  responseMessageGetProject := api.SendPost(requestGetProject, api.GetProjectBetaEndingPoint)
  var responseGetProject getProjectConfig.Response


  config := mapstructure.DecoderConfig{
    DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339Nano),
    Result:     &responseGetProject,
  }
  decoder, err := mapstructure.NewDecoder(&config)
  if err != nil {
    log.Fatalf("error to create NewDecoder: %+v", err)
  }

  err = decoder.Decode(*responseMessageGetProject)
  if err != nil {
    log.Fatalf("error to Decode: %+v", err)
  }

  assert.Equal(t, responseGetProject.Ok, expectedResponseGetProject.Ok)
  assert.Equal(t, responseGetProject.Project.ProjectId, expectedResponseGetProject.Project.ProjectId)
  assert.Equal(t, responseGetProject.Project.Content, expectedResponseGetProject.Project.Content)
  assert.Equal(t, responseGetProject.Project.AvgRating, expectedResponseGetProject.Project.AvgRating)
  assert.Equal(t, responseGetProject.Project.MilestonesInfo, expectedResponseGetProject.Project.MilestonesInfo)
}
