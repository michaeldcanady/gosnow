# gosnow

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/gosnow?style=plastic)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/michaeldcanady/gosnow/Go?style=plastic)
[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/gosnow)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/gosnow?style=plastic)
![GitHub issues](https://img.shields.io/github/issues/michaeldcanady/gosnow?style=plastic)
![GitHub](https://img.shields.io/github/license/michaeldcanady/gosnow?style=plastic)
![GitHub all releases](https://img.shields.io/github/downloads/michaeldcanady/gosnow/total?style=plastic)

### Table of Contents
* [Current Ideas](#current-ideas)
* [Installation](#installation)
* [Usage](#usage)
* [Examples](#examples)
* [Contributing](#contributing)

## Current Ideas:
- [ ] Remove need for "resource"
- [ ] Build out Tables API

GoSnow is a Golang wrapper for the Service Now API.

## Installation

```bash
go get github.com/michaeldcanady/gosnow
```

## Usage
``` golang
import "github.com/michaeldcanady/gosnow/v5/gosnow"
```

## Examples

### Creating a new client instance
``` golang
client, _ := gosnow.New(username, password, instance)
```
### Create table instance
```golang
CSTable, _ := client.Table("TableName")
```
### Get a table value
```golang
query := map[string]interface{}{"field": "value"}
respose, _ := CSTable.Get(query, )
```
### Update a table value
```golang
query := map[string]interface{}{"field": "value"}
respose, _ := CSTable.Update(query, 1, 0, true, nil)
```
### Delete a table value
```golang
query := map[string]interface{}{"field": "value"}
respose, _ := CSTable.Delete(query)
```
### Create a table value
```golang
respose, _ := CSTable.Create(query)
```
## Contributing

See [`CONTRIBUTORS.md`](CONTRIBUTORS.md) for details.

## Author

Michael Canady
