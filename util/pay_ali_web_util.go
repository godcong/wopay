package util

import "strings"

func BuildQuery(data PayData, charset string) (string, error) {
	if data.IsNil() {
		return "", nil
	}
	var sign []string
	for name, value := range data {
		value := strings.TrimSpace(value)
		if value != "" && name != "" {
			if charset == CHARSET_GBK {
				if gbk, e := Utf8ToGbk([]byte(value)); e == nil {
					value = string(gbk)
				}
			}
			sign = append(sign, strings.Join([]string{name, value}, "="))
		}
	}

	return strings.Join(sign, "&"), nil
}
