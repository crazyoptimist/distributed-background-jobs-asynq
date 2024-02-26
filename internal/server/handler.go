package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"

	"distributed-background-jobs-asynq/internal/tasks"
)

func (s *Server) HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (s *Server) SleepHandler(c *gin.Context) {
	var sleepDto SleepDto
	if err := c.BindJSON(&sleepDto); err != nil {
		fmt.Println("Bind Error: %w", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	redisAddr := os.Getenv("REDIS_URL")

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	task, err := tasks.NewSleepTask(sleepDto.Seconds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("Could not create task: %v", err))
		return
	}

	info, err := client.Enqueue(task)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("Could not enqueue task: %v", err))
		return
	}

	message := fmt.Sprintf("Your request has been enqueued successfully: id=%s queue=%s", info.ID, info.Queue)

	c.JSON(http.StatusCreated, gin.H{
		"message": message,
	})
}
