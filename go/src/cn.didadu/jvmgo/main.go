package main

import (
	"cn.didadu/jvmgo/cmd"
	"fmt"
	"cn.didadu/jvmgo/classpath"
	//"strings"
	"cn.didadu/jvmgo/classfile"
	"cn.didadu/jvmgo/rtdata"
	"strings"
	"cn.didadu/jvmgo/instructions"
	"cn.didadu/jvmgo/rtdata/heap"
)

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(command)
	}
}

// 模拟启动JVM
func startJVM(cmd *cmd.Cmd) {
	/*// 获取Classpath
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.Class, cmd.Args)

	// 将.替换成/(java.lang.String -> java/lang/String)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	// 读取class
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)*/


	/*frame := rtdata.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())*/

	// 获取Classpath
	/*cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	// 将.替换成/(java.lang.String -> java/lang/String)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	// 加载类
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		instructions.Interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class)
	}*/

	/*// 获取Classpath
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	// 创ClassLoader实例
	classloader := heap.NewClassLoader(cp)

	// class权限定名，将.替换成/(java.lang.String -> java/lang/String)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	// 加载主类
	mainClass := classloader.LoadClass(className);
	// 获取主类的main()方法
	mainMethod := mainClass.GetMainMethod()
	if mainClass != nil {
		instructions.Interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class)
	}*/

	// 获取Classpath
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	classLoader := heap.NewClassLoader(cp, cmd.VerboseClassFlag)

	// class权限定名，将.替换成/(java.lang.String -> java/lang/String)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	// 加载主类
	mainClass := classLoader.LoadClass(className)
	// 获取主类的main()方法
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		instructions.Interpret(mainMethod, cmd.VerboseInstFlag, cmd.Args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class)
	}
}

// 获取main()方法
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	// 读取class
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

// 打印出class文件的重要信息
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}

// 测试局部变量表
func testLocalVars(vars rtdata.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

// 测试操作数栈
func testOperandStack(ops *rtdata.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}