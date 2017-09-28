package alipay

import "github.com/godcong/wopay/util"

type RequestHolder struct {
	ProtocalMustParams util.PayData
	ProtocalOptParams  util.PayData
	ApplicationParams  util.PayData
}
