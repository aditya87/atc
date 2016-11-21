// This file was generated by counterfeiter
package dbfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/lock"
)

type FakeBuild struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	JobNameStub        func() string
	jobNameMutex       sync.RWMutex
	jobNameArgsForCall []struct{}
	jobNameReturns     struct {
		result1 string
	}
	PipelineNameStub        func() string
	pipelineNameMutex       sync.RWMutex
	pipelineNameArgsForCall []struct{}
	pipelineNameReturns     struct {
		result1 string
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct{}
	teamNameReturns     struct {
		result1 string
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct{}
	teamIDReturns     struct {
		result1 int
	}
	EngineStub        func() string
	engineMutex       sync.RWMutex
	engineArgsForCall []struct{}
	engineReturns     struct {
		result1 string
	}
	EngineMetadataStub        func() string
	engineMetadataMutex       sync.RWMutex
	engineMetadataArgsForCall []struct{}
	engineMetadataReturns     struct {
		result1 string
	}
	StatusStub        func() db.Status
	statusMutex       sync.RWMutex
	statusArgsForCall []struct{}
	statusReturns     struct {
		result1 db.Status
	}
	StartTimeStub        func() time.Time
	startTimeMutex       sync.RWMutex
	startTimeArgsForCall []struct{}
	startTimeReturns     struct {
		result1 time.Time
	}
	EndTimeStub        func() time.Time
	endTimeMutex       sync.RWMutex
	endTimeArgsForCall []struct{}
	endTimeReturns     struct {
		result1 time.Time
	}
	ReapTimeStub        func() time.Time
	reapTimeMutex       sync.RWMutex
	reapTimeArgsForCall []struct{}
	reapTimeReturns     struct {
		result1 time.Time
	}
	IsOneOffStub        func() bool
	isOneOffMutex       sync.RWMutex
	isOneOffArgsForCall []struct{}
	isOneOffReturns     struct {
		result1 bool
	}
	IsScheduledStub        func() bool
	isScheduledMutex       sync.RWMutex
	isScheduledArgsForCall []struct{}
	isScheduledReturns     struct {
		result1 bool
	}
	IsRunningStub        func() bool
	isRunningMutex       sync.RWMutex
	isRunningArgsForCall []struct{}
	isRunningReturns     struct {
		result1 bool
	}
	IsManuallyTriggeredStub        func() bool
	isManuallyTriggeredMutex       sync.RWMutex
	isManuallyTriggeredArgsForCall []struct{}
	isManuallyTriggeredReturns     struct {
		result1 bool
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	EventsStub        func(from uint) (db.EventSource, error)
	eventsMutex       sync.RWMutex
	eventsArgsForCall []struct {
		from uint
	}
	eventsReturns struct {
		result1 db.EventSource
		result2 error
	}
	SaveEventStub        func(event atc.Event) error
	saveEventMutex       sync.RWMutex
	saveEventArgsForCall []struct {
		event atc.Event
	}
	saveEventReturns struct {
		result1 error
	}
	GetVersionedResourcesStub        func() (db.SavedVersionedResources, error)
	getVersionedResourcesMutex       sync.RWMutex
	getVersionedResourcesArgsForCall []struct{}
	getVersionedResourcesReturns     struct {
		result1 db.SavedVersionedResources
		result2 error
	}
	GetResourcesStub        func() ([]db.BuildInput, []db.BuildOutput, error)
	getResourcesMutex       sync.RWMutex
	getResourcesArgsForCall []struct{}
	getResourcesReturns     struct {
		result1 []db.BuildInput
		result2 []db.BuildOutput
		result3 error
	}
	StartStub        func(string, string) (bool, error)
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 string
		arg2 string
	}
	startReturns struct {
		result1 bool
		result2 error
	}
	FinishStub        func(status db.Status) error
	finishMutex       sync.RWMutex
	finishArgsForCall []struct {
		status db.Status
	}
	finishReturns struct {
		result1 error
	}
	MarkAsFailedStub        func(cause error) error
	markAsFailedMutex       sync.RWMutex
	markAsFailedArgsForCall []struct {
		cause error
	}
	markAsFailedReturns struct {
		result1 error
	}
	AbortStub        func() error
	abortMutex       sync.RWMutex
	abortArgsForCall []struct{}
	abortReturns     struct {
		result1 error
	}
	AbortNotifierStub        func() (db.Notifier, error)
	abortNotifierMutex       sync.RWMutex
	abortNotifierArgsForCall []struct{}
	abortNotifierReturns     struct {
		result1 db.Notifier
		result2 error
	}
	AcquireTrackingLockStub        func(logger lager.Logger, interval time.Duration) (lock.Lock, bool, error)
	acquireTrackingLockMutex       sync.RWMutex
	acquireTrackingLockArgsForCall []struct {
		logger   lager.Logger
		interval time.Duration
	}
	acquireTrackingLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	GetPreparationStub        func() (db.BuildPreparation, bool, error)
	getPreparationMutex       sync.RWMutex
	getPreparationArgsForCall []struct{}
	getPreparationReturns     struct {
		result1 db.BuildPreparation
		result2 bool
		result3 error
	}
	SaveEngineMetadataStub        func(engineMetadata string) error
	saveEngineMetadataMutex       sync.RWMutex
	saveEngineMetadataArgsForCall []struct {
		engineMetadata string
	}
	saveEngineMetadataReturns struct {
		result1 error
	}
	SaveInputStub        func(input db.BuildInput) (db.SavedVersionedResource, error)
	saveInputMutex       sync.RWMutex
	saveInputArgsForCall []struct {
		input db.BuildInput
	}
	saveInputReturns struct {
		result1 db.SavedVersionedResource
		result2 error
	}
	SaveOutputStub        func(vr db.VersionedResource, explicit bool) (db.SavedVersionedResource, error)
	saveOutputMutex       sync.RWMutex
	saveOutputArgsForCall []struct {
		vr       db.VersionedResource
		explicit bool
	}
	saveOutputReturns struct {
		result1 db.SavedVersionedResource
		result2 error
	}
	SaveImageResourceVersionStub        func(planID atc.PlanID, identifier db.ResourceCacheIdentifier) error
	saveImageResourceVersionMutex       sync.RWMutex
	saveImageResourceVersionArgsForCall []struct {
		planID     atc.PlanID
		identifier db.ResourceCacheIdentifier
	}
	saveImageResourceVersionReturns struct {
		result1 error
	}
	GetImageResourceCacheIdentifiersStub        func() ([]db.ResourceCacheIdentifier, error)
	getImageResourceCacheIdentifiersMutex       sync.RWMutex
	getImageResourceCacheIdentifiersArgsForCall []struct{}
	getImageResourceCacheIdentifiersReturns     struct {
		result1 []db.ResourceCacheIdentifier
		result2 error
	}
	GetConfigStub        func() (atc.Config, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct{}
	getConfigReturns     struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}
	GetPipelineStub        func() (db.SavedPipeline, error)
	getPipelineMutex       sync.RWMutex
	getPipelineArgsForCall []struct{}
	getPipelineReturns     struct {
		result1 db.SavedPipeline
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuild) ID() int {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeBuild) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeBuild) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeBuild) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeBuild) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBuild) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) JobName() string {
	fake.jobNameMutex.Lock()
	fake.jobNameArgsForCall = append(fake.jobNameArgsForCall, struct{}{})
	fake.recordInvocation("JobName", []interface{}{})
	fake.jobNameMutex.Unlock()
	if fake.JobNameStub != nil {
		return fake.JobNameStub()
	} else {
		return fake.jobNameReturns.result1
	}
}

