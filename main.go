// Package main contains the entry functionality for the program
package main

import (
	"fmt"

	config "github.com/can-ek/gator/config"
)

func main() {
	configuration, err := config.Read()
	if err != nil {
		fmt.Println("Error while reading config:", err)
		return
	}

	err = configuration.SetUser("can-ek")
	if err != nil {
		fmt.Println("Error while updating username:", err)
		return
	}

	configuration_2, err := config.Read()
	if err != nil {
		fmt.Println("Error while reading config:", err)
		return
	}

	fmt.Println("Current username:", configuration_2.CurrentUsername)
}
