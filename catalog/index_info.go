package catalog

import "bustub/storage/index"

type IndexInfo struct {
	schema    Schema
	tableName string
	indexName string
	index     index.Index
	oid       index_oid_t
	keySize   uint
}

func NewIndexInfo(schema Schema, tableName string, indexName string, index index.Index, oid index_oid_t, keySize uint) *IndexInfo {
	return &IndexInfo{
		schema: schema,
		tableName: tableName,
		indexName: indexName,
		index: index,
		oid: oid,
		keySize: keySize,
	}
}
