// /////////////////////////////////////////////////////////////////////////////
// int 和 bool 类型的原子变量

package syncutil

import "sync/atomic"

// /////////////////////////////////////////////////////////////////////////////
// Int32 原子变量

// int32 类型的 原子变量
type AtomicInt32 int32

// 原子性的将 v 的值保存到 *AtomicInt32
func (ai *AtomicInt32) Store(v int32) {
	atomic.StoreInt32((*int32)(ai), v)
}

// 原子性的将 delta 的值添加到 *AtomicInt32 并返回新值。
func (ai *AtomicInt32) Add(delta int32) (new int32) {
	return atomic.AddInt32((*int32)(ai), delta)
}

// 原子性的获取 *AtomicInt32 的值。
func (ai *AtomicInt32) Load() int32 {
	return atomic.LoadInt32((*int32)(ai))
}

// 原子性的将 v 的值 与 原来的值 交换，返回原来的值
func (ai *AtomicInt32) Swap(v int32) (old int32) {
	return atomic.SwapInt32((*int32)(ai), v)
}

// /////////////////////////////////////////////////////////////////////////////
// Uint32 原子变量
type AtomicUint32 uint32

// 原子性的将 v 的值保存到 *AtomicUint32
func (ui32 *AtomicUint32) Store(v uint32) {
	atomic.StoreUint32((*uint32)(ui32), v)
}

// 原子性的将 delta 的值添加到 *AtomicUint32 并返回新值。
func (ui32 *AtomicUint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32((*uint32)(ui32), delta)
}

// 原子性的获取 *AtomicUint32 的值。
func (ui32 *AtomicUint32) Load() uint32 {
	return atomic.LoadUint32((*uint32)(ui32))
}

// 原子性的将 v 的值 与 原来的值 交换，返回原来的值
func (ui32 *AtomicUint32) Swap(v uint32) (old uint32) {
	return atomic.SwapUint32((*uint32)(ui32), v)
}

// /////////////////////////////////////////////////////////////////////////////
// int 原子变量

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

// /////////////////////////////////////////////////////////////////////////////
// int64 原子变量

// int64 类型的 原子变量
type AtomicInt64 int64

// 原子性的将 v 的值保存到 *AtomicInt64
func (ai *AtomicInt64) Store(v int64) {
	atomic.StoreInt64((*int64)(ai), v)
}

// 原子性的将 delta 的值添加到 *AtomicInt64 并返回新值。
func (ai *AtomicInt64) Add(delta int64) (new int64) {
	return atomic.AddInt64((*int64)(ai), delta)
}

// 原子性的获取 *AtomicInt64 的值。
func (ai *AtomicInt64) Load() int64 {
	return atomic.LoadInt64((*int64)(ai))
}

// 将 v 与 原来的值交换，并返回原来的值
func (ai *AtomicInt64) Swap(v int64) (old int64) {
	return atomic.SwapInt64((*int64)(ai), v)
}

// /////////////////////////////////////////////////////////////////////////////
// bool 原子变量

// bool 类型的 原子变量
type AtomicBool AtomicInt32

// 原子性的将 v 的值保存到 *AtomicBool
func (ab *AtomicBool) Store(v bool) {
	if v {
		(*AtomicInt32)(ab).Store(1)
	} else {
		(*AtomicInt32)(ab).Store(0)
	}
}

// 原子性的获取 *AtomicBool 的值。
func (ab *AtomicBool) Load() bool {
	return (*AtomicInt32)(ab).Load() != 0
}

// 将 v 与 原来的值交换，并返回原来的值
func (ab *AtomicBool) Swap(v bool) bool {
	// 原来的值
	old := ((*AtomicInt32)(ab).Load() != 0)

	// 交换
	if v {
		(*AtomicInt32)(ab).Store(0)
	} else {
		(*AtomicInt32)(ab).Store(1)
	}

	// 返回原来的值
	return old
}
