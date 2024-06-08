package monitoring

import "github.com/prometheus/client_golang/prometheus"

const (
	downloadsName = "downloads"
)

type Monitoring struct {
	downloads *prometheus.CounterVec
}

func NewMonitoring() *Monitoring {
	downloads := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: downloadsName,
		}, []string{"filename"},
	)

	prometheus.MustRegister(downloads)

	return &Monitoring{
		downloads: downloads,
	}
}

func (m *Monitoring) UpdateDownloads(filename string) {
	m.downloads.WithLabelValues(filename).Inc()
}
