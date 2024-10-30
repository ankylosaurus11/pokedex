package main

import "os"

func exit(_ *config) error {
	os.Exit(0)
	return nil
}
