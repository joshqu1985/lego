package container

import (
	"fmt"
	"testing"
)

func TestBitmap(t *testing.T) {
	bm := NewBitMap(32)

	bm.Set(11)
	ok := bm.IsSet(11)
	if !ok {
		t.Log("get bit false")
		t.FailNow()
	}
	fmt.Printf("%08b\n", bm.bits)

	bm.Clear(11)
	ok = bm.IsSet(11)
	if ok {
		t.Log("clear bit false")
		t.FailNow()
	}
	fmt.Printf("%08b\n", bm.bits)
}
