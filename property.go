package cor

import (
	"net"

	"strconv"

	"log"

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

func GetPort() string {
	if property != nil {
		return property.MustGet("port", "7777")
	}
	return "7777"
}

func GetAddr() string {
	if property != nil {
		return property.MustGet("addr", "localhost")
	}
	return "localhost"
}

func ListenType() string {
	if property != nil {
		return property.MustGet("type", "tcp")
	}
	return "tcp"
}

func TCPAddr() *net.TCPAddr {
	var addr net.TCPAddr
	var e error

	addr.Port, e = strconv.Atoi(GetPort())

	log.Println(e.Error())

}
