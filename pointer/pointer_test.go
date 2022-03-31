package pointer

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := 1
		ret := Ptr(v)
		rt := reflect.TypeOf(ret)
		assert.Equal(t, rt.Kind(), reflect.Ptr)
		assert.Equal(t, *ret, 1)
	})
	t.Run("float", func(t *testing.T) {
		v := 1.0
		ret := Ptr(v)
		rt := reflect.TypeOf(ret)
		assert.Equal(t, rt.Kind(), reflect.Ptr)
		assert.Equal(t, *ret, float64(1))
	})
	t.Run("string", func(t *testing.T) {
		v := "a"
		ret := Ptr(v)
		rt := reflect.TypeOf(ret)
		assert.Equal(t, rt.Kind(), reflect.Ptr)
		assert.Equal(t, *ret, "a")
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		v := st{1}
		ret := Ptr(v)
		rt := reflect.TypeOf(ret)
		assert.Equal(t, rt.Kind(), reflect.Ptr)
		assert.Equal(t, *ret, st{1})
	})
	t.Run("my type", func(t *testing.T) {
		type MyType int
		v := MyType(1)
		ret := Ptr(v)
		rt := reflect.TypeOf(ret)
		assert.Equal(t, rt.Kind(), reflect.Ptr)
		assert.Equal(t, *ret, MyType(1))
	})
}

func TestUnptr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		a := 1
		b := &a
		c := Unptr(b)
		assert.Equal(t, c, a)
	})
	t.Run("float", func(t *testing.T) {
		a := float64(1)
		b := &a
		c := Unptr(b)
		assert.Equal(t, c, a)
	})
	t.Run("string", func(t *testing.T) {
		a := "a"
		b := &a
		c := Unptr(b)
		assert.Equal(t, c, a)
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		a := st{1}
		b := &a
		c := Unptr(b)
		assert.Equal(t, c, a)
	})
	t.Run("my type", func(t *testing.T) {
		type MyType int
		a := MyType(1)
		b := &a
		c := Unptr(b)
		assert.Equal(t, c, a)
	})
}
