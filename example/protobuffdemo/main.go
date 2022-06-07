/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-07 21:55:49
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package main

import (
	"fmt"
	"shenzhou/example/protobuffdemo/pb/person"

	"google.golang.org/protobuf/proto"
)

func main() {
	sendPerson := &person.Person{
		Name:   "ding",
		Age:    14,
		Emails: []string{"111@qq.com", "2222@qq.com"},
		Phones: []*person.PhoneNumber{
			&person.PhoneNumber{
				Number: "1333333333",
				Type:   person.PhoneType_MOBILE,
			},
			&person.PhoneNumber{
				Number: "2626262",
				Type:   person.PhoneType_HOME,
			},
			&person.PhoneNumber{
				Number: "22222222",
				Type:   person.PhoneType_WORK,
			},
		},
	}
	// 将person对象序列化
	// 编码
	data, err := proto.Marshal(sendPerson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sendPerson)
	// 解码
	recvPerson := &person.Person{}
	err = proto.Unmarshal(data, recvPerson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(recvPerson)
}
