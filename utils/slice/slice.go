package slice

import (
	"bytes"
	"fmt"
	"math"
)

func Split[T any](list []T, size int) [][]T {
	blocks := math.Ceil(float64(len(list)) / float64(size))
	values := make([][]T, 0)
	for i := 0; i < int(blocks); i++ {
		seg := make([]T, size)
		if i == int(blocks)-1 {
			seg = list[i*size:]
		} else {
			seg = list[i*size : i*size+size]
		}
		values = append(values, seg)
	}
	return values
}

func Contain[T ~string | int | int64](list []T, one T) bool {
	for _, each := range list {
		if one == each {
			return true
		}
	}
	return false
}

func Join[T ~string | int | int64](elems []T, sep string) string {
	var buffer bytes.Buffer
	for i, elem := range elems {
		if i != len(elems)-1 {
			buffer.WriteString(fmt.Sprintf("%v%s", elem, sep))
		} else {
			buffer.WriteString(fmt.Sprintf("%v", elem))
		}
	}
	return buffer.String()
}
