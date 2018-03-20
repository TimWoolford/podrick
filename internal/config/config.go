package config

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Config struct {
	VersionLabels        []string          `yaml:"versionLabels"`
	AppNameLabel         string            `yaml:"appNameLabel"`
	LabelFile            string            `yaml:"labelFile"`
	StatusPathAnnotation string            `yaml:"statusPathAnnotation"`
	DefaultStatusPath    string            `yaml:"defaultStatusPath"`
	ClusterIdentifier    []string          `yaml:"clusterIdentifier"`
	ClusterNaming        map[string]string `yaml:"clusterNaming"`
	PodLabels            map[string]string
	ClusterName          string
}

func Load(configFile string) (*Config) {
	data, err1 := ioutil.ReadFile(configFile)
	if err1 != nil {
		log.Fatalf("error: %v", err1)
	}

	config := &Config{
		VersionLabels:     []string{"app_version"},
		AppNameLabel:      "app_name",
		DefaultStatusPath: "/status",
	}

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config.PodLabels = readPodLabels(config.LabelFile)
	config.ClusterName = clusterName(config)

	return config
}

func clusterName(c *Config) string {
	ids := make([]string, len(c.ClusterIdentifier))
	for i, id := range c.ClusterIdentifier {
		ids[i] = c.PodLabels[id]
	}

	name, present := c.ClusterNaming[strings.Join(ids, "-")]

	if present {
		return name
	}

	return "UNK"
}
