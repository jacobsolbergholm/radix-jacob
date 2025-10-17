package main

import (
	"fmt"
	"os"
	"time"
)

func printUser() {
	fmt.Printf("UID: %d\n", os.Getuid())
	fmt.Printf("GID: %d\n", os.Getgid())
}

func main() {
	printUser()

	time.Sleep(10 * time.Minute)
}
