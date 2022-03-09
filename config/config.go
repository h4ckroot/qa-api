package config

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conf *Config
	DB   *gorm.DB
)

// Config object
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	DB DBConfig `yaml:"db"`
}

// LoadConfig returns a new decoded Config struct
func LoadConfig(configFile string) (*Config, error) {
	// Create config structure
	conf := &Config{}

	// Open config file
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, conf)
	Conf = conf
	return conf, err
}

func InitDBClient(dsn string) *gorm.DB {
	gormdb, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal().Msg("failed to init DB")
	}
	DB = gormdb
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
