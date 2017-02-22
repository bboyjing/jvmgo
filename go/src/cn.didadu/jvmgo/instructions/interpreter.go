package instructions

import (
	"cn.didadu/jvmgo/rtdata"
	"fmt"
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata/heap"
)

// 传入入口方法，初始化线程
func Interpret(method *heap.Method, logInst bool) {
	// 创建一个Thread实例
	thread := rtdata.NewThread()
	// 创建栈帧
	frame := thread.NewFrame(method)
	// 栈帧推入虚拟栈
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtdata.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

// 打印Java虚拟机栈错误信息
func logFrames(thread *rtdata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(thread *rtdata.Thread, logInst bool) {
	// 实例化BytecodeReader
	reader := &base.BytecodeReader{}
	// 遍历Java虚拟机栈，直到栈顶指针为空
	for {
		// 获取Java虚拟机栈当前帧(栈顶指针)
		frame := thread.CurrentFrame()
		// 获取下个pc寄存器地址(也就是下一条将要执行的指令所在的位置)
		pc := frame.NextPC()
		// 设置线程pc寄存器为当前帧的下一个指令地址
		thread.SetPC(pc)

		// 重置reader
		reader.Reset(frame.Method().Code(), pc)
		// 获取指令操作码
		opcode := reader.ReadUint8()
		// 根据操作码常见指令，参照factory.go文件
		inst := NewInstruction(opcode)
		// 读取操作数
		inst.FetchOperands(reader)
		// 设置下一个指令起始地址(下一条指令操作码在字节码中的位置)
		frame.SetNextPC(reader.PC())

		// 是否将执行信息打印到控制台
		if logInst {
			logInstruction(frame, inst)
		}

		// 执行指令
		inst.Execute(frame)
		// 若栈顶指针为空，表示线程执行结束
		if thread.IsStackEmpty() {
			break
		}
	}
}

// 打印指令执行信息
func logInstruction(frame *rtdata.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}