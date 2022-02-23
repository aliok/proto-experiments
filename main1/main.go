package main

import (
	"github.com/aliok/proto-experiments/contract1"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	org := &contract1.Ingress{
		ContentMode: contract1.ContentMode_BINARY,
		IngressType: &contract1.Ingress_Path{Path: "/foo/bar"},
	}

	data, err := proto.Marshal(org)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("msg.bin")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := file.Write(data); err != nil {
		log.Fatalln(err)
	}
}
