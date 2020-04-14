package main

type config struct {
	AuthMethod string
	CertFile string
	KeyFile  string
	Username string
	Password string
	// MachineIDPath string
}

func defaultConfig() *config {
	return &config{
		AuthMethod: "BASIC",
		CertFile: "/etc/pki/consumer/cert.pem",
		KeyFile:  "/etc/pki/consumer/key.pem",
	}
}
/*
	CURRENT OPTIONS, CARRIED OVER

	AuthMethod: BASIC
	AutoConfig: true
	AutoUpdate: true
		turn into subcommand?
	BaseUrl: cloud.redhat.com/api
	BranchInfo: null
		if we get rid of phases we don't need to keep this as a config option
	BranchInfoUrl: null
		do we need this?
	CertVerify: True
		might need to adjust how we define/parse this option in golang
	CheckResults
		turn into subcommand
	CommandTimeout: 120
	CollectionRulesUrl
		will not need with core collection
	Compliance
		turn into subcommand
	Compressor: gz
	Conf: /etc/insights-client/insights-client.conf
	EggPath: null
		do we need this? maybe use this for egg canary service
	Debug
		bro I do not want to see or hear a phase ever again
	DisableSchedule
		turn into subcommand (schedule --disable?)
	DisplayName
		turn into subcommand maybe? or have a separate set-display-name subcommand
	EnableSchedule
		turn into subcommand (schedule --enable)
	Gpg: true
	EggGpgPath: null
		do we need this?
	Group
		deprecate maybe? we have tags now
	HttpTimeout: 120
	InsecureConnection: false
		I don't think we need this anymore?
	KeepArchive: false
	LoggingFile: /var/log/insights-client/insights.log
	LogLevel: debug
		need to adjust for golang translation
	NetDebug
	NoGpg: false
		alias for Gpg=False, get rid of it
	NoUpload: false
	Obfuscate: false
	ObfuscateHostname: false
	Offline: false
	OutputDir: null
	OutputFile: null
	Password: null
	Username: null
	Proxy: null
	Quiet: false
	Register
		turn into subcommand
	Silent: false
	RemoveFile: /etc/insights-client/remove.conf
		deprecated, get rid of it
	RedactionFile: /etc/insights-client/file-redaction.conf
	ContentRedactionFile: /etc/insights-client/file-content-redaction.conf
	ReRegister
		maybe change to "register --force"?
	Retries: 1
		should probably just change the default to 3 already
	ShowResults
		turn into subcommand
	Silent: false
	Status
		turn into subcommand
	Support
		do we still need this? if so, turn into subcommand
	SystemId
		we can probably parse this on demand and not keep it in the config
	TestConnection
		change into subcommand
	Unregister
		change into subcommand
	UploadUrl
		does this really need to be its own option?
	Username: null
	Validate
		turn into subcommand
	Verbose: false
	Version
		turn into subcommand
	LegacyUpload
		how much longer do we need to keep this around...
	Payload
		can probably refactor this to be something like an "upload --file" subcommand
	ContentType: null
		maybe change default to something generic? or advisor?
	Diagnosis
		turn into subcommand
	
	NEW OPTIONS

	MachineIDPath
		Path to the desired machine-id file. This will allow non-root users to upload if the file is outside /etc/insights-client
	ObfuscateIP
		independently choose to obfuscate IP addresses in soscleaner
	RetryTime
		configure how long to wait between upload retries (we might not want to make this configurable)
	Beta
		opt into bleeding edge enhancements? the "beta" egg/core channel?
*/