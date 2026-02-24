package config

type credsConfig struct {
	Env   string      `yaml:"env"`
	Bot   botCreds    `yaml:"bot"`
	DB    dbConfig    `yaml:"db"`
	Cache cacheConfig `yaml:"cache"`
}

type botCreds struct {
	Token string `yaml:"token"`
}

type dbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type cacheConfig struct {
	Redis redisConfig `yaml:"redis"`
}

type redisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
