package main

import (
	"net/url"
	"time"

	"log"

	"github.com/ChrisArmstrongUK/polaris-exporter/pkg/data"
	"github.com/ChrisArmstrongUK/polaris-exporter/pkg/prometheus"
	"github.com/ChrisArmstrongUK/polaris-exporter/pkg/util"
)

func main() {
	config := util.Config{}
	config.Init()
	log.Println(config.JSON())

	data := data.Data{}

	targetURL, err := url.ParseRequestURI(config.PolarisReportTarget)
	if err != nil {
		log.Println(err)
	}

	data.MonitorTarget(*targetURL, config.FetchInterval, config.FetchTimeout)

	time.Sleep(time.Second * 5)

	prometheus.SetMetrics(config.FetchInterval, &data)

	err = prometheus.ListenAndServe(config.Address)
	if err != nil {
		log.Println(err)
	}
	log.Println("Listening on", config.Address)
}
