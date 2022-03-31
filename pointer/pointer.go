package pointer

func Ptr[T any](v T) *T {
	return &v
}

func Unptr[T any](v *T) T {
	return *v
}
