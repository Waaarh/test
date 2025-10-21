package main

func main() {
	// Пример использования функции MaxInt
	result := MaxInt(10, 20)
	println("Максимальное значение:", result)
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
