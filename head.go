package cor

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type FLAG_TYPE int

const (
	FLAG_IO FLAG_TYPE = iota
	FLAG_FOO1
	FLAG_FOO2
	FLAG_FOO3
	FLAG_MAX
)

const (
	FLAG_IO_RW  = 1 << iota
	FLAG_IO_ST1 = 1 << iota
	FLAG_IO_ST2 = 1 << iota
	FLAG_IO_ST3 = 1 << iota
	FLAG_IO_ST4 = 1 << iota
)

type head struct {
	flag [4]uint8
	size uint64
	tmp  uint32
}

type Header interface {
	ReadOrWrite() bool
	Flag(int) uint8
	SetFlag(ft FLAG_TYPE, ui uint8) error
	Size() uint64
	SetSize(uint64)
	HeadWidth() int
}

const FLAG_CLIENT_READ = true
const FLAG_CLIENT_WRITE = false

func NewHead(b []byte) *head {
	h := new(head)
	if b == nil {
		return h
	}

	for i := range h.flag {
		h.flag[i] = b[i]
	}

	b_buf := bytes.NewBuffer(b[4:12])
	if e := binary.Read(b_buf, binary.BigEndian, &(h.size)); e != nil {
		log.Println(e)
		return h
	}

	b_buf = bytes.NewBuffer(b[12:16])

	if e := binary.Read(b_buf, binary.BigEndian, &h.tmp); e != nil {
		return h
	}

	log.Println("new header: ", h)
	return h
}

func (h *head) Bytes() []byte {
	b_buf := bytes.NewBuffer(h.flag[:])

	if e := binary.Write(b_buf, binary.BigEndian, &h.size); e != nil {
		log.Println(e)
	}

	if e := binary.Write(b_buf, binary.BigEndian, &h.tmp); e != nil {
		log.Println(e)
	}

	return b_buf.Bytes()
}

func (h *head) ReadOrWrite() bool {
	rw := h.flag[FLAG_IO] & FLAG_IO_RW

	if rw != 0 {
		return FLAG_CLIENT_READ
	}

	return FLAG_CLIENT_WRITE
}

func (h *head) SetFlag(ft FLAG_TYPE, ui uint8) error {
	if h.flag[ft] != ui {
		h.flag[ft] = ui
		return nil
	}
	return ERROR_HEADER_FLAG_SET_ERROR
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

func ReadHeader(reader io.Reader) (*head, error) {
	b := make([]byte, 16)
	if i, e := reader.Read(b); e != nil {
		log.Println(i, b, e)
		return nil, e

	} else {
		log.Println(i, b, e)
	}

	return NewHead(b), nil

}

func WriteHeader(writer io.Writer, h *head) error {
	log.Println(h.Bytes())
	if i, e := writer.Write(h.Bytes()); e != nil {
		log.Println("write1", i, e)
		return e
	} else {
		log.Println("write2", i, e)
	}

	return nil
}
