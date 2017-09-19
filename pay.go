package wopay

import "container/list"

const url = "https://api.mch.weixin.qq.com/pay/unifiedorder"

type payment struct {
	PayType string
}

func T1() {
	list.New()
}
