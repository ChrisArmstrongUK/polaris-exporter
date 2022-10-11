package prometheus

import (
	"time"

	"github.com/ChrisArmstrongUK/polaris-exporter/pkg/data"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	polarisOverallScore = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "polaris_overall_score_percentage",
		Help: "The overall score percentage",
	})
	polarisSuccessCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "polaris_overall_success_count",
		Help: "The overall success count",
	})
	polarisWarningCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "polaris_overall_warning_count",
		Help: "The overall warning count",
	})
	polarisDangerCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "polaris_overall_danger_count",
		Help: "The overall danger count",
	})
)

func SetMetrics(interval time.Duration, d *data.Data) {
	go func() {
		for {
			polarisOverallScore.Set(float64(d.AuditData.GetSummary().GetScore()))
			polarisSuccessCount.Set(float64(d.AuditData.GetSummary().Successes))
			polarisWarningCount.Set(float64(d.AuditData.GetSummary().Warnings))
			polarisDangerCount.Set(float64(d.AuditData.GetSummary().Dangers))

			time.Sleep(interval)
		}
	}()
}
