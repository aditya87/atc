// This file was generated by counterfeiter
package gcngfakes

import (
	"sync"

	"github.com/concourse/atc/gcng"
)

type FakeDBContainer struct {
	FindBuildContainersToDeleteStub        func() ([]*gcng.CreatedContainer, error)
	findBuildContainersToDeleteMutex       sync.RWMutex
	findBuildContainersToDeleteArgsForCall []struct{}
	findBuildContainersToDeleteReturns     struct {
		result1 []*gcng.CreatedContainer
		result2 error
	}
	FindContainersMarkedForDeletionStub        func() ([]*gcng.DestroyingContainer, error)
	findContainersMarkedForDeletionMutex       sync.RWMutex
	findContainersMarkedForDeletionArgsForCall []struct{}
	findContainersMarkedForDeletionReturns     struct {
		result1 []*gcng.DestroyingContainer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBContainer) FindBuildContainersToDelete() ([]*gcng.CreatedContainer, error) {
	fake.findBuildContainersToDeleteMutex.Lock()
	fake.findBuildContainersToDeleteArgsForCall = append(fake.findBuildContainersToDeleteArgsForCall, struct{}{})
	fake.recordInvocation("FindBuildContainersToDelete", []interface{}{})
	fake.findBuildContainersToDeleteMutex.Unlock()
	if fake.FindBuildContainersToDeleteStub != nil {
		return fake.FindBuildContainersToDeleteStub()
	} else {
		return fake.findBuildContainersToDeleteReturns.result1, fake.findBuildContainersToDeleteReturns.result2
	}
}

func (fake *FakeDBContainer) FindBuildContainersToDeleteCallCount() int {
	fake.findBuildContainersToDeleteMutex.RLock()
	defer fake.findBuildContainersToDeleteMutex.RUnlock()
	return len(fake.findBuildContainersToDeleteArgsForCall)
}

func (fake *FakeDBContainer) FindBuildContainersToDeleteReturns(result1 []*gcng.CreatedContainer, result2 error) {
	fake.FindBuildContainersToDeleteStub = nil
	fake.findBuildContainersToDeleteReturns = struct {
		result1 []*gcng.CreatedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeDBContainer) FindContainersMarkedForDeletion() ([]*gcng.DestroyingContainer, error) {
	fake.findContainersMarkedForDeletionMutex.Lock()
	fake.findContainersMarkedForDeletionArgsForCall = append(fake.findContainersMarkedForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("FindContainersMarkedForDeletion", []interface{}{})
	fake.findContainersMarkedForDeletionMutex.Unlock()
	if fake.FindContainersMarkedForDeletionStub != nil {
		return fake.FindContainersMarkedForDeletionStub()
	} else {
		return fake.findContainersMarkedForDeletionReturns.result1, fake.findContainersMarkedForDeletionReturns.result2
	}
}

func (fake *FakeDBContainer) FindContainersMarkedForDeletionCallCount() int {
	fake.findContainersMarkedForDeletionMutex.RLock()
	defer fake.findContainersMarkedForDeletionMutex.RUnlock()
	return len(fake.findContainersMarkedForDeletionArgsForCall)
}

func (fake *FakeDBContainer) FindContainersMarkedForDeletionReturns(result1 []*gcng.DestroyingContainer, result2 error) {
	fake.FindContainersMarkedForDeletionStub = nil
	fake.findContainersMarkedForDeletionReturns = struct {
		result1 []*gcng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeDBContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findBuildContainersToDeleteMutex.RLock()
	defer fake.findBuildContainersToDeleteMutex.RUnlock()
	fake.findContainersMarkedForDeletionMutex.RLock()
	defer fake.findContainersMarkedForDeletionMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDBContainer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gcng.DBContainer = new(FakeDBContainer)
