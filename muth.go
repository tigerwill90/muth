package muth

import (
	"net/http"
	"sync/atomic"
	"unsafe"
)

var _ http.Handler = (*MutH[http.Handler])(nil)

type MutH[T http.Handler] struct {
	p unsafe.Pointer
}

func Handler[T http.Handler](r T) *MutH[T] {
	m := &MutH[T]{
		p: unsafe.Pointer(&r),
	}
	return m
}

func (h *MutH[T]) Update(new T) {
	atomic.StorePointer(&h.p, unsafe.Pointer(&new))
}

func (h *MutH[T]) Swap(new T) (old T) {
	return *(*T)(atomic.SwapPointer(&h.p, unsafe.Pointer(&new)))
}

func (h *MutH[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(*(*T)(atomic.LoadPointer(&h.p))).ServeHTTP(w, r)
}
