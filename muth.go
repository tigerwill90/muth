package muth

import (
	"net/http"
	"sync/atomic"
)

var _ http.Handler = (*MutH[http.Handler])(nil)

type MutH[T http.Handler] struct {
	p *atomic.Pointer[T]
}

func Handler[T http.Handler](r T) *MutH[T] {
	var p atomic.Pointer[T]
	p.Store(&r)
	m := &MutH[T]{
		p: &p,
	}
	return m
}

func (h *MutH[T]) Update(new T) {
	h.p.Store(&new)
}

func (h *MutH[T]) Swap(new T) (old T) {
	return *h.p.Swap(&new)
}

func (h *MutH[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(*(*T)(h.p.Load())).ServeHTTP(w, r)
}
