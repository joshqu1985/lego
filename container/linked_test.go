package container

import (
	"testing"
)

type TestData struct {
	Val int
}

func TestLinked(t *testing.T) {
	linked := NewLinked()
	linked.Put(&TestData{Val: 1})
	linked.Put(&TestData{Val: 2})
	linked.Put(&TestData{Val: 3})

	if linked.Size() != 3 {
		t.Log("linked size error")
		t.FailNow()
	}

	val, err := linked.Get()
	if err != nil {
		t.Log("linked get failed:", err)
		t.FailNow()
	}
	t.Log(val.(*TestData).Val)
}

func BenchmarkLinked(b *testing.B) {
	linked := NewLinked()

	b.Run("linked put", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i <= b.N; i++ {
			linked.Put(&TestData{Val: i})
		}
	})
}
