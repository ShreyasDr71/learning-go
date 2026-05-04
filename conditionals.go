package main

import "fmt"

func main() {
	if donut := "flake"; donut != "flake" {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}
}
