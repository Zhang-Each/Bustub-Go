package catalog

import "bustub/storage/table"

type TableInfo struct {
	schema    Schema // Table Schema
	name      string // Table Name
	tableHeap *table.TableHeap
	oid       table_oid_t
}

func NewTableInfo(schema Schema, name string, tableHeap *table.TableHeap, oid table_oid_t) *TableInfo {
	return &TableInfo{
		schema: schema,
		name: name,
		tableHeap: tableHeap,
		oid: oid,
	}
}
