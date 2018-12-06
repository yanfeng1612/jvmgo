package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operantStack *OperandStack
	thread       *Thread
	nextPC       int
}

func newFrame(thread *Thread, maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operantStack: newOperandStack(maxStack),
	}
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operantStack
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
