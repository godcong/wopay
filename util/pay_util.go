package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/xml"

	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

const CUSTOM_HEADER = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`

type CDATA struct {
	Value string `xml:",cdata"`
}

func GenerateNonceStr() string {
	return GenerateUUID()
}

func GenerateUUID() string {
	s := uuid.NewV1().String()
	s = strings.Replace(s, "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

//MapToString
func MapToString(data PayData) string {
	var keys sort.StringSlice
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	var sign []string

	for _, k := range keys {
		if k == FIELD_SIGN {
			continue
		}
		v := strings.TrimSpace(data[k])
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	log.Println(strings.Join(sign, "&"))
	return strings.Join(sign, "&")
}

func ToUrlParams(data PayData) string {
	keys := data.SortKeys()
	var sign []string
	for _, k := range keys {
		if k == FIELD_SIGN {
			continue
		}
		v := strings.TrimSpace(data[k])
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}

	return strings.Join(sign, "&")

}

//MakeSignMD5 make sign with md5
func MakeSignMD5(data string) string {
	m := md5.New()
	io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

//MakeSignHMACSHA256 make sign with hmac-sha256
func MakeSignHMACSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// MapToXml Convert MAP to XML
func MapToXml(reqData PayData) (string, error) {
	return mapToXml(reqData, false)
}

func mapToXml(reqData PayData, needHeader bool) (string, error) {

	buff := bytes.NewBuffer([]byte(CUSTOM_HEADER))
	if needHeader {
		buff.Write([]byte(xml.Header))
	}

	enc := xml.NewEncoder(buff)

	enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "xml"}})
	for k, v := range reqData {
		if _, err := strconv.ParseInt(v, 10, 0); err != nil {
			enc.EncodeElement(
				CDATA{Value: v}, xml.StartElement{Name: xml.Name{Local: k}})
		} else {
			enc.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: k}})
		}

	}
	enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "xml"}})
	enc.Flush()
	return buff.String(), nil
}

// XmlToMap Convert XML to MAP
func XmlToMap(contentXml string) PayData {
	return xmlToMap(contentXml, false)
}

func xmlToMap(contentXml string, hasHeader bool) PayData {
	data := make(PayData)
	dec := xml.NewDecoder(strings.NewReader(contentXml))
	ele, val := "", ""

	for t, err := dec.Token(); err == nil; t, err = dec.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			ele = token.Name.Local
			//fmt.Printf("This is the sta: %s\n", ele)
			if strings.ToLower(ele) == "xml" {
				//xmlFlag = true
				continue
			}

			// 处理元素结束（标签）
		case xml.EndElement:
			name := token.Name.Local
			//fmt.Printf("This is the end: %s\n", name)
			if strings.ToLower(name) == "xml" {
				break
			}
			if ele == name && ele != "" {
				data.Set(ele, val)
				ele = ""
				val = ""
			}
			// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			//content := string(token)
			//fmt.Printf("This is the content: %v\n", content)
			val = string(token)
			//异常处理(Log输出）
		default:
			log.Println(token)
		}

	}

	return data
}

//CurrentTimeStampMS get current time with millisecond
func CurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / time.Millisecond.Nanoseconds()
}

//CurrentTimeStampNS get current time with nanoseconds
func CurrentTimeStampNS() int64 {
	return time.Now().UnixNano()
}

//CurrentTimeStamp get current time with unix
func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func CurrentTimeStampString() string {
	return strconv.FormatInt(CurrentTimeStamp(), 10)
}

func SHA1(s string) string {
	m := sha1.New()
	m.Write([]byte(s))
	return fmt.Sprintf("%x", m.Sum(nil))
}
