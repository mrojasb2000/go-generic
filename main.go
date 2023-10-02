package main

import "github.com/mrojasb2000/go-generic/genericlist"

func main() {
	glist := genericlist.New[string]()

	glist.Insert("Jhon Doe")
	glist.Insert("Sara Doe")
	glist.Insert("Francis Doe")
	glist.Insert("Scolly Doe")
	glist.Insert("Fox Doe")
}
