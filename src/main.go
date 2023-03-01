package main

import (
	"fmt"
	"my-ddns/src/myapi"
	"my-ddns/src/myconf"
)

func main() {
	fmt.Println(myapi.GetInterNetIp())
	list, _ := myapi.RecordList(myconf.App.SecretId, myconf.App.SecretKey, myconf.App.Domain)
	for i := 0; i < len(list); i++ {
		fmt.Printf("\n %s %s %v", *list[i].Name, *list[i].Value, *list[i].RecordId)
	}

}
