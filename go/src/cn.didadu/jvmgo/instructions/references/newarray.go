package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// newarray指令结构体
type NEW_ARRAY struct {
	atype uint8
}

// atype常量，对应8种基本类型
const (
	AT_BOOLEAN = 4
	AT_CHAR = 5
	AT_FLOAT = 6
	AT_DOUBLE = 7
	AT_BYTE = 8
	AT_SHORT = 9
	AT_INT = 10
	AT_LONG = 11
)

// 读取第一个单字节操作数
func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *rtdata.Frame) {
	// 获取当前帧的操作数栈
	stack := frame.OperandStack()
	// 弹出栈顶元素，作为数组的长度
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// 获取当前类的类加载器
	classLoader := frame.Method().Class().Loader()
	// 加载数组类
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	// 创建数组
	arr := arrClass.NewArray(uint(count))
	// 将创建的数组指针推入操作数栈
	stack.PushRef(arr)
}

// 根据atype加载对应的Class
func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}