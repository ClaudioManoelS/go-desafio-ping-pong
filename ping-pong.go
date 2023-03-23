package main

import (
	"fmt"
	"time"
)

func pingpong(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		text := <-c
		fmt.Println(text)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	c := make(chan string)

	go pingpong(c)
	go imprimir(c)

	var escreva string
	fmt.Scanln(&escreva)
}
