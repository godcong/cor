package cor

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestHead_ReadOrWrite(t *testing.T) {
	log.Println(FLAG_IO)
	log.Println(FLAG_ST)

	head := NewHead(nil)
	head.SetFlag(FLAG_IO, 254)
	head.SetSize(123456789)
	log.Println(head.ReadOrWrite())
	log.Println("width", head.HeadWidth())

}
func TestHead_Bytes(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_IO, 254)
	head.SetFlag(FLAG_ST, 254)
	head.SetSize(123456789)
	log.Println(head.Bytes())

}

func TestNewHead(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_IO, 254)
	head.SetFlag(FLAG_ST, 254)
	head.SetSize(123456789)
	log.Println(head.Bytes())

	head2 := NewHead(head.Bytes())
	log.Println(head2.Bytes())

	fmt.Println(strconv.ParseInt("255", 2, 10))
}

func TestHead_SetFlag(t *testing.T) {

}
