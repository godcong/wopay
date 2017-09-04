package wopay

import (
	"log"
	"testing"
	"time"

	"fmt"

	"github.com/godcong/wopay/wx"
	"github.com/satori/go.uuid"
)

var out_trade_no = "201613091059590000003433-asd002"

func TestDoUnifiedOrder(t *testing.T) {
	data := make(wx.PayData)
	data.Set("body", "腾讯充值中心-QQ会员充值")
	data.Set("out_trade_no", out_trade_no)
	data.Set("device_info", "")
	data.Set("fee_type", "CNY")
	data.Set("total_fee", "1")
	data.Set("spbill_create_ip", "123.12.12.123")
	data.Set("notify_url", "http://test.letiantian.me/wxpay/notify")
	data.Set("trade_type", "NATIVE")
	data.Set("product_id", "12")

	rdata, err := wx.UnifiedOrder(data)
	log.Println(rdata, err)

}

func BenchmarkLoadConfig(b *testing.B) {

	m := make(map[string]string)
	for i := 0; i < 50; i++ {
		m[uuid.NewV1().String()] = time.Now().String()
	}
	var temp *map[string]string
	now := time.Now().Nanosecond()
	for i := 0; i < 1000; i++ {
		fmt.Print(i)
		temp = &m
	}
	log.Println(len(*temp))
	mid := time.Now().Nanosecond()
	log.Println("sta", mid-now)
	var temp2 map[string]string
	for i := 0; i < 1000; i++ {
		fmt.Print(i)
		temp2 = m
	}
	log.Println(len(temp2))
	end := time.Now().Nanosecond()
	log.Println("end", end-mid)

}
