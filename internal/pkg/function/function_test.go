package function

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Function
	}{
		{name: "Simple", want: &Function{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Run(t *testing.T) {
	type args struct {
		ctx     context.Context
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		f       *Function
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{name: "Bad Request", f: nil, args: args{ctx: context.Background(), request: events.APIGatewayProxyRequest{}}, want: events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Credentials": "true",
				"P3P":                              "CP=\"NON DSP LAW CUR ADM DEV TAI PSA PSD HIS OUR DEL IND UNI PUR COM NAV INT DEM CNT STA POL HEA PRE LOC IVD SAM IVA OTC\"",
				"Access-Control-Allow-Headers":     "X-Requested-With,X-Prototype-Version,Authorization,Content-Type,Cache-Control,Pragma,Origin,Accept,Cookie",
				"Content-Type":                     "application/json",
				"Content-Security-Policy":          "upgrade-insecure-requests",
				"Strict-Transport-Security":        "max-age=31536000; includeSubDomains; preload",
			}, Body: "{\"code\":400,\"message\":\"Invalid request body\"}", IsBase64Encoded: false,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{}
			got, err := f.Run(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Function.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
