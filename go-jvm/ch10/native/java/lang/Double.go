package lang

import (
	"jvmgo/ch10/native"
	"jvmgo/ch10/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}

func longBitsToDouble(frame *rtda.Frame) {
	value := frame.LocalVars().GetLong(0)
	bits := math.Float64frombits(uint64(value))
	frame.OperandStack().PushDouble(bits)
}
