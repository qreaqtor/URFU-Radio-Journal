package monitoring

import "github.com/prometheus/client_golang/prometheus"

const (
	downloadsName = "downloads"
)

type Monitoring struct {
	downloads prometheus.CounterVec
}

func NewMonitoring() *Monitoring {
	return &Monitoring{
		downloads: *prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: downloadsName,
			}, []string{"filename"},
		),
	}
}

func (m *Monitoring) Update(filename string) {
	m.downloads.WithLabelValues(filename).Inc()
}
