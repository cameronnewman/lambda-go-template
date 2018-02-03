# lambda-go-template

## Purpose ##

Barebones Lambda Go template for interacting with APIGatewayProxy requests.

[![GoDoc][1]][2]
[![GoCard][3]][4]

[1]: https://godoc.org/github.com/cameronnewman/lambda-go-template?status.svg
[2]: https://godoc.org/github.com/cameronnewman/lambda-go-template
[3]: https://goreportcard.com/badge/github.com/cameronnewman/lambda-go-template
[4]: https://goreportcard.com/report/github.com/cameronnewman/lambda-go-template


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
