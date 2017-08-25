package wx

import (
	"log"
	"testing"
)

var data PayData = map[string]string{
	"first":  "1",
	"second": "2",
	"aecond": "3",
	"becond": "4",
}

func TestPayData_Get(t *testing.T) {
	log.Println(data.Get("first"))
}

func TestPayData_Set(t *testing.T) {
	data.Set("third", "3")
	log.Println(data)
}

func TestPayData_IsExist(t *testing.T) {
	log.Println(data.IsExist("first"))
}
