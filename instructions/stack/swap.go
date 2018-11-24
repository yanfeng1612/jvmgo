package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	slot0 := frame.OperandStack().PopSlot()
	slot1 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot0)
	frame.OperandStack().PushSlot(slot1)
}
