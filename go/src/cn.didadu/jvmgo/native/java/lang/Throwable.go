package lang

import (
	"cn.didadu/jvmgo/native"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
	"fmt"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtdata.Frame) {
	// 获取局部变量表第0位，this引用
	this := frame.LocalVars().GetThis()
	// 将this引用入栈
	frame.OperandStack().PushRef(this)
	// 创建完整的堆栈打印信息
	stes := createStackTraceElements(this, frame.Thread())
	// 将完整的堆栈信息赋值到异常类实例(this)的extra字段中，athrow指令中打印了extra信息
	this.SetExtra(stes)
}

// 该结构体用于记录Java虚拟机栈信息
type StackTraceElement struct {
	// 类所在的文件名
	fileName   string
	// 申明方法的类名
	className  string
	// 方法名
	methodName string
	// 帧正在执行哪行代码
	lineNumber int
}


func createStackTraceElements(tObj *heap.Object, thread *rtdata.Thread) []*StackTraceElement {
	// 由于栈顶两帧正在执行fillInStackTrace(int)和fillInStackTrace()，所以这两帧也需要跳过
	skip := distanceToObject(tObj.Class()) + 2
	// 获取跳过的帧之后的完整的Java虚拟机栈
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// 计算需要跳过的帧(正在执行的异常类的构造函数)
func distanceToObject(class *heap.Class) int {
	distance := 0
	// 如果有超类，跳帧 + 1
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

// 构造记录Java虚拟机栈信息的结构体
func createStackTraceElement(frame *rtdata.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}