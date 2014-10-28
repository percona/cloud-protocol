package proto

type EmailReportsInstanceSetting struct {
	InstanceId uint `json:"-"`
	Hostname   string
	Frequency  string // daily or weekly
	Day        string `json:"-"`
	Slot       string `json:"-"`
	Active     bool
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
