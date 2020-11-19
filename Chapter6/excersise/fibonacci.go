package main

// Fibonacci 斐波那契数列
func Fibonacci(n int) (pos int, res int) {
	defer func() {
		// fmt.Printf("Fibonacci(%d) = %d, %d\n", n, pos, res)
	}()

	if n == 0 {
		pos, res = 0, 1
	} else if n == 1 {
		pos, res = 1, 1
	} else {
		pos, res = Fibonacci(n - 1)
		_, res2 := Fibonacci(n - 2)
		pos++
		res += res2
	}

	return
}
