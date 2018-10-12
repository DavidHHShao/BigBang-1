package update_actor_rating_votes

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/TCR/update_actor_rating_votes/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_update_actor_rating_votes_config.Request
    response lambda_update_actor_rating_votes_config.Response
    err    error
  }{
    {
      request: lambda_update_actor_rating_votes_config.Request {
        Actor: test_constants.Actor1,
        ProjectId: test_constants.ProjectId1,
        AvailableVotes: 50,
        ReceivedVotes: 60,
      },
      response: lambda_update_actor_rating_votes_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }

  for _, test := range tests {
    result, err := lambda_update_actor_rating_votes_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Ok, result.Ok)
  }
}