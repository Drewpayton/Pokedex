package main

import "os"

// Waits for user to input "exit" and then uses Exit method to end the program.
func commandExit(cfg *config, areaName string) error {
	os.Exit(0)
	return nil
}