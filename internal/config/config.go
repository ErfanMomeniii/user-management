package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"time"
)

const AppName = "user-management"

type Config struct {
	Timezone   *time.Location `yaml:"timezone" validate:"required"`
	Logger     Logger         `yaml:"logger" validate:"required"`
	Database   Database       `yaml:"database" validate:"required"`
	HTTPServer HTTPServer     `yaml:"http_server" validate:"required"`
}

type Logger struct {
	Level string `yaml:"level"  validate:"required,oneof=debug info warn error fatal panic"`
}

type Database struct {
	Driver        string        `yaml:"driver" validate:"required"`
	Host          string        `yaml:"host" validate:"required"`
	Port          int           `yaml:"port" validate:"required"`
	Name          string        `yaml:"name" validate:"required"`
	User          string        `yaml:"user" validate:"required"`
	Password      string        `yaml:"password" validate:""`
	MaxConn       int           `yaml:"max_conn" validate:"required"`
	IdleConn      int           `yaml:"idle_conn" validate:"required"`
	Timeout       time.Duration `yaml:"timeout" validate:"required"`
	DialRetry     int           `yaml:"dial_retry" validate:"required"`
	DialTimeout   time.Duration `yaml:"dial_timeout" validate:"required"`
	ReadTimeout   time.Duration `yaml:"read_timeout" validate:"required"`
	WriteTimeout  time.Duration `yaml:"write_timeout" validate:"required"`
	UpdateTimeout time.Duration `yaml:"update_timeout" validate:"required"`
	DeleteTimeout time.Duration `yaml:"delete_timeout" validate:"required"`
	QueryTimeout  time.Duration `yaml:"query_timeout" validate:"required"`
}

func (d Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true&interpolateParams=true&collation=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		"utf8mb4_general_ci",
	)
}

type HTTPServer struct {
	Listen            string        `yaml:"listen" validate:"required"`
	ReadTimeout       time.Duration `yaml:"read_Timeout" validate:"required"`
	WriteTimeout      time.Duration `yaml:"write_timeout" validate:"required"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout" validate:"required"`
	IdleTimeout       time.Duration `yaml:"idle_timeout" validate:"required"`
}

func Init(path string) (*Config, error) {
	cfg := new(Config)
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvPrefix(strings.ToLower(AppName))
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	if path == "" {
		path = "config.defaults.yaml"
	}

	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(cfg, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			TimeLocationDecodeHook(),
		)
	}); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) Validate() error {
	return validator.New().Struct(cfg)
}

func TimeLocationDecodeHook() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		var timeLocation *time.Location
		if t != reflect.TypeOf(timeLocation) {
			return data, nil
		}

		return time.LoadLocation(data.(string))
	}
}
