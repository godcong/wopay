package wxpay

import (
	"encoding/json"
	"sort"
	"strings"
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

type RequestFunc func(url string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error)

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
	data, err := pay.DownloadBill(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Report(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.Report(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AuthCodeToOpenid(data PayData) (PayData, error) {
	pay := NewPay(PayConfigInstance())
	data, err := pay.AuthCodeToOpenid(data)
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

func (pay *Pay) SetSandBox(useSandbox bool) *Pay {
	pay.signType = SIGN_TYPE_HMACSHA256
	if useSandbox {
		pay.signType = SIGN_TYPE_MD5
	}
	pay.useSanBox = useSandbox
	return pay
}

func (pay *Pay) ApplySandBox(url string) string {
	if pay.useSanBox {
		return SANDBOX_URL_SUFFIX + url
	}
	return url
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

func (pay *Pay) fillRequest(requestFunc RequestFunc, data PayData, suffix string) (string, error) {
	return pay.fillRequestTimeout(requestFunc, suffix, data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) fillRequestTimeout(requestFunc RequestFunc, suffix string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error) {
	usb := pay.ApplySandBox(suffix)
	m, err := pay.FillRequestData(data)
	if err != nil {
		return "", err
	}
	return requestFunc(usb, m, connectTimeoutMs, readTimeoutMs)
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

// UnifiedOrder 统一下单
func (pay *Pay) UnifiedOrderTimeout(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
	return pay.unifiedOrderTimeout(data, connectTimeoutMs, readTimeoutMs)
}

// UnifiedOrder 统一下单
func (pay *Pay) UnifiedOrder(data PayData) (PayData, error) {
	return pay.unifiedOrderTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) unifiedOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {

	if pay.notifyUrl != "" {
		data.Set("notify_url", pay.notifyUrl)
	}
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, UNIFIEDORDER_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
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
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, CLOSEORDER_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
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
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, ORDERQUERY_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
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
	resp, err := pay.fillRequestTimeout(pay.RequestWithCertTimeout, REVERSE_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}

	return XmlToMap(resp), nil
}

/** Refund
* 作用：申请退款<br>
* 场景：刷卡支付、公共号支付、扫码支付、APP支付<br>
* 其他：需要证书
* @param data 向wxpay post的请求数据
* @return PayData, error API返回数据
 */
func (pay *Pay) Refund(data PayData) (PayData, error) {
	return pay.refundTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

/** RefundTimeout
* 作用：申请退款<br>
* 场景：刷卡支付、公共号支付、扫码支付、APP支付<br>
* 其他：需要证书
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error API返回数据
 */
func (pay *Pay) RefundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.refundTimeout(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) refundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	resp, err := pay.fillRequestTimeout(pay.RequestWithCertTimeout, REFUND_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}

	return XmlToMap(resp), nil
}

/** ShortUrl
* 作用：转换短链接
* 场景：刷卡支付、扫码支付
* @param data 向wxpay post的请求数据
* @return PayData, error API返回数据
 */
func (pay *Pay) ShortUrl(data PayData) (PayData, error) {
	return pay.shortUrl(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

/** ShortUrlTimeout
* 作用：转换短链接
* 场景：刷卡支付、扫码支付
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error API返回数据
 */
func (pay *Pay) ShortUrlTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.shortUrl(data, connectTimeoutMs, readTimeoutMs)
}

func (pay *Pay) shortUrl(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	resp, err := pay.fillRequestTimeout(pay.RequestWithCertTimeout, SHORTURL_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (pay *Pay) QueryRefund(data PayData) (PayData, error) {
	return pay.queryRefund(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

func (pay *Pay) QueryRefundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.queryRefund(data, connectTimeoutMs, readTimeoutMs)
}

func (pay *Pay) queryRefund(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, REFUNDQUERY_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}

	return XmlToMap(resp), nil
}

/** DownloadBill
* 作用：对账单下载
* 场景：刷卡支付、公共号支付、扫码支付、APP支付<br>
* 其他：无论是否成功都返回Map。若成功，返回的Map中含有return_code、return_msg、data，
*      其中return_code为`SUCCESS`，data为对账单数据。
* @param data 向wxpay post的请求数据
* @return PayData, error 经过封装的API返回数据
 */
func (pay *Pay) DownloadBill(data PayData) (PayData, error) {
	return pay.downloadBill(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

/** DownloadBillTimeout
* 作用：对账单下载
* 场景：刷卡支付、公共号支付、扫码支付、APP支付<br>
* 其他：无论是否成功都返回Map。若成功，返回的Map中含有return_code、return_msg、data，
*      其中return_code为`SUCCESS`，data为对账单数据。
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error 经过封装的API返回数据
 */
func (pay *Pay) DownloadBillTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.downloadBill(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) downloadBill(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, DOWNLOADBILL_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	var ret PayData
	if strings.Index(resp, "<") == 0 {
		ret = XmlToMap(resp)
	} else {
		ret = make(PayData)
		ret.Set("return_code", SUCCESS)
		ret.Set("return_msg", "ok")
		ret.Set("data", resp)
	}

	return ret, nil
}

/** Report
* 作用：交易保障<br>
* 场景：刷卡支付、公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @return PayData, error API返回数据
 */
func (pay *Pay) Report(data PayData) (PayData, error) {
	return pay.report(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

/** ReportTimeout
* 作用：交易保障<br>
* 场景：刷卡支付、公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error API返回数据
 */
func (pay *Pay) ReportTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.report(data, connectTimeoutMs, readTimeoutMs)
}
func (pay *Pay) report(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, REPORT_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

/** AuthCodeToOpenid
 * 作用: 授权码查询OPENID接口
 * 场景：刷卡支付
 * @param data 向wxpay post的请求数据
 * @return PayData, error API返回数据
 */
func (pay *Pay) AuthCodeToOpenid(data PayData) (PayData, error) {
	return pay.authCodeToOpenidTimeout(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
}

/** AuthCodeToOpenidTimeout
 * 作用: 授权码查询OPENID接口
 * 场景：刷卡支付
 * @param data 向wxpay post的请求数据
 * @param connectTimeoutMs 连接超时时间，单位是毫秒
 * @param readTimeoutMs 读超时时间，单位是毫秒
 * @return PayData, error API返回数据
 */
func (pay *Pay) AuthCodeToOpenidTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	return pay.authCodeToOpenidTimeout(data, connectTimeoutMs, readTimeoutMs)
}

func (pay *Pay) authCodeToOpenidTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, AUTHCODETOOPENID_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}