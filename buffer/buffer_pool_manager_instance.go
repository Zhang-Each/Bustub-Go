package buffer

import (
	"bustub/common"
	"bustub/storage/disk"
	"bustub/storage/page"
	"sync"
)

type BufferPoolManagerInstance struct {
	poolSize      int
	numInstance   uint32
	instanceIndex uint32
	nextPageId    page.Page
	pages         []page.Page
	diskManager   disk.DiskManager
	pageTable     map[common.PageId]common.FrameId
	replacer      *Replacer
	freeList      []common.FrameId
	latch         sync.Mutex
}

func (buffer *BufferPoolManagerInstance) GetPoolSize() int {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) FetchPage(pageId common.PageId) *page.Page {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) UnpinPage(pageId common.PageId, isDirty bool) bool {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) FlushPage(pageId common.PageId) bool {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) FlushAllPage() {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) NewPage(pageId common.PageId) *page.Page {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) DeletePage(pageId common.PageId) bool {
	panic("implement me")
}

func (buffer *BufferPoolManagerInstance) AllocatePage() common.PageId {
	panic("implement me")
}
