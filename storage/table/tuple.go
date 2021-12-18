package table

import (
	"bustub/catalog"
	"bustub/common"
	"bustub/valuetype"
)

type Tuple struct {
	allocated bool
	rid       common.RID
	size      uint32
	data      []byte
}

func NewTuple(rid common.RID) *Tuple {
	return &Tuple{
		allocated: false,
		rid: rid,
		size: 0,
		data: make([]byte, 0),
	}
}

func NewTuple2(tuple *Tuple) *Tuple {
	return &Tuple{
		allocated: tuple.allocated,
		rid: tuple.rid,
		size: tuple.size,
		data: tuple.data,
	}
}

func (t *Tuple) SerializeTo(storage []byte) {

}

func (t *Tuple) DeserializeFrom(storage []byte)  {

}

func (t *Tuple) GetRID() common.RID {
	return t.rid
}

func (t *Tuple) GetData() []byte {
	return t.data
}

func (t *Tuple) GetSize() uint32 {
	return t.size
}

func (t *Tuple) GetValue(schema *catalog.Schema, columnIndex uint32) valuetype.Value {
	panic("implement me!")
}

func (t *Tuple) KeyFromTuple(schema *catalog.Schema, keySchema *catalog.Schema, keyAttrs []uint32) Tuple {
	panic("implement me!")
}

func (t *Tuple) IsNull(schema *catalog.Schema, columnIndex uint32) bool {
	return false
}

func (t *Tuple) IsAllocated() bool {
	return t.allocated
}

func (t *Tuple) GetDataPtr(schema *catalog.Schema, columnIndex uint32) []byte {
	panic("implement me!")
}

