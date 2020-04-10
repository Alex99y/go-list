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
	//data, _ := a.GetElement(2)
	//fmt.Print(data, "\n")
	_ = a.RemoveElement(0)
	//data0, _ := a.GetElement(2)
	//fmt.Print(data0, "\n")
	_ = a.RemoveElement(2)
	//data1, _ := a.GetElement(2)
	//fmt.Print(data1, "\n")
	N := a.ItemsCount()
	for i := 0; i < N; i++ {
		element, _ := a.GetElement(i)
		fmt.Print(element, "\n")
	}
}
