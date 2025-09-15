package main

import "todolist-go/internal/bootstrap"

func main() {
	err := bootstrap.Start()
	if err != nil {
		panic(err)
	}
}
