package main

import (
	"dev/lamoda_test/internal/app"
	"dev/lamoda_test/internal/config"
	"log"
)

var cfg config.Config

func init() {
	err := cfg.InitCfg()
	if err != nil {
		panic(err)
	}
}

func main() {
	err := app.Run(cfg)
	if err != nil {
		log.Println(err)
	}
}
