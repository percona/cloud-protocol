package qan

type QueryProfileTag struct {
	LabelID int64  `json:"label_id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
}

type QueryProfileClass struct {
	ClassID         int64 `json:"class_id"`
	Distillate      string
	Checksum        string
	Fingerprint     string
	TotalMetric     float64 `json:"total_metric"`
	QueryCount      int64   `json:"query_count"`
	ShowCreateTable string  `json:"show_c_t"`
	ShowTableStatus string  `json:"show_t_s"`
	PctOverTotal    int64   `json:"pct_over_total"`
	CountOverTotal  int64   `json:"count_over_total"`
	Qps             float64
	Load            float64
	LoadPct         float64 `json:"load_pct"`
	Avg             float64
	Pct95           float64 `json:"pct_95"`
	Max             float64
	Comments        int64
	FirstSeen       int64 `json:"first_seen"`
	LastSeen        int64 `json:"last_seen"`
	Tags            []QueryProfileTag
	Status          string
}

type QueryProfileServer struct {
	QueryCount int64   `json:"querycount"`
	TotalSum   float64 `json:"totalsum"`
	RowCount   int64   `json:"rowcount"`
	Qps        float64
	Load       float64
	LoadPct    int64   `json:"loadpct"`
	TotalAvg   float64 `json:"totalavg"`
	TotalPct95 float64 `json:"totalpct95"`
	TotalMax   float64 `json:"totalmax"`
	RateLimit  float64 `json:"ratelimit"`
}

type QueryProfile struct {
	Server  QueryProfileServer
	Classes []QueryProfileClass
}
