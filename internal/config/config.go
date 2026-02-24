package config

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	v      *viper.Viper
	Logger loggerConfig `yaml:"logger"`
	Bot    botConfig    `yaml:"bot"`
	Creds  credsConfig  `yaml:"creds"`
	Worker workerConfig `yaml:"worker"`
}

func NewConfig() (*Config, error) {

	v := viper.New()

	_ = godotenv.Load()

	configWorkspaces := []struct {
		ConfigName   string
		WorkspaceKey string
		Type         string
		Path         string
	}{
		{"logger", "logger", "yaml", "config"},
		{"bot", "bot", "yaml", "config"},
		{"creds", "creds", "yaml", "config"},
	}

	for _, ws := range configWorkspaces {

		var filePath string
		possible := []string{
			ws.Path + "/" + ws.ConfigName + ".yaml",
			ws.Path + "/" + ws.ConfigName + ".yml",
		}

		for _, p := range possible {
			if _, err := os.Stat(p); err == nil {
				filePath = p
				break
			}
		}

		if filePath == "" {
			log.Printf("Config file not found: %s.(yaml|yml)", ws.ConfigName)
			continue
		}

		raw, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read %s: %v", filePath, err)
			continue
		}

		replaced := replaceEnvInConfig(raw)

		temp := viper.New()
		temp.SetConfigType("yaml")

		if err := temp.ReadConfig(strings.NewReader(string(replaced))); err != nil {
			log.Printf("Failed to parse %s: %v", filePath, err)
			continue
		}

		for key, value := range temp.AllSettings() {
			v.Set(ws.WorkspaceKey+"."+key, value)
		}
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetConfigFile(".env")
	v.SetConfigType("env")

	for _, key := range v.AllKeys() {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		_ = v.BindEnv(key, envKey)
	}

	config := &Config{v: v}

	if err := v.UnmarshalKey("logger", &config.Logger); err != nil {
		log.Fatalf("Error parsing logger configuration: %v", err)
	}
	if err := v.UnmarshalKey("bot", &config.Bot); err != nil {
		log.Fatalf("Error parsing bot configuration: %v", err)
	}
	if err := v.UnmarshalKey("creds", &config.Creds); err != nil {
		log.Fatalf("Error parsing bot configuration: %v", err)
	}

	return config, nil
}

func replaceEnvInConfig(body []byte) []byte {
	re := regexp.MustCompile(`\$\{([^}]+)}`)

	return re.ReplaceAllFunc(body, func(match []byte) []byte {
		expr := string(match[2 : len(match)-1]) // внутри ${}

		// ${VAR!}
		if strings.HasSuffix(expr, "!") {
			key := strings.TrimSuffix(expr, "!")
			val := os.Getenv(key)
			if val == "" {
				log.Fatalf("config error: required env %s not set", key)
			}
			return []byte(val)
		}

		// ${VAR?custom error}
		if strings.Contains(expr, "?") {
			parts := strings.SplitN(expr, "?", 2)
			key := parts[0]
			msg := parts[1]

			val := os.Getenv(key)
			if val == "" {
				log.Fatalf("config error: %s (env %s)", msg, key)
			}
			return []byte(val)
		}

		// ${VAR:default}
		if strings.Contains(expr, ":") {
			parts := strings.SplitN(expr, ":", 2)
			key := parts[0]
			def := parts[1]

			val := os.Getenv(key)
			if val == "" {
				val = def
			}
			return []byte(val)
		}

		// ${VAR}
		val := os.Getenv(expr)
		return []byte(val)
	})
}
