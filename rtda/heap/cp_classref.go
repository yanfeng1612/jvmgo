package heap

import (
	"jvmgo/classfile"
)

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classinfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classinfo.Name()
	return ref
}
