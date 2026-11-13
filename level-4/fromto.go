package main

// import (
// 	"fmt"
// )

func pad(n int)string {
	if n<10 {
		return "0" + string(rune('0'+n))
	}
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	result := ""
	if from <=to{
		for i:= from ; i<= to ; i++ {
			result += pad(i)
			if i != to {
				result += ","
			}
		}
	} else {
		for i := from ; i>= to ; i-- {
			result+= pad(i)
			if i != to {
				result += ","
			}
		}
	}
	return result +"\n"
}
// func main() {
// 	fmt.Print(FromTo(1, 10))
// 	fmt.Print(FromTo(10, 1))
// 	fmt.Print(FromTo(10, 10))
// 	fmt.Print(FromTo(100, 10))
// }