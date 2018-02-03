# lambda-go-template

## Purpose ##

Barebones Lambda Go template for interacting with APIGatewayProxy requests.

## Usage ##

```
make
```
make will build, test and create a zip file that can be uploaded directly to AWS Lambda. All built using  Golang docker image.

```
make tests
```
will run all tests in the repo excluding the vendor directory.


## Further Information ##

```
github.com/cameronnewman/
  cmd/
    lambda-go-template/
      main.go
  function/
    function.go <- Your function lives here
```
