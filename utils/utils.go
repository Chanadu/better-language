package utils

func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsAlpha(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

type Number interface {
	~int | ~float64
}

type operator[T Number] func(a T, b T) T

func operateNumber[T Number](a, b T, fn operator[T]) T {
	return fn(a, b)
}
