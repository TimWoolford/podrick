package config

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	VersionLabels []string `yaml:"versionLabels"`
	AppNameLabel  string   `yaml:"appNameLabel"`
}

func Load() (*Config) {
	config := Config{
		VersionLabels: []string{"app_version"},
		AppNameLabel:  "app_name",
	}

	data, err1 := ioutil.ReadFile("/config/config.yaml")
	if err1 != nil {
		log.Fatalf("error: %v", err1)
	}

	err2 := yaml.Unmarshal([]byte(data), &config)
	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}

	return &config
}
