package main

import (
	bampCli "github.com/Issei0804-ie/bamp/cli"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	godotenv.Load("./.env")
	mode := os.Getenv("DEVELOP_MODE")
	log.SetOutput(os.Stdout)
	if mode == "TRUE" {
		log.SetFlags(log.Llongfile)
	}
	app := &cli.App{
		Name:    "bamp. bamp is backup software.",
		Usage:   "bamp {settings.json}",
		Action:  bampCli.Compress,
		Version: "v0.0.1",
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln(err)
	}
}
