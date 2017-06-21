package cor

import (
	"log"
	"strconv"
	"testing"
)

func TestHead_ReadOrWrite(t *testing.T) {
	log.Println(FLAG_STABLE)
	log.Println(FLAG_SERIALIZE)

	head := NewHead(nil)
	head.SetFlag(FLAG_STABLE, 254)
	head.SetSize(123456789)
	log.Println(head.Reader())
	log.Println("width", head.HeadWidth())
}
func TestHead_Bytes(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_STABLE, 254)
	head.SetFlag(FLAG_SERIALIZE, 254)
	head.SetSize(4294967297)
	log.Println(head.Bytes())

}

func TestNewHead(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_STABLE, 254)
	head.SetFlag(FLAG_SERIALIZE, 254)
	head.SetSize(4294967297)
	log.Println(head.Bytes())

	head2 := NewHead(head.Bytes())
	log.Println(head2.Bytes())
}

func TestHead_SetFlag(t *testing.T) {

	log.Println(strconv.ParseInt("11111111", 2, 32))
	log.Println(strconv.FormatInt(254, 2))
}

func TestHead_IO(t *testing.T) {
	head := NewHead(nil)
	head.SetFlag(FLAG_STABLE, 0)
	head.SetRW(true)
	head.SetRW(false)
	head.SetFlag(FLAG_STABLE, 255)
	head.SetRW(true)
	head.SetRW(false)
}

func TestFLAG_TAG_BitSet(t *testing.T) {

	for i := uint(0); i < 8; i++ {
		ft := FLAG_TAG(0xFF)
		ft.BitSet(true, i)
		ft.BitSet(false, i)

	}
}
func TestParseUint64to8(t *testing.T) {
	log.Println(ParseUint64to8(4294967297))
}

func TestParseUint8to64(t *testing.T) {
	u, _ := ParseUint64to8(4294967297)
	log.Println(ParseUint8to64(u))

}
