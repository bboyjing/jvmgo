package base

import (
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

func InvokeMethod(invokerFrame *rtdata.Frame, method *heap.Method) {
	// 获取当前线程
	thread := invokerFrame.Thread()
	// 创建新的帧
	newFrame := thread.NewFrame(method)
	// 将新创建的帧推入Java虚拟机栈
	thread.PushFrame(newFrame)

	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}
