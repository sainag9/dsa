package Arrays

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {

	a := make([]int, 0)
	arr := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	s := make([]string, 0)
	for i := 0; i < len(digits); i++ {

		d := digits[i] - 48
		if d == 9 || d == 7 {
			a = append(a, 4)
		} else {
			a = append(a, 3)
		}
		s = append(s, arr[d-2])
	}
	fmt.Println(a, s)

	sarr := make([]string, 0)

	for i := len(s) - 1; i >= 0; i-- {

		sarr = getPerm(s[i], sarr)
		fmt.Println(sarr)

	}
	return sarr
}

func getPerm(a string, s []string) []string {

	if len(s) == 0 {
		//fmt.Println(strings.Split(a, ""))
		return strings.Split(a, "")
	}

	arr := make([]string, 0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(s); j++ {
			str := string(a[i]) + s[j]
			arr = append(arr, str)
		}
	}
	//fmt.Println(arr)
	return arr
}
