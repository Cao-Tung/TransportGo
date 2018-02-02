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

	// init Message Queue
	models.Queue = make(chan models.Message, models.MaxLenQueue)

	// init Worker pool
	models.Worker = make(chan int, 10)
	for id := 0 ; id < 3 ; id ++{
		models.Worker <-id
	}
	//for len(models.Worker) > 0{
	//	w := <- models.Worker
	//	go models.WriteToDisk(w)
	//}

	// Worker execute message in pool, write to disk
	go func() {
		for {
			w := <- models.Worker
			models.WriteToDisk(w)
		}
	}()
	beego.Run()
}
