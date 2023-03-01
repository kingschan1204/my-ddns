package myapi

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"log"
)

// 获取主域名下的所有解析记录
func RecordList(sid string, skey string, domain string) (list []*dnspod.RecordListItem, err error) {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		sid,
		skey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(domain)
	// 返回的resp是一个DescribeRecordListResponse的实例，与请求对象对应
	response, err := client.DescribeRecordList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, err
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	/*fmt.Printf("%s", response.ToJsonString())

	for i := 0; i < len(response.Response.RecordList); i++ {
		name := response.Response.RecordList[i].Name
		fmt.Printf("\n %s", *name)
	}*/
	return response.Response.RecordList, nil
}

// 更新域名记录
// sid : tencent SecretId
// skey : tencent SecretKey
// domain : 一级域名
// target : 要修改的二级域名，如果是一级域名传@
// ip : 本机的公网地址
// id : 要修改的二级域名的RecordId
func ModifyIp(sid string, skey string, domain string, target string, ip string, id uint64) {
	credential := common.NewCredential(
		sid,
		skey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyRecordRequest()

	////////////////
	request.Domain = common.StringPtr(domain)
	request.SubDomain = common.StringPtr(target)
	request.RecordType = common.StringPtr("A")
	request.RecordLine = common.StringPtr("默认")
	request.Value = common.StringPtr(ip)
	request.RecordId = common.Uint64Ptr(id)
	// 返回的resp是一个ModifyRecordResponse的实例，与请求对象对应
	response, err := client.ModifyRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	log.Print("修改完成：%s", response.ToJsonString())
}
