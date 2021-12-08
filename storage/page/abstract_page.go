package page

import "bustub/common"

type AbstractPage interface {
	GetData() []byte
	GetPageID() common.PageId
	GetPinCount() int
	GetLSN() common.Lsn
	IsDirty() bool
	SetLSN(lsn common.Lsn)
	WLatch()
	WUnlatch()
	RLatch()
	RUnlatch()
}
