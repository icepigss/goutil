package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{1, 2, 3}
		assert.True(t, Contains(s, 1))
		assert.False(t, Contains(s, 4))
	})
	t.Run("float", func(t *testing.T) {
		s := []int{1.0, 2.0, 3.0}
		assert.True(t, Contains(s, 1.0))
		assert.False(t, Contains(s, 4.0))
	})
	t.Run("string", func(t *testing.T) {
		s := []string{"a", "b", "c"}
		assert.True(t, Contains(s, "a"))
		assert.False(t, Contains(s, "d"))
	})
	t.Run("bool", func(t *testing.T) {
		s := []bool{false}
		assert.True(t, Contains(s, false))
		assert.False(t, Contains(s, true))
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		s := []st{
			st{
				a: 1,
			},
		}
		assert.True(t, Contains(s, st{a: 1}))
		assert.False(t, Contains(s, st{a: 2}))
	})
	t.Run("my slice", func(t *testing.T) {
		type MySlice []int
		s := MySlice{1, 2, 3}
		assert.True(t, Contains(s, 1))
		assert.False(t, Contains(s, 4))
	})
}

func TestUnique(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{1, 2, 2, 3}
		ret := Unique(s)
		assert.Equal(t, ret, []int{1, 2, 3})
	})
	t.Run("float", func(t *testing.T) {
		s := []float32{1.0, 2.0, 2.0, 3.0}
		ret := Unique(s)
		assert.Equal(t, ret, []float32{1.0, 2.0, 3.0})
	})
	t.Run("string", func(t *testing.T) {
		s := []string{"a", "b", "b", "c"}
		ret := Unique(s)
		assert.Equal(t, ret, []string{"a", "b", "c"})
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		s := []st{{1}, {2}, {2}, {3}}
		ret := Unique(s)
		assert.Equal(t, ret, []st{{1}, {2}, {3}})
	})
	t.Run("my slice", func(t *testing.T) {
		type MySlice []int
		s := MySlice{1, 2, 2, 3}
		ret := Unique(s)
		assert.Equal(t, ret, MySlice{1, 2, 3})
	})
}

func TestRemove(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{1, 2, 3}
		ret := Remove(s, 2)
		assert.Equal(t, ret, []int{1, 3})
	})
	t.Run("float", func(t *testing.T) {
		s := []float32{1.0, 2.0, 3.0}
		ret := Remove(s, 2.0)
		assert.Equal(t, ret, []float32{1.0, 3.0})
	})
	t.Run("string", func(t *testing.T) {
		s := []string{"a", "b", "b", "c"}
		ret := Remove(s, "b")
		assert.Equal(t, ret, []string{"a", "c"})
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		s := []st{{1}, {2}, {3}}
		ret := Remove(s, st{2})
		assert.Equal(t, ret, []st{{1}, {3}})
	})
	t.Run("MySlice", func(t *testing.T) {
		type MySlice []int
		s := MySlice{1, 2, 3}
		ret := Remove(s, 2)
		assert.Equal(t, ret, MySlice{1, 3})
	})
}

func TestEqualWithoutOrder(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{3, 2, 1}
		s3 := []int{2, 1}
		assert.True(t, EqualWithoutOrder(s1, s2))
		assert.False(t, EqualWithoutOrder(s1, s3))
	})
	t.Run("float", func(t *testing.T) {
		s1 := []float32{1.0, 2.0, 3.0}
		s2 := []float32{3.0, 2.0, 1.0}
		s3 := []float32{2.0, 1.0}
		assert.True(t, EqualWithoutOrder(s1, s2))
		assert.False(t, EqualWithoutOrder(s1, s3))
	})
	t.Run("string", func(t *testing.T) {
		s1 := []string{"a", "b", "c"}
		s2 := []string{"c", "b", "a"}
		s3 := []string{"c", "b"}
		assert.True(t, EqualWithoutOrder(s1, s2))
		assert.False(t, EqualWithoutOrder(s1, s3))
	})
	t.Run("struct", func(t *testing.T) {
		type st struct {
			a int
		}
		s1 := []st{{1}, {2}, {3}}
		s2 := []st{{3}, {2}, {1}}
		s3 := []st{{3}, {2}}
		assert.True(t, EqualWithoutOrder(s1, s2))
		assert.False(t, EqualWithoutOrder(s1, s3))
	})
	t.Run("my slice", func(t *testing.T) {
		type MySlice []int
		s1 := MySlice{1, 2, 3}
		s2 := MySlice{3, 2, 1}
		s3 := MySlice{2, 1}
		assert.True(t, EqualWithoutOrder(s1, s2))
		assert.False(t, EqualWithoutOrder(s1, s3))
	})
}

func TestMerge(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s1 := []int{1, 2}
		s2 := []int{3, 4}
		s3 := []int{5}
		ret := Merge(s1, s2, s3)
		assert.Equal(t, ret, []int{1, 2, 3, 4, 5})
	})
	t.Run("string", func(t *testing.T) {
		s1 := []string{"a", "b"}
		s2 := []string{"c", "d"}
		s3 := []string{"e"}
		ret := Merge(s1, s2, s3)
		assert.Equal(t, ret, []string{"a", "b", "c", "d", "e"})
	})

}

func TestSum(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{1, 2, 3}
		ret := Sum(s)
		assert.Equal(t, ret, 6)
	})
	t.Run("float", func(t *testing.T) {
		s := []float64{1.0, 2.0, 3.0}
		ret := Sum(s)
		assert.Equal(t, ret, 6.0)
	})
}

func TestSortAsc(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{2, 3, 1}
		SortAsc(s)
		assert.Equal(t, s, []int{1, 2, 3})
	})
	t.Run("float", func(t *testing.T) {
		s := []float64{2.0, 3.0, 1.0}
		SortAsc(s)
		assert.Equal(t, s, []float64{1.0, 2.0, 3.0})
	})
}

func TestSortDesc(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{2, 3, 1}
		SortDesc(s)
		assert.Equal(t, s, []int{3, 2, 1})
	})
	t.Run("float", func(t *testing.T) {
		s := []float64{2.0, 3.0, 1.0}
		SortDesc(s)
		assert.Equal(t, s, []float64{3.0, 2.0, 1.0})
	})
}
