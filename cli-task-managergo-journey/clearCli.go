package main

import "fmt"

func ClearCli() {
	fmt.Print("\033[H\033[2J")
}
