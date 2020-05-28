package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(2, 6))
}

//试题地址:https://leetcode-cn.com/problems/super-egg-drop/
func superEggDrop(K int, N int) int {
	stepMax := make([][]int, K+1)
	for stepK := 1; stepK <= K; stepK++ {
		stepMax[stepK] = make([]int, N+1)

		if stepK == 1 {
			for step := 1; step <= N; step++ {
				stepMax[stepK][step] = step
			}
			continue
		}

		step := 1
		for step = 1; step < N; step++ {
			if stepK <= step {
				stepMax[stepK][step] = stepMax[stepK][step-1] + stepMax[stepK-1][step-1] + 1
			} else {
				stepMax[stepK][step] = stepMax[stepK-1][step]
			}
			if stepMax[stepK][step] >= N { //剪枝
				break
			}
		}
		if stepK == K {
			fmt.Println(stepMax)
			return step
		}
	}
	return N
}
