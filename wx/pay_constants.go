package wx

type SignType int

const (
	SIGN_TYPE_MD5        SignType = iota
	SIGN_TYPE_HMACSHA256 SignType = iota
)

func (t SignType) ToString() string {
	if t == SIGN_TYPE_HMACSHA256 {
		return HMACSHA256
	}
	return MD5

}

const DOMAIN_API = "api.mch.weixin.qq.com"
const DOMAIN_API2 = "api2.mch.weixin.qq.com"
const DOMAIN_APIHK = "apihk.mch.weixin.qq.com"
const DOMAIN_APIUS = "apius.mch.weixin.qq.com"

const FAIL = "FAIL"
const SUCCESS = "SUCCESS"
const HMACSHA256 = "HMAC-SHA256"
const MD5 = "MD5"

const FIELD_SIGN = "sign"
const FIELD_SIGN_TYPE = "sign_type"

const MICROPAY_URL_SUFFIX = "/pay/micropay"
const UNIFIEDORDER_URL_SUFFIX = "/pay/unifiedorder"
const ORDERQUERY_URL_SUFFIX = "/pay/orderquery"
const REVERSE_URL_SUFFIX = "/secapi/pay/reverse"
const CLOSEORDER_URL_SUFFIX = "/pay/closeorder"
const REFUND_URL_SUFFIX = "/secapi/pay/refund"
const REFUNDQUERY_URL_SUFFIX = "/pay/refundquery"
const DOWNLOADBILL_URL_SUFFIX = "/pay/downloadbill"
const REPORT_URL_SUFFIX = "/payitil/report"
const SHORTURL_URL_SUFFIX = "/tools/shorturl"
const AUTHCODETOOPENID_URL_SUFFIX = "/tools/authcodetoopenid"

const SANDBOX_URL_SUFFIX = "/sandboxnew"

//const SANDBOX_MICROPAY_URL_SUFFIX = "/sandboxnew/pay/micropay"
//const SANDBOX_UNIFIEDORDER_URL_SUFFIX = "/sandboxnew/pay/unifiedorder"
//const SANDBOX_ORDERQUERY_URL_SUFFIX = "/sandboxnew/pay/orderquery"
//const SANDBOX_REVERSE_URL_SUFFIX = "/sandboxnew/secapi/pay/reverse"
//const SANDBOX_CLOSEORDER_URL_SUFFIX = "/sandboxnew/pay/closeorder"
//const SANDBOX_REFUND_URL_SUFFIX = "/sandboxnew/secapi/pay/refund"
//const SANDBOX_REFUNDQUERY_URL_SUFFIX = "/sandboxnew/pay/refundquery"
//const SANDBOX_DOWNLOADBILL_URL_SUFFIX = "/sandboxnew/pay/downloadbill"
//const SANDBOX_REPORT_URL_SUFFIX = "/sandboxnew/payitil/report"
//const SANDBOX_SHORTURL_URL_SUFFIX = "/sandboxnew/tools/shorturl"
//const SANDBOX_AUTHCODETOOPENID_URL_SUFFIX = "/sandboxnew/tools/authcodetoopenid"
