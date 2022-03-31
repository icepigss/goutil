package boolean

func ToBool[T comparable](v T, zero T) bool {
	return v != zero
}

func ToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
