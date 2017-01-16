package main

import (
	"cn.didadu/jvmgo/cmd"
	"fmt"
	"cn.didadu/jvmgo/classpath"
	"strings"
	"cn.didadu/jvmgo/classfile"
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
	// 获取Classpath
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.Class, cmd.Args)

	// 将.替换成/(java.lang.String -> java/lang/String)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	// 读取class
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)
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