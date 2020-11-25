// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	context "context"

	job "github.com/smartcontractkit/chainlink/core/services/job"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	null "gopkg.in/guregu/null.v4"

	packr "github.com/gobuffalo/packr"

	store "github.com/smartcontractkit/chainlink/core/store"

	synchronization "github.com/smartcontractkit/chainlink/core/services/synchronization"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// AddJob provides a mock function with given fields: _a0
func (_m *Application) AddJob(_a0 models.JobSpec) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.JobSpec) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddJobV2 provides a mock function with given fields: ctx, _a1, name
func (_m *Application) AddJobV2(ctx context.Context, _a1 job.Spec, name null.String) (int32, error) {
	ret := _m.Called(ctx, _a1, name)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, job.Spec, null.String) int32); ok {
		r0 = rf(ctx, _a1, name)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, job.Spec, null.String) error); ok {
		r1 = rf(ctx, _a1, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddServiceAgreement provides a mock function with given fields: _a0
func (_m *Application) AddServiceAgreement(_a0 *models.ServiceAgreement) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ServiceAgreement) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArchiveJob provides a mock function with given fields: _a0
func (_m *Application) ArchiveJob(_a0 *models.ID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AwaitRun provides a mock function with given fields: ctx, runID
func (_m *Application) AwaitRun(ctx context.Context, runID int64) error {
	ret := _m.Called(ctx, runID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cancel provides a mock function with given fields: runID
func (_m *Application) Cancel(runID *models.ID) (*models.JobRun, error) {
	ret := _m.Called(runID)

	var r0 *models.JobRun
	if rf, ok := ret.Get(0).(func(*models.ID) *models.JobRun); ok {
		r0 = rf(runID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobRun)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ID) error); ok {
		r1 = rf(runID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: jobSpecID, initiator, creationHeight, runRequest
func (_m *Application) Create(jobSpecID *models.ID, initiator *models.Initiator, creationHeight *big.Int, runRequest *models.RunRequest) (*models.JobRun, error) {
	ret := _m.Called(jobSpecID, initiator, creationHeight, runRequest)

	var r0 *models.JobRun
	if rf, ok := ret.Get(0).(func(*models.ID, *models.Initiator, *big.Int, *models.RunRequest) *models.JobRun); ok {
		r0 = rf(jobSpecID, initiator, creationHeight, runRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobRun)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ID, *models.Initiator, *big.Int, *models.RunRequest) error); ok {
		r1 = rf(jobSpecID, initiator, creationHeight, runRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateErrored provides a mock function with given fields: jobSpecID, initiator, err
func (_m *Application) CreateErrored(jobSpecID *models.ID, initiator models.Initiator, err error) (*models.JobRun, error) {
	ret := _m.Called(jobSpecID, initiator, err)

	var r0 *models.JobRun
	if rf, ok := ret.Get(0).(func(*models.ID, models.Initiator, error) *models.JobRun); ok {
		r0 = rf(jobSpecID, initiator, err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JobRun)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ID, models.Initiator, error) error); ok {
		r1 = rf(jobSpecID, initiator, err)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJobV2 provides a mock function with given fields: ctx, jobID
func (_m *Application) DeleteJobV2(ctx context.Context, jobID int32) error {
	ret := _m.Called(ctx, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetStatsPusher provides a mock function with given fields:
func (_m *Application) GetStatsPusher() synchronization.StatsPusher {
	ret := _m.Called()

	var r0 synchronization.StatsPusher
	if rf, ok := ret.Get(0).(func() synchronization.StatsPusher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(synchronization.StatsPusher)
		}
	}

	return r0
}

// GetStore provides a mock function with given fields:
func (_m *Application) GetStore() *store.Store {
	ret := _m.Called()

	var r0 *store.Store
	if rf, ok := ret.Get(0).(func() *store.Store); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.Store)
		}
	}

	return r0
}

// NewBox provides a mock function with given fields:
func (_m *Application) NewBox() packr.Box {
	ret := _m.Called()

	var r0 packr.Box
	if rf, ok := ret.Get(0).(func() packr.Box); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(packr.Box)
	}

	return r0
}

// ResumeAllInProgress provides a mock function with given fields:
func (_m *Application) ResumeAllInProgress() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResumeAllPendingConnection provides a mock function with given fields:
func (_m *Application) ResumeAllPendingConnection() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResumeAllPendingNextBlock provides a mock function with given fields: currentBlockHeight
func (_m *Application) ResumeAllPendingNextBlock(currentBlockHeight *big.Int) error {
	ret := _m.Called(currentBlockHeight)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int) error); ok {
		r0 = rf(currentBlockHeight)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResumePendingBridge provides a mock function with given fields: runID, input
func (_m *Application) ResumePendingBridge(runID *models.ID, input models.BridgeRunResult) error {
	ret := _m.Called(runID, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ID, models.BridgeRunResult) error); ok {
		r0 = rf(runID, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunJobV2 provides a mock function with given fields: ctx, jobID, meta
func (_m *Application) RunJobV2(ctx context.Context, jobID int32, meta map[string]interface{}) (int64, error) {
	ret := _m.Called(ctx, jobID, meta)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int32, map[string]interface{}) int64); ok {
		r0 = rf(ctx, jobID, meta)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32, map[string]interface{}) error); ok {
		r1 = rf(ctx, jobID, meta)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Application) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Application) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WakeSessionReaper provides a mock function with given fields:
func (_m *Application) WakeSessionReaper() {
	_m.Called()
}
