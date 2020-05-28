package main

func main() {

}

//试题地址：https://leetcode-cn.com/problems/powx-n/
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return doPow1(x, n)
	}
	return 1.0 / doPow1(x, n)
}

//分治
func doPow1(x float64, n int) float64 {
	result := 1.0
	tmp := x
	for n > 0 {
		if n%2 == 0 {
			result *= tmp
		}
		tmp *= tmp
		n /= 2
	}
	return result
}

//递归
func doPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := doPow2(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
