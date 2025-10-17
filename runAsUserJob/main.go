package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func printUser() {
	fmt.Println("UID: " + strconv.Itoa(os.Getuid()))
	fmt.Println("GID: " + strconv.Itoa(os.Getgid()))

	panic("test")
}

func main() {
	printUser()

	time.Sleep(10 * time.Minute)
}
