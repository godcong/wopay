package wx

import (
	"log"
	"testing"
)

type pay_util_test struct {
	Data       string
	Key        string
	MD5        string
	HMACSHA256 string
}

var td = []pay_util_test{
	pay_util_test{"1234", "1111", "81DC9BDB52D04DC20036DBD8313ED055", "388472000D021C1F7D089C6472750CDF80B91775DD7FF69A59D324216CDFF15E"},
	pay_util_test{"abcd", "aaaa", "E2FC714C4727EE9395F324CD2E7F331F", "6FA028B8CC17AE66B0EBF03E26F5E6587C114D1ECDB439D8A578E4D47DDED3EF"},
}

func TestGenerateSignature(t *testing.T) {
	GenerateSignature(data, "key", 0)
}

func TestMakeSignMD5(t *testing.T) {
	for _, v := range td {
		md5 := MakeSignMD5(v.Data)
		if v.MD5 != md5 {
			t.Error(v.Data, md5, v.MD5)
		}
	}
}

func TestMakeSignHMACSHA256(t *testing.T) {
	for _, v := range td {
		hmacsha256 := MakeSignHMACSHA256(v.Data, v.Key)
		if v.HMACSHA256 != hmacsha256 {
			t.Error(v.Data, hmacsha256, v.HMACSHA256)
		}
	}
}

var input = `<xml>
    <return_code><![CDATA[SUCCESS]]></return_code>
    <return_msg><![CDATA[OK]]></return_msg>
    <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
    <mch_id><![CDATA[10000100]]></mch_id>
    <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
    <openid><![CDATA[oUpF8uMuAJO_M2pxb1Q9zNjWeS6o]]></openid>
    <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
    <result_code><![CDATA[SUCCESS]]></result_code>
    <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
    <trade_type><![CDATA[JSAPI]]></trade_type>
 </xml> `

func TestMapToXml(t *testing.T) {
	output := XmlToMap(input)
	r, e := MapToXml(output)
	log.Println(r, e)
}

func TestXmlToMap(t *testing.T) {

	log.Println(XmlToMap(input))
}

func TestGetSignKey(t *testing.T) {
	log.Println(GetSignKey())
}
