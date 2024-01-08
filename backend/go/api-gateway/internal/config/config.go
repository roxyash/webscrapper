package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"strings"
	"webscrapper/pkg/logging"
)

type Settings struct {
	App struct {
		Mode     string `mapstructure:"mode"`
		Host     string `mapstructure:"auth_host"`
		Port     string `mapstructure:"auth_port"`
		BasePath string `mapstructure:"base_path"`
	}

	Auth struct {
		Host string `mapstructure:"auth_host"`
		Port string `mapstructure:"auth_port"`
	}

	Postgres struct {
		Port     string `mapstructure:"postgres_port"`
		Host     string `mapstructure:"postgres_host"`
		User     string `mapstructure:"postgres_user"`
		Password string `mapstructure:"postgres_password"`
		Database string `mapstructure:"postgres_database"`
		Schema   string `mapstructure:"postgres_schema"`
	}
}

var (
	config *Settings
	logger = logging.GetLogger()
)

func init() {
	v := createViper()

	var c Settings

	c.App.Mode = v.GetString("MODE")
	c.App.Host = v.GetString("API_GATEWAY_HOST")
	c.App.Port = v.GetString("API_GATEWAY_PORT")
	c.App.BasePath = v.GetString("BASE_PATH")

	c.Auth.Host = v.GetString("AUTH_HOST")
	c.Auth.Port = v.GetString("AUTH_PORT")

	c.Postgres.Port = v.GetString("POSTGRES_PORT")
	c.Postgres.User = v.GetString("POSTGRES_USER")
	c.Postgres.Password = v.GetString("POSTGRES_PASSWORD")
	c.Postgres.Database = v.GetString("POSTGRES_DATABASE")
	c.Postgres.Schema = v.GetString("POSTGRES_SCHEMA")

	config = &c
}

func GetConfig() *Settings {
	return config
}

func createViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("env")
	v.AutomaticEnv()
	err := godotenv.Load("../../config/.env")
	if err != nil {
		logger.Fatalf("load envs with err - %s", err.Error())
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}

func (s *Settings) IsDevMod() bool {
	return s.App.Mode == "development"
}

func (s *Settings) ToYAML() (string, error) {
	bs, err := yaml.Marshal(s)
	return string(bs), err
}

func (s *Settings) ToJSON() (string, error) {
	bs, err := json.Marshal(s)
	return string(bs), err
}
