package function

var (
	ErrInternalServerError = ErrorResponse{Code: 500, Message: "Sorry, we've run into a problem. Please try again or contact support"}

	ErrServiceUnavailable = ErrorResponse{Code: 500, Message: "Service is temporarily unavailable"}

	ErrGenericInvalidPayload = ErrorResponse{Code: 400, Message: "Invalid request body"}

	ErrMethodNotAllowed = ErrorResponse{Code: 405, Message: "Sorry that type of HTTP request is not allowed"}

	ErrQueryStringParametersRequired = ErrorResponse{Code: 403, Message: "You must supply a query string parameters"}

	ErrForbiddenIPAddress = ErrorResponse{Code: 403, Message: "Your request has been rate limited"}

	ErrSecureConnectionRequired = ErrorResponse{Code: 403, Message: "You must connect via https"}
)

//ErrorResponse struct
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
