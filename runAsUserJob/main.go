package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func printUser() {
	log.Println("TEST")
	log.Println("UID: " + strconv.Itoa(os.Getuid()))
	log.Println("GID: " + strconv.Itoa(os.Getgid()))
}

func main() {
	printUser()

	time.Sleep(10 * time.Minute)
}
