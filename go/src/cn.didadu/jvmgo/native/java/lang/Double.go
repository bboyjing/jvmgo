package lang

import (
	"cn.didadu/jvmgo/native"
	"cn.didadu/jvmgo/rtdata"
	"math"
)

const jlDouble = "java/lang/Double"

func init() {
	native.Register(jlDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register(jlDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtdata.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value) // todo
	frame.OperandStack().PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
// (J)D
func longBitsToDouble(frame *rtdata.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits)) // todo
	frame.OperandStack().PushDouble(value)
}