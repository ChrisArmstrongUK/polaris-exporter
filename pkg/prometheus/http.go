package prometheus

import (
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func helpMessage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Polaris Exporter is running! Metrics are hosted at /metrics ")
}

func ListenAndServe(address string) error {
	http.HandleFunc("/", helpMessage)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
