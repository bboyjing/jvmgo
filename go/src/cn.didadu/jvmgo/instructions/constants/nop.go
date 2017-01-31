package constants

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// 定义NOP结构体，继承NoOperandsInstruction
type NOP struct{ base.NoOperandsInstruction }

// 什么都不执行
func (self *NOP) Execute(frame *rtdata.Frame) {
}
