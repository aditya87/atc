package gcng_test

import (
	"code.cloudfoundry.org/garden"
	"github.com/concourse/atc/dbng"
	. "github.com/concourse/atc/gcng"

	"errors"

	"code.cloudfoundry.org/garden/gardenfakes"
	"github.com/concourse/atc/dbng/dbngfakes"
	"github.com/concourse/atc/gcng/gcngfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContainerCollector", func() {

	var (
		fakeContainerProvider *gcngfakes.FakeDBContainer
		workerProvider        *dbngfakes.FakeWorkerFactory
		gardenClientFactory   func(w *dbng.Worker) garden.Client

		fakeGardenClient *gardenfakes.FakeClient

		c *ContainerCollector
	)

	BeforeEach(func() {
		fakeContainerProvider = new(gcngfakes.FakeDBContainer)
		workerProvider = new(dbngfakes.FakeWorkerFactory)
		fakeGardenClient = new(gardenfakes.FakeClient)

		gardenClientFactory = func(w *dbng.Worker) garden.Client {
			return fakeGardenClient
		}

		c = &ContainerCollector{
			ContainerProvider:   fakeContainerProvider,
			WorkerProvider:      workerProvider,
			GardenClientFactory: gardenClientFactory,
		}
	})

	Describe("Run", func() {
		var (
			err error
		)

		JustBeforeEach(func() {
			err = c.Run()
		})

		BeforeEach(func() {
			fakeContainerProvider.FindBuildContainersToDeleteReturns(nil, nil)
		})

		It("looks for containers to transition to the deleting state", func() {
			Expect(fakeContainerProvider.FindBuildContainersToDeleteCallCount()).To(Equal(1))
		})

		Context("when looking for containers to delete fails", func() {
			BeforeEach(func() {
				fakeContainerProvider.FindBuildContainersToDeleteReturns(nil, errors.New("oh no"))
			})
		})

		Context("when containers to delete are found", func() {

			var (
				fakeCreatedContainer *gcngfakes.FakeCreatedContainer
			)

			BeforeEach(func() {
				fakeCreatedContainer = new(gcngfakes.FakeCreatedContainer)
				containersToDelete := []*CreatedContainer{fakeCreatedContainer}
				fakeContainerProvider.FindBuildContainersToDeleteReturns(containersToDelete, nil)
			})

			It("changes all the containers to deleting state", func() {
				Expect(fakeCreatedContainer.DestroyingCallCount()).To(Equal(1))
			})

		})

		It("finds all containers in deleting state", func() {
			Expect(fakeContainerProvider.FindBuildContainersToDeleteCallCount()).To(Equal(1))
		})

		Context("when deleting containers are found", func() {

			var (
				fakeDestroyingContainer *gcngfakes.FakeDestroyingContainer
			)

			BeforeEach(func() {
				fakeDestroyingContainer = new(gcngfakes.FakeDestroyingContainer)
				deletingContainers := []*DestroyingContainer{fakeDestroyingContainer}
				fakeContainerProvider.FindContainersMarkedForDeletionReturns(deletingContainers, nil)
			})

			Context("given a garden client for the worker can be found", func() {
				It("for each container it tells the garden client to delete it", func() {
					Expect(fakeGardenClient.DestroyCallCount()).To(Equal(1))
				})

				Context("when the container is removed by garden", func() {
					It("removes the container from the db", func() {

					})
				})

				Context("when removing the container from garden fails", func() {
					It("doesn't remove the container from the db", func() {

					})
				})

			})

			Context("when the worker cannot be found", func() {
				// some stuff happens to the worker maybe?
				// or nothing..
			})
		})

	})

})
