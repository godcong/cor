package main

import (
	"encoding/json"
	"math"

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

}
