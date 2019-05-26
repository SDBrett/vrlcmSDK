package types

type NetworkStatus struct {
	Hostname           string `json:"hostName"`
	DiskSize           string `json:"diskSize"`
	DiskUsedPercentage string `json:"diskUsedPercentage"`
	Netmask            string `json:"netmask"`
	DiskUsed           string `json:"diskUsed"`
	DiskAvail          string `json:"diskAvail"`
	IPv4Address        string `json:"ipv4Address"`
	PreferredDns       string `json:"preferredDns"`
	Type               string `json:"type"`
	Gateway            string `json:"gateway"`
}
