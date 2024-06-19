package main

import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		list:      list.New(),
		keyToNode: map[int]*list.Element{},
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.keyToNode[key]
	//如果不存在，返回-1
	if node == nil {
		return -1
	}
	//如果存在，返回并且放到最上方
	this.list.MoveToFront(node)
	return node.Value.(entry).value
}

func (this *LRUCache) Put(key int, value int) {
	//如果存在，更新，并且移到最上方
	if node := this.keyToNode[key]; node != nil {
		node.Value = entry{key, value}
		this.list.MoveToFront(node)
		return
	}
	//如果不存在，放置在最上面，输入值
	this.keyToNode[key] = this.list.PushFront(entry{key, value})
	//大于负载，删除最底层的
	if len(this.keyToNode) > this.capacity {
		delete(this.keyToNode, this.list.Remove(this.list.Back()).(entry).key)
	}
}
