package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	result := ""
	all := os.Args[1] + os.Args[2]
	for _, r := range all {
		found := false 
		for _, l := range result{
			if r ==l {
				found = true
			}
		}
		if !found {
			result += string(r)
		}
	}
	fmt.Println(result)
}
