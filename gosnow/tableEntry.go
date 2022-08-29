package gosnow

import "fmt"

type TableEntry ResponseEntry

func (T *TableEntry) Update() {
	fmt.Println(T.resource)
}
