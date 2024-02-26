package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"distributed-background-jobs-asynq/internal/server"
)

func main() {

	svr := server.NewServer()

	err := svr.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Could not start server: %s", err))
	}
}
