package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	Address                string
	FetchTimeout           time.Duration
	FetchInterval          time.Duration
	PolarisReportTarget    string
	PolarisResultsPath     string
	PolarisResultsFilePath string
}

const (
	DEFAULT_ADDRESS        = "127.0.0.1:2112"
	DEFAULT_FETCH_TIMEOUT  = time.Second * 20
	DEFAULT_FETCH_INTERVAL = time.Second * 30
	DEFAULT_REPORT_TARGET  = "http://127.0.0.1:8080/results.json"
)

func (c *Config) Init() {
	c.setValues()
}

func (c *Config) setValues() {
	environmentVariablePrefix := "POLARIS_EXPORTER_"

	address, isSet := os.LookupEnv(fmt.Sprintf("%sADDRESS", environmentVariablePrefix))
	if isSet {
		c.Address = address
	} else {
		c.Address = DEFAULT_ADDRESS
	}

	fetchTimeout, isSet := os.LookupEnv(fmt.Sprintf("%sFETCH_TIMEOUT", environmentVariablePrefix))
	if isSet {
		parsedFetchTimeout, err := time.ParseDuration(fetchTimeout)
		if err != nil {
			log.Println(err, "Falling back to default fetch timeout", DEFAULT_FETCH_INTERVAL)
		}
		c.FetchInterval = parsedFetchTimeout
	} else {
		c.FetchInterval = DEFAULT_FETCH_TIMEOUT
	}

	fetchInterval, isSet := os.LookupEnv(fmt.Sprintf("%sFETCH_INTERVAL", environmentVariablePrefix))
	if isSet {
		parsedFetchInterval, err := time.ParseDuration(fetchInterval)
		if err != nil {
			log.Println(err, "Falling back to default fetch interval", DEFAULT_FETCH_INTERVAL)
		}
		c.FetchInterval = parsedFetchInterval
	} else {
		c.FetchInterval = DEFAULT_FETCH_INTERVAL
	}

	reportTarget, isSet := os.LookupEnv(fmt.Sprintf("%sREPORT_TARGET", environmentVariablePrefix))
	if isSet {
		c.PolarisReportTarget = reportTarget
	} else {
		c.PolarisReportTarget = DEFAULT_REPORT_TARGET
	}
}

func (c *Config) JSON() (string, error) {
	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", fmt.Errorf("STRING: %v", err)
	}
	return string(json), nil
}
