// SignMd5_test.go
package SignMd5

import (
	"fmt"
	"testing"
)

func TestSignByKey(t *testing.T) {

	var m map[string]string
	m = make(map[string]string)

	m["goodsId"] = "201121213 "
	m["goodsName"] = "鼠标"
	m["goodsDesc"] = ""
	m["amount"] = "100"
	m["signMsg"] = "8Z9W5ZGMWCUJMPCOU9SIA7AC5YM58XUS"
	m["HMAC"] = "8Z9W5ZGMWCUJMPCOU9SIA7AC5YM58XUS"
	result := SignByKey(m, "8Z9W5ZGMWCUJMPCOU9SIA7AC5YM58XUS", "UTF8")

	fmt.Printf("assert equal [b337c6d7c2aaaed7a85f11f345e015b0]=[%s]\n", result)
}
