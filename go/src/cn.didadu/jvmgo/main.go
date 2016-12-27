package main

import (
	"cn.didadu/jvmgo/cmd"
	"fmt"
	"cn.didadu/jvmgo/classpath"
	"strings"
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
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}