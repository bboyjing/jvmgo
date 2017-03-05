package control

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

/*
	lookupswitch
	<0-3 byte pad>
	defaultbyte1
	defaultbyte2
	defaultbyte3
	defaultbyte4
	npairs1
	npairs2
	npairs3
	npairs4
	match-offset pairs...
*/

type LOOKUP_SWITCH struct {
	// 默认情况下执行跳转所需的字节码偏移量
	defaultOffset int32
	// case分支的数量
	npairs        int32
	// 存放跳转指令偏移量
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 跳过Padding字节
	reader.SkipPadding()
	// 读取4个字节default跳转字节码偏移量
	self.defaultOffset = reader.ReadInt32()
	// 读取4个字节case分值数量
	self.npairs = reader.ReadInt32()
	/*
		读取各个case跳转指令偏移量
		matchOffsets中每个case存储8个字节，前4个字节是case的值，后4个字节是跳转偏移量，有点像Map
	 */
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtdata.Frame) {
	// 弹出栈顶int型变量，作为需要匹配case分支的值
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs * 2; i += 2 {
		// 循环判断给定值是否命中case分值
		if self.matchOffsets[i] == key {
			// 若命中，则跳转到指定偏移量
			offset := self.matchOffsets[i + 1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}
