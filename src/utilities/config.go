package utilities

import (
	"flag"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ReadFile returns Configuration struct from a file path
func ReadFile(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// LoadConfigFromPath reads from path for local development
func LoadConfigFromPath(pathToConfig string) Configuration {
	cfgPath := flag.String("p", pathToConfig, "Path to config file")
	flag.Parse()
	cfg, err := ReadFile(*cfgPath)
	if err != nil {
		panic(err.Error())
	}
	return *cfg
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
	JWT    *JWT         `yaml:"jwt,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

// Database holds data necessery for database configuration
type Database struct {
	User       string `yaml:"user,omitempty"`
	Password   string `yaml:"password,omitempty"`
	Host       string `yaml:"host,omitempty"`
	Database   string `yaml:"database,omitempty"`
	Port       int    `yaml:"port,omitempty"`
	LogQueries bool   `yaml:"log_queries,omitempty"`
	Timeout    int    `yaml:"timeout_seconds,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret string `yaml:"secret,omitempty"`
	Key    string `yaml:"key,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
}
