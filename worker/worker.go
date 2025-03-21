package worker

import (
	"fmt"
	"main/models"
	"os/exec"
	"time"

	"gorm.io/gorm"
)

var ConfigChannel = make(chan models.CommandLineConfig, 10)
//var db *gorm.DB

// func StartWorker() {
// 	go func() {
// 		for config := range ConfigChannel {
// 			go ProcessConfig(db,config)
// 		}
// 	}()
// }

func ProcessConfig(db *gorm.DB, config models.CommandLineConfig) {
	for i := config.CurrentCount; i < config.Limit; i++ {
		fmt.Printf("Executing Config [%s] - %d/%d times\n", config.Name, i+1, config.Limit)
		time.Sleep(1 * time.Second)

		cmd := exec.Command("sh", "-c", config.Command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing Config [%s]: %v\n", config.Name, err)
		} else {
			fmt.Printf("Output: %s\n", output)
		}

		time.Sleep(time.Duration(config.Interval) * time.Second)
	}

	fmt.Printf("Config [%s] finished execution\n", config.Name)
}