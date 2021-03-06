package rtdata

import "cn.didadu/jvmgo/rtdata/heap"

type Frame struct {
	// 实现链表数据结构
	lower        *Frame
	// 局部变量
	localVars    LocalVars
	// 操作数栈指针
	operandStack *OperandStack
	// 线程指针
	thread       *Thread
	// 下个pc寄存器地址(为了实现跳转)
	nextPC       int
	// 方法区的方法信息指针
	method       *heap.Method
}

// 实例化栈帧
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// 实例化栈帧
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// localVars和operandStackGet方法
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

// 重置nextPC
func (self *Frame) RevertNextPC() {
	// 因为Thread的pc寄存器字段始终指向当前指令地址
	self.nextPC = self.thread.pc
}