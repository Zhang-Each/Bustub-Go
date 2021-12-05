package valuetype

type TypeId uint

const (
	INVALID   TypeId = 0
	BOOLEAN   TypeId = 1
	TINYINT   TypeId = 2
	SMALLINT  TypeId = 3
	INTEGER   TypeId = 4
	BIGINT    TypeId = 5
	DECIMAL   TypeId = 6
	VARCHAR   TypeId = 7
	TIMESTAMP TypeId = 8
)
