package base

import (
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// 类初始化方法
func InitClass(thread *rtdata.Thread, class *heap.Class) {
	// 设置类初始化开始标志
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtdata.Thread, class *heap.Class) {
	// 获取<clinit>方法
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// 创建新的帧
		newFrame := thread.NewFrame(clinit)
		// 将新建的帧推入Java虚拟机栈
		thread.PushFrame(newFrame)
	}
}

// 初始化超类
func initSuperClass(thread *rtdata.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			/*
				递归调用InitClass()方法
				这样就可以保证超类的初始化方法对应的帧在子类上面，
				使超类初始化方法先于子类
			 */
			InitClass(thread, superClass)
		}
	}
}