// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// ComplianceStore is an autogenerated mock type for the ComplianceStore type
type ComplianceStore struct {
	mock.Mock
}

// ComplianceExport provides a mock function with given fields: compliance, cursor, limit
func (_m *ComplianceStore) ComplianceExport(compliance *model.Compliance, cursor model.ComplianceExportCursor, limit int) ([]*model.CompliancePost, model.ComplianceExportCursor, error) {
	ret := _m.Called(compliance, cursor, limit)

	var r0 []*model.CompliancePost
	var r1 model.ComplianceExportCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(*model.Compliance, model.ComplianceExportCursor, int) ([]*model.CompliancePost, model.ComplianceExportCursor, error)); ok {
		return rf(compliance, cursor, limit)
	}
	if rf, ok := ret.Get(0).(func(*model.Compliance, model.ComplianceExportCursor, int) []*model.CompliancePost); ok {
		r0 = rf(compliance, cursor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.CompliancePost)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Compliance, model.ComplianceExportCursor, int) model.ComplianceExportCursor); ok {
		r1 = rf(compliance, cursor, limit)
	} else {
		r1 = ret.Get(1).(model.ComplianceExportCursor)
	}

	if rf, ok := ret.Get(2).(func(*model.Compliance, model.ComplianceExportCursor, int) error); ok {
		r2 = rf(compliance, cursor, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Get provides a mock function with given fields: id
func (_m *ComplianceStore) Get(id string) (*model.Compliance, error) {
	ret := _m.Called(id)

	var r0 *model.Compliance
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Compliance, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Compliance); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Compliance)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: offset, limit
func (_m *ComplianceStore) GetAll(offset int, limit int) (model.Compliances, error) {
	ret := _m.Called(offset, limit)

	var r0 model.Compliances
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (model.Compliances, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) model.Compliances); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Compliances)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessageExport provides a mock function with given fields: ctx, cursor, limit
func (_m *ComplianceStore) MessageExport(ctx context.Context, cursor model.MessageExportCursor, limit int) ([]*model.MessageExport, model.MessageExportCursor, error) {
	ret := _m.Called(ctx, cursor, limit)

	var r0 []*model.MessageExport
	var r1 model.MessageExportCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, model.MessageExportCursor, int) ([]*model.MessageExport, model.MessageExportCursor, error)); ok {
		return rf(ctx, cursor, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.MessageExportCursor, int) []*model.MessageExport); ok {
		r0 = rf(ctx, cursor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.MessageExport)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.MessageExportCursor, int) model.MessageExportCursor); ok {
		r1 = rf(ctx, cursor, limit)
	} else {
		r1 = ret.Get(1).(model.MessageExportCursor)
	}

	if rf, ok := ret.Get(2).(func(context.Context, model.MessageExportCursor, int) error); ok {
		r2 = rf(ctx, cursor, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Save provides a mock function with given fields: compliance
func (_m *ComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, error) {
	ret := _m.Called(compliance)

	var r0 *model.Compliance
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Compliance) (*model.Compliance, error)); ok {
		return rf(compliance)
	}
	if rf, ok := ret.Get(0).(func(*model.Compliance) *model.Compliance); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Compliance)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Compliance) error); ok {
		r1 = rf(compliance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: compliance
func (_m *ComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, error) {
	ret := _m.Called(compliance)

	var r0 *model.Compliance
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Compliance) (*model.Compliance, error)); ok {
		return rf(compliance)
	}
	if rf, ok := ret.Get(0).(func(*model.Compliance) *model.Compliance); ok {
		r0 = rf(compliance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Compliance)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Compliance) error); ok {
		r1 = rf(compliance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewComplianceStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewComplianceStore creates a new instance of ComplianceStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewComplianceStore(t mockConstructorTestingTNewComplianceStore) *ComplianceStore {
	mock := &ComplianceStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
