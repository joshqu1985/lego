package container

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(1024 * 10)

	bf.Add([]byte("hello"))
	ok := bf.Exist([]byte("hello"))
	if !ok {
		t.Log("bloom filter false")
		t.FailNow()
	}
	ok = bf.Exist([]byte("world"))
	if ok {
		t.Log("bloom filter false")
		t.FailNow()
	}
}

func BenchmarkBloomFilter(b *testing.B) {
	var lines []string
	for len(lines) < b.N {
		lines = append(lines, generate(10))
	}

	bf := NewBloomFilter(uint64(b.N))
	count := 0
	b.ResetTimer()

	for l := 0; l < b.N; l++ {
		bf.Add([]byte(lines[l]))
		if !bf.Exist([]byte(lines[l])) {
			count++
		}
	}

	b.StopTimer()
	fmt.Println("----------", count)
}

var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generate(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	return string(b)
}
