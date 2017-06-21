package cor

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"strconv"
)

type FLAG_TYPE int
type FLAG_TAG uint8

const (
	FLAG_STABLE FLAG_TYPE = iota
	FLAG_SERIALIZE
	FLAG_CUSTOM
	FLAG_FOO3
	FLAG_MAX
)

const (
	STABLE_IO     = 1 << iota
	STABLE_CUSTOM = 1 << iota
	STABLE_UNDEF1 = 1 << iota
	STABLE_UNDEF2 = 1 << iota
	STABLE_UNDEF3 = 1 << iota
	STABLE_UNDEF4 = 1 << iota
	STABLE_UNDEF5 = 1 << iota
	STABLE_UNDEF6 = 1 << iota
)

const (
	SERIALIZE_JSON = iota
	SERIALIZE_PROTOC
	SERIALIZE_GOB
)

type head struct {
	stable [4]uint8
	size   uint64
	custom [4]uint8
}

type Header interface {
	SetCustom(bool)
	Custom() bool
	RW() bool
	SetRW(bool)
	Serialize() uint8
	SetSerialize(s uint8)
	Flag(uint) uint8
	SetFlag(ft FLAG_TYPE, ui uint8)
	Size() uint64
	SetSize(uint64)
	HeadWidth() int
	Bytes() []byte
}

const IO_READ = true
const IO_WRITE = false

var _ Header = NewHead(nil)

func init() {

}

func NewHead(b []byte) *head {
	h := new(head)
	if b == nil {
		return h
	}
	for i := range h.stable {
		h.stable[i] = b[i]
	}

	b_buf := bytes.NewBuffer(b[4:12])
	if e := binary.Read(b_buf, binary.BigEndian, &(h.size)); e != nil {
		log.Println(e)
		return h
	}

	b_buf = bytes.NewBuffer(b[12:16])

	if e := binary.Read(b_buf, binary.BigEndian, &h.custom); e != nil {
		return h
	}

	log.Println("new header: ", h)
	return h
}

func (h *head) Bytes() []byte {

	b_buf := bytes.NewBuffer(h.stable[:])
	if e := binary.Write(b_buf, binary.BigEndian, &h.size); e != nil {
		log.Println(e)
	}

	if e := binary.Write(b_buf, binary.BigEndian, &h.custom); e != nil {
		log.Println(e)
	}

	return b_buf.Bytes()
}

func (h *head) RW() bool {
	ft := FLAG_TAG(h.stable[FLAG_STABLE])
	return ft.BitGet(STABLE_IO)
}

func (h *head) Flag(uint) uint8 {
	return 0
}

func (h *head) SetRW(b bool) {
	ft := FLAG_TAG(h.stable[FLAG_STABLE])
	h.stable[FLAG_STABLE] = ft.BitSet(b, STABLE_IO).Uint8()
}

func (h *head) SetSerialize(s uint8) {
	h.SetFlag(FLAG_SERIALIZE, s)
}

func (h *head) Serialize() uint8 {
	return h.stable[FLAG_SERIALIZE]
}

func (h *head) SetFlag(ft FLAG_TYPE, ui uint8) {
	if h.stable[ft] != ui {
		h.stable[ft] = ui
	}
}

func (h *head) Custom() bool {
	ft := FLAG_TAG(h.stable[FLAG_STABLE])
	return ft.BitGet(STABLE_CUSTOM)
}

func (h *head) SetCustom(b bool) {
	ft := FLAG_TAG(h.stable[FLAG_STABLE])
	h.stable[FLAG_STABLE] = ft.BitSet(b, STABLE_CUSTOM).Uint8()
}

func (h *head) SetSize(size uint64) {
	h.size = size
}

func (h *head) Size() uint64 {
	return h.size
}

func (h *head) HeadWidth() int {
	return binary.Size(*h)
}

func (ft *FLAG_TYPE) Int() int {
	return int(*ft)
}

func ReadHeader(reader io.Reader) (Header, error) {
	b := make([]byte, 16)
	if i, e := reader.Read(b); e != nil {
		log.Println(i, b, e)
		return nil, e

	} else {
		log.Println(i, b, e)
	}

	return NewHead(b), nil

}

func WriteHeader(writer io.Writer, h Header) error {
	log.Println(h.Bytes())
	if i, e := writer.Write(h.Bytes()); e != nil {
		log.Println("w,e:", i, e)
		return e
	} else {
		log.Println("w:", i, e)
	}

	return nil
}

//b:true false
//bit: 0 1 2 3 4 5 6 7
func (f FLAG_TAG) BitSet(b bool, bit uint) FLAG_TAG {

	if bit > 7 {
		return 0
		panic(ERROR_BITS_SET_OVERFLOW)
	}
	bits := uint8(^(1 << bit))
	f &= FLAG_TAG(bits)

	if b {
		f |= 1 << bit
	}
	log.Println("set", strconv.FormatUint(uint64(f), 2))
	return f

}

func (f FLAG_TAG) BitGet(bit uint) bool {
	if bit > 7 {
		return false
	}

	bits := uint8(1 << bit)
	b := uint8(f) & bits >> bit

	log.Println("get", strconv.FormatUint(uint64(f), 2))
	return b == 1
}

func (f *FLAG_TAG) SetUints([]uint8) {

}

func (f FLAG_TAG) Uint8() uint8 {
	return uint8(f)
}

func ParseUint64to8(u uint64) ([]uint8, error) {
	b_buf := bytes.NewBuffer(nil)
	if e := binary.Write(b_buf, binary.BigEndian, &u); e != nil {
		log.Println(e)
		return nil, ERROR_PARSE_TO_UINT8
	}
	return b_buf.Bytes(), nil
}

func ParseUint8to64(u []uint8) (uint64, error) {
	var u64 uint64

	b_buf := bytes.NewBuffer(u)
	if e := binary.Read(b_buf, binary.BigEndian, &u64); e != nil {
		log.Println(e)
		return 0, ERROR_PARSE_TO_UINT64
	}
	return u64, nil
}

//func (f *FLAG_TAG) Uint16() uint16 {
//	return uint16(*f)
//}
//
//func (f *FLAG_TAG) Uint32() uint32 {
//	return uint32(*f)
//}
//
//func (f *FLAG_TAG) Uint64() uint64 {
//	return uint64(*f)
//}
