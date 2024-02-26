package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// Task type(s)
const (
	TypeSleep = "sleep"
)

type SleepPayload struct {
	Seconds uint
}

func NewSleepTask(seconds uint) (*asynq.Task, error) {
	payload, err := json.Marshal(SleepPayload{Seconds: seconds})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSleep, payload, asynq.MaxRetry(3)), nil
}

func HandleSleepTask(ctx context.Context, t *asynq.Task) error {
	var payload SleepPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Println("Unmarshal payload failed:", err)
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	// Perform task
	time.Sleep(time.Duration(payload.Seconds) * time.Second)
	taskId := t.ResultWriter().TaskID()
	log.Printf("Task %s finished. Slept %d seconds.", taskId, payload.Seconds)

	return nil
}
