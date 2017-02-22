package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// invokespecial指令结构体
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	// 获取当前class
	currentClass := frame.Method().Class()
	// 获取当前class的运行时常量池
	cp := currentClass.ConstantPool()
	// 通过索引获取方法符号引用
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	// 解析需要调用的方法所属的类
	resolvedClass := methodRef.ResolvedClass()
	// 解析需要调用的方法
	resolvedMethod := methodRef.ResolvedMethod()

	// 如果方法是构造函数(init(...)方法)，则声明该方法的类必须是之前通过方法符号引用解析出来的类
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	// 判断方法是否是static
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取距离操作数栈顶n个slot的引用变量，也就是this引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	// 如果this引用为空，抛出NullPointerException异常
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// 确保protected方法只能被申明该方法的类或者子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	// 如果调用的是超类的方法，需要额外的过程查找最终调用的方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {

		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	// 如果被调用的方法为inl或者是抽象方法，报错
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}
