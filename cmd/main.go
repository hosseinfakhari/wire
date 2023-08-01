package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/hosseinfakhari/wire/internal/service"
	"github.com/hosseinfakhari/wire/pkg/loadbalancer"
	"github.com/hosseinfakhari/wire/pkg/server"
)

func main() {
	lg := log.New(os.Stdout, "[WIRE] ", log.Default().Flags())

	u1, _ := url.Parse("http://localhost:3000")
	s1 := &server.Backend{
		URL:          u1,
		Alive:        true,
		ReverseProxy: httputil.NewSingleHostReverseProxy(u1),
	}

	u2, _ := url.Parse("http://localhost:3002")
	s2 := &server.Backend{
		URL:          u2,
		Alive:        true,
		ReverseProxy: httputil.NewSingleHostReverseProxy(u2),
	}

	u3, _ := url.Parse("http://localhost:3003")
	s3 := &server.Backend{
		URL:          u3,
		Alive:        true,
		ReverseProxy: httputil.NewSingleHostReverseProxy(u3),
	}

	service.Pool = server.NewServerPool(lg)

	service.Pool.AddBackend(s1)
	service.Pool.AddBackend(s2)
	service.Pool.AddBackend(s3)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 80),
		Handler: http.HandlerFunc(loadbalancer.LB),
	}

	lg.Println("Simple RP/LB Started...")
	if err := server.ListenAndServe(); err != nil {
		lg.Fatal(err)
	}
}
