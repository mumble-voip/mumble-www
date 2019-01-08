package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host                string `yaml:"host"`
	Port                int    `yaml:port`
	PublicPath          string `yaml:publicpath`
	GitHubCacheTimeMins int    `yaml:githubcachetimemins`
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfig(path string) (c *Config, err error) {
	c = NewConfig()
	err = ReadConfigInto(path, c)
	return
}

func ReadConfigInto(path string, c *Config) (err error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, c)
	return
}
