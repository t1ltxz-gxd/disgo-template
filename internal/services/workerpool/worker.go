package workerpool

import (
	"context"
	"fmt"
	"math"
	"time"

	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func (p *WorkerPool) worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case task, ok := <-p.tasks:
			if !ok {
				return
			}

			p.executeTask(task)
		}
	}
}

func (p *WorkerPool) executeTask(task Task) {
	var err error

	for attempt := 0; attempt <= p.maxRetries; attempt++ {
		// Panic safety
		func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic: %v", r)
				}
			}()

			err = task.Execute()
		}()

		if err == nil {
			return
		}
		logger.Warn("task execution failed",
			zap.Error(err),
			zap.Int("attempt", attempt),
		)

		// Retry delay
		if attempt < p.maxRetries {
			backoff := float64(p.baseBackoff) * math.Pow(2, float64(attempt))

			time.Sleep(time.Duration(backoff))
		}
	}
	logger.Error("task failed after retries",
		zap.Error(err),
		zap.Int("max_retries", p.maxRetries),
	)
}
