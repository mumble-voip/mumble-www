package main

import (
	"flag"
	"log"
)

type Flags struct {
	ConfigPath          string
	Host                string
	Port                int
	PublicPath          string
	GitHubCacheTimeMins int
}

func NewFlags() *Flags {
	return &Flags{}
}

func ParseFlags() (f *Flags) {
	f = NewFlags()
	flag.StringVar(&f.ConfigPath, "config", "", "path to yaml config file")
	flag.StringVar(&f.Host, "host", "", "bind host address")
	flag.IntVar(&f.Port, "port", 0, "bind port number")
	flag.StringVar(&f.PublicPath, "publicpath", "", "path to public web files folder")
	flag.IntVar(&f.GitHubCacheTimeMins, "githubcachetimemins", 0, "Cache time in minutes to cache GitHub content")
	flag.Parse()
	return
}

func ParseInit() (config *Config) {
	flags := ParseFlags()

	if flags.ConfigPath == "" && (flags.Host == "" || flags.Port == 0 || flags.PublicPath == "" || flags.GitHubCacheTimeMins == 0) {
		flag.Usage()
		log.Fatalln("Missing parameter config or complete configuration parameters")
	}

	config = NewConfig()
	if flags.ConfigPath != "" {
		err := ReadConfigInto(flags.ConfigPath, config)
		if err != nil {
			log.Fatalf("Failed to read config file at %s. Error: %s", flags.ConfigPath, err)
		}
	}

	if flags.Host != "" {
		config.Host = flags.Host
	}
	if flags.Port != 0 {
		config.Port = flags.Port
	}
	if flags.PublicPath != "" {
		config.PublicPath = flags.PublicPath
	}
	if flags.GitHubCacheTimeMins != 0 {
		config.GitHubCacheTimeMins = flags.GitHubCacheTimeMins
	}
	return
}
