package main

import (
	"log"
	"os"

	"github.com/hibiken/asynq"
	_ "github.com/joho/godotenv/autoload"

	"distributed-background-jobs-asynq/internal/tasks"
)

func main() {
	redisAddr := os.Getenv("REDIS_URL")

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// No concurrency in a container, for demonstration purpose
			Concurrency: 1,
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	// register handlers
	mux.HandleFunc(tasks.TypeSleep, tasks.HandleSleepTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run a worker server: %v", err)
	}
}
