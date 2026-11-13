package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func dup(s string, c rune) int {
	count := 0
	for _, r := range s {
		if c == r {
			count++
		}
	}
	return count
}

func ZipString(s string) string {
	result := ""
	for _, r := range s {
		times := 0
		times=dup(s,r)
		for i:=1; i<=times;i++ {
			result+=string(r)
		}
	}
	return result
}
