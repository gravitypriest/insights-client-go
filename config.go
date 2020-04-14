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
