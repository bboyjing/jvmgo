package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// multianewarray指令结构体
type MULTI_ANEW_ARRAY struct {
	// 第一个操作数
	index      uint16
	// 第二个操作数
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 获取数组类符号引用(这里获取的直接就是数组类符号引用，所以下面可以直接解析)
	classRef := cp.GetConstant(uint(self.index)).(*heap.ClassRef)
	// 解析数组类符号引用
	arrClass := classRef.ResolvedClass()

	// 获取操作数栈
	stack := frame.OperandStack()
	// 获取每一维度数组的长度
	counts := popAndCheckCounts(stack, int(self.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

// 从操作数栈中弹出n个int值
func popAndCheckCounts(stack *rtdata.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		// 弹出每一维度数组的长度
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

// 创建多维数组
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}