# gosnow

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/gosnow?style=plastic)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/michaeldcanady/gosnow/Go?style=plastic)
[![GoDoc](https://img.shields.io/static/v1?style=plastic&label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/michaeldcanady/gosnow)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/michaeldcanady/gosnow?style=plastic)
![GitHub issues](https://img.shields.io/github/issues/michaeldcanady/gosnow?style=plastic)
![GitHub](https://img.shields.io/github/license/michaeldcanady/gosnow?style=plastic)
![GitHub all releases](https://img.shields.io/github/downloads/michaeldcanady/gosnow/total?style=plastic)

# DEPRECATION NOTICE: This module has been deprecated in favor of [servicenow-sdk-go](https://github.com/michaeldcanady/servicenow-sdk-go)

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
import "github.com/michaeldcanady/gosnow/v6/gosnow"
```

## Examples

### Getting a table value
``` golang
package main

import(
    "fmt"

    "github.com/michaeldcanady/gosnow/v6/gosnow"
)

client, _ := gosnow.New(username, password, instance)
CSTable, _ := client.Table("TableName")
query := map[string]interface{}{"field": "value"}
respose, _ := CSTable.Get(query, 0, 0, true, nil)
fmt.Println(respose.First())
```
### Update a table value
```golang
package main

import(
    "fmt"

    "github.com/michaeldcanady/gosnow/v6/gosnow"
)

client, _ := gosnow.New(username, password, instance)
CSTable, _ := client.Table("TableName")

// map of values to update
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
package main

import(
    "fmt"

    "github.com/michaeldcanady/gosnow/v6/gosnow"
)

client, _ := gosnow.New(username, password, instance)

CSTable, _ := client.Table("TableName")

query := map[string]interface{}{}

respose, _ := CSTable.Create(query)
```
## Contributing

See [`CONTRIBUTORS.md`](CONTRIBUTORS.md) for details.

## Author

Michael Canady
