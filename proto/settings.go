package proto

type EmailReportsInstanceSetting struct {
	Hostname  string
	Frequency string // daily or weekly
	Active    bool
}

type EmailReportsSetting struct {
	Active bool
	Auto   bool
	// Instances []EmailReportInstanceSetting
}

type EmailSetting struct {
	OptOut bool
	News   bool
}
