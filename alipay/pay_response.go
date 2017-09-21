package alipay

import "strings"

type PayResponseImple struct {
	Code    string
	Msg     string
	SubCode string
	SubMsg  string
	Body    string
	Params  map[string]string
}

func (resp *PayResponseImple) IsSuccess() bool {
	if strings.TrimSpace(resp.SubCode) == "" {
		return false
	}
	return true
}

//var pay (PayResponse) = PayResponseImple{}

type PayResponse interface {

	/**
	 * Getter method for property <tt>code</tt>.
	 *
	 * @return property value of code
	 */
	GetCode() string

	/**
	 * Setter method for property <tt>code</tt>.
	 *
	 * @param code value to be assigned to property code
	 */
	SetCode(code string)
	GetMsg() string
	SetMsg(msg string)
	GetSubCode() string
	SetSubCode(subCode string)
	GetSubMsg() string
	SetSubMsg(subMsg string)
	GetBody() string
	SetBody(body string)
	GetParams() map[string]string
	SetParams(params map[string]string)
	IsSuccess() bool
}
