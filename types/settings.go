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

