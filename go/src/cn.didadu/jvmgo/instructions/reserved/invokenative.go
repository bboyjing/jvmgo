package reserved

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/native"
	"cn.didadu/jvmgo/rtdata"
	_ "cn.didadu/jvmgo/native/java/lang"
	_ "cn.didadu/jvmgo/native/sun/misc"
)

// 0xFE指令结构体
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtdata.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	// 从本地方法注册表中查找方法
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	// 执行本地方法
	nativeMethod(frame)
}
