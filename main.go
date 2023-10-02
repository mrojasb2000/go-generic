package main

import (
	"fmt"

	"github.com/mrojasb2000/go-generic/genericlist"
)

func main() {
	glist := genericlist.New[string]()

	glist.Insert("Jhon Doe")
	glist.Insert("Sara Doe")
	glist.Insert("Francis Doe")
	glist.Insert("Scolly Doe")
	glist.Insert("Fox Doe")

	fmt.Printf("Value '%s' for index %d\n", glist.Get(3), 3)
	fmt.Printf("%+v\n", glist)

	glist.Remove(3)

	fmt.Printf("%+v\n", glist)
}
