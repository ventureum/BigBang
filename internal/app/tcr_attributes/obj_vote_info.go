package tcr_attributes

type ObjVoteInfo struct {
  ProjectId     string         `json:"projectId,required"`
  MilestoneId   int64          `json:"milestoneId,required"`
  ObjectiveId   int64          `json:"objId,required"`
  RatingVotes   *[]RatingVote  `json:"ratingVotes,required"`
}
