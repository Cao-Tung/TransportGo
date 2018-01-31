package models

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
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

func WriteToDisk(id int) bool{
	Worker <-id
	message := <-Queue
	if _, err := os.Stat("output.json"); err == nil {
		f, _ := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY, 0600)
		rs, _ := json.Marshal(message)
		if _, err := f.Write(rs); err != nil {
			panic(err)
		}
		return true
	}else {
		jsonData, _  := json.Marshal(message)
		ioutil.WriteFile("output.json", jsonData, 0600)
		return true
	}
	return false
	//f, err := os.OpenFile("output.json", os.O_APPEND|os.O_WRONLY, 0600)
	//if err != nil {
	//	//jsonData, _  := json.Marshal(<-Queue)
	//	//ioutil.WriteFile("output.json", jsonData, 0600)
	//	panic(err)
	//}
	//defer f.Close()
	//
	//if err == nil {
	//	message := <-Queue
	//	rs, _ := json.Marshal(message)
	//	if _, err = f.Write(rs); err != nil {
	//		panic(err)
	//	}
	//}
}


