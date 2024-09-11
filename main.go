package main

import "os"

func main() {
	err := runServer()
	if err != nil {
		os.Exit(1)
	}
}
