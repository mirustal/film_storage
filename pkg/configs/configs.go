package configs

import (
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigHttp struct {
	Type   string `yaml:"type" env-default:port`
	BindIp string `yaml:"bind_ip" env-default:0.0.0.0`
	Port   string `yaml:"port" env-default:8081`
}

type ConfigPostgressDB struct {
	Host string `yaml:"host" env-default:localhost`
	Port string `yaml:"port" env-default:27017`
	Database string `yaml:"database"`
	User string `yaml:"user" env-default:admin`
	Name string `yaml:"name" env-default:admin`
	Password string `yaml:"password" env-default:admin`
}

type Config struct {
	ModeLog string `yaml: modelog env-default: jsonInfo`
	Http  ConfigHttp `yaml: http`
	PostgresDB ConfigPostgressDB  `yaml:"postgressdb"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("config_path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not open: ", configPath)
	}
	once.Do(func () {
		cfg = &Config{}
		
		if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
			cleanenv.GetDescription(cfg, nil)
		}

		if mongoHost := os.Getenv("MONGODB_HOST"); mongoHost != "" {
			cfg.PostgresDB.Host = mongoHost
		}

	})
	return cfg
}