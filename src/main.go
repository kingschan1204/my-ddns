package main

import (
	"fmt"
	"my-ddns/src/myapi"
	"my-ddns/src/myconf"
)

func main() {
	list, _ := myapi.RecordList(myconf.App.SecretId, myconf.App.SecretKey, myconf.App.Domain)
	for i := 0; i < len(list); i++ {
		fmt.Printf("\n %s", *list[i].Name)
	}

}
