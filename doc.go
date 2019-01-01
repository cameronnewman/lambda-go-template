/*
Package lambda is an Golang is a template repo for golang lambda apps
*/
package lambda

// Simple example:
//
// 	//Run executes the function on the Lambda request
// func (f *Function) Run(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//
//	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
//	fmt.Printf("Body size = %d.\n", len(request.Body))
//
//	fmt.Println("Headers:")
//	for key, value := range request.Headers {
//		fmt.Printf("    %s: %s\n", key, value)
//	}
//
//	fmt.Println("Body:")
//	fmt.Printf("    %s\n", request.Body)
//
//	if len(request.Body) == 0 {
//		return f.httpErrorWriter(ErrGenericInvalidPayload)
//	}
//
//	return f.httpWriter(OkResponse{}, http.StatusOK), nil
// }
