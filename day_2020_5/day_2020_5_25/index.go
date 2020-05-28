package main

import "fmt"

func main() {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))                  // 返回  1
	cache.Put(3, 3)                            // 该操作会使得关键字 2 作废
	fmt.Println(cache.Get(2), len(cache.mLru)) // 返回 -1 (未找到)
	cache.Put(4, 4)                            // 该操作会使得关键字 1 作废
	fmt.Println(cache.Get(1))                  // 返回 -1 (未找到)
	fmt.Println(cache.Get(3))                  // 返回  3
	fmt.Println(cache.Get(4))                  // 返回  4
}

//试题地址：https://leetcode-cn.com/problems/lru-cache/
type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	mLru     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		head:     &Node{},
		tail:     &Node{},
		mLru:     make(map[int]*Node, capacity),
	}
	cache.head.next = cache.tail
	cache.tail.next = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mLru[key]; !ok {
		return -1
	} else {
		//删除这个节点
		node.pre.next = node.next
		node.next.pre = node.pre
		//将该节点插入头部
		node.next = this.head.next
		node.pre = this.head
		this.head.next.pre = node
		this.head.next = node
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	//该节点存在
	if node, ok := this.mLru[key]; ok {
		//更新数据
		node.value = value
		//删除这个节点
		node.pre.next = node.next
		node.next.pre = node.pre
		//将该节点插入头部
		node.next = this.head.next
		node.pre = this.head
		this.head.next.pre = node
		this.head.next = node
	} else {
		//超出长度
		if len(this.mLru) >= this.capacity {
			//删除最后一个节点
			delete(this.mLru, this.tail.pre.key)
			//移除该节点
			tmp := this.tail.pre
			tmp.pre.next = this.tail
			this.tail.pre = tmp.pre
		}
		//创建新节点，在节点头部
		newNode := &Node{key: key, value: value}
		//将该节点插入头部
		newNode.next = this.head.next
		newNode.pre = this.head
		this.head.next.pre = newNode
		this.head.next = newNode
		this.mLru[key] = newNode
	}
}
