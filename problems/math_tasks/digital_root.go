package math_tasks

func Digital_root(n int) int {
	if n == 0 {
		return 0
	}
	if n%9 == 0 {
		return 9
	}
	return n % 9
}
