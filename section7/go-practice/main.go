package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/complexMess"
	enumpb "github.com/enum"

	"google.golang.org/protobuf/encoding/protojson"
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

	jsonDemo(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		Name: "bahaa",
		MyMessage: &complexpb.DummyMessage{
			Title:   "HELLO",
			Content: "this is my message",
		},
		MyHist: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Title:   "HELLO2",
				Content: "this is my message",
			},
			&complexpb.DummyMessage{
				Title:   "HELLO3",
				Content: "this is my message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	ep := enumpb.PersonMessage{
		Id:       6,
		EyeColor: enumpb.EyeColor_EYE_BROWN,
	}

	fmt.Println(ep)
}

func jsonDemo(sm proto.Message) {
	simpleMessageAsJson := toJSON(sm)
	fmt.Println(simpleMessageAsJson)

	myMessagePointer := &simplepb.SimpleMessage{}
	fromJson(simpleMessageAsJson, myMessagePointer)
	fmt.Println("transformed text: ", myMessagePointer)
}

func fromJson(in string, pb proto.Message) {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("couldn't encode the Pb to string", err)
	}
}

func toJSON(pb proto.Message) string {
	// marshaler := jsonpb.Marshaler{}
	// out, err := marshaler.MarshalToString(pb)
	b, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("couldn't encode the Pb to string", err)
		return ""
	}
	return string(b)
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
