# simplehttp

[![CI](https://github.com/rpunt/dcc/actions/workflows/ci.yml/badge.svg)](https://github.com/rpunt/dcc/actions/workflows/ci.yml)

## Features

* A simple-to-use HTTP client library. Skip the "create client/create request/client does request" dance.
* Request parameters are created as client attributes.

## Get Started

### Install

You first need [Go](https://go.dev/) installed; I targeted 1.19, YMMV with versions before that. You can install simplehttp with the below command:

``` sh
go get github.com/rpunt/simplehttp
```

### Import

Import req to your code:

```go
import "github.com/rpunt/simplehttp"
```

### Basic Usage

#### GET

```go
client := simplehttp.New("https://icanhazdadjoke.com")
client.Headers["Accept"] = "application/json"

// OPTIONAL: add query parameters
client.Params["key"] = "value"

response, err := client.Get("/")
// response is an HttpResponse object
```

#### POST

```go
client := simplehttp.New("https://yoururl.here")

// OPTIONAL: add a request body
client.Data["key"] = "value"

response, err := client.Post("/")
```

### Supported methods

* [x] GET
* [x] POST
* [x] PATCH
* [x] PUT
* [x] DELETE
* [x] HEAD
* [ ] CONNECT
* [ ] OPTIONS
* [ ] TRACE
