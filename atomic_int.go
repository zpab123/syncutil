// /////////////////////////////////////////////////////////////////////////////
// int 和 bool 类型的原子变量

package syncutil

import "sync/atomic"

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// int 类型的 原子变量 （用 int64 表示的）
type AtomicInt int64

// 原子性的将 v 的值保存到 *AtomicInt
func (ai *AtomicInt) Store(v int) {
	atomic.StoreInt64((*int64)(ai), int64(v))
}

// 原子性的将 delta 的值添加到 *AtomicInt 并返回新值。
func (ai *AtomicInt) Add(delta int) {
	atomic.AddInt64((*int64)(ai), int64(delta))
}

// 原子性的获取 *AtomicInt 的值。
func (ai *AtomicInt) Load() int {
	return int(atomic.LoadInt64((*int64)(ai)))
}

// int64 类型的 原子变量
type AtomicInt64 int64

// 原子性的将 v 的值保存到 *AtomicInt64
func (ai *AtomicInt64) Store(v int64) {
	atomic.StoreInt64((*int64)(ai), v)
}

// 原子性的将 delta 的值添加到 *AtomicInt64 并返回新值。
func (ai *AtomicInt64) Add(delta int64) {
	atomic.AddInt64((*int64)(ai), delta)
}

// 原子性的获取 *AtomicInt64 的值。
func (ai *AtomicInt64) Load() int64 {
	return atomic.LoadInt64((*int64)(ai))
}

// int32 类型的 原子变量
type AtomicInt32 int32

// 原子性的将 v 的值保存到 *AtomicInt32
func (ai *AtomicInt32) Store(v int32) {
	atomic.StoreInt32((*int32)(ai), v)
}

// 原子性的将 delta 的值添加到 *AtomicInt32 并返回新值。
func (ai *AtomicInt32) Add(delta int32) {
	atomic.AddInt32((*int32)(ai), delta)
}

// 原子性的获取 *AtomicInt32 的值。
func (ai *AtomicInt32) Load() int32 {
	return atomic.LoadInt32((*int32)(ai))
}

// bool 类型的 原子变量
type AtomicBool AtomicInt

// 原子性的将 v 的值保存到 *AtomicBool
func (ab *AtomicBool) Store(v bool) {
	if v {
		(*AtomicInt)(ab).Store(1)
	} else {
		(*AtomicInt)(ab).Store(0)
	}
}

// 原子性的获取 *AtomicBool 的值。
func (ab *AtomicBool) Load() bool {
	return (*AtomicInt)(ab).Load() != 0
}
