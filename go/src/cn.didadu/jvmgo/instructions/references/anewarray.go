package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// anewarray指令结构体
type ANEW_ARRAY struct{
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtdata.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 获取类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类符号引用
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	// 从操作数栈中弹出元素，作为数组长度
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// 加载数组类
	arrClass := componentClass.ArrayClass()
	// 创建数组
	arr := arrClass.NewArray(uint(count))
	// 将数组指针推入操作数栈
	stack.PushRef(arr)
}