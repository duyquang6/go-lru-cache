package lru

import "container/list"

type LRUCache struct {
	cap            int
	m              map[string]*list.Element
	leastUsedQueue list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, m: make(map[string]*list.Element)}
}

func (this *LRUCache) Get(key string) interface{} {
	e, ok := this.m[key]
	if !ok {
		return -1
	}
	this.leastUsedQueue.MoveToFront(e)
	return e.Value.(Node).Val
}

type Node struct {
	Key string
	Val interface{}
}

func (this *LRUCache) Put(key string, value interface{}) {
	elem, ok := this.m[key]
	if !ok {
		if this.leastUsedQueue.Len() == this.cap {
			e := this.leastUsedQueue.Back()
			delete(this.m, e.Value.(Node).Key)
			this.leastUsedQueue.Remove(e)
			this.m[key] = this.leastUsedQueue.PushFront(Node{Key: key, Val: value})
		} else if this.leastUsedQueue.Len() < this.cap {
			this.m[key] = this.leastUsedQueue.PushFront(Node{Key: key, Val: value})
		}
	} else {
		this.leastUsedQueue.Remove(elem)
		this.m[key] = this.leastUsedQueue.PushFront(Node{Key: key, Val: value})
	}
}
