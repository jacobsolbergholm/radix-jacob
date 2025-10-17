package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func printUser() {
	ctx := context.Background()
	log.Ctx(ctx).Info().Msg("test")
	log.Ctx(ctx).Info().Msg("UID: " + strconv.Itoa(os.Getuid()))
	log.Ctx(ctx).Info().Msg("GID: " + strconv.Itoa(os.Getgid()))
}

func main() {
	printUser()

	time.Sleep(10 * time.Minute)
}
