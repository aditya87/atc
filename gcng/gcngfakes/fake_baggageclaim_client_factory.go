// This file was generated by counterfeiter
package gcngfakes

import (
	"sync"

	"github.com/concourse/atc/gcng"
	"github.com/concourse/baggageclaim/client"
)

type FakeBaggageclaimClientFactory struct {
	NewClientStub        func(apiURL string) client.Client
	newClientMutex       sync.RWMutex
	newClientArgsForCall []struct {
		apiURL string
	}
	newClientReturns struct {
		result1 client.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBaggageclaimClientFactory) NewClient(apiURL string) client.Client {
	fake.newClientMutex.Lock()
	fake.newClientArgsForCall = append(fake.newClientArgsForCall, struct {
		apiURL string
	}{apiURL})
	fake.recordInvocation("NewClient", []interface{}{apiURL})
	fake.newClientMutex.Unlock()
	if fake.NewClientStub != nil {
		return fake.NewClientStub(apiURL)
	} else {
		return fake.newClientReturns.result1
	}
}

func (fake *FakeBaggageclaimClientFactory) NewClientCallCount() int {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return len(fake.newClientArgsForCall)
}

func (fake *FakeBaggageclaimClientFactory) NewClientArgsForCall(i int) string {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return fake.newClientArgsForCall[i].apiURL
}

func (fake *FakeBaggageclaimClientFactory) NewClientReturns(result1 client.Client) {
	fake.NewClientStub = nil
	fake.newClientReturns = struct {
		result1 client.Client
	}{result1}
}

func (fake *FakeBaggageclaimClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBaggageclaimClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ gcng.BaggageclaimClientFactory = new(FakeBaggageclaimClientFactory)
