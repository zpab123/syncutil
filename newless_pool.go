// /////////////////////////////////////////////////////////////////////////////
// less 对象池

package syncutil

import "sync"

// /////////////////////////////////////////////////////////////////////////////
// 构造函数

// less 对象池
type NewlessPool struct {
	lock      sync.Mutex // 互斥锁
	availCond *sync.Cond // 某事件的发生
	items     []interface{}
}

// 创建新1个的 NewlessPool 对象
func NewNewlessPool() *NewlessPool {
	pool := &NewlessPool{
		items: nil,
	}
	pool.availCond = sync.NewCond(&pool.lock)
	return pool
}

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// 从 NewlessPool 获取1个 对象
func (pool *NewlessPool) Get() (v interface{}) {
	pool.lock.Lock()
	for len(pool.items) == 0 {
		pool.availCond.Wait()
	}
	// pop the last item
	n := len(pool.items)
	v, pool.items[n-1] = pool.items[n-1], nil
	pool.items = pool.items[:n-1]
	pool.lock.Unlock()
	return
}

// 从 NewlessPool 尝试获取1个 对象
func (pool *NewlessPool) TryGet() (v interface{}) {
	pool.lock.Lock()
	n := len(pool.items)
	if n > 0 {
		// pop the last item
		v, pool.items[n-1] = pool.items[n-1], nil
		pool.items = pool.items[:n-1]
	}

	pool.lock.Unlock()
	return
}

// 从 NewlessPool 放回1个 对象
func (pool *NewlessPool) Put(v interface{}) {
	pool.lock.Lock()
	pool.items = append(pool.items, v)
	pool.availCond.Signal()
	pool.lock.Unlock()
}
