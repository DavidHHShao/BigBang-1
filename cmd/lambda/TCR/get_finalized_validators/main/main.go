package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_finalized_validators/config"
)

func main() {
  lambda.Start(lambda_get_finalized_validators_config.Handler)
}
