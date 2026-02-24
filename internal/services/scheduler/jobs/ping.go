package jobs

import (
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type PingJob struct {
	name string
}

// NewPingJob creates a new sample job
func NewPingJob(name string) *PingJob {
	return &PingJob{name: name}
}

// Name returns the job name
func (e *PingJob) Name() string {
	return e.name
}

// CronExpression returns a cron expression - every minute
func (e *PingJob) CronExpression() string {
	return "* * * * *"
}

// Execute completes the task
func (e *PingJob) Execute() error {
	logger.Debug("Example job executed every minute", zap.String("job_name", e.name))
	return nil
}
