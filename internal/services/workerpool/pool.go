package workerpool

import (
	"context"
	"time"

	"go.t1ltxz.ninja/disgo-template/internal/config"
)

type WorkerPool struct {
	tasks       chan Task
	workers     int
	maxRetries  int
	baseBackoff time.Duration
}

func NewWorkerPool(cfg *config.Config) *WorkerPool {
	return &WorkerPool{
		tasks:       make(chan Task, cfg.Worker.QueueBuffer),
		workers:     cfg.Worker.Workers,
		maxRetries:  cfg.Worker.MaxRetries,
		baseBackoff: cfg.Worker.BaseBackoff,
	}
}

func (p *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		go p.worker(ctx)
	}
}

func (p *WorkerPool) Submit(task Task) {
	p.tasks <- task
}
