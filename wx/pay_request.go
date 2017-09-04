package wx

import (
	"crypto/tls"

	"net/http"

	"bytes"

	"errors"
	"io/ioutil"
	"log"
)

type PayRequest struct {
	config PayConfig
}

var (
	ErrorNilDomain = errors.New("PayConfig.PayDomain().getDomain() is empty or null")
)

func NewPayRequest(config PayConfig) *PayRequest {
	return &PayRequest{config: config}
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
func (request *PayRequest) RequestOnce(domain, urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert bool) (string, error) {
	return requestOnce(request, domain, urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, useCert)

	var tr *http.Transport
	if useCert {
		//key, cert, err := pkcs12.Decode(request.config.cert, request.config.MchID)
		//cert, err := tls.LoadX509KeyPair(SSLCERT_PATH, SSLKEY_PATH)
		//if err != nil {
		//	return "", err
		//}
		tlsConfig := &tls.Config{
			//Certificates:       []tls.Certificate(cert),
			InsecureSkipVerify: false,
		}
		tlsConfig.BuildNameToCertificate()
		tr = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	} else {
		//c := &http.Client{
		//	Transport: &http.Transport{
		//Dial: (&net.Dialer{
		//Timeout:   30 * time.Second,
		//KeepAlive: 30 * time.Second,
		//}).Dial,
		//TLSHandshakeTimeout:   10 * time.Second,
		//ResponseHeaderTimeout: 10 * time.Second,
		//ExpectContinueTimeout: 1 * time.Second,
		//},
		//}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	url := "https://" + domain + urlSuffix

	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", "wxpay sdk go v1.0 "+request.config.MchID())

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func requestOnce(request *PayRequest, domain, urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert bool) (string, error) {
	var tr *http.Transport
	if useCert {
		//key, cert, err := pkcs12.Decode(request.config.cert, request.config.MchID)
		//cert, err := tls.LoadX509KeyPair(SSLCERT_PATH, SSLKEY_PATH)
		//if err != nil {
		//	return "", err
		//}
		tlsConfig := &tls.Config{
			//Certificates:       []tls.Certificate(cert),
			InsecureSkipVerify: false,
		}
		tlsConfig.BuildNameToCertificate()
		tr = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	} else {
		//c := &http.Client{
		//	Transport: &http.Transport{
		//Dial: (&net.Dialer{
		//Timeout:   30 * time.Second,
		//KeepAlive: 30 * time.Second,
		//}).Dial,
		//TLSHandshakeTimeout:   10 * time.Second,
		//ResponseHeaderTimeout: 10 * time.Second,
		//ExpectContinueTimeout: 1 * time.Second,
		//},
		//}
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	url := "https://" + domain + urlSuffix
	log.Println(urlSuffix)
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", "wxpay sdk go v1.0 "+request.config.MchID())

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func (request *PayRequest) RequestWithoutCert(urlSuffix, uuid, data string, autoReport bool) (string, error) {
	//elapsedTimeMillis := int64(0)
	//startTimestampMs := CurrentTimeStampMS()
	//firstHasDnsErr, firstHasConnectTimeout, firstHasReadTimeout := false, false, false
	//domainInfo := request.config.PayDomain().GetDomain()
	//
	//result, err := request.requestOnce(domainInfo.Domain, url, uuid, body, connect, read)
	//if err == nil {
	//	elapsedTimeMillis = CurrentTimeStampMS() - startTimestampMs
	//	request.config.PayDomain().Report(domainInfo.Domain, elapsedTimeMillis, nil)
	//
	//}
	return request.request(urlSuffix, uuid, data, request.config.ConnectTimeoutMs(), request.config.ReadTimeoutMs(), false, autoReport)
}

func (request *PayRequest) RequestWithoutCertTimeout(urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, autoReport bool) (string, error) {
	//elapsedTimeMillis := int64(0)
	//startTimestampMs := CurrentTimeStampMS()
	//firstHasDnsErr, firstHasConnectTimeout, firstHasReadTimeout := false, false, false
	//domainInfo := request.config.PayDomain().GetDomain()
	//
	//result, err := request.requestOnce(domainInfo.Domain, url, uuid, body, connect, read)
	//if err == nil {
	//	elapsedTimeMillis = CurrentTimeStampMS() - startTimestampMs
	//	request.config.PayDomain().Report(domainInfo.Domain, elapsedTimeMillis, nil)
	//
	//}
	return request.request(urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, false, autoReport)
}

func (request *PayRequest) request(urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert, autoReport bool) (string, error) {
	startTimestampMs := CurrentTimeStampMS()
	firstHasDnsErr, firstHasConnectTimeout, firstHasReadTimeout := false, false, false
	domainInfo := request.config.PayDomainInstance().GetDomainInfo()
	if domainInfo == nil {
		return "", ErrorNilDomain
	}
	result, err := requestOnce(request, domainInfo.Domain, urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, useCert)
	elapsedTimeMillis := CurrentTimeStampMS() - startTimestampMs
	request.config.PayDomainInstance().Report(domainInfo.Domain, elapsedTimeMillis, nil)

	PayReportInstance(request.config).Report(uuid,
		elapsedTimeMillis,
		domainInfo.Domain,
		domainInfo.PrimaryDomain,
		connectTimeoutMs,
		readTimeoutMs,
		firstHasDnsErr,
		firstHasConnectTimeout,
		firstHasReadTimeout)

	return result, err
}
