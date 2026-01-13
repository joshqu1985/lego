package container

import "errors"

type BitMap struct {
	bits []byte
	size uint64 // 比特位数量
}

var ErrOutOfRange = errors.New("out of range")

func NewBitMap(size uint64) *BitMap {
	return &BitMap{
		bits: make([]byte, (size+7)/8),
		size: size,
	}
}

func (bm *BitMap) Set(i uint64) error {
	if i > bm.size {
		return ErrOutOfRange
	}

	bm.bits[i/8] |= 1 << (i % 8)

	return nil
}

func (bm *BitMap) Clear(i uint64) error {
	if i > bm.size {
		return ErrOutOfRange
	}

	bm.bits[i/8] &= ^(1 << (i % 8))

	return nil
}

func (bm *BitMap) IsSet(i uint64) bool {
	if i > bm.size {
		return false
	}

	return bm.bits[i/8]&(1<<(i%8)) != 0
}
