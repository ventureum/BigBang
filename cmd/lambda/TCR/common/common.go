package common

import (
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/client_config"
)

func ConstructMilestoneFromMilestoneRecordTx(
	milestoneRecord *milestone_config.MilestoneRecord, postgresBigBangClient *client_config.PostgresBigBangClient) *tcr_attributes.Milestone {
	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}

	milestone := &tcr_attributes.Milestone{
		ProjectId:      milestoneRecord.ProjectId,
		MilestoneId:    milestoneRecord.MilestoneId,
		Content:        milestoneRecord.Content,
		StartTime:      milestoneRecord.StartTime,
		EndTime:        milestoneRecord.EndTime,
		BlockTimestamp: milestoneRecord.BlockTimestamp,
		NumObjectives:  milestoneRecord.NumObjectives,
		State:          milestoneRecord.State,
		AvgRating:      milestoneRecord.AvgRating,
	}

	objectiveRecords := objectiveExecutor.GetObjectiveRecordsByProjectIdAndMilestoneIdTx(
		milestoneRecord.ProjectId, milestoneRecord.MilestoneId)

	var objectives []tcr_attributes.Objective
	for _, objectiveRecord := range *objectiveRecords {
		objective := tcr_attributes.Objective{
			ProjectId:      objectiveRecord.ProjectId,
			MilestoneId:    objectiveRecord.MilestoneId,
			ObjectiveId:    objectiveRecord.ObjectiveId,
			Content:        objectiveRecord.Content,
			BlockTimestamp: objectiveRecord.BlockTimestamp,
			AvgRating:      objectiveRecord.AvgRating,
		}
		objectives = append(objectives, objective)
	}
	milestone.Objectives = &objectives
	return milestone
}

func ConstructProjectFromProjectRecordTx(
	projectRecord *project_config.ProjectRecord, postgresBigBangClient *client_config.PostgresBigBangClient) *tcr_attributes.Project {
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

	project := &tcr_attributes.Project{
		ProjectId:      projectRecord.ProjectId,
		Admin:          projectRecord.Admin,
		Content:        projectRecord.Content,
		BlockTimestamp: projectRecord.BlockTimestamp,
		AvgRating:      projectRecord.AvgRating,
	}

	milestonesInfo := &tcr_attributes.MilestonesInfo{
		CurrentMilestone:       projectRecord.CurrentMilestone,
		NumMilestones:          projectRecord.NumMilestones,
		NumMilestonesCompleted: projectRecord.NumMilestonesCompleted,
	}

	milestoneRecords := milestoneExecutor.GetMilestonesRecordsByProjectIdTx(project.ProjectId)

	var milestones []tcr_attributes.Milestone
	for _, milestoneRecord := range *milestoneRecords {
		milestone := ConstructMilestoneFromMilestoneRecordTx(&milestoneRecord, postgresBigBangClient)
		milestones = append(milestones, *milestone)
	}
	milestonesInfo.Milestones = &milestones
	project.MilestonesInfo = milestonesInfo

	return project
}
