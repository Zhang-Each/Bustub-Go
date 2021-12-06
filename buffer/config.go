package buffer

type frame_id_t = int32
type page_id_t = int32
type rxn_id_t = int32
type lsn_t = int32
type slot_offset_t = int
type oid_t = uint16

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

