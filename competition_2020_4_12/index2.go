package main

// func main() {
// 	m := []int{7, 5, 5, 8, 3}
// 	fmt.Println(processQueries(m, 8))
// }

func processQueries(queries []int, m int) []int {
	result := []int{}

	record := map[int]int{}
	for i := 0; i < m; i++ {
		record[i+1] = i
	}

	for _, v := range queries {
		result = append(result, record[v])
		adjustMap(&record, v)
		// fmt.Println(record)
	}
	return result
}

func adjustMap(record *map[int]int, m int) {
	temp := (*record)[m]
	// fmt.Println(temp)
	(*record)[m] = 0
	for k, v := range *record {
		if k != m && v < temp {
			(*record)[k] = v + 1
		}
	}
}
