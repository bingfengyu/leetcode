package main

import (
	"fmt"
)

func main() {
	// nums := []int{-3, 2, -3, 4, 2}
	// nums := []int{1, -2, -3}
	// fmt.Println(minStartValue(nums))

	// fmt.Println(findMinFibonacciNumbers(5))

	fmt.Println(getHappyString(10, 100))
}

func minStartValue(nums []int) int {
	value := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i] > 0 {
				continue
			}
			value = 1 - nums[i]
			continue
		}
		nums[i] = nums[i-1] + nums[i]
		if nums[i] < 0 {
			if value < 1-nums[i] {
				value = 1 - nums[i]
			}
		}
	}
	return value
}

func findMinFibonacciNumbers(k int) int {
	fibMap := map[int]bool{}
	fibList := []int{}
	getFibonacciMap(&fibMap, &fibList, k)
	if fibMap[k] {
		return 1
	}
	result := 0
	flag := false
	rescursion(&fibMap, &fibList, k, &result, &flag)
	return result
}

func getFibonacciMap(fibMap *map[int]bool, fibList *[]int, k int) {
	f1, f2 := 1, 2
	(*fibMap)[f1] = true
	(*fibMap)[f2] = true
	(*fibList) = append((*fibList), f1)
	(*fibList) = append((*fibList), f2)

	for f1+f2 <= k {
		f1, f2 = f2, f1+f2
		(*fibMap)[f2] = true
		(*fibList) = append((*fibList), f2)
	}
}

func rescursion(fibMap *map[int]bool, fibList *[]int, k int, result *int, flag *bool) {
	fmt.Println(*fibMap, *fibList, k, *result, *flag)
	if *flag {
		return
	}
	for i := len(*fibList) - 1; i < len(*fibList); i-- {
		if (*fibMap)[k] {
			(*result)++
			*flag = true
			return
		}
		if k-(*fibList)[i] < 0 {
			break
		}
		(*result)++
		rescursion(fibMap, fibList, k-(*fibList)[i], result, flag)
		if !(*flag) {
			(*result)--
		} else {
			break
		}
	}
	return
}

func getHappyString(n int, k int) string {
	happyStr := []string{"a", "b", "c"}
	happyString := [][]string{}
	happyString = append(happyString, happyStr)
	i := 0
	for n-i > 1 {
		temp := []string{}
		for _, v1 := range happyString[i] {
			for _, v2 := range happyStr {
				if string(v1[len(v1)-1]) != v2 {
					temp = append(temp, v1+v2)
				}
			}
		}
		happyString = append(happyString, temp)
		i++
	}
	if len(happyString[n-1]) >= k {
		return happyString[n-1][k-1]
	}
	return ""
}

// func numberOfArrays(s string, k int) int {
// 	strList := strings.Split(s, "0")

// 	for i := 0; i < len(strList)-2; i++ {
// 		if strList[i] == "" {
// 			continue
// 		}
// 		strList[i] += "0"
// 		num, _ := strconv.Atoi(strList[i])
// 	}

// 	return 0
// }
