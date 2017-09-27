package util_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/godcong/wopay/util"
)

func TestPayData_Set(t *testing.T) {
	data := util.PayData{}
	data.Set("ax", "ax")
	log.Println(data)

	log.Println(strconv.FormatBool(false))
	loc, _ := time.LoadLocation("America/New_York")
	log.Println(time.Now().String())
	log.Println(time.Now().In(loc).String())
	log.Println(time.Now().In(loc).Format(util.DATE_TIME_FORMAT))
}
