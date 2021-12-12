package common

type RID struct {
	pageId  PageId
	slotNum uint32
}

func NewRID(pageId PageId, slotNum uint32) *RID {
	return &RID{
		pageId: pageId,
		slotNum: slotNum,
	}
}

func NewRIDByCopy(rid int64) *RID {
	return &RID{
		pageId: PageId(rid >> 32),
		slotNum: uint32(rid),
	}
}

func (rid *RID) Get() int64 {
	return int64(rid.pageId << 32) | int64(rid.slotNum)
}

func (rid *RID) GetPageID() PageId {
	return rid.pageId
}

func (rid *RID) GetSlotNum() uint32 {
	return rid.slotNum
}

func (rid *RID) Set(pageId PageId, slotNum uint32)  {
	rid.pageId = pageId
	rid.slotNum = slotNum
}

func (rid *RID) Equals(other *RID) bool {
	return rid.pageId == other.pageId && rid.slotNum == other.slotNum
}

