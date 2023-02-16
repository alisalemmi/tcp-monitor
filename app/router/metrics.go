package router

import (
	"fmt"
	"net/http"

	"tcp-monitor/app/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetTcpMetrics(w http.ResponseWriter, r *http.Request) {
	tcp, err := service.GetTcpStatus()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "an error has been occurred when try to get tcp status")
	}

	service.SetTcpMetrics(tcp)
	promhttp.Handler().ServeHTTP(w, r)
}
