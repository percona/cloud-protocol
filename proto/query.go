package proto

import (
	"time"
)

type ExplainQuery struct {
	ServiceInstance
	Db    string
	Query string
}

type ExplainResult struct {
	Classic []*ExplainRow
	JSON    string // since MySQL 5.6.5
}

type ExplainRow struct {
	Id           NullInt64
	SelectType   NullString
	Table        NullString
	Partitions   NullString // split by comma; since MySQL 5.1
	Type         NullString
	PossibleKeys NullString // split by comma
	Key          NullString
	KeyLen       NullString // https://jira.percona.com/browse/PCT-863
	Ref          NullString
	Rows         NullInt64
	Extra        NullString // split by semicolon
}

type Table struct {
	Db    string
	Table string
}

type TableInfoQuery struct {
	ServiceInstance
	Create []Table // SHOW CREATE TABLE Db.Table
	Index  []Table // SHOW INDEXES FROM Db.Table
	Status []Table // SHOW TABLE STATUS FROM Db LIKE 'Table'
}

type ShowIndexRow struct {
	Table        string
	NonUnique    bool
	KeyName      string
	SeqInIndex   int
	ColumnName   string
	Collation    NullString
	Cardinality  NullInt64
	SubPart      NullInt64
	Packed       NullString
	Null         NullString
	IndexType    string
	Comment      NullString
	IndexComment NullString
}

type ShowTableStatus struct {
	Name          string
	Engine        string
	Version       string
	RowFormat     string
	Rows          NullInt64
	AvgRowLength  NullInt64
	DataLength    NullInt64
	MaxDataLength NullInt64
	IndexLength   NullInt64
	DataFree      NullInt64
	AutoIncrement NullInt64
	CreateTime    time.Time
	UpdateTime    time.Time
	CheckTime     NullString // todo: need NullTime
	Collation     NullString
	Checksum      NullString
	CreateOptions NullString
	Comment       NullString
}

type TableInfo struct {
	Create string          `json:",omitempty"`
	Index  []ShowIndexRow  `json:",omitempty"`
	Status ShowTableStatus `json:",omitempty"`
	Errors []string        `json:",omitempty"`
}

type TableInfoResult map[string]*TableInfo
