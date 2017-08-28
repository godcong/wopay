package wx

import (
	"strings"

	"sort"

	"errors"

	"crypto/md5"
	"io"

	"crypto/hmac"
	"crypto/sha256"

	"fmt"

	"encoding/xml"

	"bytes"

	"time"

	uuid "github.com/satori/go.uuid"
)

var ErrorSignType = errors.New("sign type error")

func GenerateUUID() string {
	s := uuid.NewV1().String()
	s = strings.Replace(s, "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

func GenerateSignature(reqData PayData, key string, signType SignType) (string, error) {
	var keys sort.StringSlice
	for k := range reqData {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	var sign []string

	for _, k := range keys {
		if k == FIELD_SIGN {
			continue
		}
		v := strings.TrimSpace(reqData[k])
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	if signType == SIGN_TYPE_MD5 {
		sb = MakeSignMD5(sb)
		return sb, nil
	} else if signType == SIGN_TYPE_HMACSHA256 {
		sb = MakeSignHMACSHA256(sb, key)
		return sb, nil
	} else {
		return "", ErrorSignType
	}
}

func MakeSignMD5(data string) string {
	m := md5.New()
	io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

func MakeSignHMACSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

func MapToXml(reqData PayData) (string, error) {
	return mapToXml(reqData, false)
}

func mapToXml(reqData PayData, needHeader bool) (string, error) {

	buff := bytes.NewBuffer(nil)
	if needHeader {
		buff.Write([]byte(xml.Header))
	}

	enc := xml.NewEncoder(buff)

	enc.EncodeToken(xml.StartElement{xml.Name{"", "xml"}, nil})
	for k, v := range reqData {
		enc.EncodeElement(v, xml.StartElement{xml.Name{"", k}, nil})
	}
	enc.EncodeToken(xml.EndElement{xml.Name{"", "xml"}})
	enc.Flush()
	return buff.String(), nil
}

func CurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / time.Millisecond.Nanoseconds()
}

func CurrentTimeStampNS() int64 {
	return time.Now().UnixNano()
}

func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}
