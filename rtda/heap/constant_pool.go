package heap

import (
	"jvmgo/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfcp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfcp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfcp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		}
	}
	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	return self.consts[index]
}

func newConstantPoll(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldInfo)
		case *classfile.ConstantMethodrefInfo:
			methodInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodInfoRef(rtCp, methodInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfaceMethodInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodInfoRef(rtCp, interfaceMethodInfo)
		}
	}
	return rtCp
}
