package get_feed_post_test

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/cmd/lambda/feed/get_feed_post/config"
  "BigBang/test/constants"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_feed_post_config.Request
    response lambda_get_feed_post_config.Response
    err    error
  }{
    {
      request: lambda_get_feed_post_config.Request{
        PrincipalId: test_constants.Actor1,
        Body: lambda_get_feed_post_config.RequestContent{
          PostHash:  test_constants.PostHash1,
          Requestor: test_constants.Actor1,
        },
      },
      response: lambda_get_feed_post_config.Response {
        Post: &lambda_get_feed_post_config.ResponseContent {
          Actor: test_constants.Actor1,
          Username: test_constants.UserName1,
          PhotoUrl: "http://123.com",
          BoardId: test_constants.BoardId1,
          ParentHash: test_constants.EmptyParentHash,
          PostHash: test_constants.PostHash1,
          PostType: string(feed_attributes.PostPostType),
          Content: &test_constants.PostContent1,
          DeltaFuel: -50,
          DeltaReputation: 0,
          DeltaMilestonePoints: 0,
          WithdrawableMPs: 0,
          RepliesLength: 0,
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 0,
            UpVoteCount: 0,
            TotalVoteCount: 0,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 0,
            UpVoteCount: 0,
            TotalVoteCount: 0,
          },
        },
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_get_feed_post_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
