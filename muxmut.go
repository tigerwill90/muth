package muxmut

import (
	"net/http"
	"sync/atomic"
	"unsafe"
)

var _ http.Handler = (*MutMux)(nil)

type MutMux[T http.Handler] struct {
	p *unsafe.Pointer
}

func New[T http.Handler](r T) *MutMux[T] {
	m := &MutMux[T]{
		p: (*unsafe.Pointer)(unsafe.Pointer(new(T))),
	}
	m.Update(r)
	return m
}

func (mux *MutMux[T]) Update(r T) {
	atomic.StorePointer(mux.p, unsafe.Pointer(&r))
}

func (mux *MutMux[T]) Swap(new T) (old T) {
	return *(*T)(atomic.SwapPointer(mux.p, unsafe.Pointer(&new)))
}

func (mux *MutMux[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := *(*T)(atomic.LoadPointer(mux.p))
	router.ServeHTTP(w, r)
}
