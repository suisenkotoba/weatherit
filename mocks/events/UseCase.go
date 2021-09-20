// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	events "weatherit/usecases/events"

	mock "github.com/stretchr/testify/mock"

	time "time"

	weatherforecast "weatherit/usecases/weatherforecast"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CancelEvent provides a mock function with given fields: ctx, eventId, userId
func (_m *UseCase) CancelEvent(ctx context.Context, eventId int, userId int) error {
	ret := _m.Called(ctx, eventId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, eventId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForecastEvent provides a mock function with given fields: event, mode, dt1, dt2
func (_m *UseCase) ForecastEvent(event events.Domain, mode string, dt1 int64, dt2 int64) weatherforecast.Domain {
	ret := _m.Called(event, mode, dt1, dt2)

	var r0 weatherforecast.Domain
	if rf, ok := ret.Get(0).(func(events.Domain, string, int64, int64) weatherforecast.Domain); ok {
		r0 = rf(event, mode, dt1, dt2)
	} else {
		r0 = ret.Get(0).(weatherforecast.Domain)
	}

	return r0
}

// GetAllEventByDateRange provides a mock function with given fields: ctx, from, to
func (_m *UseCase) GetAllEventByDateRange(ctx context.Context, from time.Time, to time.Time) ([]events.Domain, error) {
	ret := _m.Called(ctx, from, to)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []events.Domain); ok {
		r0 = rf(ctx, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUserEvents provides a mock function with given fields: ctx, userId, from, to, month
func (_m *UseCase) GetAllUserEvents(ctx context.Context, userId int, from string, to string, month string) ([]events.Domain, error) {
	ret := _m.Called(ctx, userId, from, to, month)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, string) []events.Domain); ok {
		r0 = rf(ctx, userId, from, to, month)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string, string, string) error); ok {
		r1 = rf(ctx, userId, from, to, month)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUserEventsByDateRange provides a mock function with given fields: ctx, userId, from, to
func (_m *UseCase) GetAllUserEventsByDateRange(ctx context.Context, userId int, from time.Time, to time.Time) ([]events.Domain, error) {
	ret := _m.Called(ctx, userId, from, to)

	var r0 []events.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, time.Time, time.Time) []events.Domain); ok {
		r0 = rf(ctx, userId, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, time.Time, time.Time) error); ok {
		r1 = rf(ctx, userId, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScheduleEvent provides a mock function with given fields: ctx, event
func (_m *UseCase) ScheduleEvent(ctx context.Context, event *events.Domain) (int, error) {
	ret := _m.Called(ctx, event)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *events.Domain) int); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *events.Domain) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: ctx, event
func (_m *UseCase) UpdateEvent(ctx context.Context, event *events.Domain) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *events.Domain) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
