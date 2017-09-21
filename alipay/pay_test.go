package alipay_test

import (
	"testing"

	"github.com/godcong/wopay/alipay"
)

func TestPay(t *testing.T) {
	client := alipay.PayClientImpl{
		ServerUrl:       alipay.URL,
		AppId:           alipay.APPID,
		PrivateKey:      alipay.RSA_PRIVATE_KEY,
		Format:          alipay.FORMAT_JSON,
		Charset:         alipay.CHARSET_GBK,
		AlipayPublicKey: alipay.ALIPAY_PUBLIC_KEY,
	}
	
}
