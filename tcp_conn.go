package cor

type TCP struct {
	header *head
}

type Sender interface {
}

type Receiver interface {
}

func (tcp *TCP) Send(b []byte) {

}

func (tcp *TCP) Receive() []byte {
	return nil
}
