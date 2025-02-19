package conf

var Config *Configuration

type Configuration struct {
	Interval        int64            `yaml:"interval"`
	IpSniffer       IpSniffer        `yaml:"ipSniffer"`
	DomainRegisters []DomainRegister `yaml:"domainRegisters"`
}

type IpSniffer struct {
	Addr string `yaml:"addr"`
}

type DomainRegister struct {
	RegisterName   string `yaml:"registerName"`
	Key            string `yaml:"key"`
	Secret         string `yaml:"secret"`
	TopLevelDomain string `yaml:"topLevelDomain"`
}
