package proto

type EmailReportSetting struct {
	Hostname string
	Frequency string // daily or weekly
	Enabled bool
}

