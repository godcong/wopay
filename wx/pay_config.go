package wx

type PayConfig struct {
	AppID                string
	MchID                string
	Key                  string
	HttpConnectTimeoutMs int
	HttpReadTimeoutMs    int
	AutoReport           bool
	ReportWorkerNum      int
	ReportQueueMaxSize   int
	ReportBatchSize      int
	cert                 []byte
}

func NewPayConfig() PayConfig {
	return PayConfig{
		HttpConnectTimeoutMs: 6000,
		HttpReadTimeoutMs:    8000,
		AutoReport:           true,
		ReportWorkerNum:      6,
		ReportQueueMaxSize:   10000,
		ReportBatchSize:      10,
	}
}
