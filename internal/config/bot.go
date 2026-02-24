package config

type botConfig struct {
	Color struct {
		Primary   string `yaml:"primary"`
		Secondary string `yaml:"secondary"`
		Tertiary  string `yaml:"tertiary"`
		Success   string `yaml:"success"`
		Warning   string `yaml:"warning"`
		Error     string `yaml:"error"`
	} `yaml:"color"`
	Status   string `yaml:"status"`
	Activity struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		URL  string `yaml:"url,omitempty"`
	}
	TestGuildID string `yaml:"testGuildID"`
	Sharding    struct {
		Enabled     bool  `yaml:"enabled"`
		ShardCount  int   `yaml:"shardCount"`
		ShardIDs    []int `yaml:"shardIDs"`
		AutoScaling bool  `yaml:"autoScaling"`
	} `yaml:"sharding"`
}
