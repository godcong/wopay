package alipay

type PayParser interface {
	Parse(rsp string)(PayResponse,error)
	
}
