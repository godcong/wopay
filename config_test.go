package wopay_test

import (
	"log"
	"testing"

	"github.com/godcong/wopay"
)

func TestGetWechat(t *testing.T) {
	log.Println(wopay.GetWechat())
}
