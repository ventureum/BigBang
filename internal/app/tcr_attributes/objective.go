package tcr_attributes

type Objective struct {
  ObjectiveId   string         `json:"objectiveId,required"`
  Content       string         `json:"content,required"`
  TotalRating   int64          `json:"totalRating,required"`
  TotalWeight   int64          `json:"totalWeight,required"`
}
