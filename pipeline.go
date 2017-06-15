package cor

import (
	"container/list"
	"sync"
)

type Pipe list.List

var (
	pipe  Pipe
	mutex sync.Mutex
)

func init() {
	pipe = list.New()
}

func (p *Pipe) Push(v interface{}) {
	if p.Inn.Back() == nil {
		pipe.PushBack(v)
	}
}
