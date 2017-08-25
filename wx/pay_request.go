package wx

import (
	"log"

	"golang.org/x/crypto/pkcs12"
)

type PayRequest struct {
	config PayConfig
}

func NewPayRequest(config PayConfig) PayRequest {
	return PayRequest{config: config}
}

/**
 * 请求，只请求一次，不做重试
 * @param domain
 * @param urlSuffix
 * @param uuid
 * @param data
 * @param connectTimeoutMs
 * @param readTimeoutMs
 * @param useCert 是否使用证书，针对退款、撤销等操作
 * @return
 * @throws Exception
 */
func (request *PayRequest) requestOnce(domain, urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert bool) error {
	if useCert {
		key, cert, err := pkcs12.Decode(request.config.cert, request.config.MchID)
		log.Println(key, cert, err)
	}
	return nil
}
