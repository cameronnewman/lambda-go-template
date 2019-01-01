package function

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestFunction_httpErrorWriter(t *testing.T) {
	type args struct {
		e ErrorResponse
	}
	tests := []struct {
		name    string
		f       *Function
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{name: "empty", f: nil, args: args{e: ErrorResponse{Code: 403, Message: "Your request has been rate limited"}}, want: events.APIGatewayProxyResponse{
			StatusCode: 403,
			Headers: map[string]string{
				"Access-Control-Allow-Credentials": "true",
				"P3P": "CP=\"NON DSP LAW CUR ADM DEV TAI PSA PSD HIS OUR DEL IND UNI PUR COM NAV INT DEM CNT STA POL HEA PRE LOC IVD SAM IVA OTC\"",
				"Access-Control-Allow-Headers": "X-Requested-With,X-Prototype-Version,Authorization,Content-Type,Cache-Control,Pragma,Origin,Accept,Cookie",
				"Content-Type":                 "application/json",
				"Content-Security-Policy":      "upgrade-insecure-requests",
				"Strict-Transport-Security":    "max-age=31536000; includeSubDomains; preload",
			}, Body: "{\"code\":403,\"message\":\"Your request has been rate limited\"}", IsBase64Encoded: false,
		}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{}
			got, err := f.httpErrorWriter(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Function.httpErrorWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.httpErrorWriter() = %v, want %v", got.Body, tt.want.Body)
			}
		})
	}
}

func TestFunction_httpWriter(t *testing.T) {
	type args struct {
		body       interface{}
		statusCode int
	}
	tests := []struct {
		name string
		f    *Function
		args args
		want events.APIGatewayProxyResponse
	}{
		// TODO: Add test cases.
		{name: "error", f: nil, args: args{body: ErrorResponse{Code: 403, Message: "Your request has been rate limited"}, statusCode: 403}, want: events.APIGatewayProxyResponse{
			StatusCode: 403,
			Headers: map[string]string{
				"Access-Control-Allow-Credentials": "true",
				"P3P": "CP=\"NON DSP LAW CUR ADM DEV TAI PSA PSD HIS OUR DEL IND UNI PUR COM NAV INT DEM CNT STA POL HEA PRE LOC IVD SAM IVA OTC\"",
				"Access-Control-Allow-Headers": "X-Requested-With,X-Prototype-Version,Authorization,Content-Type,Cache-Control,Pragma,Origin,Accept,Cookie",
				"Content-Type":                 "application/json",
				"Content-Security-Policy":      "upgrade-insecure-requests",
				"Strict-Transport-Security":    "max-age=31536000; includeSubDomains; preload",
			}, Body: "{\"code\":403,\"message\":\"Your request has been rate limited\"}", IsBase64Encoded: false,
		},
		},
		{name: "struct", f: nil, args: args{body: "", statusCode: 200}, want: events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Credentials": "true",
				"P3P": "CP=\"NON DSP LAW CUR ADM DEV TAI PSA PSD HIS OUR DEL IND UNI PUR COM NAV INT DEM CNT STA POL HEA PRE LOC IVD SAM IVA OTC\"",
				"Access-Control-Allow-Headers": "X-Requested-With,X-Prototype-Version,Authorization,Content-Type,Cache-Control,Pragma,Origin,Accept,Cookie",
				"Content-Type":                 "application/json",
				"Content-Security-Policy":      "upgrade-insecure-requests",
				"Strict-Transport-Security":    "max-age=31536000; includeSubDomains; preload",
			}, Body: "\"\"", IsBase64Encoded: false,
		},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{}
			if got := f.httpWriter(tt.args.body, tt.args.statusCode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.httpWriter() = %v, want *%v*", got.Body, tt.want.Body)
			}
		})
	}
}
