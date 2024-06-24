package LeetCode705_DesignHashSet

import "container/list"

const base = 769

type MyHashSet struct {
	data []list.List
}

func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (this *MyHashSet) hash(key int) int {
	return key & base
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		k := this.hash(key)
		this.data[k].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	k := this.hash(key)
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[k].Remove(e)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	k := this.hash(key)
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
