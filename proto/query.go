package proto

type ExplainQuery struct {
	ServiceInstance
	Query string
}

type Explain struct {
	Result []ExplainRow
}

type ExplainRow struct {
	Id           NullInt64
	SelectType   NullString
	Table        NullString
	CreateTable  NullString // @todo
	Type         NullString
	PossibleKeys NullString // maybe map (split by comma)
	Key          NullString
	KeyLen       NullInt64
	Ref          NullString
	Rows         NullInt64
	Extra        NullString // maybe map (split by semicolon)
}
