package models

import (
	"fmt"
	//"encoding/json"
	//"os"
	"time"
)

type Message struct {
	Content   string
}

var Queue chan(Message)

var Worker chan(int)

func RecMessage(message Message) {
	Queue <- message
	fmt.Println(len(Queue))
}

func WriteToDisk(id int){
	fmt.Println("worker ", id, <- Queue)
	time.Sleep(time.Second)
	Worker <-id
	//f, err := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY, 0600)
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//time.Sleep(time.Second)
	//message := <- Queue
	//rs,_  := json.Marshal(message)
	//if _, err = f.Write(rs); err != nil {
	//	panic(err)
	//}
}


