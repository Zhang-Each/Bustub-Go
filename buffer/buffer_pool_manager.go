package buffer

import (
	"bustub/common"
	"bustub/storage/page"
)

// BufferPoolManager Basic API for a Buffer Pool Manager
type BufferPoolManager interface {
	GetPoolSize() int
	FetchPage(id common.PageId) *page.Page
	UnpinPage(id common.PageId, isDirty bool) bool
	FlushPage(id common.PageId) bool
	FlushAllPage()
	NewPage(id common.PageId) *page.Page
	DeletePage(id common.PageId) bool
}