func (fake *FakeBuild) JobNameCallCount() int {
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	return len(fake.jobNameArgsForCall)
}

func (fake *FakeBuild) JobNameReturns(result1 string) {
	fake.JobNameStub = nil
	fake.jobNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) PipelineName() string {
	fake.pipelineNameMutex.Lock()
	fake.pipelineNameArgsForCall = append(fake.pipelineNameArgsForCall, struct{}{})
	fake.recordInvocation("PipelineName", []interface{}{})
	fake.pipelineNameMutex.Unlock()
	if fake.PipelineNameStub != nil {
		return fake.PipelineNameStub()
	} else {
		return fake.pipelineNameReturns.result1
	}
}

func (fake *FakeBuild) PipelineNameCallCount() int {
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	return len(fake.pipelineNameArgsForCall)
}

func (fake *FakeBuild) PipelineNameReturns(result1 string) {
	fake.PipelineNameStub = nil
	fake.pipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) TeamName() string {
	fake.teamNameMutex.Lock()
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct{}{})
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if fake.TeamNameStub != nil {
		return fake.TeamNameStub()
	} else {
		return fake.teamNameReturns.result1
	}
}

func (fake *FakeBuild) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeBuild) TeamNameReturns(result1 string) {
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) TeamID() int {
	fake.teamIDMutex.Lock()
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct{}{})
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if fake.TeamIDStub != nil {
		return fake.TeamIDStub()
	} else {
		return fake.teamIDReturns.result1
	}
}

func (fake *FakeBuild) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeBuild) TeamIDReturns(result1 int) {
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeBuild) Engine() string {
	fake.engineMutex.Lock()
	fake.engineArgsForCall = append(fake.engineArgsForCall, struct{}{})
	fake.recordInvocation("Engine", []interface{}{})
	fake.engineMutex.Unlock()
	if fake.EngineStub != nil {
		return fake.EngineStub()
	} else {
		return fake.engineReturns.result1
	}
}

