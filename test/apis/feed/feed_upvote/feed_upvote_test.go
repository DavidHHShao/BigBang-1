package feed_upvote

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/error_config"
  "BigBang/cmd/lambda/feed/feed_upvote/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_feed_upvote_config.Request
    response lambda_feed_upvote_config.Response
    err    error
  }{
    {
      request: lambda_feed_upvote_config.Request {
        Actor: test_constants.Actor2,
        PostHash: test_constants.PostHash1,
        Value: -1,
      },
      response: lambda_feed_upvote_config.Response {
        VoteInfo: &feed_attributes.VoteInfo{
          Actor: test_constants.Actor2,
          PostHash: test_constants.PostHash1,
          FuelCost: 20,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 80,
            Reputation: 100,
            MilestonePoints: 100,
          },
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_feed_upvote_config.Request {
        Actor: test_constants.Actor2,
        PostHash: test_constants.PostHash1,
        Value: 0,
      },
      response: lambda_feed_upvote_config.Response {
        VoteInfo: &feed_attributes.VoteInfo{
          Actor: test_constants.Actor2,
          PostHash: test_constants.PostHash1,
          FuelCost: 17,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 80,
            Reputation: 100,
            MilestonePoints: 100,
          },
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 0,
            TotalVoteCount : 1,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_feed_upvote_config.Request {
        Actor: test_constants.Actor2,
        PostHash: test_constants.PostHash1,
        Value: 1,
      },
      response: lambda_feed_upvote_config.Response {
        VoteInfo: &feed_attributes.VoteInfo{
          Actor: test_constants.Actor2,
          PostHash: test_constants.PostHash1,
          FuelCost: 17,
          RewardsInfo: &feed_attributes.RewardsInfo{
            Fuel: 63,
            Reputation: 100,
            MilestonePoints: 100,
          },
          PostVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 1,
            TotalVoteCount : 2,
          },
          RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
            DownVoteCount: 1,
            UpVoteCount: 1,
            TotalVoteCount : 2,
          },
        },
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_feed_upvote_config.Request {
        Actor: test_constants.Actor2,
        PostHash: test_constants.PostHash1,
        Value: 1,
      },
      response: lambda_feed_upvote_config.Response {
        Message: &error_config.ErrorInfo{
          ErrorCode: "ExceedingUpvoteLimit",
          ErrorData: error_config.ErrorData {
            "actor": test_constants.Actor2,
            "postHash": test_constants.PostHash1,
          },
          ErrorLocation: "ActorVotesCountersRecordLocation",
        },
        Ok: false,
      },
      err: nil,
    },
    {
      request: lambda_feed_upvote_config.Request {
        Actor: test_constants.Actor2,
        PostHash: test_constants.PostHash1,
        Value: -1,
      },
      response: lambda_feed_upvote_config.Response {
        Message: &error_config.ErrorInfo{
          ErrorCode: "ExceedingDownvoteLimit",
          ErrorData: error_config.ErrorData {
            "actor": test_constants.Actor2,
            "postHash": test_constants.PostHash1,
          },
          ErrorLocation: "ActorVotesCountersRecordLocation",
        },
        Ok: false,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    responseMessageFeedUpvote := api.SendPost(test.request, api.FeedUpvoteAlphaEndingPoint)
    var responseFeedUpvote lambda_feed_upvote_config.Response
    mapstructure.Decode(*responseMessageFeedUpvote , &responseFeedUpvote)
    assert.Equal(t, test.response, responseFeedUpvote)
  }
}
