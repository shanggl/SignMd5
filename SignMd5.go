// SignMd5 project SignMd5.go
package SignMd5

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/axgle/mahonia"
)

func createLinkString(m map[string]string) string {
	//create the link string seprated by |key=
	//sort
	var i int = 0
	sortedKeys := make([]string, len(m))
	for k := range m {
		sortedKeys[i] = k
		i++
	}
	sort.Strings(sortedKeys)
	//combin with | to get  key=value|
	i = len(m)
	c := 0
	var buffer bytes.Buffer
	for j := 0; j < i; j++ {
		buffer.WriteString(sortedKeys[j])
		buffer.WriteString("=")
		buffer.WriteString(m[sortedKeys[j]])
		if c < i-1 {
			buffer.WriteString("|")
		}
		c++
	}
	return buffer.String()

}

func doFilterParam(m map[string]string) map[string]string {
	//filter the k,v to remove hmac\signMsg\cert
	//igore null or empty string
	var nm map[string]string
	nm = make(map[string]string)
	for k, v := range m {
		if strings.EqualFold(k, "hmac") || strings.EqualFold(k, "signMsg") || strings.EqualFold(k, "cert") {
			continue
		} else if len(v) == 0 {
			continue
		} else {
			nm[k] = strings.Trim(v, " ")
		}
	}

	return nm
}

func SignByKey(m map[string]string, key string, charSet string) string {
	//sign k,v by key using charSet default utf8
	s := createLinkString(doFilterParam(m))

	//fmt.Printf("begin sign String[%s]\n", s)

	s = s + "|key=" + key

	//fmt.Printf("key =[%s] after [%s]\n", key, s)

	//convert to utf8
	enc := mahonia.NewEncoder(charSet)

	h := md5.New()
	h.Write([]byte(enc.ConvertString(s)))
	return hex.EncodeToString(h.Sum(nil))
}
