package main

import (
	"log"
	"video-factory/cmd"
)

func main() {
	command := cmd.NewRootCmd()
	err := command.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
