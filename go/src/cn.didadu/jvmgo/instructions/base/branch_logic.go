package base

import "cn.didadu/jvmgo/rtdata"

func Branch(frame *rtdata.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}