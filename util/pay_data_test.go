package util

import (
	"log"
	"testing"
)

func TestPayData_Set(t *testing.T) {
	data := PayData{}
	data.Set("ax", "ax")
	log.Println(data)

	st := ST{""}
	st.Set("1234")
	log.Println(st)
}
