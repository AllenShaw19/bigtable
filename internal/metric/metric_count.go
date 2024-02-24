package metric

import "github.com/allen-shaw/bigtable/internal/util"

type MetricCounter struct {
	couter     *util.Counter
	registered bool
	metricID   MetricID
	// typeList
}
