package wx

type Pay struct {
	config     PayConfig
	signType   SignType
	autoReport bool
	useSanBox  bool
	notifyUrl  string
	payRequest PayRequest
}

type PayData map[string]string

// UnifiedOrder 统一下单
func (p *Pay) UnifiedOrder(reqData PayData, to ...int) error {
	conn := p.config.HttpConnectTimeoutMs
	read := p.config.HttpReadTimeoutMs
	if len(to) == 2 {
		conn = to[0]
		read = to[1]
	}
	return p.unifiedOrderTimeout(reqData, conn, read)
}

func (p *Pay) unifiedOrderTimeout(reqData PayData, connect int, read int) error {
	url := DOMAIN_API + UNIFIEDORDER_URL_SUFFIX
	if p.useSanBox {
		url = DOMAIN_API + SANDBOX_URL_SUFFIX + UNIFIEDORDER_URL_SUFFIX
	}

	if p.notifyUrl != "" {
		reqData.Set("notify_url", p.notifyUrl)
	}
	_, err := p.RequestWithoutCert(url)

	return err
}

func (p *Pay) RequestWithoutCert(url string) (string, error) {
	var resp string

	return resp, nil
}

func (p *Pay) FillRequestData(reqData PayData) {
	reqData.Set("appid", p.config.AppID)
	reqData.Set("mch_id", p.config.MchID)
	reqData.Set("nonce_str", GenerateUUID())
	reqData.Set("sign_type", p.signType.ToString())
	reqData.Set("sign", "") //TODO

}

func (data PayData) Set(key, val string) {
	data[key] = val
}

func (data PayData) Get(key string) string {
	return data[key]
}

func (data PayData) IsExist(key string) bool {
	_, b := data[key]
	return b
}