func (fake *FakeBuild) EngineCallCount() int {
	fake.engineMutex.RLock()
	defer fake.engineMutex.RUnlock()
	return len(fake.engineArgsForCall)
}

func (fake *FakeBuild) EngineReturns(result1 string) {
	fake.EngineStub = nil
	fake.engineReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) EngineMetadata() string {
	fake.engineMetadataMutex.Lock()
	fake.engineMetadataArgsForCall = append(fake.engineMetadataArgsForCall, struct{}{})
	fake.recordInvocation("EngineMetadata", []interface{}{})
	fake.engineMetadataMutex.Unlock()
	if fake.EngineMetadataStub != nil {
		return fake.EngineMetadataStub()
	} else {
		return fake.engineMetadataReturns.result1
	}
}

func (fake *FakeBuild) EngineMetadataCallCount() int {
	fake.engineMetadataMutex.RLock()
	defer fake.engineMetadataMutex.RUnlock()
	return len(fake.engineMetadataArgsForCall)
}

func (fake *FakeBuild) EngineMetadataReturns(result1 string) {
	fake.EngineMetadataStub = nil
	fake.engineMetadataReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) Status() db.Status {
	fake.statusMutex.Lock()
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct{}{})
	fake.recordInvocation("Status", []interface{}{})
	fake.statusMutex.Unlock()
	if fake.StatusStub != nil {
		return fake.StatusStub()
	} else {
		return fake.statusReturns.result1
	}
}

func (fake *FakeBuild) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeBuild) StatusReturns(result1 db.Status) {
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 db.Status
	}{result1}
}

func (fake *FakeBuild) StartTime() time.Time {
	fake.startTimeMutex.Lock()
	fake.startTimeArgsForCall = append(fake.startTimeArgsForCall, struct{}{})
	fake.recordInvocation("StartTime", []interface{}{})
	fake.startTimeMutex.Unlock()
	if fake.StartTimeStub != nil {
		return fake.StartTimeStub()
	} else {
		return fake.startTimeReturns.result1
	}
}

func (fake *FakeBuild) StartTimeCallCount() int {
	fake.startTimeMutex.RLock()
	defer fake.startTimeMutex.RUnlock()
	return len(fake.startTimeArgsForCall)
}

func (fake *FakeBuild) StartTimeReturns(result1 time.Time) {
	fake.StartTimeStub = nil
	fake.startTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeBuild) EndTime() time.Time {
	fake.endTimeMutex.Lock()
	fake.endTimeArgsForCall = append(fake.endTimeArgsForCall, struct{}{})
	fake.recordInvocation("EndTime", []interface{}{})
	fake.endTimeMutex.Unlock()
	if fake.EndTimeStub != nil {
		return fake.EndTimeStub()
	} else {
		return fake.endTimeReturns.result1
	}
}

func (fake *FakeBuild) EndTimeCallCount() int {
	fake.endTimeMutex.RLock()
	defer fake.endTimeMutex.RUnlock()
	return len(fake.endTimeArgsForCall)
}

func (fake *FakeBuild) EndTimeReturns(result1 time.Time) {
	fake.EndTimeStub = nil
	fake.endTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeBuild) ReapTime() time.Time {
	fake.reapTimeMutex.Lock()
	fake.reapTimeArgsForCall = append(fake.reapTimeArgsForCall, struct{}{})
	fake.recordInvocation("ReapTime", []interface{}{})
	fake.reapTimeMutex.Unlock()
	if fake.ReapTimeStub != nil {
		return fake.ReapTimeStub()
	} else {
		return fake.reapTimeReturns.result1
	}
}

func (fake *FakeBuild) ReapTimeCallCount() int {
	fake.reapTimeMutex.RLock()
	defer fake.reapTimeMutex.RUnlock()
	return len(fake.reapTimeArgsForCall)
}

func (fake *FakeBuild) ReapTimeReturns(result1 time.Time) {
	fake.ReapTimeStub = nil
	fake.reapTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeBuild) IsOneOff() bool {
	fake.isOneOffMutex.Lock()
	fake.isOneOffArgsForCall = append(fake.isOneOffArgsForCall, struct{}{})
	fake.recordInvocation("IsOneOff", []interface{}{})
	fake.isOneOffMutex.Unlock()
	if fake.IsOneOffStub != nil {
		return fake.IsOneOffStub()
	} else {
		return fake.isOneOffReturns.result1
	}
}

func (fake *FakeBuild) IsOneOffCallCount() int {
	fake.isOneOffMutex.RLock()
	defer fake.isOneOffMutex.RUnlock()
	return len(fake.isOneOffArgsForCall)
}

