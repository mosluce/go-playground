package main

import (
	"fmt"
)

type A interface{}

type B struct {
	Title string
}

type C struct {
	Title string
}

func main() {
	var oa A

	oa = &B{"yo"}

	if ob, ok := oa.(*B); ok {
		fmt.Println(ob.Title)
	} else {
		fmt.Println("to *B", "not ok")
	}

	if ob, ok := oa.(B); ok {
		fmt.Println(ob.Title)
	} else {
		fmt.Println("to B", "not ok")
	}

	if oc, ok := oa.(*C); ok {
		fmt.Println(oc.Title)
	} else {
		fmt.Println("to C", "not ok")
	}
}
