package update

import (
  "testing"
  "BigBang/cmd/lambda/TCR/update_delegate_votes/config"
  "BigBang/test/constants"
  "github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_update_delegate_votes_config.Request
    response lambda_update_delegate_votes_config.Response
    err    error
  }{
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor1,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor2,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor3,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor4,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor5,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor6,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_update_delegate_votes_config.Request {
        Actor:                  test_constants.Actor7,
        ProjectId:              test_constants.ProjectId1,
        AvailableDelegateVotes: 50,
        ReceivedDelegateVotes:  60,
      },
      response: lambda_update_delegate_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  for _, test := range tests {
    result, err := lambda_update_delegate_votes_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
  }
}
