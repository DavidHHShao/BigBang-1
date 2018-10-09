package tcr_attributes

type Project struct {
  ProjectId      string          `json:"projectId,required"`
  Admin          string          `json:"admin,required"`
  Content        string          `json:"content,required"`
  AvgRating      int64           `json:"avgRating,required"`
  MilestonesInfo *MilestonesInfo `json:"milestonesInfo,required"`
}
