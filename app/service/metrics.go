package service

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var sentGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "tcp_sent",
	Help: "tcp sent segments",
})

var retransmissionGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "tcp_retransmited",
	Help: "tcp retransmited segments",
})

var ratioGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "tcp_retransmission_ratio",
	Help: "tcp retransmission ratio",
})

func SetTcpMetrics(tcp map[string]int) {
	sentGauge.Set(float64(tcp["OutSegs"]))
	retransmissionGauge.Set(float64(tcp["RetransSegs"]))
	ratioGauge.Set(GetRetransmissionRatio(tcp))
}
