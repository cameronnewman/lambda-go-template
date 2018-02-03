package main

import (
	"function"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	version string
)

func main() {

	f := function.New()
	// Let Lambda handle the rest!
	lambda.Start(f.Run)
}
