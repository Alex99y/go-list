package main

import (
	"fmt"

	"github.com/alex99y/go-list/list"
)

func main() {
	var a list.List
	a.AddElement(0)
	a.AddElement(5)
	a.AddElement(10)
	a.AddElement(15)
	a.AddElement(20)
	a.UpdateElement(2, 100)
	_ = a.RemoveElement(0)
	_ = a.RemoveElement(2)
	_ = a.RemoveElement(2)
	N := a.ItemsCount()
	for i := 0; i < N; i++ {
		element, _ := a.GetElement(i)
		fmt.Print(element, "\n")
	}
}
