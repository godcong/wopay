package wx

import "strconv"

type PayDomain interface {
	Report(string, int64, error)
	GetDomain(PayConfig) DomainInfo
}

type DomainInfo struct {
	Domain        string //域名
	PrimaryDomain bool   //该域名是否为主域名。例如:api.mch.weixin.qq.com为主域名

}

func NewDomainInfo(domain string, primary bool) DomainInfo {
	return DomainInfo{
		Domain:        domain,
		PrimaryDomain: primary,
	}
}

func (info *DomainInfo) String() string {
	return "DomainInfo{" + "domain='" + info.Domain + "'" + ", primaryDomain=" + strconv.FormatBool(info.PrimaryDomain) + "}"
}
