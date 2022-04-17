package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func getConfig() (*Config, error) {
	yfile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		fmt.Println(err)
	}
	config := &Config{}

	err = yaml.Unmarshal(yfile, config)

	if err != nil {
		fmt.Println(err)
	}

	return config, err
}

type Config struct {
	AppConfig AppConfig `yaml:"app"`
	DBConfig  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	AppPort string `yaml:"port"`
	AppName string `yaml:"name"`
}

type DBConfig struct {
	DBHost       string `yaml:"host"`
	DBPort       string `yaml:"port"`
	DBName       string `yaml:"name"`
	DBUsername   string `yaml:"username"`
	DBPassword   string `yaml:"password"`
	RunMigration bool   `yaml:"run_migration"`
}

type RedisConfig struct {
	RedisHost     string `yaml:"host"`
	RedisPort     string `yaml:"port"`
	RedisUsername string `yaml:"username"`
	RedisPassword string `yaml:"password"`
}
