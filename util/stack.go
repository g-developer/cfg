package util

import (
	"sync"
)

type el interface{}

type node struct {
	data *el
	next *node
	prev *node
}

type Stack struct {
	size  int
	data  *node
	mutex *sync.Mutex
}

func NewStack() *Stack {
	return &Stack{0, nil, &sync.Mutex{}}
}

func (self *Stack) Top() el {
	if self.size > 0 {
		self.mutex.Lock()
		index := self.size - 1
		tmp := self.data
		for i := 0; i < index; i++ {
			tmp = tmp.next
		}
		self.mutex.Unlock()
		return *(tmp.data)
	} else {
		return nil
	}
}

func (self *Stack) Pop() el {
	if self.size > 0 {
		var ret el
		self.mutex.Lock()
		var n *node = self.data
		index := self.size - 1
		for i := 0; i < index; i++ {
			n = n.next
		}
		ret = *(n.data)
		n.prev.next = nil
		self.size = self.size - 1
		self.mutex.Unlock()
		return ret
	} else {
		return nil
	}
}

func (self *Stack) Push(one el) {
	self.mutex.Lock()
	n := &node{&one, nil, nil}
	if 0 == self.size {
		self.data = n
	} else {
		index := self.size - 1
		var tmp *node = self.data
		for i := 0; i < index; i++ {
			tmp = tmp.next
		}
		tmp.next = n
		n.prev = tmp
	}
	self.size += 1
	self.mutex.Unlock()
}

func (self *Stack) Size() int {
	return self.size
}
