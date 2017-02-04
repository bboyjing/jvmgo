package rtdata

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
	// pc寄存器
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