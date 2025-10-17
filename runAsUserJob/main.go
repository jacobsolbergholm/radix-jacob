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
	fmt.Println("test bare")
}

func main() {
	printUser()

	time.Sleep(10 * time.Minute)
}
