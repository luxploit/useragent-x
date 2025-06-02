package uax

type UserAgent struct {
	RawString        string
	Name             string
	AgentPlatform    UserAgentAgentPlatform
	HostPlatform     UserAgentHostPlatform
	QurikBits        Quriks
	CrawlerURL       string
	ActualType       PlatformType
	ImpersonatedType PlatformType
	WasImpersonated  bool
}

type UserAgentAgentPlatform struct {
	BrowserEngine      BrowserEngine
	PlatformVersion    SemVer
	RawPlatformVersion string
	IsPlatformBeta     bool
}

type UserAgentHostPlatform struct {
	OS          PlatformOS
	Version     SemVer
	RawVersion  string
	ProductLine string // eg. NT or 9x
	Device      string
	HostISA     string // Host Instruction Set Architecture
}