func (fake *FakeBuild) IsOneOffReturns(result1 bool) {
	fake.IsOneOffStub = nil
	fake.isOneOffReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsScheduled() bool {
	fake.isScheduledMutex.Lock()
	fake.isScheduledArgsForCall = append(fake.isScheduledArgsForCall, struct{}{})
	fake.recordInvocation("IsScheduled", []interface{}{})
	fake.isScheduledMutex.Unlock()
	if fake.IsScheduledStub != nil {
		return fake.IsScheduledStub()
	} else {
		return fake.isScheduledReturns.result1
	}
}

func (fake *FakeBuild) IsScheduledCallCount() int {
	fake.isScheduledMutex.RLock()
	defer fake.isScheduledMutex.RUnlock()
	return len(fake.isScheduledArgsForCall)
}

func (fake *FakeBuild) IsScheduledReturns(result1 bool) {
	fake.IsScheduledStub = nil
	fake.isScheduledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsRunning() bool {
	fake.isRunningMutex.Lock()
	fake.isRunningArgsForCall = append(fake.isRunningArgsForCall, struct{}{})
	fake.recordInvocation("IsRunning", []interface{}{})
	fake.isRunningMutex.Unlock()
	if fake.IsRunningStub != nil {
		return fake.IsRunningStub()
	} else {
		return fake.isRunningReturns.result1
	}
}

func (fake *FakeBuild) IsRunningCallCount() int {
	fake.isRunningMutex.RLock()
	defer fake.isRunningMutex.RUnlock()
	return len(fake.isRunningArgsForCall)
}

func (fake *FakeBuild) IsRunningReturns(result1 bool) {
	fake.IsRunningStub = nil
	fake.isRunningReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) IsManuallyTriggered() bool {
	fake.isManuallyTriggeredMutex.Lock()
	fake.isManuallyTriggeredArgsForCall = append(fake.isManuallyTriggeredArgsForCall, struct{}{})
	fake.recordInvocation("IsManuallyTriggered", []interface{}{})
	fake.isManuallyTriggeredMutex.Unlock()
	if fake.IsManuallyTriggeredStub != nil {
		return fake.IsManuallyTriggeredStub()
	} else {
		return fake.isManuallyTriggeredReturns.result1
	}
}

func (fake *FakeBuild) IsManuallyTriggeredCallCount() int {
	fake.isManuallyTriggeredMutex.RLock()
	defer fake.isManuallyTriggeredMutex.RUnlock()
	return len(fake.isManuallyTriggeredArgsForCall)
}

func (fake *FakeBuild) IsManuallyTriggeredReturns(result1 bool) {
	fake.IsManuallyTriggeredStub = nil
	fake.isManuallyTriggeredReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBuild) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	} else {
		return fake.reloadReturns.result1, fake.reloadReturns.result2
	}
}

func (fake *FakeBuild) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeBuild) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) Events(from uint) (db.EventSource, error) {
	fake.eventsMutex.Lock()
	fake.eventsArgsForCall = append(fake.eventsArgsForCall, struct {
		from uint
	}{from})
	fake.recordInvocation("Events", []interface{}{from})
	fake.eventsMutex.Unlock()
	if fake.EventsStub != nil {
		return fake.EventsStub(from)
	} else {
		return fake.eventsReturns.result1, fake.eventsReturns.result2
	}
}

func (fake *FakeBuild) EventsCallCount() int {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	return len(fake.eventsArgsForCall)
}

func (fake *FakeBuild) EventsArgsForCall(i int) uint {
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	return fake.eventsArgsForCall[i].from
}

