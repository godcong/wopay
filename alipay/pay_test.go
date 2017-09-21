package alipay_test

import (
	"testing"

	"log"

	"github.com/godcong/wopay/alipay"
)

func TestPay(t *testing.T) {
	//client := alipay.PayClientImpl{
	//	ServerUrl:       alipay.URL,
	//	AppId:           alipay.APPID,
	//	PrivateKey:      alipay.RSA_PRIVATE_KEY,
	//	Format:          alipay.FORMAT_JSON,
	//	Charset:         alipay.CHARSET_UTF8,
	//	AlipayPublicKey: alipay.ALIPAY_PUBLIC_KEY,
	//}

	request := alipay.PayRequest(alipay.PayRequestPage{})
	log.Println(request)
}
