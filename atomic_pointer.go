// /////////////////////////////////////////////////////////////////////////////
// 跳过安全检测的 原子指针

package syncutil

import (
	"sync/atomic"
	"unsafe"
)

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// Pointer 类型用于表示任意类型的指针
type AtomicPointer struct {
	ptr unsafe.Pointer
}

// 原子性的将 pointer 的值保存到 *AtomicPointer
func (ap *AtomicPointer) Store(pointer unsafe.Pointer) {
	atomic.StorePointer(&ap.ptr, pointer)
}

// 原子性的获取 *AtomicPointer 的值
func (ap *AtomicPointer) Load() unsafe.Pointer {
	return atomic.LoadPointer(&ap.ptr)
}
