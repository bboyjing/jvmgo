package cmd

import (
	"flag"
	"fmt"
	"os"
)

/*
	定义结构体，用于存储命令参数。
	java [-options] class [args...]
*/
type Cmd struct {
	HelpFlag         bool
	VersionFlag      bool
	VerboseClassFlag bool
	VerboseInstFlag  bool
	CpOption         string
	Class            string
	Args             []string
	XjreOption       string
}

/*
	ParseCmd()方法返回值为*Cmd，是指向Cmd的值的指针。
	Go语言中，常量、函数的首字母大写表示对外公开的相当于Java的public，小写表示私有的相当于Java的private。
*/
func ParseCmd() *Cmd {
	//声明cmd为指向空的Cmd对象的指针
	cmd := &Cmd{}

	/*
		Usage是一个函数，默认输出所有定义了的命令行参数和帮助信息
		一般，当命令行参数解析出错时，该函数会被调用
		这里我们指定了自己的Usage函数，即printUsage()
	*/
	flag.Usage = PrintUsage
	//flag.XxxVar()，将flag绑定到第一个参数指定的指针上，并设置默认值和提示信息
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.VerboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.VerboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.VerboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	//在所有的flag定义完成之后，可以通过调用flag.Parse()进行解析。
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

// 提示输入正确的命令格式
func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
