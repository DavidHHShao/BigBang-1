package tcr_attributes

import "time"

type RatingVote struct {
  Voter string `json:"voter,required"`
  Rating int64 `json:"rating,required"`
  Weight int64 `json:"weight,required"`
  VotedAt time.Time `json:"votedAt,required"`
}
