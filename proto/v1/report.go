package proto

type MySQLInstancePerf struct {
	AgentUuid  string  // agent running QAN for this MySQL instance
	Id         uint    // MySQL instance
	Hostname   string  // MySQL instance
	Alias      string  // MySQL instance
	LastReport int64   // last time data was received, UTC
	ClockTime  int64   // seconds, when there was data
	QueryTime  float64 // seconds, SUM(Query_time_sum)
	QueryCount int64   // total, not unique
	Avg95th    float64 // AVG(Query_time_pct_95)
	QPS        float64 // QueryCount / ClockTime
	Load       float64 // QueryTime / ClockTime
	// Remove after migrating to PCT v3
	UUID string // v3 instance ID
}
