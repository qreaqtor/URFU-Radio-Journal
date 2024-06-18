package monitoring

import "github.com/prometheus/client_golang/prometheus"

const (
	downloadsName = "downloads"
)

type Monitoring struct {
	downloads       *prometheus.CounterVec
	monitoringTypes []string
}

// monitoringTypes is a slice of Content-Type headers, which would be monitoring for download
func NewMonitoring(monitoringTypes ...string) *Monitoring {
	downloads := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: downloadsName,
		}, []string{"filename"},
	)

	prometheus.MustRegister(downloads)

	return &Monitoring{
		downloads:       downloads,
		monitoringTypes: monitoringTypes,
	}
}

func (m *Monitoring) UpdateDownloads(filename, contentType string) {
	for _, content := range m.monitoringTypes {
		if content == contentType {
			m.downloads.WithLabelValues(filename).Inc()
			return
		}
	}
}
