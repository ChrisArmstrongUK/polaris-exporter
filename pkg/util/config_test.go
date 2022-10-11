package util

import (
	"os"
	"testing"
)

func TestSetValues(t *testing.T) {
	type test struct {
		name string
		env  map[string]string
		want Config
	}
	testTable := []test{
		{
			"empty environment",
			map[string]string{},
			Config{
				Address:             DEFAULT_ADDRESS,
				FetchInterval:       DEFAULT_FETCH_INTERVAL,
				PolarisReportTarget: DEFAULT_REPORT_TARGET,
			},
		},
		{
			"environment",
			map[string]string{
				"POLARIS_EXPORTER_ADDRESS": "127.0.0.1:2112",
			},
			Config{
				Address:             "127.0.0.1:2112",
				FetchInterval:       DEFAULT_FETCH_INTERVAL,
				PolarisReportTarget: DEFAULT_REPORT_TARGET,
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			for key := range tt.env {
				os.Setenv(key, tt.env[key])
			}
			c := Config{}
			c.setValues()
			if c != tt.want {
				t.Errorf("got %+v, want %+v", c, tt.want)
			}
			for key := range tt.env {
				os.Unsetenv(key)
			}
		},
		)
	}
}
