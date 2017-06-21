package main

import (
	"github.com/godcong/cor"
	"github.com/godcong/cor/corlog"
)

func main() {
	corlog.LogToFile("client")
	cli := cor.NewClient()
	cli.Head().SetRW(true)
	cli.WriteCallback(func(h cor.Header) ([]byte, error) {
		return []byte("654321"), nil
	})
	cli.Start()

}
