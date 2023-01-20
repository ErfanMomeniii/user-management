package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
	"strings"
	"time"
)

// AppName is the application name.
const AppName = "Translator"

// C is the global configuration instance.
var C *Config

// Config is the project root level configuration.
type Config struct {
	Timezone   *time.Location `yaml:"timezone"`
	Logger     Logger         `yaml:"logger"`
	Database   Database       `yaml:"database"`
	HTTPServer HTTPServer     `yaml:"http_server"`
}

// Logger is the logging configuration.
type Logger struct {
	Level string `yaml:"level"`
}

// Database is the database (sql) configuration.
type Database struct {
	Driver        string        `yaml:"driver"`
	Host          string        `yaml:"host"`
	Port          int           `yaml:"port"`
	Name          string        `yaml:"name"`
	User          string        `yaml:"user"`
	Password      string        `yaml:"password"`
	MaxConn       int           `yaml:"max_conn"`
	IdleConn      int           `yaml:"idle_conn"`
	Timeout       time.Duration `yaml:"timeout"`
	DialRetry     int           `yaml:"dial_retry"`
	DialTimeout   time.Duration `yaml:"dial_timeout"`
	ReadTimeout   time.Duration `yaml:"read_timeout"`
	WriteTimeout  time.Duration `yaml:"write_timeout"`
	UpdateTimeout time.Duration `yaml:"update_timeout"`
	DeleteTimeout time.Duration `yaml:"delete_timeout"`
	QueryTimeout  time.Duration `yaml:"query_timeout"`
}

// DSN returns the database DSN.
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

// HTTPServer is http server configuration.
type HTTPServer struct {
	Listen            string        `yaml:"listen"`
	ReadTimeout       time.Duration `yaml:"read_Timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}

// Init initialize the global instance.
func Init(path string) error {
	C = new(Config)
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
		return err
	}

	return v.Unmarshal(C, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			TimeLocationDecodeHook(),
		)
	})
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
