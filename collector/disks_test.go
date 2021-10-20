package collector

// //TestDisksCollector initializes a disk collector and then checks
// //that the scraped metrics output have the correct values, units, and labels.
// //This is different from TestDesc which is checking the collector's Describe function.
// //This is a test of the Collect function.
// func TestDisksCollector(t *testing.T) {
// 	assert := assert.New(t)
// 	value := float32(3.14)
// 	mock := &MockClient{}
// 	mock.givenDisksMeasurements = getGivenMeasurements(&value)
// 	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

// 	disks, err := NewDisks(logger, mock)
// 	assert.NotNil(disks)
// 	assert.NoError(err)

// 	metricsCh := make(chan prometheus.Metric, 99)
// 	defer close(metricsCh)
// 	expectedMetrics := getExpectedDisksMetrics(float64(value))
// 	disks.Collect(metricsCh)
// 	var resultingMetrics []prometheus.Metric
// 	for len(metricsCh) > 0 {
// 		resultingMetrics = append(resultingMetrics, <-metricsCh)
// 	}
// 	assert.Equal(len(expectedMetrics), len(resultingMetrics))
// 	expectedMetricMap := convertMetrics(expectedMetrics)
// 	actualMetricMap := convertMetrics(resultingMetrics)

// 	for key, value := range expectedMetricMap {
// 		assert.Equal(value, actualMetricMap[key], "key %s", key)
// 	}
// }

// func getExpectedDisksMetrics(value float64) []prometheus.Metric {
// 	measurer := measurer.Disk{
// 		PartitionName: "testPartition",
// 	}
// 	fqNames := []string{
// 		prometheus.BuildFQName(namespace, "disks_stats", "disk_partition_iops_read_ratio"),
// 		prometheus.BuildFQName(namespace, disksPrefix, "up"),
// 		prometheus.BuildFQName(namespace, disksPrefix, "scrapes_total"),
// 		prometheus.BuildFQName(namespace, disksPrefix, "scrape_failures_total"),
// 		prometheus.BuildFQName(namespace, disksPrefix, "measurement_transformation_failures_total"),
// 	}

// 	variableLabels := [][]string{
// 		nil,
// 		nil,
// 		nil,
// 		nil,
// 		//measurement_transformation_failures_total has variable labels
// 		//to indicate the metric and the error.
// 		{"atlas_metric", "error"},
// 	}

// 	variableLabelValues := [][]string{
// 		nil,
// 		nil,
// 		nil,
// 		nil,
// 		//these correspond to "atlas_metric" and "error" variable labels
// 		//for measurement_transformation_failures_total
// 		{"DISK_PARTITION_SPACE_USED", "no_data"},
// 	}

// 	help := []string{
// 		"Original measurements.name: 'DISK_PARTITION_IOPS_READ'. " + defaultHelp,
// 		upHelp,
// 		totalScrapesHelp,
// 		scrapeFailuresHelp,
// 		measurementTransformationFailuresHelp,
// 	}

// 	values := []float64{
// 		value,
// 		1,
// 		1,
// 		3,
// 		1,
// 	}

// 	results := make([]prometheus.Metric, len(values))

// 	for i := range results {
// 		//measurer.PromConstLabels ensures that all the identifying fields
// 		//for a particular instance of a disk are added to every metric.
// 		results[i] = prometheus.MustNewConstMetric(
// 			prometheus.NewDesc(fqNames[i], help[i], variableLabels[i], measurer.PromConstLabels()),
// 			prometheus.GaugeValue,
// 			values[i],
// 			variableLabelValues[i]...,
// 		)
// 	}
// 	return results
// }
