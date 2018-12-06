package main

import (
	"jvmgo/classfile"
	"jvmgo/rtda"
)

func interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	// bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		panic(r)
	}
}
