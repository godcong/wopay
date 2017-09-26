package alipay

import "reflect"

type PayRequest interface {
	PayResponse
	/**
	 * 获取TOP的API名称。
	 *
	 * @return API名称
	 */
	GetApiMethodName() string

	/**
	 * 获取所有的Key-Value形式的文本请求参数集合。其中：
	 * <ul>
	 * <li>Key: 请求参数名</li>
	 * <li>Value: 请求参数值</li>
	 * </ul>
	 *
	 * @return 文本请求参数集合
	 */
	GetTextParams() map[string]string

	/**
	 * 得到当前接口的版本
	 *
	 * @return API版本
	 */
	GetApiVersion() string

	/**
	 * 设置当前API的版本信息
	 *
	 * @param apiVersion API版本
	 */
	SetApiVersion(apiVersion string)

	/**
	 * 获取终端类型
	 *
	 * @return 终端类型
	 */
	GetTerminalType() string

	/**
	 * 设置终端类型
	 *
	 * @param terminalType 终端类型
	 */
	SetTerminalType(terminalType string)

	/**
	 * 获取终端信息
	 *
	 * @return 终端信息
	 */
	GetTerminalInfo() string

	/**
	 * 设置终端信息
	 *
	 * @param terminalInfo 终端信息
	 */
	SetTerminalInfo(terminalInfo string)

	/**
	 * 获取产品码
	 *
	 * @return 产品码
	 */
	GetProdCode() string

	/**
	 * 设置产品码
	 *
	 * @param prodCode 产品码
	 */
	SetProdCode(prodCode string)

	/**
	 * 返回通知地址
	 *
	 * @return
	 */
	GetNotifyUrl() string

	/**
	 *  设置通知地址
	 *
	 * @param notifyUrl
	 */
	SetNotifyUrl(notifyUrl string)

	/**
	 * 返回回跳地址
	 *
	 * @return
	 */
	GetReturnUrl() string

	/**
	 *  设置回跳地址
	 *
	 * @param notifyUrl
	 */
	SetReturnUrl(returnUrl string)

	/**
	 * 得到当前API的响应结果类型
	 *
	 * @return 响应类型
	 */
	GetResponseType() reflect.Type

	/**
	 * 判断是否需要加密
	 *
	 * @return
	 */
	IsNeedEncrypt() bool

	/**
	 * 设置请求是否需要加密
	 *
	 * @param needEncrypt
	 */
	SetNeedEncrypt(needEncrypt bool)

	GetBizModel() PayObject

	/**
	 * 设置业务实体，如需使用此方法，请勿直接设置biz_content
	 *
	 * @param bizModel
	 */
	SetBizModel(bizModel PayObject)
}

type PayRequestWap struct {
	payRequest
}

type PayRequestApp struct {
	payRequest
}

type PayRequestPage struct {
	payRequest
}

type payRequest struct {
	/**
	* add user-defined text parameters
	 */
	udfParams map[string]string
	/**
	* 统一收单下单并支付页面接口
	 */
	bizContent   string
	terminalType string
	terminalInfo string
	prodCode     string
	notifyUrl    string
	returnUrl    string
	needEncrypt  bool
	bizModel     interface{}
}

func (*payRequest) GetCode() string {
	panic("implement me")
}

func (*payRequest) SetCode(code string) {
	panic("implement me")
}

func (*payRequest) GetMsg() string {
	panic("implement me")
}

func (*payRequest) SetMsg(msg string) {
	panic("implement me")
}

func (*payRequest) GetSubCode() string {
	panic("implement me")
}

func (*payRequest) SetSubCode(subCode string) {
	panic("implement me")
}

func (*payRequest) GetSubMsg() string {
	panic("implement me")
}

func (*payRequest) SetSubMsg(subMsg string) {
	panic("implement me")
}

func (*payRequest) GetBody() string {
	panic("implement me")
}

func (*payRequest) SetBody(body string) {
	panic("implement me")
}

func (*payRequest) GetParams() map[string]string {
	panic("implement me")
}

func (*payRequest) SetParams(params map[string]string) {
	panic("implement me")
}

func (*payRequest) IsSuccess() bool {
	panic("implement me")
}

func (*payRequest) GetApiMethodName() string {
	panic("implement me")
}

func (*payRequest) GetTextParams() map[string]string {
	panic("implement me")
}

func (*payRequest) GetApiVersion() string {
	panic("implement me")
}

func (*payRequest) SetApiVersion(apiVersion string) {
	panic("implement me")
}

func (*payRequest) GetTerminalType() string {
	panic("implement me")
}

func (*payRequest) SetTerminalType(terminalType string) {
	panic("implement me")
}

func (*payRequest) GetTerminalInfo() string {
	panic("implement me")
}

func (*payRequest) SetTerminalInfo(terminalInfo string) {
	panic("implement me")
}

func (*payRequest) GetProdCode() string {
	panic("implement me")
}

func (*payRequest) SetProdCode(prodCode string) {
	panic("implement me")
}

func (*payRequest) GetNotifyUrl() string {
	panic("implement me")
}

func (*payRequest) SetNotifyUrl(notifyUrl string) {
	panic("implement me")
}

