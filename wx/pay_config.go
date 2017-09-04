package wx

type PayConfigInstance struct {
	appID              string
	mchID              string
	key                string
	connectTimeoutMs   int
	readTimeoutMs      int
	autoReport         bool
	reportWorkerNum    int
	reportQueueMaxSize int
	reportBatchSize    int
	payDomain          PayDomain
	cert               []byte
}

type PayConfig interface {
	AppID() string
	MchID() string
	Key() string
	Cert() []byte
	ConnectTimeoutMs() int
	ReadTimeoutMs() int
	PayDomainInstance() PayDomain
	AutoReport() bool
	ReportWorkNum() int
	ReportQueueMaxSize() int
	ReportBatchSize() int
}

var config PayConfig

func init() {
	PayConfigImpl()
}

func PayConfigImpl() PayConfig {
	if config == nil {
		config = NewPayConfig()
	}
	return config
}

func NewPayConfig() PayConfig {
	return &PayConfigInstance{
		connectTimeoutMs:   6000,
		readTimeoutMs:      8000,
		autoReport:         true,
		reportWorkerNum:    6,
		reportQueueMaxSize: 10000,
		reportBatchSize:    10,
	}
}

func (instance *PayConfigInstance) AppID() string {
	return "wx426b3015555a46be"
}
func (instance *PayConfigInstance) MchID() string {
	return "1225312702"
}
func (instance *PayConfigInstance) Key() string {
	return "e10adc3949ba59abbe56e057f20f883e"
}
func (instance *PayConfigInstance) Cert() []byte {
	//TODO
	return []byte("")
}
func (instance *PayConfigInstance) ConnectTimeoutMs() int {
	return instance.connectTimeoutMs
}
func (instance *PayConfigInstance) ReadTimeoutMs() int {
	return instance.readTimeoutMs
}
func (instance *PayConfigInstance) PayDomainInstance() PayDomain {
	return instance.payDomain
}
func (instance *PayConfigInstance) AutoReport() bool {
	return instance.autoReport
}
func (instance *PayConfigInstance) ReportWorkNum() int {
	return instance.reportWorkerNum
}
func (instance *PayConfigInstance) ReportQueueMaxSize() int {
	return instance.reportQueueMaxSize
}
func (instance *PayConfigInstance) ReportBatchSize() int {
	return instance.reportBatchSize
}
