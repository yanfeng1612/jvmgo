package heap

import (
	"math"
)

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = bits
}

func (self Slots) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(self[index].num))
}

func (self Slots) SetLong(index uint, val int64) {
	low := uint32(self[index].num)
	high := uint32(self[index].num)

}
