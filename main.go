package main

import "fmt"

func main() {
	price := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}
	for key, volume := range price {
		if volume >= 500 {
			fmt.Println(key)
		}
	}
	zakaz := []string{"хлеб", "буженина", "сыр", "огурцы"}
	cost := 0
	for _, volume := range zakaz {
		cost += price[volume]
	}
	fmt.Println(cost)

	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(find(arr, 3))
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println(RemuveDuplicate(input))
}

func find(arr []int, k int) []int {

	m := make(map[int]int)
	for index, volume := range arr {
		if indexMap, volumeMap := m[k-volume]; volumeMap {
			return []int{index, indexMap}
		}
		m[volume] = index
	}

	return nil
}

func RemuveDuplicate(input []string) []string {
	m := make(map[string]int)
	output := make([]string, len(input))
	counter := 0
	for _, vol := range input {
		if _, z := m[vol]; !z {
			output[counter] = vol
			counter++
		}
		m[vol] += 1
	}
	fmt.Print(m)
	return output

}
