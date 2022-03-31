package types

type Singed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsinged interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Singed | Unsinged
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

type Complex interface {
	~complex64 | ~complex128
}

type SC[T any] interface {
	~[]T
}

type MC[K comparable, V any] interface {
	~map[K]V
}

type Addable interface {
	Integer | Float
}
