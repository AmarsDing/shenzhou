/*
 * @Description: ##  描述文件功能  ##
 * @Author: AmarsDing
 * @Date: 2022-06-02 22:13:52
 * @Copyright: 北京迈特力德信息技术有限公司, METLED@2021
 */
package snet

import "shenzhou/tools/swift/siface"

type BaseRouter struct {
}

// 处理业务之前HOOK
func (br *BaseRouter) BeforeHandle(request siface.IRequest) {

}

// 处理业务
func (br *BaseRouter) Handle(request siface.IRequest) {

}

// 处理业务之后HOOK
func (br *BaseRouter) AfterHandle(request siface.IRequest) {

}
