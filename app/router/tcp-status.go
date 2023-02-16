package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tcp-monitor/app/dto"
	"tcp-monitor/app/service"
)

func GetTcpStatus(w http.ResponseWriter, r *http.Request) {
	tcp, err := service.GetTcpStatus()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "an error has been occurred when try to get tcp status")
	}

	res := dto.RetransmissionRatio{
		Sent:         tcp["OutSegs"],
		Retransmited: tcp["RetransSegs"],
		Ratio:        service.GetRetransmissionRatio(tcp),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
