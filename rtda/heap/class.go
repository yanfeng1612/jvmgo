package heap

import (
	"jvmgo/classfile"
)

type Class struct {
	accessFlags        uint16
	name               string
	superClassName     string
	interfaceNames     []string
	constantPool       *ConstantPool
	fields             []*Field
	methods            []*Method
	loader             *ClassLoader
	superClass         *Class
	interfaces         []*Class
	instantceSlotCount uint
	staticSlotCount    uint
	staticVars         *Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	// class.constantPool =
}

func (self *Class) IsPublic() bool {
	return self.accessFlags&ACC_PUBLIC != 0
}
