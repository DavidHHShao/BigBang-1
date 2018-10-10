package milestone_config

import (
  "time"
  "BigBang/internal/app/tcr_attributes"
)


type MilestoneRecord struct {
  ProjectId     string         `json:"projectId" db:"project_id"`
  MilestoneId   int64          `json:"milestoneId" db:"milestone_id"`
  Content       string         `json:"content" db:"content"`
  BlockTimestamp int64         `json:"block_timestamp" db:"block_timestamp"`
  StartTime     int64          `json:"startTime" db:"start_time"`
  EndTime       int64          `json:"endTime" db:"end_time"`
  State         tcr_attributes.MileStoneState         `json:"state" db:"state"`
  NumObjs       int64          `json:"numObjs" db:"num_objs"`
  AvgRating     int64          `json:"avgRating" db:"avg_rating"`
  TotalRating   int64          `json:"totalRating" db:"total_rating"`
  TotalWeight   int64          `json:"totalWeight" db:"total_weight"`
  CreatedAt     time.Time      `json:"createdAt" db:"created_at"`
  UpdatedAt     time.Time      `json:"updatedAt" db:"updated_at"`
}
