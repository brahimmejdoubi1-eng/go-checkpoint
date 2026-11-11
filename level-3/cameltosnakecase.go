package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorld"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("hey2"))
// }


func iscamelcase(s string)bool {
	for i, r:= range s {
		if i>0 {
			if (s[i]>='A' && s[i]<='Z') && (s[i-1]>='A' && s[i-1]<='Z') {
				return false
			}
		}
		
		if !(r>='a' && r<='z' || r>='A' && r<='Z') {
			return false
		}
		if i==len(s)-1 && (r>='A' && r<='Z') {
			return false
		}
	}
	return true
}

func CamelToSnakeCase(s string) string{
	result := ""
	if s==""{
		return ""
	}
	if iscamelcase(s) {
		for i, r := range s {
			if r>='A' && r<='Z' && i>0 {
				result+="_"
			}
			result+=string(r)
		}
	} else {
		result=s
	}
	return result
}