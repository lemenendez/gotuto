package main

import (
	"log"
	"runtime"
)

const info = `Application % starting. The binary was build by GO: %s`

func main() {
	log.Printf(info, "Version", runtime.Version())
}