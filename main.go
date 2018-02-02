package main

import (
	_ "svGo/routers"
	"github.com/astaxie/beego"
	"svGo/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	models.Queue = make(chan models.Message, 600)
	models.Worker = make(chan int, 100)

	for id := 0 ; id < 100 ; id ++{
		models.Worker <-id
	}
	//for len(models.Worker) > 0{
	//	w := <- models.Worker
	//	go models.WriteToDisk(w)
	//}

	go func() {
		for {
			w := <- models.Worker
			models.WriteToDisk(w)
		}
	}()
	beego.Run()
}
