package main

import (
	"github.com/aliok/proto-experiments/contract1"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	message := &contract1.Ingress{
		ContentMode: contract1.ContentMode_BINARY,
		IngressType: &contract1.Ingress_Path{Path: "/foo/bar"},
	}

	bindata, err := proto.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	jsondata, err := protojson.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	binfile, err := os.Create("msg.bin")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := binfile.Write(bindata); err != nil {
		log.Fatalln(err)
	}

	jsonfile, err := os.Create("msg.json")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := jsonfile.Write(jsondata); err != nil {
		log.Fatalln(err)
	}

}
