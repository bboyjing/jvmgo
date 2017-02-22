package control

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)


// return指令结构体
type RETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtdata.Frame) {
	// 无返回值，直将当前帧从Java虚拟机栈中弹出
	frame.Thread().PopFrame()
}

// areturn指令结构体
type ARETURN struct{ base.NoOperandsInstruction }

func (self *ARETURN) Execute(frame *rtdata.Frame) {
	// 获取当前线程
	thread := frame.Thread()
	// 弹出当前帧
	currentFrame := thread.PopFrame()
	// 获取前一帧，也就是调用方栈帧
	invokerFrame := thread.TopFrame()
	// 弹出当前帧的操作数栈顶引用变量的值
	ref := currentFrame.OperandStack().PopRef()
	// 将返回值推入前一帧的操作数栈顶
	invokerFrame.OperandStack().PushRef(ref)
}

// dreturn指令结构体
type DRETURN struct{ base.NoOperandsInstruction }

func (self *DRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}


// freturn指令结构体
type FRETURN struct{ base.NoOperandsInstruction }

func (self *FRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// ireturn指令结构体
type IRETURN struct{ base.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// lreturn指令结构体
type LRETURN struct{ base.NoOperandsInstruction }

func (self *LRETURN) Execute(frame *rtdata.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}