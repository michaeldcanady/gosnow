# gosnow

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/michaeldcanady/gosnow)

GoSnow is a Golang wrapper for the Service Now API.

```golang

package main

import (
	"fmt"

	"github.com/michaeldcanady/gosnow"
)

func main() {
	FetchExample()
	UpdateExample()
	CreateExample()
	QueryBuilderExample()
}

const (
	username = ""
	password = ""
	instance = ""
)

func QueryBuilderExample() {
	client, _ := gosnow.New(username, password, instance)

	// Create a table connection
	CSTable, _ := client.Resource("/table/u_computer_support")

	query := gosnow.Filter("number") + gosnow.IS("CS0012345")
	// Fetch a ticket with the values in the dictionary
	respose, _ := CSTable.Get(query, 0, 0, true, nil)

	fmt.Println(respose.First())
}

func FetchExample() {
	client, _ := gosnow.New(username, password, instance)

	// Create a table connection
	CSTable, _ := client.Resource("/table/u_computer_support")

	// Fetch a ticket with the values in the dictionary
	respose, _ := CSTable.Get(map[string]interface{}{"number": "CS0086712"}, 0, 0, true, nil)

	fmt.Println(respose.First())
}

func UpdateExample() {
	client, _ := gosnow.New(username, password, instance)

	// Create a table connection
	CSTable, _ := client.Resource("/table/u_computer_support")

	update := map[string]string{"short_description": "New short description 1"}

	updated_record, _ := CSTable.Update(map[string]interface{}{"number": "CS0086712"}, update)

	fmt.Println(updated_record.First())
}

func CreateExample() {
	client, _ := gosnow.New(username, password, instance)

	// Create a table connection
	CSTable, _ := client.Resource("/table/u_computer_support")

	//# Set the payload for creation
	new_record := map[string]string{
		"short_description": "GoSnow created incident",
		"description":       "This is awesome",
		"u_depot_location":  "Example location",
		"requested_for":     username,
	}

	createdticket, _ := CSTable.Create(new_record)

	fmt.Println(createdticket.First())
}


```
