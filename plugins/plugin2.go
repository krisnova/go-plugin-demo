package main

import "fmt"

func Run() {
	fmt.Printf("(plugin2) Running an even better function in a Go plugin!\n", )
}

var I = 41

func PrintI() {
	fmt.Printf("(plugin2) Printing new value: %d\n", I)
}
