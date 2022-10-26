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

	//locate, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	panic(err)
	//}
	//c := cron.New(cron.WithLocation(locate), cron.WithSeconds())
	//_, err = c.AddFunc("*/3 * * * * *", test)
	//if err != nil {
	//	panic(err)
	//}
	//c.Start()
	//defer c.Stop()
	//select {}
}

//func test() {
//	fmt.Println("every 3 seconds running ...")
//}
