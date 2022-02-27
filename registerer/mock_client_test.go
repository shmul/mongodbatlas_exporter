package registerer

import (
	"github.com/shmul/mongodbatlas_exporter/measurer"
	"github.com/shmul/mongodbatlas_exporter/model"

	internal "github.com/shmul/mongodbatlas_exporter/mongodbatlas"

	"go.mongodb.org/atlas/mongodbatlas"
)

type MockClient struct {
	processes []*mongodbatlas.Process
}

func (c *MockClient) GetDiskMeasurements(*measurer.Process, *measurer.Disk) error {
	return nil
}
func (c *MockClient) GetProcessMeasurements(measurer.Process) (map[model.MeasurementID]*model.Measurement, error) {
	return nil, nil
}
func (c *MockClient) GetDiskMeasurementsMetadata(*measurer.Process, *measurer.Disk) (map[model.MeasurementID]*model.MeasurementMetadata, error) {
	return nil, nil
}
func (c *MockClient) GetProcessMeasurementsMetadata(*measurer.Process) *internal.HTTPError {
	return nil
}
func (c *MockClient) ListProcesses() ([]*mongodbatlas.Process, *internal.HTTPError) {
	return c.processes, nil
}
func (c *MockClient) ListDisks(*mongodbatlas.Process) ([]*mongodbatlas.ProcessDisk, *internal.HTTPError) {
	return nil, nil
}
