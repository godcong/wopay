package wx

type PayConfig struct {
	AppID              string
	MchID              string
	Key                string
	ConnectTimeoutMs   int
	ReadTimeoutMs      int
	AutoReport         bool
	ReportWorkerNum    int
	ReportQueueMaxSize int
	ReportBatchSize    int
	payDomain          PayDomain
	cert               []byte
}

var config *PayConfig

func init() {
	PayConfigImpl()
}

func PayConfigImpl() *PayConfig {
	if config == nil {
		config = NewPayConfig()
	}
	return config
}

func NewPayConfig() *PayConfig {
	return &PayConfig{
		ConnectTimeoutMs:   6000,
		ReadTimeoutMs:      8000,
		AutoReport:         true,
		ReportWorkerNum:    6,
		ReportQueueMaxSize: 10000,
		ReportBatchSize:    10,
	}
}

func (config *PayConfig) PayDomain() PayDomain {
	return config.payDomain
}
