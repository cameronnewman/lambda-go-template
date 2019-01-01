package function

var (

	//ErrInternalServerError is an internal error
	ErrInternalServerError = ErrorResponse{Code: 500, Message: "Sorry, we've run into a problem. Please try again or contact support"}

	//ErrServiceUnavailable is an error when the service can not respond
	ErrServiceUnavailable = ErrorResponse{Code: 500, Message: "Service is temporarily unavailable"}

	//ErrGenericInvalidPayload is an error when the payload is incorrect
	ErrGenericInvalidPayload = ErrorResponse{Code: 400, Message: "Invalid request body"}

	//ErrMethodNotAllowed is an error when the request method isnt allowed
	ErrMethodNotAllowed = ErrorResponse{Code: 405, Message: "Sorry that type of HTTP request is not allowed"}

	//ErrQueryStringParametersRequired is an error when a query string was included in the request
	ErrQueryStringParametersRequired = ErrorResponse{Code: 403, Message: "You must supply a query string parameters"}

	//ErrForbiddenIPAddress is an error when the connecting IP is banned OR forbidden
	ErrForbiddenIPAddress = ErrorResponse{Code: 403, Message: "Your request has been rate limited"}

	//ErrSecureConnectionRequired is an error when the connection is unsecure
	ErrSecureConnectionRequired = ErrorResponse{Code: 403, Message: "You must connect via https"}
)

//ErrorResponse struct
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
