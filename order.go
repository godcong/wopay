package wopay

type Order struct {
	Body           string
	Detail         string
	Attach         string
	OutTradeNo     string
	FeeType        string
	TotalFee       string
	SpBillCreateIp string
	TimeStart      string
	TimeExpire     string
	Goods_tag      string
	NotifyUrl      string
	TradeType      string
	ProductId      string
	LimitPay       string
	Openid         string
	SubOpenid      string
	AuthCode       string
}

const JSAPI = "JSAPI"
const NATIVE = "NATIVE"
const APP = "APP"
const MICROPAY = "MICROPAY"
