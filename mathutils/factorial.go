package mathutils

func Factorial(a int) uint64 {
	if a > 1 {
		var res uint64 = 1
		for i := 1; i < a+1; i++ {
			res *= res * uint64(i)
		}
		return res
	}
	return 0
}
