package main

import (
	"fmt"
)

func main() {
	for l := 'a'; l <= 'z'; l++ {
		fmt.Printf("%c", l) //%c iterates all characters from a to z in ASCII
	}
}
