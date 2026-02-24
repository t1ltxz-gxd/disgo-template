package config

import "time"

type workerConfig struct {
	MaxRetries  int           `yaml:"maxRetries"`
	BaseBackoff time.Duration `yaml:"baseBackoff"`
	QueueBuffer int           `yaml:"queueBuffer"`
	Workers     int           `yaml:"workers"`
}
