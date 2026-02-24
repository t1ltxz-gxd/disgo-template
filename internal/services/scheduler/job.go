package scheduler

type Job interface {
	// Name returns the job name
	Name() string
	// CronExpression Returns an expression in cron format (e.g., "0 0 * * *" - every day at midnight)
	CronExpression() string
	// Execute completes the task
	Execute() error
}
