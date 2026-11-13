package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(IsCapitalized("Hello! How are you?"))
// 	fmt.Println(IsCapitalized("Hello How Are You"))
// 	fmt.Println(IsCapitalized("Whats 4this 100K?"))
// 	fmt.Println(IsCapitalized("Whatsthis4"))
// 	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
// 	fmt.Println(IsCapitalized(""))
// }


func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	foundword:=false 
	startofword := true

	for i:= 0; i<len(s); i++ {
		if s[i] == ' ' {
			startofword= true
			continue
		}
		if startofword {
			foundword=true
			if s[i]>='a' && s[i]<='z' {
				return false
			}
			startofword=false
		}
	}
	return foundword
}