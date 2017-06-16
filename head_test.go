package cor

import (
	"log"
	"testing"
)

func TestHead_ReadOrWrite(t *testing.T) {
	log.Println(FLAG_IO_RW)
	log.Println(FLAG_IO_ST1)
	log.Println(FLAG_IO_ST2)
	log.Println(FLAG_IO_ST3)
	log.Println(FLAG_IO_ST4)
	head := NewHead(nil)
	head.SetFlag(FLAG_IO, 254)
	head.SetSize(123456789)
	log.Println(head.ReadOrWrite())
	log.Println("width", head.HeadWidth())

}
func TestHead_Bytes(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_IO, 254)
	head.SetFlag(FLAG_FOO1, 254)
	head.SetSize(123456789)
	log.Println(head.Bytes())

}

func TestNewHead(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_IO, 254)
	head.SetFlag(FLAG_FOO1, 254)
	head.SetSize(123456789)
	log.Println(head.Bytes())

	head2 := NewHead(head.Bytes())
	log.Println(head2.Bytes())
}
