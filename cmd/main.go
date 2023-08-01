package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hosseinfakhari/wire/pkg/loadbalancer"
)

func main() {
	lg := log.New(os.Stdout, "[WIRE] ", log.Default().Flags())

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 80),
		Handler: http.HandlerFunc(loadbalancer.LB),
	}

	lg.Println("Simple RP/LB Started...")
	if err := server.ListenAndServe(); err != nil {
		lg.Fatal(err)
	}
}