func (fake *FakeBuild) EventsReturns(result1 db.EventSource, result2 error) {
	fake.EventsStub = nil
	fake.eventsReturns = struct {
		result1 db.EventSource
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) SaveEvent(event atc.Event) error {
	fake.saveEventMutex.Lock()
	fake.saveEventArgsForCall = append(fake.saveEventArgsForCall, struct {
		event atc.Event
	}{event})
	fake.recordInvocation("SaveEvent", []interface{}{event})
	fake.saveEventMutex.Unlock()
	if fake.SaveEventStub != nil {
		return fake.SaveEventStub(event)
	} else {
		return fake.saveEventReturns.result1
	}
}

func (fake *FakeBuild) SaveEventCallCount() int {
	fake.saveEventMutex.RLock()
	defer fake.saveEventMutex.RUnlock()
	return len(fake.saveEventArgsForCall)
}

func (fake *FakeBuild) SaveEventArgsForCall(i int) atc.Event {
	fake.saveEventMutex.RLock()
	defer fake.saveEventMutex.RUnlock()
	return fake.saveEventArgsForCall[i].event
}

func (fake *FakeBuild) SaveEventReturns(result1 error) {
	fake.SaveEventStub = nil
	fake.saveEventReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) GetVersionedResources() (db.SavedVersionedResources, error) {
	fake.getVersionedResourcesMutex.Lock()
	fake.getVersionedResourcesArgsForCall = append(fake.getVersionedResourcesArgsForCall, struct{}{})
	fake.recordInvocation("GetVersionedResources", []interface{}{})
	fake.getVersionedResourcesMutex.Unlock()
	if fake.GetVersionedResourcesStub != nil {
		return fake.GetVersionedResourcesStub()
	} else {
		return fake.getVersionedResourcesReturns.result1, fake.getVersionedResourcesReturns.result2
	}
}

func (fake *FakeBuild) GetVersionedResourcesCallCount() int {
	fake.getVersionedResourcesMutex.RLock()
	defer fake.getVersionedResourcesMutex.RUnlock()
	return len(fake.getVersionedResourcesArgsForCall)
}

func (fake *FakeBuild) GetVersionedResourcesReturns(result1 db.SavedVersionedResources, result2 error) {
	fake.GetVersionedResourcesStub = nil
	fake.getVersionedResourcesReturns = struct {
		result1 db.SavedVersionedResources
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) GetResources() ([]db.BuildInput, []db.BuildOutput, error) {
	fake.getResourcesMutex.Lock()
	fake.getResourcesArgsForCall = append(fake.getResourcesArgsForCall, struct{}{})
	fake.recordInvocation("GetResources", []interface{}{})
	fake.getResourcesMutex.Unlock()
	if fake.GetResourcesStub != nil {
		return fake.GetResourcesStub()
	} else {
		return fake.getResourcesReturns.result1, fake.getResourcesReturns.result2, fake.getResourcesReturns.result3
	}
}

func (fake *FakeBuild) GetResourcesCallCount() int {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return len(fake.getResourcesArgsForCall)
}

func (fake *FakeBuild) GetResourcesReturns(result1 []db.BuildInput, result2 []db.BuildOutput, result3 error) {
	fake.GetResourcesStub = nil
	fake.getResourcesReturns = struct {
		result1 []db.BuildInput
		result2 []db.BuildOutput
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuild) Start(arg1 string, arg2 string) (bool, error) {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Start", []interface{}{arg1, arg2})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1, arg2)
	} else {
		return fake.startReturns.result1, fake.startReturns.result2
	}
}

func (fake *FakeBuild) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeBuild) StartArgsForCall(i int) (string, string) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].arg1, fake.startArgsForCall[i].arg2
}

func (fake *FakeBuild) StartReturns(result1 bool, result2 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) Finish(status db.Status) error {
	fake.finishMutex.Lock()
	fake.finishArgsForCall = append(fake.finishArgsForCall, struct {
		status db.Status
	}{status})
	fake.recordInvocation("Finish", []interface{}{status})
	fake.finishMutex.Unlock()
	if fake.FinishStub != nil {
		return fake.FinishStub(status)
	} else {
		return fake.finishReturns.result1
	}
}

func (fake *FakeBuild) FinishCallCount() int {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return len(fake.finishArgsForCall)
}

func (fake *FakeBuild) FinishArgsForCall(i int) db.Status {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return fake.finishArgsForCall[i].status
}

func (fake *FakeBuild) FinishReturns(result1 error) {
	fake.FinishStub = nil
	fake.finishReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) MarkAsFailed(cause error) error {
	fake.markAsFailedMutex.Lock()
	fake.markAsFailedArgsForCall = append(fake.markAsFailedArgsForCall, struct {
		cause error
	}{cause})
	fake.recordInvocation("MarkAsFailed", []interface{}{cause})
	fake.markAsFailedMutex.Unlock()
	if fake.MarkAsFailedStub != nil {
		return fake.MarkAsFailedStub(cause)
	} else {
		return fake.markAsFailedReturns.result1
	}
}

func (fake *FakeBuild) MarkAsFailedCallCount() int {
	fake.markAsFailedMutex.RLock()
	defer fake.markAsFailedMutex.RUnlock()
	return len(fake.markAsFailedArgsForCall)
}

func (fake *FakeBuild) MarkAsFailedArgsForCall(i int) error {
	fake.markAsFailedMutex.RLock()
	defer fake.markAsFailedMutex.RUnlock()
	return fake.markAsFailedArgsForCall[i].cause
}

func (fake *FakeBuild) MarkAsFailedReturns(result1 error) {
	fake.MarkAsFailedStub = nil
	fake.markAsFailedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) Abort() error {
	fake.abortMutex.Lock()
	fake.abortArgsForCall = append(fake.abortArgsForCall, struct{}{})
	fake.recordInvocation("Abort", []interface{}{})
	fake.abortMutex.Unlock()
	if fake.AbortStub != nil {
		return fake.AbortStub()
	} else {
		return fake.abortReturns.result1
	}
}

