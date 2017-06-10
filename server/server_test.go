package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestWrite(t *testing.T) {
	msg := &Helloworld{
		Id1: proto.Int32(101),
		Str: proto.String("hello"),
	} //msg init

	path := string("./log.txt")
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer f.Close()
	buffer, err := proto.Marshal(msg) //SerializeToOstream
	f.Write(buffer)

}

func TestRead(t *testing.T) {
	path := string("./log.txt")
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}
	b := make([]byte, 1024)

	f.Read(b)
	fmt.Println(b)
	var hello Helloworld
	proto.Unmarshal(b, &hello)

	fmt.Println(*hello.Id1, hello.Opt, *hello.Str)
}
