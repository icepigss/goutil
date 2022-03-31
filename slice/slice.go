package slice

import (
	"sort"

	"github.com/icepigss/goutil/types"
)

// check whether the specified element exists in the slice
func Contains[T comparable](s []T, v T) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}

	return false
}

// remove duplicate elements from slice
func Unique[S types.SC[T], T comparable](s S) S {
	l := len(s)
	if l <= 0 {
		return s
	}

	out := make(S, 0, l)
	exists := make(map[T]struct{}, l)
	for _, item := range s {
		if _, ok := exists[item]; ok {
			continue
		}
		exists[item] = struct{}{}
		out = append(out, item)
	}

	return out
}

// remove the pecified element from slice
func Remove[S types.SC[T], T comparable](s S, v T) S {
	out := make(S, 0, len(s))
	for _, item := range s {
		if item == v {
			continue
		}
		out = append(out, item)
	}

	return out
}

// determine whether two slices are equal without considering the order
func EqualWithoutOrder[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, item1 := range s1 {
		if !Contains(s2, item1) {
			return false
		}
	}

	return true
}

// combine multiple slices into one
func Merge[S types.SC[T], T any](ss ...S) S {
	var out S
	for _, s := range ss {
		out = append(out, s...)
	}
	return out
}

// calculate the sum of slice elements
func Sum[T types.Addable](s []T) T {
	var out T
	for _, item := range s {
		out += item
	}

	return out
}

// sort the slice in ascending order
func SortAsc[T types.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// sort the slice in descending order
func SortDesc[T types.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
