package lambda_finalize_milestone_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "log"
)


type Request struct {
  ProjectId   string                  `json:"projectId,required"`
  MilestoneId int64                   `json:"milestoneId,required"`
  BlockTimestamp  int64               `json:"blockTimestamp,required"`
  EndTime int64                       `json:"endTime,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId
  milestoneId := request.MilestoneId
  postgresBigBangClient.Begin()

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}

  existing := milestoneExecutor.CheckMilestoneRecordExistingTx(projectId, milestoneId)

  if !existing {
    errorInfo := error_config.ErrorInfo{
      ErrorCode: error_config.NoMilestoneIdExisting,
      ErrorData: map[string]interface{} {
        "milestoneId": milestoneId,
        "projectId": projectId,
      },
      ErrorLocation: error_config.MilestoneRecordLocation,
    }
    log.Printf("No milestone record for projectId %s and milestoneId %d", projectId, milestoneId)
    log.Panicln(errorInfo.Marshal())
  }

   milestoneExecutor.FinalizeMilestoneTx(projectId, milestoneId, request.BlockTimestamp, request.EndTime)
   projectExecutor.SetCurrentMilestoneTx(projectId, 0)
   projectExecutor.IncreaseNumMilestonesCompletedTx(projectId)

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
