package page

import (
	"bustub/common"
	"bustub/common/latch"
	"unsafe"
)

const (
	SIZE_PAGE_HEADER  int = 8
	OFFSET_PAGE_START int = 0
	OFFSET_LSN        int = 4
)

type Page struct {
	data     [common.PAGE_SIZE]byte
	pageId   common.PageId
	pinCount int
	isDirty  bool
	latch    latch.ReaderWriterLatch
}

func (page *Page) GetData() []byte {
	return page.data[:]
}

func (page *Page) GetPageID() common.PageId {
	return page.pageId
}

func (page *Page) GetPinCount() int {
	return page.pinCount
}

func (page *Page) GetLSN() common.Lsn {
	data := page.GetData()
	ptr := unsafe.Pointer(&data[OFFSET_LSN])
	return *(*common.Lsn)(ptr)
}

func (page *Page) IsDirty() bool {
	return page.isDirty
}

func (page *Page) SetLSN(lsn common.Lsn) {
	data := page.GetData()
	ptr := unsafe.Pointer(&data[OFFSET_LSN])
	*(*common.Lsn)(ptr) = lsn
}

func (page *Page) WLatch() {
	page.latch.WLock()
}

func (page *Page) WUnlatch() {
	page.latch.WUnlock()
}

func (page *Page) RLatch() {
	page.latch.RLock()
}

func (page *Page) RUnlatch() {
	page.latch.RUnlock()
}
