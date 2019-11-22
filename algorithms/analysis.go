package algorithms

import (
	"math"
)

// Analyse - an example of evaluating the time efficiency of algorithms
func Analyse (f func(int) int, x int) int {

return f(x)

}

func fun1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m++
	}
	return m
}

func fun2(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m++
		}
	}
	return m
}

func fun3(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m++
		}
	}
	return m
}

func fun4(n int) int {
	m := 0
	i := 1
	for i < n {
		m++
		i = i * 2
	}
	return m
}

func fun5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m++
		i = i / 2
	}
	return m
}

func fun6(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				m++
			}
		}
	}
	return m
}

func fun7(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m++
		}
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			m++
		}
	}
	return m
}

func fun8(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		sq := math.Sqrt(float64(n))
		for j := 0; j < int(sq); j++ {
			m++
		}
	}
	return m
}

func fun9(n int) int {
	m := 0
	for i := n; i > 0; i /= 2 {
		for j := 0; j < i; j++ {
			m++
		}
	}
	return m
}

func fun10(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			m++
		}
	}
	return m
}

func fun11(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				m++
			}
		}
	}
	return m
}

func fun12(n int) int {
	j := 0
	m := 0
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			m++
		}
	}
	return m
}

func fun13(n int) int {
	m := 0
	for i := 1; i <= n; i *= 2 {
		for j := 0; j <= i; j++ {
			m++
		}
	}
	return m
}
