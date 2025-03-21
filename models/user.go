package models

import "gorm.io/gorm"

type CommandLineConfig struct {
	gorm.Model
	Name     string `yaml:"NAME" json:"name"`
	Command  string `yaml:"COMMAND" json:"command"`
	Interval int    `yaml:"INTERVAL" json:"interval"`
	Limit    int    `yaml:"LIMIT" json:"limit"`
}