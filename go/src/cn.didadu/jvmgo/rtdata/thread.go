package rtdata

import "cn.didadu/jvmgo/rtdata/heap"

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/

// Thread结构体
type Thread struct {
	// pc寄存器，存放当前正在执行的Java虚拟机指令的地址
	pc    int
	// 虚拟机栈结构体指针
	stack *Stack
}

// 初始化线程
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// pc寄存器的Get、Set方法
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

// 入栈，调用虚拟机栈的方法
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 出栈，调用虚拟机栈的方法
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 查看栈顶指针，调用虚拟机栈的方法
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

// 为线程创建新的栈帧
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

// 获取栈顶指针
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

// 判断虚拟机栈是否为空
func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}

// 获取完整的Java虚拟机栈
func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}