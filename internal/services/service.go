package services

import (
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.t1ltxz.ninja/disgo-template/internal/services/scheduler"
	"go.t1ltxz.ninja/disgo-template/internal/services/scheduler/jobs"
	"go.uber.org/zap"
)

func initializeService(s *scheduler.Scheduler) {
	jobs := []scheduler.Job{
		jobs.NewPingJob("ping-job"),
	}

	if err := s.RegisterJobs(jobs...); err != nil {
		logger.Error("Failed to initialize jobs", zap.Error(err))
	}
}
