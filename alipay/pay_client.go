package alipay

type PayClient struct {
	serverUrl       string
	appId           string
	privateKey      string
	prodCode        string
	format          string //= AlipayConstants.FORMAT_JSON
	sign_type       string //= AlipayConstants.SIGN_TYPE_RSA
	encryptType     string //= AlipayConstants.ENCRYPT_TYPE_AES
	encryptKey      string
	alipayPublicKey string
	charset         string
	connectTimeout  int //= 3000
	readTimeout     int //= 15000
}
