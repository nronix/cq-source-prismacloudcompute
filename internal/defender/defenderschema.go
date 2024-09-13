package defender

import "time"

type Defenders struct {
	Hostname     string    `json:"hostname"`
	Version      string    `json:"version"`
	LastModified time.Time `json:"lastModified"`
	Type         string    `json:"type"`
	Category     string    `json:"category"`
	Connected    bool      `json:"connected"`
	Features     struct {
		ProxyListenerType string `json:"proxyListenerType"`
	} `json:"features"`
	Status struct {
		LastModified time.Time `json:"lastModified"`
		Filesystem   struct {
			Enabled bool `json:"enabled"`
		} `json:"filesystem"`
		Network struct {
			Enabled bool `json:"enabled"`
		} `json:"network"`
		Process struct {
			Enabled bool `json:"enabled"`
		} `json:"process"`
		AppFirewall struct {
			Enabled bool `json:"enabled"`
		} `json:"appFirewall"`
		OutOfBandAppFirewall struct {
			Enabled bool `json:"enabled"`
		} `json:"outOfBandAppFirewall"`
		ContainerNetworkFirewall struct {
		} `json:"containerNetworkFirewall"`
		HostNetworkFirewall struct {
			Enabled bool `json:"enabled"`
		} `json:"hostNetworkFirewall"`
		Features struct {
		} `json:"features"`
		HostCustomCompliance struct {
			Enabled bool `json:"enabled"`
		} `json:"hostCustomCompliance"`
	} `json:"status"`
	Fqdn          string   `json:"fqdn"`
	Collections   []string `json:"collections"`
	CloudMetadata struct {
		ResourceID string `json:"resourceID"`
		Provider   string `json:"provider"`
		Type       string `json:"type"`
		Region     string `json:"region"`
		AccountID  string `json:"accountID"`
		Labels     []struct {
			SourceType string    `json:"sourceType"`
			SourceName string    `json:"sourceName"`
			Timestamp  time.Time `json:"timestamp"`
			Key        string    `json:"key"`
			Value      string    `json:"value"`
		} `json:"labels"`
		AwsExecutionEnv string `json:"awsExecutionEnv"`
		Image           string `json:"image"`
	} `json:"cloudMetadata,omitempty"`
	TasBlobstoreScanner   bool      `json:"tasBlobstoreScanner"`
	CertificateExpiration time.Time `json:"certificateExpiration"`
	SystemInfo            struct {
		CPUCount         int     `json:"cpuCount"`
		KernelVersion    string  `json:"kernelVersion"`
		TotalDiskSpaceGB int     `json:"totalDiskSpaceGB"`
		FreeDiskSpaceGB  int     `json:"freeDiskSpaceGB"`
		MemoryGB         float64 `json:"memoryGB"`
	} `json:"systemInfo,omitempty"`
	RemoteMgmtSupported    bool `json:"remoteMgmtSupported"`
	RemoteLoggingSupported bool `json:"remoteLoggingSupported"`
	FirewallProtection     struct {
		Enabled       bool   `json:"enabled"`
		Supported     bool   `json:"supported"`
		OutOfBandMode string `json:"outOfBandMode"`
	} `json:"firewallProtection"`
	Proxy struct {
		HTTPProxy string `json:"httpProxy"`
		NoProxy   string `json:"noProxy"`
		Ca        string `json:"ca"`
		User      string `json:"user"`
		Password  struct {
			Encrypted string `json:"encrypted"`
		} `json:"password"`
	} `json:"proxy"`
	Port              int  `json:"port"`
	CompatibleVersion bool `json:"compatibleVersion"`
	UsingOldCA        bool `json:"usingOldCA"`
	IsARM64           bool `json:"isARM64"`
	CloudMetadata0    struct {
		ResourceID string `json:"resourceID"`
		Provider   string `json:"provider"`
		Name       string `json:"name"`
		Region     string `json:"region"`
		AccountID  string `json:"accountID"`
		Labels     []struct {
			SourceType string    `json:"sourceType"`
			SourceName string    `json:"sourceName"`
			Timestamp  time.Time `json:"timestamp"`
			Key        string    `json:"key"`
			Value      string    `json:"value,omitempty"`
		} `json:"labels"`
		VMID      string `json:"vmID"`
		Image     string `json:"image"`
		VMImageID string `json:"vmImageID"`
	} `json:"cloudMetadata,omitempty"`
	SystemInfo0 struct {
		CPUCount         int     `json:"cpuCount"`
		KernelVersion    string  `json:"kernelVersion"`
		TotalDiskSpaceGB int     `json:"totalDiskSpaceGB"`
		MemoryGB         float64 `json:"memoryGB"`
	} `json:"systemInfo,omitempty"`
}

type VmsDefender struct {
	ID           string    `json:"_id"`
	AccountID    string    `json:"accountID"`
	Architecture string    `json:"architecture"`
	Arn          string    `json:"arn"`
	AwsSubnetID  string    `json:"awsSubnetID"`
	AwsVPCID     string    `json:"awsVPCID"`
	Cluster      string    `json:"cluster"`
	Collections  []string  `json:"collections"`
	CreatedAt    time.Time `json:"createdAt"`
	Fqdn         string    `json:"fqdn"`
	HasDefender  bool      `json:"hasDefender"`
	Hostname     string    `json:"hostname"`
	ImageID      string    `json:"imageID"`
	ImageName    string    `json:"imageName"`
	Name         string    `json:"name"`
	Os           string    `json:"os"`
	OsInfo       struct {
		Distro                  string `json:"distro"`
		DistroRelease           string `json:"distroRelease"`
		FullName                string `json:"fullName"`
		UnderlyingDistro        string `json:"underlyingDistro"`
		UnderlyingDistroRelease string `json:"underlyingDistroRelease"`
		Version                 string `json:"version"`
	} `json:"osInfo"`
	Provider string `json:"provider"`
	Region   string `json:"region"`
	Tags     []struct {
		Key        string    `json:"key"`
		SourceName string    `json:"sourceName"`
		SourceType string    `json:"sourceType"`
		Timestamp  time.Time `json:"timestamp"`
		Value      string    `json:"value"`
	} `json:"tags"`
	Timestamp time.Time `json:"timestamp"`
}