func (fake *FakeBuild) AbortCallCount() int {
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	return len(fake.abortArgsForCall)
}

func (fake *FakeBuild) AbortReturns(result1 error) {
	fake.AbortStub = nil
	fake.abortReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) AbortNotifier() (db.Notifier, error) {
	fake.abortNotifierMutex.Lock()
	fake.abortNotifierArgsForCall = append(fake.abortNotifierArgsForCall, struct{}{})
	fake.recordInvocation("AbortNotifier", []interface{}{})
	fake.abortNotifierMutex.Unlock()
	if fake.AbortNotifierStub != nil {
		return fake.AbortNotifierStub()
	} else {
		return fake.abortNotifierReturns.result1, fake.abortNotifierReturns.result2
	}
}

func (fake *FakeBuild) AbortNotifierCallCount() int {
	fake.abortNotifierMutex.RLock()
	defer fake.abortNotifierMutex.RUnlock()
	return len(fake.abortNotifierArgsForCall)
}

func (fake *FakeBuild) AbortNotifierReturns(result1 db.Notifier, result2 error) {
	fake.AbortNotifierStub = nil
	fake.abortNotifierReturns = struct {
		result1 db.Notifier
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) AcquireTrackingLock(logger lager.Logger, interval time.Duration) (lock.Lock, bool, error) {
	fake.acquireTrackingLockMutex.Lock()
	fake.acquireTrackingLockArgsForCall = append(fake.acquireTrackingLockArgsForCall, struct {
		logger   lager.Logger
		interval time.Duration
	}{logger, interval})
	fake.recordInvocation("AcquireTrackingLock", []interface{}{logger, interval})
	fake.acquireTrackingLockMutex.Unlock()
	if fake.AcquireTrackingLockStub != nil {
		return fake.AcquireTrackingLockStub(logger, interval)
	} else {
		return fake.acquireTrackingLockReturns.result1, fake.acquireTrackingLockReturns.result2, fake.acquireTrackingLockReturns.result3
	}
}

func (fake *FakeBuild) AcquireTrackingLockCallCount() int {
	fake.acquireTrackingLockMutex.RLock()
	defer fake.acquireTrackingLockMutex.RUnlock()
	return len(fake.acquireTrackingLockArgsForCall)
}

func (fake *FakeBuild) AcquireTrackingLockArgsForCall(i int) (lager.Logger, time.Duration) {
	fake.acquireTrackingLockMutex.RLock()
	defer fake.acquireTrackingLockMutex.RUnlock()
	return fake.acquireTrackingLockArgsForCall[i].logger, fake.acquireTrackingLockArgsForCall[i].interval
}

func (fake *FakeBuild) AcquireTrackingLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireTrackingLockStub = nil
	fake.acquireTrackingLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuild) GetPreparation() (db.BuildPreparation, bool, error) {
	fake.getPreparationMutex.Lock()
	fake.getPreparationArgsForCall = append(fake.getPreparationArgsForCall, struct{}{})
	fake.recordInvocation("GetPreparation", []interface{}{})
	fake.getPreparationMutex.Unlock()
	if fake.GetPreparationStub != nil {
		return fake.GetPreparationStub()
	} else {
		return fake.getPreparationReturns.result1, fake.getPreparationReturns.result2, fake.getPreparationReturns.result3
	}
}

func (fake *FakeBuild) GetPreparationCallCount() int {
	fake.getPreparationMutex.RLock()
	defer fake.getPreparationMutex.RUnlock()
	return len(fake.getPreparationArgsForCall)
}

