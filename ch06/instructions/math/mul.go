package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type DMUL struct{ base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

type FMUL struct{ base.NoOperandsInstruction }

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

type IMUL struct{ base.NoOperandsInstruction }

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

type LMUL struct{ base.NoOperandsInstruction }

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
