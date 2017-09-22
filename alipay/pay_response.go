package alipay

import "strings"

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

type PayResponseImple struct {
	Code    string
	Msg     string
	SubCode string
	SubMsg  string
	Body    string
	Params  map[string]string
}

func (resp *PayResponseImple) GetCode() string {
	panic("implement me")
}

func (resp *PayResponseImple) SetCode(code string) {
	panic("implement me")
}

func (resp *PayResponseImple) GetMsg() string {
	panic("implement me")
}

func (resp *PayResponseImple) SetMsg(msg string) {
	panic("implement me")
}

func (resp *PayResponseImple) GetSubCode() string {
	panic("implement me")
}

func (resp *PayResponseImple) SetSubCode(subCode string) {
	panic("implement me")
}

func (resp *PayResponseImple) GetSubMsg() string {
	panic("implement me")
}

func (resp *PayResponseImple) SetSubMsg(subMsg string) {
	panic("implement me")
}

func (resp *PayResponseImple) GetBody() string {
	panic("implement me")
}

func (resp *PayResponseImple) SetBody(body string) {
	panic("implement me")
}

func (resp *PayResponseImple) GetParams() map[string]string {
	panic("implement me")
}

func (resp *PayResponseImple) SetParams(params map[string]string) {
	panic("implement me")
}

func (resp *PayResponseImple) IsSuccess() bool {
	if strings.TrimSpace(resp.SubCode) == "" {
		return false
	}
	return true
}
