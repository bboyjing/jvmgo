package rtdata

type Frame struct {
	// 实现链表数据结构
	lower        *Frame
	// 局部变量
	localVars    LocalVars
	// 操作数栈指针
	operandStack *OperandStack
}

// 实例化栈帧
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// localVars和operandStackGet方法
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}