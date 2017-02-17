package instructions

import (
	"cn.didadu/jvmgo/rtdata"
	"fmt"
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata/heap"
)

// 传入入口方法，初始化线程
func Interpret(method *heap.Method) {
	// 创建一个Thread实例
	thread := rtdata.NewThread()
	// 创建栈帧
	frame := thread.NewFrame(method)
	// 栈帧推入虚拟栈
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtdata.Frame) {
	if r := recover(); r != nil {
		/*fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)*/
	}
}

func loop(thread *rtdata.Thread, bytecode []byte) {
	// 获取虚拟机栈栈顶指针
	frame := thread.PopFrame()
	// 实例化BytecodeReader
	reader := &base.BytecodeReader{}
	for {
		// 获取下个pc寄存器地址
		pc := frame.NextPC()
		// 设置线程pc寄存器
		thread.SetPC(pc)
		// 重置
		reader.Reset(bytecode, pc)
		// 获取操作码
		opcode := reader.ReadUint8()
		// 根据操作码常见指令，参照factory.go文件
		inst := NewInstruction(opcode)
		inst.FetchOperands(reader)
		// 设置下一个指令起始地址
		frame.SetNextPC(reader.PC())

		// 执行指令
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}