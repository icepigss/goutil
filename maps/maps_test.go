package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := map[int]int{1: 1, 2: 2, 3: 3}
		ret := Keys(s)
		assertSliceEqualWithoutOrder(t, ret, []int{1, 2, 3})
	})
	t.Run("float", func(t *testing.T) {
		s := map[float32]int{1.0: 1, 2.0: 2, 3.0: 3}
		ret := Keys(s)
		assertSliceEqualWithoutOrder(t, ret, []float32{1.0, 2.0, 3.0})
	})
	t.Run("string", func(t *testing.T) {
		s := map[string]int{"a": 1, "b": 2, "c": 3}
		ret := Keys(s)
		assertSliceEqualWithoutOrder(t, ret, []string{"a", "b", "c"})
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		s := map[st]int{{1}: 1, {2}: 2, {3}: 3}
		ret := Keys(s)
		assertSliceEqualWithoutOrder(t, ret, []st{{1}, {2}, {3}})
	})
	t.Run("my map", func(t *testing.T) {
		type MyInt int
		type MyMap map[MyInt]int
		s := MyMap{1: 1, 2: 2, 3: 3}
		ret := Keys(s)
		assertSliceEqualWithoutOrder(t, ret, []MyInt{1, 2, 3})
	})

}

func TestValues(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := map[int]int{1: 1, 2: 2, 3: 3}
		ret := Values(s)
		assertSliceEqualWithoutOrder(t, ret, []int{1, 2, 3})
	})
	t.Run("float", func(t *testing.T) {
		s := map[int]float32{1: 1.0, 2: 2.0, 3: 3.0}
		ret := Values(s)
		assertSliceEqualWithoutOrder(t, ret, []float32{1.0, 2.0, 3.0})
	})
	t.Run("string", func(t *testing.T) {
		s := map[int]string{1: "a", 2: "b", 3: "c"}
		ret := Values(s)
		assertSliceEqualWithoutOrder(t, ret, []string{"a", "b", "c"})
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		s := map[int]st{1: {1}, 2: {2}, 3: {3}}
		ret := Values(s)
		assertSliceEqualWithoutOrder(t, ret, []st{{1}, {2}, {3}})
	})
	t.Run("my map", func(t *testing.T) {
		type MyInt int
		type MyMap map[int]MyInt
		s := MyMap{1: 1, 2: 2, 3: 3}
		ret := Values(s)
		assertSliceEqualWithoutOrder(t, ret, []MyInt{1, 2, 3})
	})

}

func TestMerge(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		m1 := map[int]int{1: 1, 2: 2}
		m2 := map[int]int{3: 3, 4: 4}
		m3 := map[int]int{5: 5}
		ret := Merge(m1, m2, m3)
		assert.Equal(t, ret, map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5})
	})
	t.Run("float", func(t *testing.T) {
		m1 := map[float64]float64{1: 1, 2: 2}
		m2 := map[float64]float64{3: 3, 4: 4}
		m3 := map[float64]float64{5: 5}
		ret := Merge(m1, m2, m3)
		assert.Equal(t, ret, map[float64]float64{1: 1, 2: 2, 3: 3, 4: 4, 5: 5})
	})
	t.Run("string", func(t *testing.T) {
		m1 := map[string]string{"a": "a", "b": "b"}
		m2 := map[string]string{"c": "c", "d": "d"}
		m3 := map[string]string{"e": "e"}
		ret := Merge(m1, m2, m3)
		assert.Equal(t, ret, map[string]string{"a": "a", "b": "b", "c": "c", "d": "d", "e": "e"})
	})
}

func assertSliceEqualWithoutOrder[T comparable](t *testing.T, s1, s2 []T) {
	if len(s1) != len(s2) {
		t.Error("len not equal")
	}
	for _, item := range s1 {
		flag := false
		for _, item2 := range s2 {
			if item2 == item {
				flag = true
				break
			}
		}
		if !flag {
			t.Errorf("%v not exists in s2:%v", item, s2)
		}
	}
}
