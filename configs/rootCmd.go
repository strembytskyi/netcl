package configs

// RootConfig ...
type RootConfig struct {
	Name    string
	Short   string
	Silence bool
}

// GetRootConfig ...
func GetRootConfig() RootConfig {
	config := RootConfig{
		Name:    "netcli",
		Short:   "Let's your query IPs, CNAMEs, MX records and Name Servers",
		Silence: true,
	}

	return config
}
