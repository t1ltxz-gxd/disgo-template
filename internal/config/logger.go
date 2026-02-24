package config

// loggerConfig contains logging settings, including the logger name and file synchronization configuration.
type loggerConfig struct {
	Name       string           `yaml:"name"`
	FileSyncer FileSyncerConfig `yaml:"fileSyncer"`
}

// FileSyncerConfig contains settings for synchronizing logs with a file,
// including file name, maximum size, number of backups, etc.
type FileSyncerConfig struct {
	Filename   string `yaml:"filename"`
	MaxSize    string `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	Compress   bool   `yaml:"compress"`
	MaxAge     int    `yaml:"maxAge"`
}
