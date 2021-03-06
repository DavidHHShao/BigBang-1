package set_actor_private_key_test

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/test/constants"
  "BigBang/cmd/lambda/feed/set_actor_private_key/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_set_actor_private_key_config.Request
    response lambda_set_actor_private_key_config.Response
    err    error
  }{
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor1,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor1,
          PrivateKey: test_constants.PrivateKey1,
        },
      },
      response: lambda_set_actor_private_key_config.Response {
          Ok: true,
      },
       err: nil,
    },
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor2,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor2,
          PrivateKey: test_constants.PrivateKey2,
        },
      },
      response: lambda_set_actor_private_key_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor3,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor3,
          PrivateKey: test_constants.PrivateKey3,
        },
      },
      response: lambda_set_actor_private_key_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor4,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor4,
          PrivateKey: test_constants.PrivateKey4,
        },
      },
      response: lambda_set_actor_private_key_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor5,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor5,
          PrivateKey: test_constants.PrivateKey5,
        },
      },
      response: lambda_set_actor_private_key_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor6,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor6,
          PrivateKey: test_constants.PrivateKey6,
        },
      },
      response: lambda_set_actor_private_key_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_set_actor_private_key_config.Request{
        PrincipalId: test_constants.Actor6,
        Body: lambda_set_actor_private_key_config.RequestContent{
          Actor:      test_constants.Actor6,
          PrivateKey: "",
        },
      },
      response: lambda_set_actor_private_key_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_set_actor_private_key_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
