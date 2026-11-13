package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(WeAreUnique("foo", "boo"))
// 	fmt.Println(WeAreUnique("", ""))
// 	fmt.Println(WeAreUnique("abc", "def"))
// }




func contains(str string, c rune) bool {
	if str == "" {
		return false
	}
	for _, r := range str {
		if c == r {
			return true
		}
	}
	return false
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2==""{
		return -1
	}
	
	count := 0
	for i, r := range str1 {
		if !contains(str2, r) && !contains(str1[:i],r) {
			count++
		}
	}
	for i, r:= range str2 {
		if !contains(str1, r) && !contains(str2[:i],r){
			count++
		}
	}
	return count
}
