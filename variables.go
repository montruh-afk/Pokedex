package main

import (
	"fmt"
	"os"
) 


func getWorkingDir() (string, error) {
	wdir, err := os.Getwd()
	if err != nil {
		return "./", fmt.Errorf("Error accessing current working directory: %v", err)
	}
	return wdir, nil
}
