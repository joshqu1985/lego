package container

import (
	"github.com/spaolacci/murmur3"
)

const BLOOM_HASH_K = 10

type BloomFilter struct {
	bm   *BitMap
	size uint64
}

/*
size计算:

	  当前 k = 10
		m: size 内存大小、多少个bit位 n: 能容纳的去重对象数量

		假阳率小于0.0012时 m / n = 14
		  当申请size = 256M(21.5亿)时 能容纳 21.5亿 / 14 = 1.53亿
		假阳率小于0.00013时 m / n = 19
		  当申请size = 256M(21.5亿)时 能容纳 21.5亿 / 19 = 1.13亿
*/
func NewBloomFilter(size uint64) *BloomFilter {
	return &BloomFilter{
		bm:   NewBitMap(size),
		size: size,
	}
}

func (bf *BloomFilter) Add(data []byte) {
	positions := bf.hash(data)
	for _, position := range positions {
		_ = bf.bm.Set(position)
	}
}

func (bf *BloomFilter) Exist(data []byte) bool {
	positions := bf.hash(data)
	for _, position := range positions {
		if !bf.bm.IsSet(position) {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) hash(data []byte) []uint64 {
	positions := make([]uint64, BLOOM_HASH_K)
	for i := 0; i < int(BLOOM_HASH_K); i++ {
		positions[i] = murmur3.Sum64(append(data, byte(i))) % bf.size
	}

	return positions
}
