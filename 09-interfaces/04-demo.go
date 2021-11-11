package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, "20"))
	fmt.Println(sum(10, 20, 30, 40, 50, "abc"))
	fmt.Println(sum(10))
	fmt.Println(sum())
	fmt.Println(sum(10, 20, []int{30, 40, 50}))
	fmt.Println(sum(10, 20, []interface{}{30, "40", []int{50, 60}}))
}

func sum(data ...interface{}) int {
	result := 0
	for _, v := range data {
		switch value := v.(type) {
		case int:
			result += value
		case string:
			if i, err := strconv.Atoi(value); err == nil {
				result += i
			}
		case []interface{}:
			result += sum(value...)
		case []int:
			intfList := make([]interface{}, len(value))
			for i, v := range value {
				intfList[i] = v
			}
			result += sum(intfList...)
		}
	}
	return result
}
