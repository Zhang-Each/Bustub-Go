package catalog

type Schema struct {
	length         uint32
	columns        []Column
	isTupleInlined bool
	unlinedColumns []int32
}

// NewSchema New a schema object using several columns
func NewSchema(columns []Column) *Schema {
	var offset uint32 = 0
	isInline := true
	unInlined := make([]int32, 0)
	schemaColumns := make([]Column, 0)
	for i := 0; i < len(columns); i++ {
		var column = columns[i]
		if !column.IsInline() {
			isInline = false
			unInlined = append(unInlined, int32(i))
		}
		column.columnOffset = offset
		offset += column.GetFixedLength()
		schemaColumns = append(schemaColumns, column)
	}
	var length uint32 = offset
	return &Schema{
		length:         length,
		columns:        schemaColumns,
		isTupleInlined: isInline,
		unlinedColumns: unInlined,
	}
}

func (s Schema) GetColumns() *[]Column {
	return &s.columns
}

func (s Schema) GetColumn(idx uint32) *Column {
	return &s.columns[idx]
}

func (s Schema) GetColumnIndex(name string) uint32 {
	for i := 0; i < len(s.columns); i++ {
		if s.columns[i].columnName == name {
			return uint32(i)
		}
	}
	return uint32(len(s.columns))
}

func (s Schema) GetInlinedColumns() *[]int32 {
	return &s.unlinedColumns
}

func (s Schema) GetInlinedColumnCount() uint32 {
	return uint32(len(s.unlinedColumns))
}

func (s Schema) GetLength() uint32 {
	return s.length
}

func (s Schema) IsInlined() bool {
	return s.isTupleInlined
}
