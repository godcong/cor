package main

import (
	"github.com/godcong/cor"
	"github.com/godcong/cor/corlog"
)

func main() {
	corlog.LogToFile("server")
	serv := cor.NewServer()
	serv.Head().SetRW(true)
	serv.WriteCallback(func(h cor.Header) ([]byte, error) {
		return []byte("123456"), nil
	})
	serv.Start()

}