func (*payRequest) GetReturnUrl() string {
	panic("implement me")
}

func (*payRequest) SetReturnUrl(returnUrl string) {
	panic("implement me")
}

func (*payRequest) GetResponseType() reflect.Type {
	panic("implement me")
}

func (*payRequest) IsNeedEncrypt() bool {
	panic("implement me")
}

func (*payRequest) SetNeedEncrypt(needEncrypt bool) {
	panic("implement me")
}

func (*payRequest) GetBizModel() PayObject {
	panic("implement me")
}

func (*payRequest) SetBizModel(bizModel PayObject) {
	panic("implement me")
}

func NewPayRequest() PayRequest {
	return &PayRequestPage{}
}

//
///**
// * 获取TOP的API名称。
// *
// * @return API名称
// */
//func (page *PayRequestPage) GetApiMethodName() string {
//	return ""
//}
//
///**
// * 获取所有的Key-Value形式的文本请求参数集合。其中：
// * <ul>
// * <li>Key: 请求参数名</li>
// * <li>Value: 请求参数值</li>
// * </ul>
// *
// * @return 文本请求参数集合
// */
//func (page *PayRequestPage) GetTextParams() map[string]string {
//	return nil
//}
//
///**
// * 得到当前接口的版本
// *
// * @return API版本
// */
//func (page *PayRequestPage) GetApiVersion() string {
//	return ""
//}
//
///**
// * 设置当前API的版本信息
// *
// * @param apiVersion API版本
// */
//func (page *PayRequestPage) SetApiVersion(apiVersion string) {
//
//}
//
///**
// * 获取终端类型
// *
// * @return 终端类型
// */
//func (page *PayRequestPage) GetTerminalType() string {
//	return ""
//}
//
///**
// * 设置终端类型
// *
// * @param terminalType 终端类型
// */
//func (page *PayRequestPage) SetTerminalType(terminalType string) {
//
//}
//
///**
// * 获取终端信息
// *
// * @return 终端信息
// */
//func (page *PayRequestPage) GetTerminalInfo() string {
//	return ""
//}
//
///**
// * 设置终端信息
// *
// * @param terminalInfo 终端信息
// */
//func (page *PayRequestPage) SetTerminalInfo(terminalInfo string) {}
//
///**
// * 获取产品码
// *
// * @return 产品码
// */
//func (page *PayRequestPage) GetProdCode() string {
//	return ""
//}
//
///**
// * 设置产品码
// *
// * @param prodCode 产品码
// */
//func (page *PayRequestPage) SetProdCode(prodCode string) {}
//
///**
// * 返回通知地址
// *
// * @return
// */
//func (page *PayRequestPage) GetNotifyUrl() string {
//	return ""
//}
//
///**
// *  设置通知地址
// *
// * @param notifyUrl
// */
//func (page *PayRequestPage) SetNotifyUrl(notifyUrl string) {}
//
///**
// * 返回回跳地址
// *
// * @return
// */
//func (page *PayRequestPage) GetReturnUrl() string {
//	return ""
//}
//
///**
// *  设置回跳地址
// *
// * @param notifyUrl
// */
//func (page *PayRequestPage) SetReturnUrl(returnUrl string) {}
//
///**
// * 得到当前API的响应结果类型
// *
// * @return 响应类型
// */
//func (page *PayRequestPage) GetResponse() PayResponse {
//	return nil
//}
//
///**
// * 判断是否需要加密
// *
// * @return
// */
//func (page *PayRequestPage) IsNeedEncrypt() bool {
//	return false
//}
//
///**
// * 设置请求是否需要加密
// *
// * @param needEncrypt
// */
//func (page *PayRequestPage) SetNeedEncrypt(needEncrypt bool) {}
//
//func (page *PayRequestPage) GetBizModel() PayObject {
//	return nil
//}
//
///**
// * 设置业务实体，如需使用此方法，请勿直接设置biz_content
// *
// * @param bizModel
// */
//func (page *PayRequestPage) SetBizModel(bizModel PayObject) {}
//
///**
// * Getter method for property <tt>code</tt>.
// *
// * @return property value of code
// */
//func (page *PayRequestPage) GetCode() string {
//	return ""
//}
//
///**
// * Setter method for property <tt>code</tt>.
// *
// * @param code value to be assigned to property code
// */
//func (page *PayRequestPage) SetCode(code string) {}
//
//func (page *PayRequestPage) GetMsg() string                     { return "" }
//func (page *PayRequestPage) SetMsg(msg string)                  {}
//func (page *PayRequestPage) GetSubCode() string                 { return "" }
//func (page *PayRequestPage) SetSubCode(subCode string)          {}
//func (page *PayRequestPage) GetSubMsg() string                  { return "" }
//func (page *PayRequestPage) SetSubMsg(subMsg string)            {}
//func (page *PayRequestPage) SetBody(body string)                {}
//func (page *PayRequestPage) GetBody() string                    { return "" }
//func (page *PayRequestPage) GetParams() map[string]string       { return nil }
//func (page *PayRequestPage) SetParams(params map[string]string) {}
//func (page *PayRequestPage) IsSuccess() bool                    { return false }
