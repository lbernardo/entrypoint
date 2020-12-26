# Entrypoint

### Install
```shell
go get github.com/lbernardo/entrypoint
```

### Use

Convert Http Request  to Generic Request
```go
request := entrypoint.NewRequestByHttp(r)  // r is *http.Request and request is entrypoint.Request
```

Convert Http Response  to Generic Response
```go
response := entrypoint.NewResponseToHttp(w, http.StatusOK, responseGen) // response is entrypoint.Response 
```

Convert APIGateway Request  to Generic Request
```go
entrypoint.NewRequestByApiGateway(entrypoint.RequestAPIGateway)
```

Convert Generic Response  to API Gateway response
```go
entrypoint.NewResponseToApiGateway(status, response) // status is int and response is entrypoint.Response
```