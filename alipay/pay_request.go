package alipay

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
	GetResponse() PayResponse

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

//type AlipayRequest struct {
//}

type PayRequestWap struct {
}

type PayRequestApp struct {
}
type PayRequestPage struct {
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

/**
 * 获取TOP的API名称。
 *
 * @return API名称
 */
func (PayRequestPage) GetApiMethodName() string {
	return ""
}

/**
 * 获取所有的Key-Value形式的文本请求参数集合。其中：
 * <ul>
 * <li>Key: 请求参数名</li>
 * <li>Value: 请求参数值</li>
 * </ul>
 *
 * @return 文本请求参数集合
 */
func (PayRequestPage) GetTextParams() map[string]string {
	return nil
}

/**
 * 得到当前接口的版本
 *
 * @return API版本
 */
func (PayRequestPage) GetApiVersion() string {
	return ""
}

/**
 * 设置当前API的版本信息
 *
 * @param apiVersion API版本
 */
func (PayRequestPage) SetApiVersion(apiVersion string) {

}

/**
 * 获取终端类型
 *
 * @return 终端类型
 */
func (PayRequestPage) GetTerminalType() string {
	return ""
}

/**
 * 设置终端类型
 *
 * @param terminalType 终端类型
 */
func (PayRequestPage) SetTerminalType(terminalType string) {

}

/**
 * 获取终端信息
 *
 * @return 终端信息
 */
func (PayRequestPage) GetTerminalInfo() string {
	return ""
}

/**
 * 设置终端信息
 *
 * @param terminalInfo 终端信息
 */
func (PayRequestPage) SetTerminalInfo(terminalInfo string) {}

/**
 * 获取产品码
 *
 * @return 产品码
 */
func (PayRequestPage) GetProdCode() string {
	return ""
}

/**
 * 设置产品码
 *
 * @param prodCode 产品码
 */
func (PayRequestPage) SetProdCode(prodCode string) {}

/**
 * 返回通知地址
 *
 * @return
 */
func (PayRequestPage) GetNotifyUrl() string {
	return ""
}

/**
 *  设置通知地址
 *
 * @param notifyUrl
 */
func (PayRequestPage) SetNotifyUrl(notifyUrl string) {}

/**
 * 返回回跳地址
 *
 * @return
 */
func (PayRequestPage) GetReturnUrl() string {
	return ""
}

/**
 *  设置回跳地址
 *
 * @param notifyUrl
 */
func (PayRequestPage) SetReturnUrl(returnUrl string) {}

/**
 * 得到当前API的响应结果类型
 *
 * @return 响应类型
 */
func (PayRequestPage) GetResponse() PayResponse {
	return nil
}

/**
 * 判断是否需要加密
 *
 * @return
 */
func (PayRequestPage) IsNeedEncrypt() bool {
	return false
}

/**
 * 设置请求是否需要加密
 *
 * @param needEncrypt
 */
func (PayRequestPage) SetNeedEncrypt(needEncrypt bool) {}

func (PayRequestPage) GetBizModel() PayObject {
	return nil
}

/**
 * 设置业务实体，如需使用此方法，请勿直接设置biz_content
 *
 * @param bizModel
 */
func (PayRequestPage) SetBizModel(bizModel PayObject) {}

/**
 * Getter method for property <tt>code</tt>.
 *
 * @return property value of code
 */
func (PayRequestPage) GetCode() string {
	return ""
}

/**
 * Setter method for property <tt>code</tt>.
 *
 * @param code value to be assigned to property code
 */
func (PayRequestPage) SetCode(code string)                {}
func (PayRequestPage) GetMsg() string                     { return "" }
func (PayRequestPage) SetMsg(msg string)                  {}
func (PayRequestPage) GetSubCode() string                 { return "" }
func (PayRequestPage) SetSubCode(subCode string)          {}
func (PayRequestPage) GetSubMsg() string                  { return "" }
func (PayRequestPage) SetSubMsg(subMsg string)            {}
func (PayRequestPage) SetBody(body string)                {}
func (PayRequestPage) GetBody() string                    { return "" }
func (PayRequestPage) GetParams() map[string]string       { return nil }
func (PayRequestPage) SetParams(params map[string]string) {}
func (PayRequestPage) IsSuccess() bool                    { return false }
