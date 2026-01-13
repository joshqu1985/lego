package routine

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// Safe 对函数调用包装 防止panic.
func Safe(fn func()) (err error) {
	defer func() {
		rec := recover()
		if rec == nil {
			return
		}
		glog.Errorf("[PANIC] recover stack:%s err:%v\n", string(debug.Stack()), rec)
		if _, ok := rec.(error); ok {
			err, _ = rec.(error)
		} else {
			err = fmt.Errorf("%+v", rec)
		}
		err = errors.Wrap(err, FunctionName(fn))
	}()

	fn()

	return err
}

func FunctionName(v any) string {
	return runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
}
