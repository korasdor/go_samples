package main

import (
	"fmt"
)

//default не блокирует до получения сообщения по каналу, а вызывает сразу
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//обычный прием сообщения, не будет блокироваться, выйдет на default
	select {
	case msg := <-messages:
		fmt.Println("recieved message ", msg)
	default:
		fmt.Println("no message received")
	}

	//обычный посылка сообщения, не будет блокироваться, выйдет на default
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message ", msg)
	default:
		fmt.Println("no message received")
	}

	//обычный прием двух и более сообщений, не будет блокироваться, выйдет на default
	select {
	case msg := <-messages:
		fmt.Println("recieved message ", msg)
	case sig := <-signals:
		fmt.Println("recieved signals ", sig)
	default:
		fmt.Println("no message received")
	}

}
