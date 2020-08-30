package main

import (
	"fmt"

	simplepb "github.com/ahmed-bahaa/protobuf/section7/go-practice/src/simple"
)

func main() {
	// fmt.Println("hello")
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(sm)
	fmt.Println("This is the ID: ", sm.GetId())
}
