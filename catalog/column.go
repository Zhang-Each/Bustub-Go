package catalog

import (
	"bustub/execution/expressions"
	"bustub/valuetype"
)

type Column struct {
	// properties of a column in the database table
	columnName     string
	columnType     valuetype.TypeId
	fixedLength    uint32
	variableLength uint32
	columnOffset   uint32
	absExp         *expressions.AbstractExpression
}

func NewColumn(columnName string, columnType valuetype.TypeId, absExp *expressions.AbstractExpression) *Column {
	return &Column{
		columnName:  columnName,
		columnType:  columnType,
		fixedLength: uint32(TypeSize(columnType)),
		absExp:      absExp,
	}
}

func (c Column) GetName() string {
	return c.columnName
}

func (c Column) GetType() valuetype.TypeId {
	return c.columnType
}

func (c Column) GetFixedLength() uint32 {
	return c.fixedLength
}

func (c Column) GetVariableLength() uint32 {
	return c.variableLength
}

func (c Column) GetColumnOffset() uint32 {
	return c.columnOffset
}

func (c *Column) GetAbstractExpression() expressions.AbstractExpression {
	return c.absExp
}

func (c Column) IsInline() bool {
	return c.columnType != valuetype.VARCHAR
}

func (c Column) GetLength() uint32 {
	if c.IsInline() {
		return c.GetFixedLength()
	}
	return c.GetVariableLength()
}

func TypeSize(type_id valuetype.TypeId) uint8 {
	switch type_id {
	case valuetype.BOOLEAN:
		return 1
	case valuetype.TINYINT:
		return 1
	case valuetype.SMALLINT:
		return 2
	case valuetype.INTEGER:
		return 4
	case valuetype.BIGINT:
	case valuetype.DECIMAL:
	case valuetype.TIMESTAMP:
		return 8
	case valuetype.VARCHAR:
		return 12
	default:
		return 0
	}
	return 0
}
