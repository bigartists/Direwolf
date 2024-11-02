package main

import (
	"github.com/bigartists/Direwolf/cmd/direwolf/inject"
	"log"
	"os"
)

const (
	greetingBanner = `
  ____   _                             _   __ 
 |  _ \ (_) _ __  ___ __      __ ___  | | / _|
 | | | || || '__|/ _ \\ \ /\ / // _ \ | || |_ 
 | |_| || || |  |  __/ \ V  V /| (_) || ||  _|
 |____/ |_||_|   \___|  \_/\_/  \___/ |_||_|  

`
)

func main() {
	print(greetingBanner)
	app, err := inject.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	cmd := app.NewApiServerCommand()
	err = cmd.Execute()
	if err != nil {
		os.Exit(1)
		return
	}
}
