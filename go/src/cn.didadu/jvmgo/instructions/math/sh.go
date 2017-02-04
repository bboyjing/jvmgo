package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// int左移结构体(<<)
type ISHL struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素，指定左移位数
	v2 := stack.PopInt()
	// 弹出栈顶元素，作为操作数
	v1 := stack.PopInt()
	/*
		因为int变量只有32位，所以取v2的前5位就足够表示位移位数了
		0x1f为31，二进制00011111，正好和v2进行&操作
	 */
	s := uint32(v2) & 0x1f
	// Go位移操作符右侧必须是无符号数，所以上面对v2进行了uint32转换
	result := v1 << s
	// 将移位操作结果入栈
	stack.PushInt(result)
}

// long左移结构体(<<)
type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素，指定左移位数
	v2 := stack.PopInt()
	// 弹出栈顶元素，作为操作数
	v1 := stack.PopLong()
	// long变量64位，和00111111进行&操作，足够表示位移位数
	s := uint32(v2) & 0x3f
	result := v1 << s
	// 将移位操作结果入栈
	stack.PushLong(result)
}

// int算数右移结构体(>>)
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素，指定右移位数
	v2 := stack.PopInt()
	// 弹出栈顶元素，作为操作数
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// 算数右移操作
	result := v1 >> s
	// 将移位操作结果入栈
	stack.PushInt(result)
}

// int逻辑右移结构体(>>>)
type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶元素，指定右移位数
	v2 := stack.PopInt()
	// 弹出栈顶元素，作为操作数
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// Go没有>>>运算符，需要先把v1转成无符号整数，移位操作后再转回有符号整数
	result := int32(uint32(v1) >> s)
	// 将移位操作结果入栈
	stack.PushInt(result)
}

// long算数右移结构体(>>)
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// long逻辑右移结构体(>>>)
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}