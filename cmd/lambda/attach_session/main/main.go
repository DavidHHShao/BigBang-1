package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/attach_session/config"
)

func main() {
  lambda.Start(config.Handler)
}