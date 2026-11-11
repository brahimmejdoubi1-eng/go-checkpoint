package main 

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }

func HashCode(dec string) string {
	var result string
	for _, r := range dec {
		if int(r)<33 {
			r+=33
		}
		char:= (int(r)+len(dec))%127
		result+= string(char)
	}
	return result
}