package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
	"reflect"
)

type ATHROW struct{ base.NoOperandsInstruction }

func (self *ATHROW) Execute(frame *rtdata.Frame) {
	// athrow指令操作数为异常对象，从操作数栈中弹出
	ex := frame.OperandStack().PopRef()
	// 如果引用为空，抛出NullPointerException异常
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		// 若没有找到异常处理代码块
		handleUncaughtException(thread, ex)
	}
}

// 查找并跳转到异常处理代码
func findAndGotoExceptionHandler(thread *rtdata.Thread, ex *heap.Object) bool {
	for {
		// 获取当前帧
		frame := thread.CurrentFrame()
		// 当前正在执行指令的地址为NextPC - 1
		pc := frame.NextPC() - 1

		// 查找异常处理表
		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		// 查找到异常处理代码块地址
		if handlerPC > 0 {
			// 获取当前操作数栈
			stack := frame.OperandStack()
			// 清空操作数栈
			stack.Clear()
			// 将异常对象引用入栈
			stack.PushRef(ex)
			// 设置执行下一条指令的地址为异常处理代码块的起始位置
			frame.SetNextPC(handlerPC)
			return true
		}

		// 遍历java虚拟机栈帧，若当前帧没有找到可用异常处理表，则弹出，执行上一帧
		thread.PopFrame()
		// 如果虚拟机栈空了，则退出循环
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtdata.Thread, ex *heap.Object) {
	// 清空JVM虚拟机栈
	thread.ClearStack()

	// 打印Java虚拟机栈信息
	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}