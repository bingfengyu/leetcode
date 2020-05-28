package main

import "fmt"

func main() {
	cache := Constructor(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	fmt.Println(cache.Get(4))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))

}

//题目地址：https://leetcode-cn.com/problems/lfu-cache/
type LFUCache struct {
	len      int //长度
	capacity int //容纳量
	nodes    map[int]*FreqNode
	lsFreq   *FreqNode //头节点
}

type cacheNode struct {
	key   int
	value int
}

type FreqNode struct {
	cache    *cacheNode
	value    int
	freq     int       //访问频率数
	preNode  *FreqNode //前一个结点
	nextNode *FreqNode //后一个节点
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes:    make(map[int]*FreqNode, capacity),
		lsFreq:   &FreqNode{},
	}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.nodes[key]; ok {
		v.freq += 1
		this.adjustOrder(v)
		return v.cache.value
	}

	return -1
}

func (this *LFUCache) adjustOrder(node *FreqNode) {
	for next := node.nextNode; next != nil; {
		if next.freq > node.freq {
			break
		}
		pre := node.preNode
		pre.nextNode = next
		node.nextNode = next.nextNode
		next.nextNode = node
		next.preNode = pre
		node.preNode = next
		next = node.nextNode
		if next != nil {
			node.nextNode.preNode = node
		}
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	if v, ok := this.nodes[key]; ok {
		v.cache.value = value
		v.freq += 1
		this.adjustOrder(v)
		return
	}

	if this.len < this.capacity {
		node := &FreqNode{
			cache: &cacheNode{
				key:   key,
				value: value,
			},
			freq:     1,
			nextNode: this.lsFreq.nextNode,
			preNode:  this.lsFreq,
		}

		this.lsFreq.nextNode = node
		if node.nextNode != nil {
			node.nextNode.preNode = node
		}

		this.adjustOrder(node)
		this.nodes[key] = node

		this.len++
		return
	}

	if this.len == this.capacity {
		rmNode := this.lsFreq.nextNode

		this.lsFreq.nextNode = rmNode.nextNode
		if rmNode.nextNode != nil {
			rmNode.nextNode.preNode = this.lsFreq
		}
		delete(this.nodes, rmNode.cache.key)

		this.len--

		node := &FreqNode{
			cache: &cacheNode{
				key:   key,
				value: value,
			},
			freq:     1,
			nextNode: this.lsFreq.nextNode,
			preNode:  this.lsFreq,
		}
		this.lsFreq.nextNode = node
		if node.nextNode != nil {
			node.nextNode.preNode = node
		}

		this.adjustOrder(node)
		this.nodes[key] = node
		this.len++
	}
}
