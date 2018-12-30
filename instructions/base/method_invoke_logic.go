package base

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
}
