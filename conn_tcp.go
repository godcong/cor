package cor

import (
	"log"
	"net"
)

var _ Connector = new(tcp_server)
var _ Connector = new(tcp_client)

type TCP struct {
	head   *head
	reader func(Header, []byte) error
	writer func(Header) ([]byte, error)
}

type tcp_server struct {
	TCP
	//head   *head
	//reader func(io.Reader, []byte, interface{}) error
	//writer func(io.Writer, []byte, interface{}) error
}

//init tcp_server
func NewServer() *tcp_server {
	s := new(tcp_server)
	s.head = NewHead(nil)
	return s
}

//listen start
func (s *tcp_server) Start() error {
	log.Println("tcp_server start")
	if s == nil {
		return NIL_TARGET
	}

	listen, err := net.ListenTCP(ConnType(), TCPAddr())
	if err != nil {
		panic(err)
	}

	serv := writer
	if s.Head().RW() {
		serv = reader
	}

	defer listen.Close()
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("receive connection failed: ", err.Error())
			continue
		}

		log.Println("connected from " + conn.RemoteAddr().String())
		go serv(conn, s)

	}

	return nil
}
func (s *tcp_server) Head() Header {
	return s.head
}

func (s *tcp_server) ReadAble() bool {
	return s.reader != nil
}

func (s *tcp_server) WriteAble() bool {
	return s.writer != nil
}

func (s *tcp_server) Read(h Header, b []byte) error {
	return s.reader(h, b)
}

func (s *tcp_server) Write(h Header) ([]byte, error) {
	return s.writer(h)
}

func (s *tcp_server) ReadCallback(f func(h Header, b []byte) error) {
	s.reader = f
}

func (s *tcp_server) WriteCallback(f func(h Header) ([]byte, error)) {
	s.writer = f
}

func writer(conn *net.TCPConn, c Connector) {
	log.Println("writeServer: ", conn)
	defer conn.Close()
	var b []byte
	var e error
	h := c.Head()
	//for {
	if c.WriteAble() {
		if b, e = c.Write(h); e != nil {
			log.Println("writer:", e)
			return
		}
		h.SetSize(uint64(len(b)))
	}

	e = WriteHeader(conn, h)
	if e != nil {
		log.Println("disconnect from "+conn.RemoteAddr().String(), e.Error())
	}

	i, e := conn.Write(b)
	log.Println("writed:", i, e)
	if e != nil {
		return
	}
	log.Println("header is ", h, string(b))
	//}
}

func reader(conn *net.TCPConn, c Connector) {
	log.Println("readServer: ", conn)

	var b []byte
	defer conn.Close()
	//for {
	h, e := ReadHeader(conn)
	b = make([]byte, h.Size())

	i, e := conn.Read(b)
	log.Println("read:", i, e)
	if c.ReadAble() {
		if e = c.Read(h, b); e != nil {
			log.Println("reader:", e)
		}

	}

	log.Println("header is ", h, string(b))

	return
	//}

}

type tcp_client struct {
	TCP
}

func NewClient() *tcp_client {
	c := new(tcp_client)
	c.head = NewHead(nil)
	return c
}

func (c *tcp_client) Start() error {
	log.Println("tcp_client start")
	if c == nil {
		return NIL_TARGET
	}

	serv := reader
	if c.Head().RW() {
		serv = writer
	}

	conn, e := net.DialTCP(ConnType(), nil, TCPAddr())
	if e != nil {
		log.Println(e.Error())
		panic(e)
	}
	defer conn.Close()

	serv(conn, c)
	return nil

}
func (c *tcp_client) Head() Header {
	return c.head
}

func (c *tcp_client) ReadCallback(f func(h Header, b []byte) error) {
	c.reader = f

}
func (c *tcp_client) WriteCallback(f func(h Header) ([]byte, error)) {
	c.writer = f
}
func (c *tcp_client) Read(h Header, b []byte) error {
	return c.reader(h, b)
}

func (c *tcp_client) Write(h Header) ([]byte, error) {
	return c.writer(h)
}

func (c *tcp_client) WriteAble() bool {
	return c.writer != nil
}
func (c *tcp_client) ReadAble() bool {
	return c.reader != nil
}
