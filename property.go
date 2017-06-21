package cor

import (
	"net"

	"strings"

	"github.com/g7n3/configo"
)

var property *configo.Property

//two types server/client
func Load(s string) error {
	if pro, e := configo.Get(s); e == nil {
		property = pro
	} else {
		return e
	}

	return nil
}

func Port() string {
	if property != nil {
		return property.MustGet("port", "7777")
	}
	return "7777"
}

func Addr() string {
	if property != nil {
		return property.MustGet("addr", "127.0.0.1")
	}
	return "127.0.0.1"
}

func ConnType() string {
	if property != nil {
		return property.MustGet("type", "tcp")
	}
	return "tcp"
}

func NetType() string {
	if property != nil {
		return property.MustGet("addr", "127.0.0.1")
	}
	return "127.0.0.1"
}

func TCPAddr() *net.TCPAddr {
	sadr := strings.Join([]string{Addr(), Port()},
		":")

	addr, e := net.ResolveTCPAddr("tcp4", sadr)
	if e != nil {
		return nil
	}

	return addr
}
