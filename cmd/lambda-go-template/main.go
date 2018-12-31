package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cameronnewman/lambda-go-template/internal/function"
)

var (
	version string
)

func main() {

	f := function.New()
	fmt.Println("Starting Lambda, version ", version)

	// Let Lambda handle the rest!
	lambda.Start(f.Run)
}
