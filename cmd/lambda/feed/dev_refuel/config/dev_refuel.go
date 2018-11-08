package lambda_dev_refuel_config

import (
  "log"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/feed/refuel_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
)

type Request struct {
  Actor string `json:"actor,required"`
  Fuel int64 `json:"fuel,required"`
  Reputation int64 `json:"reputation,required"`
  MilestonePoints int64 `json:"milestonePoints,required"`
}

type Response struct {
  Ok      bool   `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()

  fuel := feed_attributes.Fuel(request.Fuel)
  reputation := feed_attributes.Reputation(request.Reputation)
  milestonePoints := feed_attributes.MilestonePoint(request.MilestonePoints)
  actor := request.Actor



  postgresBigBangClient.Begin()

  refuelRecordExecutor := refuel_record_config.RefuelRecordExecutor{
    *postgresBigBangClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}


  actorProfileRecordExecutor.VerifyActorExistingTx(actor)
  actorRewardsInfoRecordExecutor.VerifyActorExistingTx(actor)

  actorRewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)
  refuelRecordExecutor.UpsertRefuelRecordTx(&refuel_record_config.RefuelRecord{
    Actor: actor,
    Fuel: fuel.SubFuels(actorRewardsInfo.Fuel),
    Reputation: reputation - actorRewardsInfo.Reputation,
    MilestonePoints: milestonePoints - actorRewardsInfo.MilestonePoints,
  })
  actorRewardsInfoRecordExecutor.UpsertActorRewardsInfoRecordTx(&actor_rewards_info_record_config.ActorRewardsInfoRecord{
    Actor: actor,
    Fuel: fuel,
    Reputation: reputation,
    MilestonePointsFromPosts: 0,
    MilestonePointsFromVotes: 0,
    MilestonePointsFromOthers: milestonePoints,
    MilestonePoints: milestonePoints,
  })

  postgresBigBangClient.Commit()

  log.Printf("Reset %d fuel, %d reputation, and %d milestonePoints to actor %s", fuel, reputation, milestonePoints, actor)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
