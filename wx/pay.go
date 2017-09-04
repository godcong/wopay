package wx

import "log"

type Pay struct {
	config     *PayConfig
	signType   SignType
	autoReport bool
	useSanBox  bool
	notifyUrl  string
	payRequest PayRequest
}

type PayData map[string]string

func UnifiedOrder(reqData PayData) (PayData, error) {
	pay := NewPay(PayConfigImpl())
	data, err := pay.UnifiedOrder(reqData)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewPay(config *PayConfig) *Pay {
	return &Pay{
		config: config,
	}
}

// UnifiedOrder 统一下单
func (p *Pay) UnifiedOrderTimeout(reqData PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
	return p.unifiedOrderTimeout(reqData, connectTimeoutMs, readTimeoutMs)
}

// UnifiedOrder 统一下单
func (p *Pay) UnifiedOrder(reqData PayData) (PayData, error) {
	return p.unifiedOrderTimeout(reqData, p.config.ConnectTimeoutMs, p.config.ReadTimeoutMs)
}

func (p *Pay) unifiedOrderTimeout(reqData PayData, connect int, read int) (PayData, error) {
	url := DOMAIN_API + UNIFIEDORDER_URL_SUFFIX
	if p.useSanBox {
		url = DOMAIN_API + SANDBOX_URL_SUFFIX + UNIFIEDORDER_URL_SUFFIX
	}

	if p.notifyUrl != "" {
		reqData.Set("notify_url", p.notifyUrl)
	}
	m, err := p.FillRequestData(reqData)
	if err != nil {
		return nil, err
	}
	resp, err := p.RequestWithoutCert(url, m)
	log.Println(resp)

	return xmlToMap(resp), err
}

func (p *Pay) RequestWithoutCert(url string, reqData PayData) (string, error) {
	//msgUUID := reqData.Get("nonce_str")
	reqBody, err := MapToXml(reqData)
	if err != nil {
		return "", err
	}
	log.Println(reqBody)
	//resp, err := p.payRequest.RequestWithoutCert(url, msgUUID, reqBody, p.config.ConnectTimeoutMs, p.config.ReadTimeoutMs, p.autoReport)
	resp := ""
	return resp, err
}

func (p *Pay) FillRequestData(reqData PayData) (PayData, error) {
	reqData.Set("appid", p.config.AppID)
	reqData.Set("mch_id", p.config.MchID)
	reqData.Set("nonce_str", GenerateUUID())
	reqData.Set("sign_type", p.signType.ToString())
	sign, e := GenerateSignature(reqData, p.config.Key, p.signType)
	if e != nil {
		return nil, e
	}
	reqData.Set("sign", sign)
	return reqData, nil
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
