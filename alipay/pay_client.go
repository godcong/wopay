package alipay

import (
	"log"

	"github.com/godcong/wopay/wxpay"
)

type payClient struct {
	serverUrl       string
	appId           string
	privateKey      string
	prodCode        string
	format          string //= AlipayConstants.FORMAT_JSON
	signType        string //= AlipayConstants.SIGN_TYPE_RSA
	encryptType     string //= AlipayConstants.ENCRYPT_TYPE_AES
	encryptKey      string
	alipayPublicKey string
	charset         string
	connectTimeout  int //= 3000
	readTimeout     int //= 15000
}

func defaultPayClient(serverUrl, appId, priviteKey string) *payClient {
	return &payClient{
		serverUrl:       serverUrl,
		appId:           appId,
		format:          FORMAT_JSON,
		alipayPublicKey: "",
		privateKey:      priviteKey,
		signType:        SIGN_TYPE_RSA,
		encryptType:     ENCRYPT_TYPE_AES,
		connectTimeout:  3000,
		readTimeout:     15000,
	}
}

func NewPayClient(serverUrl, appId, privateKey, format,
	charset, alipayPulicKey, signType,
	encryptKey, encryptType string) PayClient {
	client := defaultPayClient(serverUrl, appId, privateKey)
	client.format = format
	client.charset = charset
	client.alipayPublicKey = alipayPulicKey
	client.signType = signType
	client.encryptKey = encryptKey
	client.encryptType = encryptType
	return client
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

func (*payClient) Execute(request PayRequest) (PayResponse, error) {
	panic("implement me")
}

func (*payClient) ExecuteToken(request PayRequest, accessToken string, appToken string) (PayResponse, error) {
	panic("implement me")
}

func (*payClient) execute(request PayRequest, accessToken string, appToken string) (PayResponse, error) {
	request.GetResponseType()

	panic("implement me")
}

func (*payClient) PageExecute(request PayRequest) (PayResponse, error) {
	panic("implement me")
}

func (*payClient) PageExecuteMethod(request PayRequest, method string) (PayResponse, error) {
	panic("implement me")
}

func (*payClient) SdkExecute(request PayRequest) (PayResponse, error) {
	panic("implement me")
}

func (*payClient) ParseAppSyncResult(result map[string]string, request PayRequest) (PayResponse, error) {
	panic("implement me")
}

func (*payClient) doPost(request PayRequest,
	accessToken,
	appAuthToken string) (PayResponse, error) {
	data := wxpay.PayData{}
	//Todo
	log.Println(data)
}

func getRequestHolderWithSign(result map[string]string, request PayRequest) (PayResponse, error) {

}
