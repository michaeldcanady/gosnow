package main

import (
	"fmt"
	"github.com/michaeldcanady/GoServe/gosnow"
)

func main() {
	FetchExample()
	UpdateExample()
	CreateExample()
	QueryBuilderExample()
}

const(
	username
	password
	instance
)

func QueryBuilderExample() {
	client, err := gosnow.New(username,password,instance)
	if err != nil {
		panic(err)
	}

	// Create a table connection
	CSTable, err := client.Resource("/table/u_computer_support")
	if err != nil {
		panic(err)
	}
	query := gosnow.Filter("number")+gosnow.IS("CS0012345")
	// Fetch a ticket with the values in the dictionary
	respose := CSTable.Get(query, 0, 0, true, nil)

	fmt.Println(respose.First())
}

func FetchExample() {
	client, err := gosnow.New(username,password,instance)
	if err != nil {
		panic(err)
	}

	// Create a table connection
	CSTable, err := client.Resource("/table/u_computer_support")
	if err != nil {
		panic(err)
	}

	// Fetch a ticket with the values in the dictionary
	respose := CSTable.Get(map[string]interface{}{"number": "CS0086712"}, 0, 0, true, nil)

	fmt.Println(respose.First())
}

func UpdateExample() {
	client, err := gosnow.New(username,password,instance)
	if err != nil {
		panic(err)
	}

	// Create a table connection
	CSTable, err := client.Resource("/table/u_computer_support")
	if err != nil {
		panic(err)
	}

	update := map[string]string{"short_description": "New short description 1"}

	updated_record := CSTable.Update(map[string]interface{}{"number": "CS0086712"}, update)

	fmt.Println(updated_record.Created())
}

func CreateExample() {
	client, err := gosnow.New(username,password,instance)
	if err != nil {
		panic(err)
	}

	// Create a table connection
	CSTable, err := client.Resource("/table/u_computer_support")
	if err != nil {
		panic(err)
	}

	//# Set the payload for creation
	new_record := map[string]string{
    "short_description": "Pysnow created incident",
    "description": "This is awesome",
		"u_depot_location": "CMHS 3122A - Help Desk Location",
		"requested_for": username,
	}

	createdticket := CSTable.Create(new_record)

	fmt.Println(createdticket.Created())

	fmt.Println(client)
}
