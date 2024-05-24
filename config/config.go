package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host    string
	Port    string
	Timeout uint32
	Workdir string
}

type Service struct {
	Method   string
	Url      string
	NeedAuth bool
}

type Config struct {
	Server Server
	Api    map[string]Service
}

var config Config

func Get() *Config {
	return &config
}

func Location() string {
	return config.Server.Host + ":" + config.Server.Port
}

func init() {
	file, _ := os.ReadFile("golang_pr.yml")
	err := yaml.Unmarshal(file, &config)
	ex, _ := os.Executable()
	config.Server.Workdir = filepath.Dir(ex)
	if err != nil {
		panic("Error to read yaml.")
	} else {
		fmt.Println(config.Server.Workdir)
	}
}
