package main

import (
	"net"

	"log"

	"github.com/godcong/cor"
)

func main() {
	NewClient().Start()

}

type client struct {
}

func NewClient() *client {
	cli := new(client)
	return cli
}

func (c *client) Start() {

	//net.DialTCP(cor.ConnType(),cor.TCPAddr())
	h := cor.NewHead(nil)
	str := "123456456456fdsafdsfdsafsdafsdafdsafdsa"
	h.SetSize(uint64(len(str)))
	//h.SetFlag(cor.FLAG_STABLE, 1)
	h.SetIO(cor.FLAG_CLIENT_READ)
	log.Println("call run")
	//conn, e := net.Dial(cor.ConnType(), cor.TCPAddr().String())

	conn, e := net.DialTCP(cor.ConnType(), nil, cor.TCPAddr())
	if e != nil {
		log.Println(e.Error())
		panic(e)
	}
	defer conn.Close()

	e = cor.WriteHeader(conn, h)
	if e != nil {
		log.Println(e)
		return
	}

	b := make([]byte, h.Size())

	if h.ReadOrWrite() != cor.FLAG_CLIENT_READ {
		conn.Read(b)
		log.Println(string(b))
	} else {
		_, e = conn.Write([]byte(str))
		if e != nil {
			log.Println(e)
			return
		}

	}

}
