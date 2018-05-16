package drawer_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/maxbet1507/drawer"
)

func Example() {
	d := drawer.New()

	d.Push(123, "hello", true)
	d.Push(fmt.Errorf("world"), false, 456)

	rv1 := []int{}
	d.Pull(&rv1)
	fmt.Println(rv1)

	rv2 := []string{}
	d.Pull(&rv2)
	fmt.Println(rv2)

	rv3 := []error{}
	d.Pull(&rv3)
	fmt.Println(rv3)

	// Output:
	// [123 456]
	// [hello]
	// [world]
}

type testInterface1 interface {
	FakeFunc1()
}

type testInterface2 interface {
	FakeFunc2()
}

type testStruct1 struct {
}

func (s *testStruct1) FakeFunc1() {}

type testStruct2 struct {
}

func (s *testStruct2) FakeFunc2() {}

type testStruct3 struct {
}

func (s *testStruct3) FakeFunc1() {}
func (s *testStruct3) FakeFunc2() {}

func TestDrawerEmpty(t *testing.T) {
	r := drawer.New()
	if d := r.Dump(); !reflect.DeepEqual(d, []interface{}{}) {
		t.Fatal(d)
	}

	var v1 []interface{}
	if err := r.Pull(&v1); err != nil || !reflect.DeepEqual(v1, []interface{}{}) {
		t.Fatal(err, v1)
	}

	v2 := []interface{}{}
	if err := r.Pull(&v2); err != nil || !reflect.DeepEqual(v2, []interface{}{}) {
		t.Fatal(err, v2)
	}
}

func TestDrawerNil(t *testing.T) {
	var v interface{}

	r := drawer.New(nil)
	if d := r.Dump(); !reflect.DeepEqual(d, []interface{}{v}) {
		t.Fatal(d)
	}

	var v1 []interface{}
	if err := r.Pull(&v1); err != nil || !reflect.DeepEqual(v1, []interface{}{v}) {
		t.Fatal(err, v1)
	}

	var v2 []testInterface1
	if err := r.Pull(&v2); err != nil || !reflect.DeepEqual(v2, []testInterface1{}) {
		t.Fatal(err, v2)
	}

	var v3 []*testStruct1
	if err := r.Pull(&v3); err != nil || !reflect.DeepEqual(v3, []*testStruct1{}) {
		t.Fatal(err, v3)
	}

	var v4 []testInterface2
	if err := r.Pull(&v4); err != nil || !reflect.DeepEqual(v4, []testInterface2{}) {
		t.Fatal(err, v4)
	}

	var v5 []*testStruct2
	if err := r.Pull(&v5); err != nil || !reflect.DeepEqual(v5, []*testStruct2{}) {
		t.Fatal(err, v5)
	}

	var v6 []*testStruct3
	if err := r.Pull(&v6); err != nil || !reflect.DeepEqual(v6, []*testStruct3{}) {
		t.Fatal(err, v6)
	}
}

func TestDrawerInterface1(t *testing.T) {
	var v testInterface1

	r := drawer.New(v)
	if d := r.Dump(); !reflect.DeepEqual(d, []interface{}{v}) {
		t.Fatal(d)
	}

	var v1 []interface{}
	if err := r.Pull(&v1); err != nil || !reflect.DeepEqual(v1, []interface{}{v}) {
		t.Fatal(err, v1)
	}

	var v2 []testInterface1
	if err := r.Pull(&v2); err != nil || !reflect.DeepEqual(v2, []testInterface1{}) {
		t.Fatal(err, v2)
	}

	var v3 []*testStruct1
	if err := r.Pull(&v3); err != nil || !reflect.DeepEqual(v3, []*testStruct1{}) {
		t.Fatal(err, v3)
	}

	var v4 []testInterface2
	if err := r.Pull(&v4); err != nil || !reflect.DeepEqual(v4, []testInterface2{}) {
		t.Fatal(err, v4)
	}

	var v5 []*testStruct2
	if err := r.Pull(&v5); err != nil || !reflect.DeepEqual(v5, []*testStruct2{}) {
		t.Fatal(err, v5)
	}

	var v6 []*testStruct3
	if err := r.Pull(&v6); err != nil || !reflect.DeepEqual(v6, []*testStruct3{}) {
		t.Fatal(err, v6)
	}
}

func TestDrawerStruct1(t *testing.T) {
	var v *testStruct1

	r := drawer.New(v)
	if d := r.Dump(); !reflect.DeepEqual(d, []interface{}{v}) {
		t.Fatal(d)
	}

	var v1 []interface{}
	if err := r.Pull(&v1); err != nil || !reflect.DeepEqual(v1, []interface{}{v}) {
		t.Fatal(err, v1)
	}

	var v2 []testInterface1
	if err := r.Pull(&v2); err != nil || !reflect.DeepEqual(v2, []testInterface1{v}) {
		t.Fatal(err, v2)
	}

	var v3 []*testStruct1
	if err := r.Pull(&v3); err != nil || !reflect.DeepEqual(v3, []*testStruct1{v}) {
		t.Fatal(err, v3)
	}

	var v4 []testInterface2
	if err := r.Pull(&v4); err != nil || !reflect.DeepEqual(v4, []testInterface2{}) {
		t.Fatal(err, v4)
	}

	var v5 []*testStruct2
	if err := r.Pull(&v5); err != nil || !reflect.DeepEqual(v5, []*testStruct2{}) {
		t.Fatal(err, v5)
	}

	var v6 []*testStruct3
	if err := r.Pull(&v6); err != nil || !reflect.DeepEqual(v6, []*testStruct3{}) {
		t.Fatal(err, v6)
	}
}

func TestDrawerStruct2(t *testing.T) {
	var v *testStruct3

	r := drawer.New(v)
	if d := r.Dump(); !reflect.DeepEqual(d, []interface{}{v}) {
		t.Fatal(d)
	}

	var v1 []interface{}
	if err := r.Pull(&v1); err != nil || !reflect.DeepEqual(v1, []interface{}{v}) {
		t.Fatal(err, v1)
	}

	var v2 []testInterface1
	if err := r.Pull(&v2); err != nil || !reflect.DeepEqual(v2, []testInterface1{v}) {
		t.Fatal(err, v2)
	}

	var v3 []*testStruct1
	if err := r.Pull(&v3); err != nil || !reflect.DeepEqual(v3, []*testStruct1{}) {
		t.Fatal(err, v3)
	}

	var v4 []testInterface2
	if err := r.Pull(&v4); err != nil || !reflect.DeepEqual(v4, []testInterface2{v}) {
		t.Fatal(err, v4)
	}

	var v5 []*testStruct2
	if err := r.Pull(&v5); err != nil || !reflect.DeepEqual(v5, []*testStruct2{}) {
		t.Fatal(err, v5)
	}

	var v6 []*testStruct3
	if err := r.Pull(&v6); err != nil || !reflect.DeepEqual(v6, []*testStruct3{v}) {
		t.Fatal(err, v6)
	}
}

func TestDrawerNestPush(t *testing.T) {
	r1 := drawer.New(1)
	r2 := drawer.New(r1, 2)

	if d := r2.Dump(); !reflect.DeepEqual(d, []interface{}{1, 2}) {
		t.Fatal(d)
	}
}

func TestDrawerPullError(t *testing.T) {
	r := drawer.New()

	var v1 []interface{}
	if err := r.Pull(v1); err != drawer.ErrInvalidParameterType {
		t.Fatal(err)
	}

	var v2 interface{}
	if err := r.Pull(&v2); err != drawer.ErrInvalidParameterType {
		t.Fatal(err)
	}
}
