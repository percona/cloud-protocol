package qan

type ConfigQuery struct {
	Set    string `json:"Set,omitempty"`
	Verify string `json:"Verify,omitempty"`
	Expect string `json:"Expect,omitempty"`
}

type QanConfig struct {
	InstanceUUID      string `json:"InstanceUUID"`
	CollectFrom       string `json:"CollectFrom"`
	Interval          int    `json:"Interval"`
	ExampleQueries    bool   `json:"ExampleQueries"`
	MaxSlowLogSize    int    `json:"MaxSlowLogSize"`
	MaxWorkers        int    `json:"MaxWorkers"`
	RemoveOldSlowLogs bool   `json:"RemoveOldSlowLogs"`
	ReportLimit       int    `json:"ReportLimit"`
	WorkerRunTime     int    `json:"WorkerRunTime"`
	//
	Start []ConfigQuery `json:"Start"`
	Stop  []ConfigQuery `json:"Stop"`
}
