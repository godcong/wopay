package util

import (
	"encoding/json"
	"sort"
	"strconv"
	"time"
)

type SortArray struct {
	k interface{}
	v interface{}
}

type PayData map[string]string

func ParseDate(v map[string]string) PayData {
	return PayData(v)
}

func (data PayData) Set(key, val string) {
	data[key] = val
}

func (data PayData) SetString(key, val string) {
	data.Set(key, val)
}

func (data PayData) SetInt(key string, val int64) {
	data.Set(key, strconv.FormatInt(val, 10))
}

func (data PayData) SetFloat(key string, val float64) {
	data.Set(key, strconv.FormatFloat(val, 'f', -1, 64))
}

func (data PayData) SetBoolean(key string, val bool) {
	data.Set(key, strconv.FormatBool(val))
}
func (data PayData) SetDate(key string, val time.Time) {
	loc, _ := time.LoadLocation(DATE_TIMEZONE)
	data.Set(key, val.In(loc).Format(DATE_TIME_FORMAT))
}

func (data PayData) Get(key string) string {
	return data[key]
}

func (data PayData) IsExist(key string) bool {
	_, b := data[key]
	return b
}

func (data PayData) SortKeys() []string {
	var keys sort.StringSlice
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

func (data PayData) ToJson() string {
	b, e := json.Marshal(&data)
	if e != nil {
		return ""
	}
	return string(b)
}

func (data PayData) ToMap() map[string]string {
	return data
}

func (data PayData) IsNil() bool {
	if data == nil || len(data) == 0 {
		return true
	}
	return false
}
