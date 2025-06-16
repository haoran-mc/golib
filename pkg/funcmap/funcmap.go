package funcmap

import (
	"errors"
	"reflect"
)

var (
	ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

type Funcs map[string]reflect.Value

func New() Funcs {
	return make(Funcs, 2)
}

func (f Funcs) Bind(name string, fn interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()

	v := reflect.ValueOf(fn)

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	v.Type().NumIn() // 没有接收返回值，这条语句判断 fn 是否是函数类型
	f[name] = v
	return
}

func (f Funcs) Has(name string) bool {
	_, ok := f[name]
	return ok
}

func (f Funcs) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	if _, ok := f[name]; !ok { // 是否存在该 name
		err = errors.New(name + " does not exist.")
		return
	}
	if len(params) != f[name].Type().NumIn() { // 参数数目是否正确
		err = ErrParamsNotAdapted
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f[name].Call(in) // Call
	return
}
