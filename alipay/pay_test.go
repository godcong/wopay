package alipay_test

import (
	"testing"

	"log"

	"github.com/godcong/wopay/alipay"
)

func TestPay(t *testing.T) {
	client := alipay.NewPayClient(
		alipay.URL,
		alipay.APPID,
		alipay.RSA_PRIVATE_KEY,
		alipay.FORMAT_JSON,
		alipay.CHARSET_UTF8,
		alipay.ALIPAY_PUBLIC_KEY,
		alipay.SIGN_TYPE_RSA2,
		"",
		"")

	request := alipay.PayRequest(alipay.PayRequestPage{})
	request.SetReturnUrl("")
	request.SetNotifyUrl("")

	out_trade_no, total_amount, subject, body := "", "", "", ""
	request.SetBizModel(out_trade_no + total_amount + subject + body)
	resp, err := client.PageExecute(request)
	if err != nil {
		return
	}
	resp.GetBody()

	log.Println(request)
}
