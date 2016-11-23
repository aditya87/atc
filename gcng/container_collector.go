package gcng

import (
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/garden/client"
	"code.cloudfoundry.org/garden/client/connection"
	"github.com/concourse/atc/dbng"
)

//go:generate counterfeiter . DBContainer
type DBContainer interface {
	FindBuildContainersToDelete() ([]*CreatedContainer, error)
	DoForContainersMarkedForDeletion(func(*dbng.DestroyingContainer) error) error
}

//go:generate counterfeiter . CreatedContainer
type CreatedContainer interface {
	Destroying() (*DestroyingContainer, error)
}

//go:generate counterfeiter . DestroyingContainer
type DestroyingContainer interface {
	Destroy() (bool, error)
}

type ContainerCollector struct {
	ContainerProvider   DBContainer
	WorkerProvider      dbng.WorkerFactory
	GardenClientFactory GardenClientFactory
}

type GardenClientFactory func(*dbng.Worker) garden.Client

func NewGardenClientFactory() GardenClientFactory {
	return func(w *dbng.Worker) garden.Client {
		gconn := connection.New("tcp", w.GardenAddr)
		return client.New(gconn)
	}
}

func (c *ContainerCollector) Run() error {
	// get containers that should be marked for deletion
	// --> not used in last failed build of a job
	// --> used for a resource that doesn't exist
	// --> part of a pipeline that doesn't exist

	// mark those containers for deletion
	// find containers marked for deletion

	// FindContainersMarkedForDeletion
	// try and delete from garden
	// create client for the worker the container is on
	// gardenClient.Delete(handle)

	// once deleting from garden succeeds delete from db

	return nil
}
