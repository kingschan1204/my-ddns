package myapi

import (
	"io/ioutil"
	"net/http"
)

//获取本机公网ip的地址
//https://ipinfo.io/ip
//https://myexternalip.com/raw
//http://httpbin.org/ip
func GetInterNetIp() string {
	resp, err := http.Get("https://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//s := buf.String()
	return string(content)
}
