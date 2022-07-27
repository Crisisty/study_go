package math

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

type FDIV struct{ base.NoOperandsInstruction }

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

type IDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushInt(result)
}

type LDIV struct{ base.NoOperandsInstruction }

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushLong(result)
}
