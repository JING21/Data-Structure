package LeetCode706_DesignHashMap

import "container/list"

const base = 31

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(entry); v.key == key {
			e.Value = entry{key: key, value: value}
			return
		}
	}
	this.data[h].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if v := e.Value.(entry); v.key == key {
			return v.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[h].Remove(e)
		}
	}
}
