package latch

type AbstractLatch interface {
	WLock()
	WUnlock()
	RLock()
	RUnlock()
}
