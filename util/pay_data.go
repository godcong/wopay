package util

import (
	"encoding/json"
	"sort"
)

type PayData map[string]string

func (data PayData) Set(key, val string) {
	data[key] = val
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

type ST []string

func (data ST) Set(val string) {
	data = ST{val}
}
