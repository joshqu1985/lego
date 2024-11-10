package container

import (
	"testing"
)

func TestPriority(t *testing.T) {
	priority := NewPriority()

	priority.Put(&TestData{Val: 2}, 2)
	priority.Put(&TestData{Val: 1}, 1)
	priority.Put(&TestData{Val: 3}, 3)

	if priority.Size() != 3 {
		t.Log("priority size error")
		t.FailNow()
	}

	v1, err := priority.Get()
	if err != nil {
		t.Log("priority get failed:", err)
		t.FailNow()
	}
	t.Log(v1.(*TestData).Val)

	v2, err := priority.Get()
	if err != nil {
		t.Log("priority get failed:", err)
		t.FailNow()
	}
	t.Log(v2.(*TestData).Val)

	v3, err := priority.Get()
	if err != nil {
		t.Log("priority get failed:", err)
		t.FailNow()
	}
	t.Log(v3.(*TestData).Val)
}

func BenchmarkPriority(b *testing.B) {
	priority := NewPriority()

	b.Run("priority put", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i <= b.N; i++ {
			priority.Put(&TestData{Val: i}, int64(i))
		}
	})
}
