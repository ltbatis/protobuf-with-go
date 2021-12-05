package main

import (
	"fmt"
	"io/ioutil"
	"log"

	simplepb "github.com/ltbatista/protobuf-with-go/src/simple"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	writeToFile("simple.bin", sm)
	//TODO: readFromFile()
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes ", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file ", err)
		return err
	}
	fmt.Println("Data has been written.")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}
