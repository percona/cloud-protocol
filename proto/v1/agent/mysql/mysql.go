package mysql

type Query struct {
	Set    string // SET GLOBAL long_query_time=0
	Verify string // SELECT @@long_query_time
	Expect string // 0
}
