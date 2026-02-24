package scheduler

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"go.t1ltxz.ninja/disgo-template/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type Scheduler struct {
	cron *gocron.Scheduler
}

func New() *Scheduler {
	return &Scheduler{
		cron: gocron.NewScheduler(time.UTC),
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	logger.Info("Starting scheduler")
	s.cron.StartAsync()
	return nil
}

func (s *Scheduler) Stop(ctx context.Context) error {
	logger.Info("Stopping scheduler")
	s.cron.Stop()
	return nil
}

func (s *Scheduler) AddJob(spec string, task func()) error {
	_, err := s.cron.Cron(spec).Do(task)
	if err != nil {
		logger.Error("Failed to add job", zap.Error(err), zap.String("spec", spec))
		return err
	}
	logger.Info("Job added successfully", zap.String("spec", spec))
	return nil
}

// GetScheduler возвращает внутренний планировщик gocron
func (s *Scheduler) GetScheduler() *gocron.Scheduler {
	return s.cron
}
