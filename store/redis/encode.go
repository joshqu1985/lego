package redis

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/joshqu1985/lego/encoding/json"
)

func Format(v any) ([]byte, error) {
	switch v := v.(type) {
	case nil:
		return StringToBytes(""), nil
	case string:
		return StringToBytes(v), nil
	case *string:
		return StringToBytes(*v), nil
	case []byte:
		return v, nil
	case int:
		s := strconv.FormatInt(int64(v), 10)

		return StringToBytes(s), nil
	case *int:
		s := strconv.FormatInt(int64(*v), 10)

		return StringToBytes(s), nil
	case int8:
		s := strconv.FormatInt(int64(v), 10)

		return StringToBytes(s), nil
	case *int8:
		s := strconv.FormatInt(int64(*v), 10)

		return StringToBytes(s), nil
	case int16:
		s := strconv.FormatInt(int64(v), 10)

		return StringToBytes(s), nil
	case *int16:
		s := strconv.FormatInt(int64(*v), 10)

		return StringToBytes(s), nil
	case int32:
		s := strconv.FormatInt(int64(v), 10)

		return StringToBytes(s), nil
	case *int32:
		s := strconv.FormatInt(int64(*v), 10)

		return StringToBytes(s), nil
	case int64:
		s := strconv.FormatInt(v, 10)

		return StringToBytes(s), nil
	case *int64:
		s := strconv.FormatInt(*v, 10)

		return StringToBytes(s), nil
	case uint:
		s := strconv.FormatUint(uint64(v), 10)

		return StringToBytes(s), nil
	case *uint:
		s := strconv.FormatUint(uint64(*v), 10)

		return StringToBytes(s), nil
	case uint8:
		s := strconv.FormatUint(uint64(v), 10)

		return StringToBytes(s), nil
	case *uint8:
		s := strconv.FormatUint(uint64(*v), 10)

		return StringToBytes(s), nil
	case uint16:
		s := strconv.FormatUint(uint64(v), 10)

		return StringToBytes(s), nil
	case *uint16:
		s := strconv.FormatUint(uint64(*v), 10)

		return StringToBytes(s), nil
	case uint32:
		s := strconv.FormatUint(uint64(v), 10)

		return StringToBytes(s), nil
	case *uint32:
		s := strconv.FormatUint(uint64(*v), 10)

		return StringToBytes(s), nil
	case uint64:
		s := strconv.FormatUint(v, 10)

		return StringToBytes(s), nil
	case *uint64:
		s := strconv.FormatUint(*v, 10)

		return StringToBytes(s), nil
	case float32:
		s := strconv.FormatFloat(float64(v), 'f', -1, 64)

		return StringToBytes(s), nil
	case *float32:
		s := strconv.FormatFloat(float64(*v), 'f', -1, 64)

		return StringToBytes(s), nil
	case float64:
		s := strconv.FormatFloat(float64(v), 'f', -1, 64)

		return StringToBytes(s), nil
	case *float64:
		s := strconv.FormatFloat(float64(*v), 'f', -1, 64)

		return StringToBytes(s), nil
	case bool:
		var s string
		if v {
			s = strconv.FormatInt(int64(1), 10)
		} else {
			s = strconv.FormatInt(int64(0), 10)
		}

		return StringToBytes(s), nil
	case *bool:
		var s string
		if *v {
			s = strconv.FormatInt(int64(1), 10)
		} else {
			s = strconv.FormatInt(int64(0), 10)
		}

		return StringToBytes(s), nil
	case time.Time:
		s := strconv.FormatInt(v.UnixNano(), 10)

		return StringToBytes(s), nil
	case time.Duration:
		s := strconv.FormatInt(v.Nanoseconds(), 10)

		return StringToBytes(s), nil
	case net.IP:
		return []byte(v), nil
	default:
		return json.Marshal(v)
	}
}

func Scan(b []byte, v any) error {
	switch v := v.(type) {
	case nil:
		return errors.New("redis: Scan(nil)")
	case *string:
		*v = BytesToString(b)

		return nil
	case *[]byte:
		*v = b

		return nil
	case *int:
		var err error
		*v, err = strconv.Atoi(BytesToString(b))

		return err
	case *int8:
		n, err := strconv.ParseInt(BytesToString(b), 10, 8)
		if err != nil {
			return err
		}
		*v = int8(n)

		return nil
	case *int16:
		n, err := strconv.ParseInt(BytesToString(b), 10, 16)
		if err != nil {
			return err
		}
		*v = int16(n)

		return nil
	case *int32:
		n, err := strconv.ParseInt(BytesToString(b), 10, 32)
		if err != nil {
			return err
		}
		*v = int32(n)

		return nil
	case *int64:
		n, err := strconv.ParseInt(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = n

		return nil
	case *uint:
		n, err := strconv.ParseUint(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = uint(n)

		return nil
	case *uint8:
		n, err := strconv.ParseUint(BytesToString(b), 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(n)

		return nil
	case *uint16:
		n, err := strconv.ParseUint(BytesToString(b), 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(n)

		return nil
	case *uint32:
		n, err := strconv.ParseUint(BytesToString(b), 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(n)

		return nil
	case *uint64:
		n, err := strconv.ParseUint(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = n

		return nil
	case *float32:
		n, err := strconv.ParseFloat(BytesToString(b), 32)
		if err != nil {
			return err
		}
		*v = float32(n)

		return err
	case *float64:
		var err error
		*v, err = strconv.ParseFloat(BytesToString(b), 64)

		return err
	case *bool:
		*v = len(b) == 1 && b[0] == '1'

		return nil
	case *time.Time:
		n, err := strconv.ParseInt(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = time.Unix(n/1000000000, n%1000000000*int64(time.Nanosecond))

		return nil
	case *time.Duration:
		n, err := strconv.ParseInt(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = time.Duration(n)

		return nil
	case *net.IP:
		*v = b

		return nil
	default:
		return json.Unmarshal(b, v)
	}
}

func ScanSlice(data []string, slice any) error {
	v := reflect.ValueOf(slice)
	if !v.IsValid() {
		return errors.New("redis: ScanSlice(nil)")
	}
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("redis: ScanSlice(non-pointer %T)", slice)
	}
	v = v.Elem()
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("redis: ScanSlice(non-slice %T)", slice)
	}

	next := makeSliceNextElemFunc(v)
	for i, s := range data {
		elem := next()
		if err := Scan([]byte(s), elem.Addr().Interface()); err != nil {
			err = fmt.Errorf("redis: ScanSlice index=%d value=%q failed: %w", i, s, err)

			return err
		}
	}

	return nil
}

func makeSliceNextElemFunc(v reflect.Value) func() reflect.Value {
	elemType := v.Type().Elem()

	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()

		return func() reflect.Value {
			if v.Len() < v.Cap() {
				v.Set(v.Slice(0, v.Len()+1))
				elem := v.Index(v.Len() - 1)
				if elem.IsNil() {
					elem.Set(reflect.New(elemType))
				}

				return elem.Elem()
			}

			elem := reflect.New(elemType)
			v.Set(reflect.Append(v, elem))

			return elem.Elem()
		}
	}

	zero := reflect.Zero(elemType)

	return func() reflect.Value {
		if v.Len() < v.Cap() {
			v.Set(v.Slice(0, v.Len()+1))

			return v.Index(v.Len() - 1)
		}

		v.Set(reflect.Append(v, zero))

		return v.Index(v.Len() - 1)
	}
}

// BytesToString converts byte slice to string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to byte slice.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
