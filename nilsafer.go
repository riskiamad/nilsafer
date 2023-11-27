package nilsafer

type model struct {
	Name string
}

// ValueOrZero check if pointer of value is not nil and return its value if exists, otherwise return zero value of its type
func ValueOrZero[T any](v *T) T {
	if v == nil {
		var zeroVal T
		return zeroVal
	}

	return *v
}
