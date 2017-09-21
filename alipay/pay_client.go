package alipay

type PayClientImpl struct {
	ServerUrl      string
	AppId          string
	PrivateKey     string
	ProdCode       string
	Format         string //= AlipayConstants.FORMAT_JSON
	Sign_type      string //= AlipayConstants.SIGN_TYPE_RSA
	EncryptType    string //= AlipayConstants.ENCRYPT_TYPE_AES
	EncryptKey     string
	AlipayPublicKey      string
	Charset        string
	ConnectTimeout int //= 3000
	ReadTimeout    int //= 15000
}

type PayClient interface {
	/**
	 * @param request
	 * @return PayResponse,error
	 */
	Execute(request PayRequest) (PayResponse, error)
	/**
	 * @param request
	 * @param accessToken
	 * @return PayResponse,error
	 */
	ExecuteToken(request PayRequest, accessToken string, appToken string) (PayResponse, error)

	/**
	 * @param request
	 * @return PayResponse,error
	 */
	PageExecute(request PayRequest) (PayResponse, error)

	/**
	* @param request
	* @param method
	* @return PayResponse,error
	 */
	PageExecuteMethod(request PayRequest, method string) (PayResponse, error)

	/**
	* SDK客户端调用生成sdk字符串
	* @param request
	* @return PayResponse,error
	 */
	SdkExecute(request PayRequest) (PayResponse, error)

	/**
	 * 移动客户端同步结果返回解析的参考工具方法
	 *
	 * @param result 移动客户端SDK同步返回的结果map，一般包含resultStatus，result和memo三个key
	 * @param requsetClazz 接口请求request类，如App支付传入 AlipayTradeAppPayRequest.class
	 * @return 同步返回结果的response对象, error
	 */
	ParseAppSyncResult(result map[string]string, request PayRequest) (PayResponse, error)
}
