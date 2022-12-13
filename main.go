package main

import (
	"log"
	"video-factory/cmds"
)

func main() {
	command := cmds.NewRootCmd()
	err := command.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
