package rtdata

// 虚拟机栈结构体
type Stack struct {
	// 栈的容量(最多可以容纳多少帧)
	maxSize uint
	// 当前栈的大小
	size    uint
	// 栈顶指针
	_top    *Frame
}

// 初始化虚拟机栈
func newStack(maxSize uint) *Stack {
	return &Stack{maxSize:maxSize}
}

// 返回栈顶指针
func (self *Stack) top() *Frame {
	if self._top == nil {
		// 若栈时空的，肯定是有bug，需要panic
		panic("jvm stack is empty!")
	}

	return self._top
}

// 入栈
func (self *Stack) push(frame *Frame) {
	// 超过栈最大深度报错
	if self.size > self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	// 将栈顶指针下移
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

// 出栈
func (self *Stack) pop() *Frame {
	// 若栈顶指针为空，报错
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	// 弹出栈顶指针
	self._top = top.lower
	// 弹出指针的lower指向不再有意义
	top.lower = nil
	self.size--
	return top;
}

// 判断栈顶指针是否为空
func (self *Stack) isEmpty() bool {
	return self._top == nil
}