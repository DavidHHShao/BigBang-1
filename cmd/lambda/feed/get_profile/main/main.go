package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_profile/config"
)

func main() {
  lambda.Start(lambda_get_profile_config.Handler)
}
