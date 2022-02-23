package main

import (
	"fmt"
	"github.com/aliok/proto-experiments/contract2"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("msg.bin")
	if err != nil {
		log.Fatalln(err)
	}

	retrieved := &contract2.Ingress{}
	err = proto.Unmarshal(data, retrieved)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%#v", retrieved)
}
