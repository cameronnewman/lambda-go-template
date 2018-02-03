package function

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Function struct{}

func New() *Function {
	return &Function{}
}

func (f *Function) Run(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	fmt.Println("Body:")
	fmt.Printf("    %s\n", request.Body)

	if len(request.Body) == 0 {
		return f.httpErrorWriter(ErrGenericInvalidPayload)
	}

	return f.httpWriter(OkResponse{}, http.StatusOK), nil
}
