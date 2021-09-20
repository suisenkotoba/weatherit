// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	alterplan "weatherit/usecases/alterplan"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// GetEventAlterPlan provides a mock function with given fields: ctx, eventId
func (_m *UseCase) GetEventAlterPlan(ctx context.Context, eventId int) alterplan.Domain {
	ret := _m.Called(ctx, eventId)

	var r0 alterplan.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) alterplan.Domain); ok {
		r0 = rf(ctx, eventId)
	} else {
		r0 = ret.Get(0).(alterplan.Domain)
	}

	return r0
}

// MakeEventAlterPlan provides a mock function with given fields: ctx, plan
func (_m *UseCase) MakeEventAlterPlan(ctx context.Context, plan *alterplan.Domain) (int, error) {
	ret := _m.Called(ctx, plan)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *alterplan.Domain) int); ok {
		r0 = rf(ctx, plan)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *alterplan.Domain) error); ok {
		r1 = rf(ctx, plan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEventAlterPlan provides a mock function with given fields: ctx, plan
func (_m *UseCase) UpdateEventAlterPlan(ctx context.Context, plan *alterplan.Domain) error {
	ret := _m.Called(ctx, plan)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *alterplan.Domain) error); ok {
		r0 = rf(ctx, plan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
