package proto

type ExplainQuery struct {
	UUID  string
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
	CreateTable  NullString // @todo
	Type         NullString
	PossibleKeys NullString // split by comma
	Key          NullString
	KeyLen       NullString // https://jira.percona.com/browse/PCT-863
	Ref          NullString
	Rows         NullInt64
	Extra        NullString // split by semicolon
}
