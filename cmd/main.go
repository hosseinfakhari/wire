package main

import (
	"log"
	"os"
)

func main() {
	lg := log.New(os.Stdout, "[WIRE] ", log.Default().Flags())
	lg.Println("Simple RP/LB...")
}
