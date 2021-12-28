package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type LogConfig struct {
	StrictMode       bool   `yaml:"strict_mode"`
	ShowDebug        bool   `yaml:"show_debug"`
	ShowHTTPLog      bool   `yaml:"show_http_log"`
	ShowWebSocketLog bool   `yaml:"show_ws_log"`
	ErrorOutDir      string `yaml:"error_out_dir"`
}

type Config struct {
	Version     string `yaml:"version"`
	Application struct {
		URL string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"application"`
	Client struct {
		URL string `yaml:"url"`
		Key string `yaml:"key"`
	} `yaml:"client"`
	Logs LogConfig `yaml:"logs"`
}

var templateConfig = []byte(`
version: 0.0.1a

application:
	- url: ""
	- key: ""

client:
	- url: ""
	- key: ""

logs:
	- strict_mode: false
	- show_debug: false
	- show_http_log: false
	- show_ws_log: false
	- error_out_dir: ""
`)

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func GetConfig() (*Config, error) {
	fp := os.Getenv("SOAR_PATH")
	if fp == "" {
		return nil, errors.New("soar directories not found or set")
	}

	if !Exists(fp) {
		return nil, errors.New("soar config file not found")
	}

	raw, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, errors.New("couldn't open config file")
	}

	var config *Config
	yaml.Unmarshal(raw, &config)
	return config, nil
}

func CreateEnv(fp string) error {
	err := os.MkdirAll(fp, os.ModeDir)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return errors.New("missing permissions to create soar directories")
		} else {
			return errors.New("failed creating soar directories")
		}
	}

	config, err := os.Create(path.Join(fp, "config.yml"))
	if err != nil {
		fmt.Println("soar: failed creating config file; check that this application has the correct permissions.")
	}

	defer config.Close()

	os.Setenv("SOAR_PATH", fp)
	config.Write(templateConfig)

	return nil
}

func ClearOldEnv(fp string) error {
	if !Exists(fp) {
		return nil
	}

	err := os.RemoveAll(fp)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return errors.New("missing permissions to remove old soar directories")
		} else {
			return errors.New("failed removing old soar directories")
		}
	}

	return nil
}
