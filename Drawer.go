package drawer

import (
	"fmt"
	"reflect"
)

// -
var (
	ErrInvalidParameterType = fmt.Errorf("invalid parameter type")
)

// Drawer -
type Drawer interface {
	Pull(interface{}) error
	Push(...interface{})
	Dump() []interface{}
}

type rawDrawer struct {
	rvs []reflect.Value
}

func (s *rawDrawer) Dump() []interface{} {
	ret := []interface{}{}

	for _, i := range s.rvs {
		if i.IsValid() {
			ret = append(ret, i.Interface())
		} else {
			ret = append(ret, nil)
		}
	}

	return ret
}

func (s *rawDrawer) Push(v ...interface{}) {
	for _, i := range v {
		if d, ok := i.(Drawer); ok {
			s.Push(d.Dump()...)
		} else {
			s.rvs = append(s.rvs, reflect.ValueOf(i))
		}
	}
}

func (s *rawDrawer) Pull(r interface{}) error {
	rv1 := reflect.ValueOf(r)
	if rv1.Kind() != reflect.Ptr {
		return ErrInvalidParameterType
	}

	rv2 := rv1.Elem()
	if rv2.Kind() != reflect.Slice {
		return ErrInvalidParameterType
	}

	basetype := rv2.Type().Elem()
	retslice := reflect.MakeSlice(reflect.SliceOf(basetype), 0, 0)

	chk := func(t reflect.Type) bool {
		return basetype == t
	}
	flg := false

	if basetype.Kind() == reflect.Interface {
		chk = func(t reflect.Type) bool {
			return t.ConvertibleTo(basetype)
		}
		flg = basetype.Name() == "" // build-in interface{}
	}

	for _, rv := range s.rvs {
		if rv.IsValid() {
			if chk(rv.Type()) {
				retslice = reflect.Append(retslice, rv)
			}
		} else {
			if flg { // nil?
				// TODO: idk how to append nil to slice, directly...
				retslice = reflect.AppendSlice(retslice,
					reflect.MakeSlice(reflect.SliceOf(basetype), 1, 1))
			}
		}
	}

	rv2.Set(retslice)
	return nil
}

// New -
func New(v ...interface{}) Drawer {
	ret := &rawDrawer{}
	ret.Push(v...)
	return ret
}
