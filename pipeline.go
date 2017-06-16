package cor

import (
	"container/list"
	"sync"
)

type Pipe struct {
	*list.List
	sync.Mutex
}

var pip *Pipe

func init() {
	pip = new(Pipe)
	pip.List = list.New()
	//pip.Mutex = sync.Mutex{}
}

func (p *Pipe) Push(v interface{}) {
	p.Lock()
	defer p.Unlock()
	p.PushBack(v)
}

func (p *Pipe) Pop() interface{} {
	p.Lock()
	defer p.Unlock()
	v := p.Front()
	p.Remove(v)
	return v
}
