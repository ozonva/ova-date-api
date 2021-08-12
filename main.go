package main

import (
	"fmt"
	"ova-date-api/cmd/ova-template-api"
	"ova-date-api/internal/utils"
)

func main() {
	ova_template_api.Hello()

	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	utils.PrintSlice(testSlice)
	r := utils.SplitSlice(testSlice, 3)
	for i := range r {
		utils.PrintSlice(r[i])
	}

	var testMap = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	var r2 = utils.ReverseKeysValues(testMap)
	fmt.Println(testMap)
	fmt.Println(r2)

	testString := []string{"first", "second", "error", "third", "error", "correct"}
	filter := []string{"error", "third"}
	r3 := utils.ExcludeItems(testString, filter)
	fmt.Println(r3)
}
