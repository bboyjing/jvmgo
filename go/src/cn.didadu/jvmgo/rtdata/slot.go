package rtdata

import "cn.didadu/jvmgo/rtdata/heap"

type Slot struct {
	// 整数
	num int32
	// 引用
	ref *heap.Object
}
