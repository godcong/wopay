package wopay_test

import (
	"log"
	"testing"
	"time"

	"fmt"

	"github.com/satori/go.uuid"
)

func BenchmarkLoadConfig(b *testing.B) {

	m := make(map[string]string)
	for i := 0; i < 50; i++ {
		m[uuid.NewV1().String()] = time.Now().String()
	}
	var temp *map[string]string
	now := time.Now().Nanosecond()
	for i := 0; i < 1000; i++ {
		fmt.Print(i)
		temp = &m
	}
	log.Println(len(*temp))
	mid := time.Now().Nanosecond()
	log.Println("sta", mid-now)
	var temp2 map[string]string
	for i := 0; i < 1000; i++ {
		fmt.Print(i)
		temp2 = m
	}
	log.Println(len(temp2))
	end := time.Now().Nanosecond()
	log.Println("end", end-mid)

}

func TestT1(t *testing.T) {
	l, e := time.LoadLocation("Asia/Shanghai")
	log.Println(time.FixedZone("", 8))
	log.Println(l.String(), e)
	log.Println(time.Now().In(l))
}
