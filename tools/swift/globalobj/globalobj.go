/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-03 15:17:01
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package globalobj

import (
	"encoding/json"
	"io/ioutil"
	"shenzhou/tools/swift/siface"
)

type GlobalObj struct {
	// server对象
	TcpServer siface.IServer
	// IP
	Host string
	// PORT
	TcpPort int
	// 服务器名称
	Name string
	// 当前服务版本
	Version string
	// 最大链接数
	MaxConn int
	// 数据包的最大值
	MaxPackageSize uint32
	// 工作池的数量
	WorkPoolSize uint32
	// 工作池中任务队列允许的最大数量
	MaxWorkerTaskLen uint32
}

var GlobalConfig *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("./config/swift.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalConfig)
	if err != nil {
		panic(err)
	}

}

// init方法

func init() {
	GlobalConfig = &GlobalObj{
		Name:             "SwiftServerApp",
		Version:          "V0.0.7",
		TcpPort:          9001,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   1024,
		WorkPoolSize:     10,
		MaxWorkerTaskLen: 1024,
	}
	// 加载配置中用户自定义的属性
	GlobalConfig.Reload()
}
