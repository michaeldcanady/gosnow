# gosnow

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/gosnow?style=plastic)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/michaeldcanady/gosnow/Go?style=plastic)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/gosnow?style=plastic)
![GitHub issues](https://img.shields.io/github/issues/michaeldcanady/gosnow?style=plastic)
![GitHub](https://img.shields.io/github/license/michaeldcanady/gosnow?style=plastic)

GoSnow is a Golang wrapper for the Service Now API.

Install gosnow
```bash
go get github.com/michaeldcanady/gosnow
```
Creating a client instance
``` golang
client, _ := gosnow.New(username, password, instance)
```
Create table instance
```golang
CSTable, _ := client.Resource("/table/TableName")
```

## Contributing

## Author

Michael Canady
