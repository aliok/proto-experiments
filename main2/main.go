package main

import (
	"fmt"
	"github.com/aliok/proto-experiments/contract2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	bindata, err := ioutil.ReadFile("msg.bin")
	if err != nil {
		log.Fatalln(err)
	}

	jsondata, err := ioutil.ReadFile("msg.json")
	if err != nil {
		log.Fatalln(err)
	}

	retrievedbin := &contract2.Ingress{}
	err = proto.Unmarshal(bindata, retrievedbin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Bin %#v\n\n", retrievedbin)

	retrievedjson := &contract2.Ingress{}
	err = protojson.Unmarshal(jsondata, retrievedjson)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("JSON %#v\n\n", retrievedjson)

}
