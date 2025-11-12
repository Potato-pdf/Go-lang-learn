package main 

import "sync"

type Buffer struct{
	mu	sync.Mutex
	cond 	*sync.Cond
	item []int
}

func NewBuffer()*Buffer{
	b := &Buffer{item: []int{}}
	b.cond = sync.NewCond(&b.mu)
	return b
}