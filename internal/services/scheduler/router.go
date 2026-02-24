package scheduler

import (
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func (s *Scheduler) RegisterJobs(jobs ...Job) error {
	for _, job := range jobs {
		job := job // capture for closure
		err := s.AddJob(job.CronExpression(), func() {
			logger.Info("Executing job", zap.String("job_name", job.Name()))
			if err := job.Execute(); err != nil {
				logger.Error("Job execution failed", zap.String("job_name", job.Name()), zap.Error(err))
			}
		})
		if err != nil {
			logger.Error("Failed to register job", zap.String("job_name", job.Name()), zap.Error(err))
			return err
		}
	}
	return nil
}
