package qan

type Metric struct {
	Name        string
	Total       float32
	Avg         float32
	Max         float32
	Min         float32
	Pct_95      float32
	Stddev      float32
	Median      float32
	Query_count int
}
