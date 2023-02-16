package app

import (
	"fmt"
	"net/http"
	"os"

	"tcp-monitor/app/router"
)

func Bootstrap(port int) {
	http.HandleFunc("/tcp-status", router.GetTcpStatus)
	http.HandleFunc("/metrics", router.GetTcpMetrics)

	fmt.Printf("starting server at port %d...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("unable to start http server.")
		os.Exit(1)
	}
}
