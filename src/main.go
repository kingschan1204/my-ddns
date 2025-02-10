package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"log"
	"myddns/myapi"
	"myddns/myconf"
	"os"
)

// 记录需要修改域名的最新状态
var item dnspod.RecordListItem

func main() {

	// 检查是否提供了文件路径参数
	if len(os.Args) < 2 {
		fmt.Println("Usage: myddns <file-path>")
		return
	}
	// 获取文件路径
	filePath := os.Args[1]
	log.Println("配置文件路径:", filePath)
	// 调用函数并传入文件路径
	myconf.InitConfig(filePath)
	modifyTask()
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	task := func() {
		modifyTask()
		//fmt.Println("hello world", time.Now())
	}
	//定时任务
	//cron表达式
	spec := "*/60 * * * * ?"
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} //阻塞主线程停止

}

// 更新任务
func modifyTask() {
	if item.RecordId == nil {
		//log.Println("获取域名记录列表")
		list, _ := myapi.RecordList(myconf.App.SecretId, myconf.App.SecretKey, myconf.App.Domain)
		for i := 0; i < len(list); i++ {
			b, _ := json.Marshal(*list[i])
			if *list[i].Name == myconf.App.Target {
				item = *list[i]
				log.Printf("%s \n", b)
			}
		}
	}

	if item.RecordId == nil {
		log.Print("未找到要修改的记录！")
		return
	}

	ip := myapi.GetInterNetIp()
	if *item.Value == ip {
		//log.Print("本机公网Ip: ", ip, " -> 地址未改变,无需修改！")
		return
	}
	myapi.ModifyIp(myconf.App.SecretId, myconf.App.SecretKey, myconf.App.Domain, myconf.App.Target, ip, *item.RecordId)
	//更新ip后重置
	item.RecordId = nil
}
