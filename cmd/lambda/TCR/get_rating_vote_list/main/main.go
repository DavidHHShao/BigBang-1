package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_rating_vote_list/config"
)

func main() {
  lambda.Start(config.Handler)
}
