package muth

import (
	"net/http"
	"sync/atomic"
	"unsafe"
)

var _ http.Handler = (*MutH[http.Handler])(nil)

type MutH[T http.Handler] struct {
	p *unsafe.Pointer
}

func Handler[T http.Handler](r T) *MutH[T] {
	m := &MutH[T]{
		p: (*unsafe.Pointer)(unsafe.Pointer(new(T))),
	}
	m.Update(r)
	return m
}

func (h *MutH[T]) Update(r T) {
	atomic.StorePointer(h.p, unsafe.Pointer(&r))
}

func (h *MutH[T]) Swap(new T) (old T) {
	return *(*T)(atomic.SwapPointer(h.p, unsafe.Pointer(&new)))
}

func (h *MutH[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := *(*T)(atomic.LoadPointer(h.p))
	handler.ServeHTTP(w, r)
}
