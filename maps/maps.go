package maps

import (
	"github.com/icepigss/goutil/types"
)

// get all the keys of the map
func Keys[K comparable, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for key, _ := range m {
		out = append(out, key)
	}

	return out
}

// get all the values of the map
func Values[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, value := range m {
		out = append(out, value)
	}

	return out
}

// combine multiple maps into one
func Merge[M types.MC[K, V], K comparable, V any](ms ...M) M {
	out := make(M)
	for _, m := range ms {
		for key, value := range m {
			out[key] = value
		}
	}

	return out
}
