// 消息队列包
package queueutil

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

type esCache struct {
	putNo uint32
	getNo uint32
	value interface{}
}

// lock free queue
type EsQueue struct {
	capaciity uint32
	capMod    uint32
	putPos    uint32
	getPos    uint32
	cache     []esCache
}

// NewQueue 创建消息队列
func NewQueue(capaciity uint32) *EsQueue {
	q := new(EsQueue)
	q.capaciity = minQuantity(capaciity)
	q.capMod = q.capaciity - 1
	q.putPos = 0
	q.getPos = 0
	q.cache = make([]esCache, q.capaciity)
	for i := range q.cache {
		cache := &q.cache[i]
		cache.getNo = uint32(i)
		cache.putNo = uint32(i)
	}
	cache := &q.cache[0]
	cache.getNo = q.capaciity
	cache.putNo = q.capaciity
	return q
}

// String 获取消息队列的描述
// 如容量大小，用了多少用量等等
func (q *EsQueue) String() string {
	getPos := atomic.LoadUint32(&q.getPos)
	putPos := atomic.LoadUint32(&q.putPos)
	return fmt.Sprintf("Queue{capaciity: %v, capMod: %v, putPos: %v, getPos: %v}",
		q.capaciity, q.capMod, putPos, getPos)
}

// Capaciity 获取消息队列的容量
func (q *EsQueue) Capaciity() uint32 {
	return q.capaciity
}

// Quantity 计算消息队列的容量
func (q *EsQueue) Quantity() uint32 {
	var putPos, getPos uint32
	var quantity uint32
	getPos = atomic.LoadUint32(&q.getPos)
	putPos = atomic.LoadUint32(&q.putPos)

	if putPos >= getPos {
		quantity = putPos - getPos
	} else {
		quantity = q.capMod + (putPos - getPos)
	}

	return quantity
}

// Put 将数据存入消息队列
// 返回是否存入成功以及已经在队列的元素个数
func (q *EsQueue) Put(val interface{}) (ok bool, quantity uint32) {
	var putPos, putPosNew, getPos, posCnt uint32
	var cache *esCache
	capMod := q.capMod

	getPos = atomic.LoadUint32(&q.getPos)
	putPos = atomic.LoadUint32(&q.putPos)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt >= capMod-1 {
		runtime.Gosched()
		return false, posCnt
	}

	putPosNew = putPos + 1
	if !atomic.CompareAndSwapUint32(&q.putPos, putPos, putPosNew) {
		runtime.Gosched()
		return false, posCnt
	}

	cache = &q.cache[putPosNew&capMod]

	for {
		getNo := atomic.LoadUint32(&cache.getNo)
		putNo := atomic.LoadUint32(&cache.putNo)
		if putPosNew == putNo && getNo == putNo {
			cache.value = val
			atomic.AddUint32(&cache.putNo, q.capaciity)
			return true, posCnt + 1
		} else {
			runtime.Gosched()
		}
	}
}

// Get 从消息队列中取出数据
// 返回取到的数据、是否取出成功，队列当前的容量
func (q *EsQueue) Get() (val interface{}, ok bool, quantity uint32) {
	var putPos, getPos, getPosNew, posCnt uint32
	var cache *esCache
	capMod := q.capMod

	putPos = atomic.LoadUint32(&q.putPos)
	getPos = atomic.LoadUint32(&q.getPos)

	if putPos >= getPos {
		posCnt = putPos - getPos
	} else {
		posCnt = capMod + (putPos - getPos)
	}

	if posCnt < 1 {
		runtime.Gosched()
		return nil, false, posCnt
	}

	getPosNew = getPos + 1
	if !atomic.CompareAndSwapUint32(&q.getPos, getPos, getPosNew) {
		runtime.Gosched()
		return nil, false, posCnt
	}

	cache = &q.cache[getPosNew&capMod]

	for {
		getNo := atomic.LoadUint32(&cache.getNo)
		putNo := atomic.LoadUint32(&cache.putNo)
		if getPosNew == getNo && getNo == putNo-q.capaciity {
			val = cache.value
			atomic.AddUint32(&cache.getNo, q.capaciity)
			return val, true, posCnt - 1
		} else {
			runtime.Gosched()
		}
	}
}

// round 到最近的2的倍数
func minQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
