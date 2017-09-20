package main

import (
	"fmt"
)

type MyInterface interface {
	SayHello()
}

type Base struct{}

func (b Base) SayHello() {
	fmt.Println("Hello!!!!")
}

type MyItem struct {
	Title string

	Base
}

func main() {
	var x MyInterface

	x = &MyItem{Title: "yo"}

	if item, ok := x.(*MyItem); ok {
		fmt.Println(item.Title)
	} else {
		fmt.Println("to *MyItem", "not ok")
	}
	// => yo

	if item, ok := x.(MyItem); ok {
		fmt.Println(item.Title)
	} else {
		fmt.Println("to MyItem", "not ok")
	}
	// => to B not ok
}
