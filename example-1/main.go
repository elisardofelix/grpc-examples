package main

import (
	"log"

	"github.com/elisardofelix/grpc-examples/example-1/proto"
)

func main() {
	person := proto.Person{
		Name: "Chris",
	}

	log.Println(person.GetName())
}