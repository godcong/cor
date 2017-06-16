package main

import (
	"encoding/json"
	"log"
	"math"

	"net"

	"github.com/godcong/cor"
	"github.com/golang/protobuf/proto"
)

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Helloworld struct {
	Id1              *int32  `protobuf:"varint,1,req,name=id" json:"id1,omitempty"`
	Str              *string `protobuf:"bytes,2,req,name=str" json:"str,omitempty"`
	Opt              *int32  `protobuf:"varint,3,opt,name=opt" json:"opt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Helloworld) Reset()         { *this = Helloworld{} }
func (this *Helloworld) String() string { return proto.CompactTextString(this) }
func (*Helloworld) ProtoMessage()       {}

func (this *Helloworld) GetId1() int32 {
	if this != nil && this.Id1 != nil {
		return *this.Id1
	}
	return 0
}

func (this *Helloworld) GetStr() string {
	if this != nil && this.Str != nil {
		return *this.Str
	}
	return ""
}

func (this *Helloworld) GetOpt() int32 {
	if this != nil && this.Opt != nil {
		return *this.Opt
	}
	return 0
}

func main() {
	NewServer().Start()
}

type server struct {
}

//init server
func NewServer() *server {
	s := new(server)
	return s
}

//listen start
func (s *server) Start() error {
	if s == nil {
		return cor.NIL_TARGET
	}

	listen, err := net.ListenTCP(cor.ConnType(), cor.TCPAddr())
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("receive connection failed: ", err.Error())
			continue
		}

		log.Println("connected from " + conn.RemoteAddr().String())
		go handleClient(conn)

	}

	defer listen.Close()
	return nil
}

func handleClient(conn *net.TCPConn) {
	log.Println(conn)

	var bytes []byte
	var vb []byte
	str := "123123123123123fdasfdasfdsafdsaa"
	defer conn.Close()
	for {
		h, e := cor.ReadHeader(conn)

		if e != nil {
			log.Println("disconnect from "+conn.RemoteAddr().String(), e.Error())
			break
		}

		bytes = make([]byte, h.Size())

		if h.ReadOrWrite() == cor.FLAG_CLIENT_READ {
			i, e := conn.Write([]byte(str))
			log.Println("write:", i, e)
		} else {
			i, e := conn.Read(bytes)
			log.Println("read:", i, e)
		}

		log.Println("header is ", h, string(bytes))

		return
	}
	log.Println("value is :", string(vb), len(vb))
}
