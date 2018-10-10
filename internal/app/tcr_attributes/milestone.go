package tcr_attributes

type Milestone struct {
  MilestoneId   string         `json:"milestoneId,required"`
  Content       string         `json:"content,required"`
  StartTime     int64          `json:"startTime,required"`
  EndTime       int64          `json:"endTime,required"`
  BlockTimestamp  int64        `json:"blockTimestamp,required"`
  NumObjs       int64          `json:"numObjs,required"`
  AvgRating     int64          `json:"avgRating,required"`
  State         MileStoneState `json:"state,required"`
  Objs          *[]Objective   `json:"objs,required"`
}
