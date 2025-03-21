package worker

import (
	"fmt"
	"os/exec"
	"time"
	"main/models"
)

var ConfigChannel = make(chan models.CommandLineConfig, 10)

func StartWorker() {
	go func() {
		for config := range ConfigChannel {
			go processConfig(config)
		}
	}()
}

func processConfig(config models.CommandLineConfig) {
	for i := 0; i < config.Limit; i++ {
		fmt.Printf("Executing Config [%s] - %d/%d times\n", config.Name, i+1, config.Limit)

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