package main

// import (
// 	"fmt"

// )

// func main() {
// 	fmt.Println(RepeatAlpha("abc"))
// 	fmt.Println(RepeatAlpha("Choumi."))
// 	fmt.Println(RepeatAlpha(""))
// 	fmt.Println(RepeatAlpha("abacadaba 01!"))
// }


func RepeatAlpha(s string) string {
	result:=""
	for _, r := range s {
		count:=1
		if r>='a' && r<='z' {
			count=int(r-'a')+1
		}
		if r>='A' && r<='Z' {
			count=int(r-'A')+1
		}
		for i:=1;i<=count;i++ {
			result+=string(r)
		}
	}
	return result 
}

