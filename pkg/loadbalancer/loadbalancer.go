package loadbalancer

import (
	"net/http"

	"github.com/hosseinfakhari/wire/internal/service"
)

func LB(w http.ResponseWriter, r *http.Request) {
	peer := service.Pool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}
