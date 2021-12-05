package main

// LinkedHashMap的核心逻辑 3 * 如果是访问顺序，那put和get操作已存在的Entry时，都会把Entry移动到双向链表的表尾(其实是先删除再插入)。
// go没有LinkedHashMap，自己实现一个吧
// 维护一个hash，且维护一个操作次数的集合

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type LRUCache struct {
	limit int
	head *node
	m map[int]*node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		limit: capacity,
		head: &node{},
		m: make(map[int]*node),
	}
}

func (c *LRUCache) Get(key int) int {
	n := c.m[key]

	// 不存在
	// 直接返回-1
	if n == nil { return -1 }

	// 存在
	// 将其从双向中删除
	n.prev.next = n.next
	// 将n插入表尾
	n.prev = c.head.prev
	c.head.prev = n
	n.next = c.head

	return n.val
}

func (c *LRUCache) Put(key int, value int) {
	// 不存在，是新插入的，且目前容量已经满了
	if len(c.m) >= c.limit && c.m[key] == nil {
		// 淘汰逻辑
		// 删除map
		delete(c.m, c.head.next.key)
		// 删除头部
		c.head.next = c.head.next.next
	}

	var n *node
	if c.m[key] == nil {
		n = &node{key: key, val: value}
	} else {
		n = c.m[key]
		// 将其从双向中删除
		n.prev.next = n.next
	}

	// 将n插入表尾
	n.prev = c.head.prev
	c.head.prev = n
	n.next = c.head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */