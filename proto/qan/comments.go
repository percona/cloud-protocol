package qan

type Comment struct {
	Id      uint
	User    User
	Created int64
	Edited  int64
	Comment string
}

type User struct {
	Id    uint
	Email string
	Name  string
}
