package transformer

import (
	m "github.com/shmul/mongodbatlas_exporter/model"

	"github.com/prometheus/client_golang/prometheus"
)

// TransformType transforms MeasurementMetadata into prometheus.ValueType
func TransformType(measurement *m.MeasurementMetadata) (prometheus.ValueType, error) {
	// According to current knowledge all mongodbatlas Measurements are Gauges.
	return prometheus.GaugeValue, nil
}
