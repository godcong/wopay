package wxpay

import (
	"encoding/json"
	"sort"
)

type Pay struct {
	config     PayConfig
	payRequest *PayRequest
	signType   SignType
	autoReport bool
	useSanBox  bool
	notifyUrl  string
}

type PayData map[string]string

//UnifiedOrder
func UnifiedOrder(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.UnifiedOrder(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//CloseOrder
func CloseOrder(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.CloseOrder(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//QueryOrder
func QueryOrder(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.QueryOrder(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//ReverseOrder
func ReverseOrder(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.ReverseOrder(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//QueryRefund
func QueryRefund(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.QueryRefund(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Refund
func Refund(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.Refund(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ShortUrl(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.ShortUrl(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DownloadBill(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.ShortUrl(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewPay(config PayConfig) *Pay {
	return newPay(config, "", true, false)
}

func newPay(config PayConfig, notifyUrl string, autoReport bool, useSandbox bool) *Pay {
	pay := Pay{
		config:     config,
		notifyUrl:  notifyUrl,
		autoReport: autoReport,
		useSanBox:  useSandbox,
	}
	pay.signType = SIGN_TYPE_HMACSHA256
	if useSandbox {
		pay.signType = SIGN_TYPE_MD5
	}
	pay.payRequest = NewPayRequest(config)
	return &pay
}

func (pay *Pay) UseSandBox(url string) string {
	if pay.useSanBox {
		return SANDBOX_URL_SUFFIX + url
	}
	return url
}

// UnifiedOrder 统一下单
func (pay *Pay) UnifiedOrderTimeout(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
	return pay.unifiedOrderTimeout(data, connectTimeoutMs, readTimeoutMs)
}

// UnifiedOrder 统一下单
func (pay *Pay) UnifiedOrder(data PayData) (PayData, error) {
	return pay.unifiedOrderTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) unifiedOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	url := pay.UseSandBox(UNIFIEDORDER_URL_SUFFIX)

	if pay.notifyUrl != "" {
		data.Set("notify_url", pay.notifyUrl)
	}
	m, err := pay.FillRequestData(data)
	if err != nil {
		return nil, err
	}
	resp, err := pay.RequestWithoutCertTimeout(url, m, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (pay *Pay) CloseOrder(data PayData) (PayData, error) {
	return pay.closeOrderTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) CloseOrderTimeout(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
	return pay.closeOrderTimeout(data, connectTimeoutMs, readTimeoutMs)
}

func (pay *Pay) closeOrderTimeout(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
	url := pay.UseSandBox(CLOSEORDER_URL_SUFFIX)
	m, err := pay.FillRequestData(data)
	if err != nil {
		return nil, err
	}
	resp, err := pay.RequestWithoutCertTimeout(url, m, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (pay *Pay) QueryOrder(data PayData) (PayData, error) {
	return pay.queryOrderTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) QueryOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.queryOrderTimeout(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) queryOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	url := pay.UseSandBox(ORDERQUERY_URL_SUFFIX)
	m, err := pay.FillRequestData(data)
	if err != nil {
		return nil, err
	}
	resp, err := pay.RequestWithoutCertTimeout(url, m, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (pay *Pay) ReverseOrder(data PayData) (PayData, error) {
	return pay.reverseOrderTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) ReverseOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.reverseOrderTimeout(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) reverseOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	url := pay.UseSandBox(REVERSE_URL_SUFFIX)
	m, err := pay.FillRequestData(data)
	if err != nil {
		return nil, err
	}
	resp, err := pay.RequestWithCertTimeout(url, m, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}

	return XmlToMap(resp), nil
}

func (pay *Pay) Refund(data PayData) (PayData, error) {
	return pay.refundTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}
func (pay *Pay) RefundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.refundTimeout(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) refundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	url := pay.UseSandBox(REFUND_URL_SUFFIX)
	m, err := pay.FillRequestData(data)
	if err != nil {
		return nil, err
	}
	resp, err := pay.RequestWithCertTimeout(url, m, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}

	return XmlToMap(resp), nil
}

func (pay *Pay) ShortUrl(data PayData) (PayData, error) {
	return pay.shortUrl(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) ShortUrlTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.shortUrl(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) shortUrl(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	url := pay.UseSandBox(SHORTURL_URL_SUFFIX)
	m, err := pay.FillRequestData(data)
	if err != nil {
		return nil, err
	}
	resp, err := pay.RequestWithCertTimeout(url, m, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (pay *Pay) QueryRefund(data PayData) (PayData, error) {
	//TODO
	return nil, nil
}

func (pay *Pay) RequestWithoutCert(url string, data PayData) (string, error) {
	msgUUID := data.Get("nonce_str")
	reqBody, err := MapToXml(data)
	if err != nil {
		return "", err
	}
	resp, err := pay.payRequest.RequestWithoutCert(url, msgUUID, reqBody, pay.autoReport)
	return resp, err
}

func (pay *Pay) RequestWithoutCertTimeout(url string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error) {
	msgUUID := data.Get("nonce_str")
	reqBody, err := MapToXml(data)
	if err != nil {
		return "", err
	}
	resp, err := pay.payRequest.RequestWithoutCertTimeout(url, msgUUID, reqBody, connectTimeoutMs, readTimeoutMs, pay.autoReport)
	return resp, err
}

func (pay *Pay) RequestWithCert(url string, data PayData) (string, error) {
	msgUUID := data.Get("nonce_str")
	reqBody, err := MapToXml(data)
	if err != nil {
		return "", err
	}
	resp, err := pay.payRequest.RequestWithCert(url, msgUUID, reqBody, pay.autoReport)
	return resp, err
}

func (pay *Pay) RequestWithCertTimeout(url string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error) {
	msgUUID := data.Get("nonce_str")
	reqBody, err := MapToXml(data)
	if err != nil {
		return "", err
	}
	resp, err := pay.payRequest.RequestWithCertTimeout(url, msgUUID, reqBody, connectTimeoutMs, readTimeoutMs, pay.autoReport)
	return resp, err
}

func (pay *Pay) FillRequestData(data PayData) (PayData, error) {
	data.Set("appid", pay.config.AppID())
	data.Set("mch_id", pay.config.MchID())
	data.Set("nonce_str", GenerateUUID())
	data.Set("sign_type", pay.signType.ToString())
	sign, e := GenerateSignature(data, pay.config.Key(), pay.signType)
	if e != nil {
		return nil, e
	}
	data.Set("sign", sign)
	return data, nil
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

func (data PayData) SortKeys() []string {
	var keys sort.StringSlice
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

func (data PayData) ToJson() string {
	b, e := json.Marshal(&data)
	if e != nil {
		return ""
	}
	return string(b)
}
