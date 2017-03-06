package native

import "cn.didadu/jvmgo/rtdata"

// 本地方法
type NativeMethod func(frame *rtdata.Frame)

// 注册表
var registry = map[string]NativeMethod{}

// 注册本地方法
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// 查找本地方法
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtdata.Frame) {
	// do nothing
}

