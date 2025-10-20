package main

import (
	"fmt"
	"os"
	"strconv"
)

func printUser() {
	fmt.Println("UID: " + strconv.Itoa(os.Getuid()))
	fmt.Println("GID: " + strconv.Itoa(os.Getgid()))
}

func main() {
	printUser()
}
