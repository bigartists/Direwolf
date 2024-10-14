package main

import "github.com/bigartists/Direwolf/server"

func main() {
	cmd := server.NewApiServerCommand()
	err := cmd.Execute()
	if err != nil {
		return
	}
}
