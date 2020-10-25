package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"

	simplepb "github.com/simple"
)

func main() {
	// fmt.Println("hello")
	fmt.Println("__________Start____________>")
	sm := doSimple()
	fmt.Println("Printing the Message before serializing: ", sm)
	fmt.Println("______________________")
	readAndWriteDemo(sm)
	fmt.Println("__________End____________")
}

func readAndWriteDemo(sm proto.Message) {
	// writefun
	writeToFile("simple.bin", sm)
	// create simple message
	sm2 := &simplepb.SimpleMessage{}
	// readfunc
	readFromFile("simple.bin", sm2)
	//print the message
	fmt.Println("Printing the Message after de-serializing: ", sm2)
}

func writeToFile(fn string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cann't serialize the proto to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fn, out, 0644); err != nil {
		log.Fatalln("couldn't write to the file", err)
		return err
	}

	fmt.Println("File has been written succssefully!")
	fmt.Println("______________________")

	return nil
}

func readFromFile(fn string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalln("couldn't read from thr file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("couldn't deserialize the file", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {

	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}

	// fmt.Println(sm)
	// fmt.Println("This is the ID: ", sm.GetId())

	sm.Name = "This is the new name"
	// fmt.Println(sm)
	// fmt.Println(sm.GetId)

	return &sm
}
