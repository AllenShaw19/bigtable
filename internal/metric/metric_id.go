package metric

type MetricLabels map[string]string

type MetricID struct {
	name   string
	labels MetricLabels
	id     string
}
