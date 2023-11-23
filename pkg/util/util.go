package util

import (
	"math/rand"
	"time"
)

var randowCodes = [...]byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
}

// 随机验证码
func RandomNumeric() string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var pwd []byte = make([]byte, 6)
	for i := 0; i < 3; i++ {
		for j := 0; j < 6; j++ {
			index := r.Int() % len(randowCodes)
			pwd[j] = randowCodes[index]
		}
	}
	return string(pwd)
}

// redis缓存时间
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}
