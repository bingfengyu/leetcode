package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// fmt.Println(reformat("a0b1c2"))
	orders := [][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}
	fmt.Println(displayTable(orders))
}

func reformat(s string) string {
	strNum := []string{}
	strStr := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			strNum = append(strNum, string(s[i]))
		} else {
			strStr = append(strStr, string(s[i]))
		}
	}
	if len(strNum) == len(strStr) {
		str := ""
		for i := 0; i < len(strNum); i++ {
			str += strNum[i] + strStr[i]
		}
		return str
	} else if len(strNum)-len(strStr) == 1 {
		str := ""
		for i := 0; i < len(strStr); i++ {
			str += strNum[i] + strStr[i]
		}
		str += strNum[len(strNum)-1]
		return str
	} else if len(strStr)-len(strNum) == 1 {
		str := ""
		for i := 0; i < len(strNum); i++ {
			str += strStr[i] + strNum[i]
		}
		str += strStr[len(strStr)-1]
		return str
	}
	return ""
}

func displayTable(orders [][]string) [][]string {
	orderTable := [][]string{}
	temp := []string{}
	tempMap := map[string]bool{}
	tempMap1 := map[string]int{}
	temp = append(temp, "Table")
	for i := 0; i < len(orders); i++ {
		if !tempMap[orders[i][2]] {
			temp = append(temp, orders[i][2])
			tempMap[orders[i][2]] = true
		}
	}
	sort.Strings(temp[1:])
	orderTable = append(orderTable, temp)
	for k, v := range temp[1:] {
		tempMap1[v] = k + 1
	}

	temp2 := []int{}
	tempMap = map[string]bool{}
	tempMap2 := map[string]int{}
	for i := 0; i < len(orders); i++ {
		if !tempMap[orders[i][1]] {
			tempInt, _ := strconv.Atoi(orders[i][1])
			temp2 = append(temp2, tempInt)
			tempMap[orders[i][1]] = true
		}
	}
	sort.Ints(temp2)
	for k, v := range temp2 {
		t := []string{}
		t = append(t, strconv.Itoa(v))
		tempMap2[strconv.Itoa(v)] = k + 1
		for i := 0; i < len(orderTable[0])-1; i++ {
			t = append(t, "0")
		}
		orderTable = append(orderTable, t)
	}
	// fmt.Println(tempMap1, tempMap2)

	for i := 0; i < len(orders); i++ {
		table := tempMap2[orders[i][1]]
		dinner := tempMap1[orders[i][2]]
		// fmt.Println(table, dinner)
		tempInt, _ := strconv.Atoi(orderTable[table][dinner])
		tempInt++
		orderTable[table][dinner] = strconv.Itoa(tempInt)
	}

	// fmt.Println(orderTable)
	return orderTable
}

func minNumberOfFrogs(croakOfFrogs string) int {
	return 0
}
