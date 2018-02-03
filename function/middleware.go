package function

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func (f *Function) httpErrorWriter(e ErrorResponse) (events.APIGatewayProxyResponse, error) {
	return f.httpWriter(e, e.Code), nil
}

func (f *Function) httpWriter(body interface{}, statusCode int) events.APIGatewayProxyResponse {

	result := events.APIGatewayProxyResponse{}
	result.Headers = map[string]string{}

	b, err := json.Marshal(body)
	if err != nil {
		return result
	}

	result.Body = string(b)
	result.StatusCode = statusCode
	result.Headers[HeaderContentType] = MIMEApplicationJSON
	result.Headers[HeaderContentSecurityPolicy] = httpUpgradeInsecureRequests
	result.Headers[HeaderStrictTransportSecurity] = httpStrictTransportSecurity
	result.Headers[HeaderAccessControlAllowCredentials] = HeaderTrue
	result.Headers[HeaderPlatformPrivacyPreferences] = PlatformPrivacyPreferencesPolicy
	result.Headers[HeaderAccessControlAllowHeaders] = allowedCORSHeaders

	return result
}
