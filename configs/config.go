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

// NSConfig ...
type NSConfig struct {
	Name  string
	Short string
}

// GetNSConfig ...
func GetNSConfig() NSConfig {
	config := NSConfig{
		Name:  "ns",
		Short: "Looks up the Name Servers for a Particular Host",
	}
	return config
}

// IPConfig ...
type IPConfig struct {
	Name  string
	Short string
}

// GetIPConfig ...
func GetIPConfig() IPConfig {
	config := IPConfig{
		Name:  "ip",
		Short: "Looks up the IP addresses for a particular host",
	}
	return config
}

// CnameConfig ...
type CnameConfig struct {
	Name  string
	Short string
}

// GetCnameConfig ...
func GetCnameConfig() CnameConfig {
	config := CnameConfig{
		Name:  "cname",
		Short: "Looks up he CNAME for a particular host",
	}

	return config
}
