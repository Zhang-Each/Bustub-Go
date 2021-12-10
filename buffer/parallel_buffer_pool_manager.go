package buffer

import (
	"bustub/common"
	"bustub/storage/disk"
	"bustub/storage/page"
)

type ParallelBufferPoolManager struct {
	numInstance  int
	poolSize     int
	starterIndex int
	buffers      []*BufferPoolManagerInstance
	diskManager  disk.DiskManager
}

func (buffer *ParallelBufferPoolManager) GetPoolSize() int {
	panic("implement me")
}

func (buffer *ParallelBufferPoolManager) FetchPage(pageId common.PageId) *page.Page {
	panic("implement me")
}

func (buffer *ParallelBufferPoolManager) UnpinPage(pageID common.PageId, isDirty bool) bool {
	panic("implement me")
}

func (buffer *ParallelBufferPoolManager) FlushPage(pageId common.PageId) bool {
	panic("implement me")
}

func (buffer *ParallelBufferPoolManager) FlushAllPage() {
	panic("implement me")
}

func (buffer *ParallelBufferPoolManager) NewPage(pageId common.PageId) *page.Page {
	panic("implement me")
}

func (buffer *ParallelBufferPoolManager) DeletePage(pageID common.PageId) bool {
	panic("implement me")
}
