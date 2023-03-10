package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"mode51.dev/goserverandom/internal/app"
	"mode51.dev/goserverandom/internal/web"
	"os"
	"strconv"
)

const (
	USAGE          = "Please supply a port number as the cmd line arg"
	UNABLETODECODE = "Unable to decode the cmd line arg"
)

func init() {
	app.InitLogger()
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		log.Fatal().Msg(USAGE)
	}
	port, err := strconv.ParseUint(argsWithoutProg[0], 10, 64)
	if err != nil {
		log.Fatal().Msg(UNABLETODECODE)
	}

	log.Info().Msgf("Setting up server on port %d", port)
	r := web.SetupRouter()
	r.Run(fmt.Sprintf(":%d", port))
}
