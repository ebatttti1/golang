package models

import (
	"gorm.io/gorm"
)

type CommandLineConfig struct {
	gorm.Model
	Name     string `yaml:"NAME" json:"name"`
	Command  string `yaml:"COMMAND" json:"command"`
	Interval int    `yaml:"INTERVAL" json:"interval"`
	Limit    int    `yaml:"LIMIT" json:"limit"`
	CurrentCount int `yaml:"CURRENT_COUNT" json:"current_count"`
}

func (c *CommandLineConfig) BeforeCreate(db *gorm.DB) (err error) {
	c.CurrentCount = 0
	return nil
}

func (c *CommandLineConfig) BeforeUpdate(db *gorm.DB) (err error) {
	var before CommandLineConfig
	db.First(&before, c.ID)

	if before.CurrentCount < before.Limit {
		c.CurrentCount = before.CurrentCount + 1
	} else {
		c.CurrentCount = before.Limit
	}

	c.Interval = before.Interval
	return nil
}

