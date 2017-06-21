package cor

type Connector interface {
	//Server() Server
	//Client() Client
	Head() Header
	Start() error
	ReadCallback(func(h Header, b []byte) error)
	WriteCallback(func(h Header) ([]byte, error))
	Read(h Header, b []byte) error
	Write(h Header) ([]byte, error)
	WriteAble() bool
	ReadAble() bool
}

//
//type Server interface {
//	Head() Header
//	Start() error
//	ReadCallback(func(h Header, b []byte) error)
//	WriteCallback(func(h Header) ([]byte, error))
//	Read(h Header, b []byte) error
//	Write(h Header) ([]byte, error)
//	WriteAble() bool
//	ReadAble() bool
//}
//
//type Client interface {
//	Head() Header
//	Start() error
//	ReadCallback(func(h Header, b []byte) error)
//	WriteCallback(func(h Header) ([]byte, error))
//	Read(h Header, b []byte) error
//	Write(h Header) ([]byte, error)
//	WriteAble() bool
//	ReadAble() bool
//}
