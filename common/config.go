package common

type FrameId = int32
type PageId = int32
type RxnId = int32
type Lsn = int32
type SlotOffset = int
type Oid = uint16

const (
	INVALID_PAGE_ID  int = -1
	INVALID_TXN_ID   int = -1
	INVALID_LSN      int = -1
	HEADER_PAGE_ID   int = 0
	PAGE_SIZE        int = 4096
	BUFFER_POOL_SIZE int = 10
	LOG_BUFFER_SIZE  int = ((BUFFER_POOL_SIZE + 1) * PAGE_SIZE)
	BUCKET_SIZE      int = 50
)
