package control

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// tableswitch字节码内容
/*
	tableswitch
	<0-3 byte pad>
	defaultbyte1
	defaultbyte2
	defaultbyte3
	defaultbyte4
	lowbyte1
	lowbyte2
	lowbyte3
	lowbyte4
	highbyte1
	highbyte2
	highbyte3
	highbyte4
	jump offsets...
*/

type TABLE_SWITCH struct {
	// 默认情况下执行跳转所需的字节码偏移量
	defaultOffset int32
	// low和high记录case的取值范围
	low           int32
	high          int32
	// 索引表，存放high-low+1个int值
	jumpOffsets   []int32
}

// 读取操作数
func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 跳过Padding字节
	reader.SkipPadding()
	// 读取4个字节default跳转字节码偏移量
	self.defaultOffset = reader.ReadInt32()
	// 读取4个字节case范围最小值
	self.low = reader.ReadInt32()
	// 读取4个字节case范围最大值
	self.high = reader.ReadInt32()
	// 计算索引表数量
	jumpOffsetsCount := self.high - self.low + 1
	// 填充索引表
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtdata.Frame) {
	// 弹出栈顶int型变量，作为需要匹配case分支的值
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		// 若在case范围内，则通过index获取索引表中的值
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		// 若不再case范围内，给出默认值
		offset = int(self.defaultOffset)
	}
	// 跳转到指定offset
	base.Branch(frame, offset)
}