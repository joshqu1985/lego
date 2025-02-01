package container

import (
	"fmt"
)

type BitMap struct {
	bits []byte
	size uint64 // 比特位数量
}

func NewBitMap(size uint64) *BitMap {
	return &BitMap{
		bits: make([]byte, (size+7)/8),
		size: size,
	}
}

func (this *BitMap) Set(i uint64) error {
	if i > this.size {
		return fmt.Errorf("out of range")
	}

	this.bits[i/8] |= 1 << (i % 8)
	return nil
}

func (this *BitMap) Clear(i uint64) error {
	if i > this.size {
		return fmt.Errorf("out of range")
	}

	this.bits[i/8] &= ^(1 << (i % 8))
	return nil
}

func (this *BitMap) IsSet(i uint64) bool {
	if i > this.size {
		return false
	}

	return (this.bits[i/8] & (1 << (i % 8))) != 0
}