func (fake *FakeBuild) GetPreparationReturns(result1 db.BuildPreparation, result2 bool, result3 error) {
	fake.GetPreparationStub = nil
	fake.getPreparationReturns = struct {
		result1 db.BuildPreparation
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuild) SaveEngineMetadata(engineMetadata string) error {
	fake.saveEngineMetadataMutex.Lock()
	fake.saveEngineMetadataArgsForCall = append(fake.saveEngineMetadataArgsForCall, struct {
		engineMetadata string
	}{engineMetadata})
	fake.recordInvocation("SaveEngineMetadata", []interface{}{engineMetadata})
	fake.saveEngineMetadataMutex.Unlock()
	if fake.SaveEngineMetadataStub != nil {
		return fake.SaveEngineMetadataStub(engineMetadata)
	} else {
		return fake.saveEngineMetadataReturns.result1
	}
}

func (fake *FakeBuild) SaveEngineMetadataCallCount() int {
	fake.saveEngineMetadataMutex.RLock()
	defer fake.saveEngineMetadataMutex.RUnlock()
	return len(fake.saveEngineMetadataArgsForCall)
}

func (fake *FakeBuild) SaveEngineMetadataArgsForCall(i int) string {
	fake.saveEngineMetadataMutex.RLock()
	defer fake.saveEngineMetadataMutex.RUnlock()
	return fake.saveEngineMetadataArgsForCall[i].engineMetadata
}

func (fake *FakeBuild) SaveEngineMetadataReturns(result1 error) {
	fake.SaveEngineMetadataStub = nil
	fake.saveEngineMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) SaveInput(input db.BuildInput) (db.SavedVersionedResource, error) {
	fake.saveInputMutex.Lock()
	fake.saveInputArgsForCall = append(fake.saveInputArgsForCall, struct {
		input db.BuildInput
	}{input})
	fake.recordInvocation("SaveInput", []interface{}{input})
	fake.saveInputMutex.Unlock()
	if fake.SaveInputStub != nil {
		return fake.SaveInputStub(input)
	} else {
		return fake.saveInputReturns.result1, fake.saveInputReturns.result2
	}
}

func (fake *FakeBuild) SaveInputCallCount() int {
	fake.saveInputMutex.RLock()
	defer fake.saveInputMutex.RUnlock()
	return len(fake.saveInputArgsForCall)
}

func (fake *FakeBuild) SaveInputArgsForCall(i int) db.BuildInput {
	fake.saveInputMutex.RLock()
	defer fake.saveInputMutex.RUnlock()
	return fake.saveInputArgsForCall[i].input
}

func (fake *FakeBuild) SaveInputReturns(result1 db.SavedVersionedResource, result2 error) {
	fake.SaveInputStub = nil
	fake.saveInputReturns = struct {
		result1 db.SavedVersionedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) SaveOutput(vr db.VersionedResource, explicit bool) (db.SavedVersionedResource, error) {
	fake.saveOutputMutex.Lock()
	fake.saveOutputArgsForCall = append(fake.saveOutputArgsForCall, struct {
		vr       db.VersionedResource
		explicit bool
	}{vr, explicit})
	fake.recordInvocation("SaveOutput", []interface{}{vr, explicit})
	fake.saveOutputMutex.Unlock()
	if fake.SaveOutputStub != nil {
		return fake.SaveOutputStub(vr, explicit)
	} else {
		return fake.saveOutputReturns.result1, fake.saveOutputReturns.result2
	}
}

func (fake *FakeBuild) SaveOutputCallCount() int {
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	return len(fake.saveOutputArgsForCall)
}

func (fake *FakeBuild) SaveOutputArgsForCall(i int) (db.VersionedResource, bool) {
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	return fake.saveOutputArgsForCall[i].vr, fake.saveOutputArgsForCall[i].explicit
}

func (fake *FakeBuild) SaveOutputReturns(result1 db.SavedVersionedResource, result2 error) {
	fake.SaveOutputStub = nil
	fake.saveOutputReturns = struct {
		result1 db.SavedVersionedResource
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) SaveImageResourceVersion(planID atc.PlanID, identifier db.ResourceCacheIdentifier) error {
	fake.saveImageResourceVersionMutex.Lock()
	fake.saveImageResourceVersionArgsForCall = append(fake.saveImageResourceVersionArgsForCall, struct {
		planID     atc.PlanID
		identifier db.ResourceCacheIdentifier
	}{planID, identifier})
	fake.recordInvocation("SaveImageResourceVersion", []interface{}{planID, identifier})
	fake.saveImageResourceVersionMutex.Unlock()
	if fake.SaveImageResourceVersionStub != nil {
		return fake.SaveImageResourceVersionStub(planID, identifier)
	} else {
		return fake.saveImageResourceVersionReturns.result1
	}
}

func (fake *FakeBuild) SaveImageResourceVersionCallCount() int {
	fake.saveImageResourceVersionMutex.RLock()
	defer fake.saveImageResourceVersionMutex.RUnlock()
	return len(fake.saveImageResourceVersionArgsForCall)
}

func (fake *FakeBuild) SaveImageResourceVersionArgsForCall(i int) (atc.PlanID, db.ResourceCacheIdentifier) {
	fake.saveImageResourceVersionMutex.RLock()
	defer fake.saveImageResourceVersionMutex.RUnlock()
	return fake.saveImageResourceVersionArgsForCall[i].planID, fake.saveImageResourceVersionArgsForCall[i].identifier
}

func (fake *FakeBuild) SaveImageResourceVersionReturns(result1 error) {
	fake.SaveImageResourceVersionStub = nil
	fake.saveImageResourceVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) GetImageResourceCacheIdentifiers() ([]db.ResourceCacheIdentifier, error) {
	fake.getImageResourceCacheIdentifiersMutex.Lock()
	fake.getImageResourceCacheIdentifiersArgsForCall = append(fake.getImageResourceCacheIdentifiersArgsForCall, struct{}{})
	fake.recordInvocation("GetImageResourceCacheIdentifiers", []interface{}{})
	fake.getImageResourceCacheIdentifiersMutex.Unlock()
	if fake.GetImageResourceCacheIdentifiersStub != nil {
		return fake.GetImageResourceCacheIdentifiersStub()
	} else {
		return fake.getImageResourceCacheIdentifiersReturns.result1, fake.getImageResourceCacheIdentifiersReturns.result2
	}
}

func (fake *FakeBuild) GetImageResourceCacheIdentifiersCallCount() int {
	fake.getImageResourceCacheIdentifiersMutex.RLock()
	defer fake.getImageResourceCacheIdentifiersMutex.RUnlock()
	return len(fake.getImageResourceCacheIdentifiersArgsForCall)
}

func (fake *FakeBuild) GetImageResourceCacheIdentifiersReturns(result1 []db.ResourceCacheIdentifier, result2 error) {
	fake.GetImageResourceCacheIdentifiersStub = nil
	fake.getImageResourceCacheIdentifiersReturns = struct {
		result1 []db.ResourceCacheIdentifier
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) GetConfig() (atc.Config, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct{}{})
	fake.recordInvocation("GetConfig", []interface{}{})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub()
	} else {
		return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3
	}
}

func (fake *FakeBuild) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeBuild) GetConfigReturns(result1 atc.Config, result2 db.ConfigVersion, result3 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuild) GetPipeline() (db.SavedPipeline, error) {
	fake.getPipelineMutex.Lock()
	fake.getPipelineArgsForCall = append(fake.getPipelineArgsForCall, struct{}{})
	fake.recordInvocation("GetPipeline", []interface{}{})
	fake.getPipelineMutex.Unlock()
	if fake.GetPipelineStub != nil {
		return fake.GetPipelineStub()
	} else {
		return fake.getPipelineReturns.result1, fake.getPipelineReturns.result2
	}
}

func (fake *FakeBuild) GetPipelineCallCount() int {
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	return len(fake.getPipelineArgsForCall)
}

func (fake *FakeBuild) GetPipelineReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineStub = nil
	fake.getPipelineReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeBuild) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	fake.pipelineNameMutex.RLock()
	defer fake.pipelineNameMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.engineMutex.RLock()
	defer fake.engineMutex.RUnlock()
	fake.engineMetadataMutex.RLock()
	defer fake.engineMetadataMutex.RUnlock()
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	fake.startTimeMutex.RLock()
	defer fake.startTimeMutex.RUnlock()
	fake.endTimeMutex.RLock()
	defer fake.endTimeMutex.RUnlock()
	fake.reapTimeMutex.RLock()
	defer fake.reapTimeMutex.RUnlock()
	fake.isOneOffMutex.RLock()
	defer fake.isOneOffMutex.RUnlock()
	fake.isScheduledMutex.RLock()
	defer fake.isScheduledMutex.RUnlock()
	fake.isRunningMutex.RLock()
	defer fake.isRunningMutex.RUnlock()
	fake.isManuallyTriggeredMutex.RLock()
	defer fake.isManuallyTriggeredMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.eventsMutex.RLock()
	defer fake.eventsMutex.RUnlock()
	fake.saveEventMutex.RLock()
	defer fake.saveEventMutex.RUnlock()
	fake.getVersionedResourcesMutex.RLock()
	defer fake.getVersionedResourcesMutex.RUnlock()
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	fake.markAsFailedMutex.RLock()
	defer fake.markAsFailedMutex.RUnlock()
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	fake.abortNotifierMutex.RLock()
	defer fake.abortNotifierMutex.RUnlock()
	fake.acquireTrackingLockMutex.RLock()
	defer fake.acquireTrackingLockMutex.RUnlock()
	fake.getPreparationMutex.RLock()
	defer fake.getPreparationMutex.RUnlock()
	fake.saveEngineMetadataMutex.RLock()
	defer fake.saveEngineMetadataMutex.RUnlock()
	fake.saveInputMutex.RLock()
	defer fake.saveInputMutex.RUnlock()
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	fake.saveImageResourceVersionMutex.RLock()
	defer fake.saveImageResourceVersionMutex.RUnlock()
	fake.getImageResourceCacheIdentifiersMutex.RLock()
	defer fake.getImageResourceCacheIdentifiersMutex.RUnlock()
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	fake.getPipelineMutex.RLock()
	defer fake.getPipelineMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBuild) recordInvocation(key string, args []interface{}) {
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

var _ db.Build = new(FakeBuild)
