package types

type ConfigSettings struct {
	SshEnabled       string `json:"sshEnabled, omitempty"`
	TelemetryEnabled string `json:"telemetryEnabled, omitempty"`
}

type ActiveDirectorySettings struct {
	AdName                 string `json:"adName"`
	BaseDN                 string `json:"baseDN"`
	BindDN                 string `json:"bindDN"`
	GroupDN                string `json:"groupDN"`
	UserDN                 string `json:"userDN"`
	BindPassword           string `json:"bindPassword"`
	UberAdmin              string `json:"uberAdmin"`
	SyncNestedGroupMembers bool   `json:"syncNestedGroupMembers"`
}

type VidmSettings struct {
	Hostname      string `json:"hostName"`
	AdminUsername string `json:"adminUserName"`
	AdminPassword string `json:"adminPassword"`
}

type MyVMwareSettings struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SshPassword struct {
	SshUserPassword string `json:"sshuserPassword"`
}

type RootPassword struct {
	RootPassword string `json:"rootPassword"`
}

type ConfigDriftInterval struct {
	IntervalMinutes int `json:"intervalMinutes"`
}

// LCM appliance restart schedule
type RestartSchedule struct {
	WeeklyServerRestartEnable bool   `json:"weeklyServerRestartEnable"`
	Day                       string `json:"day"`  // 1 = Monday, 2 = Tuesday, 3 = Wednesday, 4 = Thursday, 5 = Friday, 6 = Saturday, 7 = Sunday
	Hour                      string `json:"hour"` // Hour for appliance to restart. 04 = 4 AM, 16 = 4 PM
}
